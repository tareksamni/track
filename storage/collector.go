package storage

import (
	"sync"
	"time"

	"github.com/simonz05/util/log"
)

type Queue struct {
	done *sync.WaitGroup
	buf  *InsertBuffer
	Chan chan TableRecord
	ref  int
}

var ref int

func NewInsertQueue(done *sync.WaitGroup) *Queue {
	qu := make(chan TableRecord, 256)
	var q *Queue

	for i := 0; i < 128; i++ {
		ref += 1
		q = &Queue{
			ref:  ref,
			done: done,
			Chan: qu,
			buf:  NewInsertBuffer(insertBufSize, false),
		}
		done.Add(1)
		go q.collect()
	}

	// only public field is shared among all workers
	return q
}

func (q *Queue) collect() {
	log.Printf("[%d] Queue Starting …", q.ref)
	defer q.onExit()

	if q.buf == nil {
		log.Errorln("InsertBuffer was nil")
		return
	}

	for {
		var err error

		select {
		case v, ok := <-q.Chan:
			if !ok {
				return
			}

			log.Printf("Got Table Record")
			err = q.buf.Add(v)
		case <-time.After(time.Second * 1):
			err = q.buf.Flush()
		}

		if err != nil {
			log.Errorf("Storage Err %v", err)
		}
	}
}

func (q *Queue) onExit() {
	log.Printf("[%d] Queue Exit Started", q.ref)

	if q.buf != nil {
		if err := q.buf.Flush(); err != nil {
			log.Errorf("[%d] Queue Exit ERR %v", q.ref, err)
		}
	}

	log.Printf("[%d] Queue Exit OK", q.ref)
	q.done.Done()
}
