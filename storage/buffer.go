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

type SessionBuffer struct {
	buf []*Session
}

func NewSessionBuffer(n int) *SessionBuffer {
	return &SessionBuffer{
		buf: make([]*Session, 0, n),
	}
}

func (s *SessionBuffer) Flush() error {
	err := InsertSessions(s.buf)

	for i := 0; i < len(s.buf); i++ {
		s.buf[i] = nil
	}

	s.buf = s.buf[:0]
	return err
}

func (s *SessionBuffer) Add(ev Event) error {
	v, ok := ev.(*Session)

	if !ok {
		return typeErr
	}

	if len(s.buf) == cap(s.buf) {
		if err := s.Flush(); err != nil {
			return err
		}
	}

	s.buf = append(s.buf, v)
	return nil
}

type UserBuffer struct {
	buf []*User
}

func NewUserBuffer(n int) *UserBuffer {
	return &UserBuffer{
		buf: make([]*User, 0, n),
	}
}

func (s *UserBuffer) Flush() error {
	err := InsertUsers(s.buf)

	for i := 0; i < len(s.buf); i++ {
		s.buf[i] = nil
	}

	s.buf = s.buf[:0]
	return err
}

func (s *UserBuffer) Add(ev Event) error {
	v, ok := ev.(*User)

	if !ok {
		return typeErr
	}

	if len(s.buf) == cap(s.buf) {
		if err := s.Flush(); err != nil {
			return err
		}
	}

	s.buf = append(s.buf, v)
	return nil
}

type ItemBuffer struct {
	buf []*Item
}

func NewItemBuffer(n int) *ItemBuffer {
	return &ItemBuffer{
		buf: make([]*Item, 0, n),
	}
}

func (s *ItemBuffer) Flush() error {
	err := InsertItems(s.buf)

	for i := 0; i < len(s.buf); i++ {
		s.buf[i] = nil
	}

	s.buf = s.buf[:0]
	return err
}

func (s *ItemBuffer) Add(ev Event) error {
	v, ok := ev.(*Item)

	if !ok {
		return typeErr
	}

	if len(s.buf) == cap(s.buf) {
		if err := s.Flush(); err != nil {
			return err
		}
	}

	s.buf = append(s.buf, v)
	return nil
}

type PurchaseBuffer struct {
	buf []*Purchase
}

func NewPurchaseBuffer(n int) *PurchaseBuffer {
	return &PurchaseBuffer{
		buf: make([]*Purchase, 0, n),
	}
}

func (s *PurchaseBuffer) Flush() error {
	err := InsertPurchases(s.buf)

	for i := 0; i < len(s.buf); i++ {
		s.buf[i] = nil
	}

	s.buf = s.buf[:0]
	return err
}

func (s *PurchaseBuffer) Add(ev Event) error {
	v, ok := ev.(*Purchase)

	if !ok {
		return typeErr
	}
	if len(s.buf) == cap(s.buf) {
		if err := s.Flush(); err != nil {
			return err
		}
	}

	s.buf = append(s.buf, v)
	return nil
}
