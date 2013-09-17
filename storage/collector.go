package storage

import (
	"time"

	"github.com/simonz05/track/util"
)

const bufSize = 100

type Queue struct {
	Buf	  Buffer
	Chan  chan Event
}

func newQueue() *Queue {
	return &Queue{
		Chan:  make(chan Event, 100),
	}
}

func NewEventQueue() *Queue {
	q := newQueue()
	q.Buf = NewEventBuffer(bufSize)
	go q.Collect()
	return q
}

func (q *Queue) Collect() {
	util.Logf("Queue Starting ")

	if q.Buf == nil {
		panic("Buffer was nil")
	}

	for {
		select {
		case v := <-q.Chan:
			util.Logf("got event")
			err := q.Buf.Add(v)
			if err != nil {
				util.Logf("err %v", err)
			}
		case <-time.After(time.Millisecond*500):
			util.Logf("timeout")
			q.Buf.Flush()
		}
	}
}
