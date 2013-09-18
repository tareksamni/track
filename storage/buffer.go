package storage

type Buffer interface {
	Flush() error
	Add(interface{}) error
}

type EventBuffer struct {
	buf []interface{}
}

func NewEventBuffer(n int) *EventBuffer {
	return &EventBuffer{
		buf: make([]interface{}, 0, n),
	}
}

func (s *EventBuffer) Flush() error {
	//tmp := make([]interface{}, len(s.buf))
	//copy(tmp, s.buf)
	//go InsertEvents(tmp)
	err := InsertEvents(s.buf)

	for i := 0; i < len(s.buf); i++ {
		s.buf[i] = nil
	}

	s.buf = s.buf[:0]
	return err
}

func (s *EventBuffer) Add(ev interface{}) error {
	if len(s.buf) == cap(s.buf) {
		if err := s.Flush(); err != nil {
			return err
		}
	}

	s.buf = append(s.buf, ev)
	return nil
}
