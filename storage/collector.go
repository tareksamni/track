package storage

import (
	"time"

	"github.com/simonz05/track/util"
)

type Queue struct {
	buf  *InsertBuffer
	Chan chan TableRecord
}

func NewInsertQueue() *Queue {
	q := &Queue{
		Chan: make(chan TableRecord, 100),
		buf:  NewInsertBuffer(insertBufSize),
	}

	go q.collect()
	return q
}

func (q *Queue) collect() {
	util.Logf("Queue Starting …")

	if q.buf == nil {
		panic("InsertBuffer was nil")
	}

	for {
		select {
		case v := <-q.Chan:
			util.Logf("Got Table Record")
			err := q.buf.Add(v)

			if err != nil {
				util.Logf("err %v", err)
			}
		case <-time.After(time.Second):
			q.buf.Flush()
		}
	}
}
