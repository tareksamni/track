package storage

import (
	"errors"
	"time"
)

var typeErr = errors.New("Invalid Type")

type Table interface {
	Table() string
	Columns() []string
}

type Record interface {
	Values() []interface{}
}

type TableRecord interface {
	Table
	Record
}

type Session struct {
	ProfileID   int       // 100
	Region      string    // BR
	SessionID   string    // 123ABCDFG
	RemoteIP    string    // 127.0.0.1
	SessionType string    // Web
	Message     string    // PageView
	Created     time.Time `schema:"-"`
}

func (s *Session) Table() string {
	return "SessionEvent"
}

func (s *Session) Columns() []string {
	return []string{"Region", "SessionID", "ProfileID", "RemoteIP", "SessionType", "Created", "Message"}
}

func (s *Session) Values() []interface{} {
	return []interface{}{s.Region, s.SessionID, s.ProfileID, s.RemoteIP, s.SessionType, s.Created, s.Message}
}

type User struct {
	ProfileID int
	Region    string
	Referrer  string
	Message   string
	Created   time.Time `schema:"-"`
}

func (u *User) Table() string {
	return "UserEvent"
}

func (u *User) Columns() []string {
	return []string{"Region", "ProfileID", "Referrer", "Created", "Message"}
}

func (u *User) Values() []interface{} {
	return []interface{}{u.Region, u.ProfileID, u.Referrer, u.Created, u.Message}
}

type Item struct {
	ProfileID   int
	Region      string
	ItemName    string
	ItemType    string
	IsUGC       bool
	PriceGold   int
	PriceSilver int
	Created     time.Time `schema:"-"`
}

func (i *Item) Table() string {
	return "ItemEvent"
}

func (i *Item) Columns() []string {
	return []string{"Region", "ProfileID", "ItemName", "ItemType", "IsUGC", "PriceGold", "PriceSilver", "Created"}
}

func (i *Item) Values() []interface{} {
	return []interface{}{i.Region, i.ProfileID, i.ItemName, i.ItemType, i.IsUGC, i.PriceGold, i.PriceSilver, i.Created}
}

type Purchase struct {
	ProfileID       int
	Region          string
	Currency        string
	GrossAmount     int
	NetAmount       int
	PaymentProvider string
	Product         string
	Created         time.Time `schema:"-"`
}

func (p *Purchase) Table() string {
	return "PurchaseEvent"
}

func (p *Purchase) Columns() []string {
	return []string{"Region", "ProfileID", "Currency", "GrossAmount", "NetAmount", "PaymentProvider", "Product", "Created"}
}

func (p *Purchase) Values() []interface{} {
	return []interface{}{p.Region, p.ProfileID, p.Currency, p.GrossAmount, p.NetAmount, p.PaymentProvider, p.Product, p.Created}
}
