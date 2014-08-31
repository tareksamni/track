package storage

import (
	"github.com/tideland/goas/v2/monitoring"
	"github.com/simonz05/util/log"
)

const insertBufSize = 100

type InsertBuffer struct {
	buf   []TableRecord
	async bool
}

func NewInsertBuffer(n int, async bool) *InsertBuffer {
	return &InsertBuffer{
		buf:   make([]TableRecord, 0, n),
		async: async,
	}
}

func (s *InsertBuffer) Flush() (err error) {
	if s.async {
		tmp := make([]TableRecord, len(s.buf))
		copy(tmp, s.buf)
		go func() {
			m := monitoring.BeginMeasuring("flush-insert-buffer")
			if err := InsertMulti(tmp); err != nil {
				log.Errorf("Storage Err %v", err)
			}
			m.EndMeasuring()
		}()
	} else {
		m := monitoring.BeginMeasuring("flush-insert-buffer")
		defer m.EndMeasuring()
		if err = InsertMulti(s.buf); err != nil {
			return
		}
	}

	for i := 0; i < len(s.buf); i++ {
		s.buf[i] = nil
	}

	s.buf = s.buf[:0]
	return err
}

func (s *InsertBuffer) Add(row TableRecord) error {
	if len(s.buf) == cap(s.buf) {
		if err := s.Flush(); err != nil {
			return err
		}
	}

	s.buf = append(s.buf, row)
	return nil
}
