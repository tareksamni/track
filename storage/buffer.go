package storage

const insertBufSize = 100

type InsertBuffer struct {
	buf []TableRecord
}

func NewInsertBuffer(n int) *InsertBuffer {
	return &InsertBuffer{
		buf: make([]TableRecord, 0, n),
	}
}

func (s *InsertBuffer) Flush() (err error) {
	//tmp := make([]TableRecord, len(s.buf))
	//copy(tmp, s.buf)
	//go InsertMulti(tmp)
	err = InsertMulti(s.buf)

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
