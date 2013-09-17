package storage

import (
	"errors"
)

var typeErr = errors.New("Invalid Type")

type Event interface{}

type Buffer interface{
	Flush() error
	Add(Event) error
}

type EventBuffer struct {
	buf []Event
}

func NewEventBuffer(n int) *EventBuffer {
	return &EventBuffer{
		buf: make([]Event, 0, n),
	}
}

func (s *EventBuffer) Flush() error {
	err := InsertEvents(s.buf)

	for i := 0; i < len(s.buf); i++ {
		s.buf[i] = nil
	}

	s.buf = s.buf[:0]
	return err
}

func (s *EventBuffer) Add(ev Event) error {
	if len(s.buf) == cap(s.buf) {
		if err := s.Flush(); err != nil {
			return err
		}
	}

	s.buf = append(s.buf, ev)
	return nil
}
