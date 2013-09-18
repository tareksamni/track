package storage

import (
	"time"

	"github.com/simonz05/track/util"
)

const bufSize = 64

type Queue struct {
	buf  Buffer
	Chan chan interface{}
}

func NewEventQueue() *Queue {
	q := &Queue{
		Chan: make(chan interface{}, 100),
		buf: NewEventBuffer(bufSize),
	}
	go q.collect()
	return q
}

func (q *Queue) collect() {
	util.Logf("Queue Starting ")

	if q.buf == nil {
		panic("Buffer was nil")
	}

	for {
		select {
		case v := <-q.Chan:
			util.Logf("Got Event")
			err := q.buf.Add(v)

			if err != nil {
				util.Logf("err %v", err)
			}
		case <-time.After(time.Second):
			q.buf.Flush()
		}
	}
}
