package storage

import (
	"time"

	"github.com/simonz05/track/util"
)

const bufSize = 100

type Queue struct {
	Session  chan *Session
	User     chan *User
	Item     chan *Item
	Purchase chan *Purchase
}

func NewQueue() *Queue {
	return &Queue{
		Session:  make(chan *Session, 100),
		User:     make(chan *User, 100),
		Item:     make(chan *Item, 100),
		Purchase: make(chan *Purchase, 100),
	}
}

func (q *Queue) Collect() {
	util.Logf("Queue Starting ")
	sessions := NewSessionBuffer(bufSize)
	users := NewUserBuffer(bufSize)
	items := NewItemBuffer(bufSize)
	purchases := NewPurchaseBuffer(bufSize)

	for {
		select {
		case v := <-q.Session:
			util.Logf("got session")
			err := sessions.Add(v)
			if err != nil {
				util.Logf("err %v", err)
			}
		case v := <-q.User:
			util.Logf("got user")
			err := users.Add(v)
			if err != nil {
				util.Logf("err %v", err)
			}
		case v := <-q.Item:
			util.Logf("got item")
			err := items.Add(v)
			if err != nil {
				util.Logf("err %v", err)
			}
		case v := <-q.Purchase:
			util.Logf("got purchase")
			err := purchases.Add(v)
			if err != nil {
				util.Logf("err %v", err)
			}
		case <-time.After(time.Second):
			util.Logf("timeout")
			sessions.Flush()
			users.Flush()
			items.Flush()
			purchases.Flush()
		}
	}
}
