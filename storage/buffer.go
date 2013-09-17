package storage

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

func (s *SessionBuffer) Add(ses *Session) error {
	if len(s.buf) == cap(s.buf) {
		if err := s.Flush(); err != nil {
			return err
		}
	}

	s.buf = append(s.buf, ses)
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

func (s *UserBuffer) Add(ses *User) error {
	if len(s.buf) == cap(s.buf) {
		if err := s.Flush(); err != nil {
			return err
		}
	}

	s.buf = append(s.buf, ses)
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

func (s *ItemBuffer) Add(ses *Item) error {
	if len(s.buf) == cap(s.buf) {
		if err := s.Flush(); err != nil {
			return err
		}
	}

	s.buf = append(s.buf, ses)
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

func (s *PurchaseBuffer) Add(ses *Purchase) error {
	if len(s.buf) == cap(s.buf) {
		if err := s.Flush(); err != nil {
			return err
		}
	}

	s.buf = append(s.buf, ses)
	return nil
}
