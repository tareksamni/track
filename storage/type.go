package storage

import (
	"errors"
	"regexp"
	"time"
)

var typeErr = errors.New("Invalid Type")
var InvalidRegionErr = errors.New("Invalid Region")
var RequiredFieldErr = errors.New("Required Field was empty")

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

type Validator interface {
	Validate() error
}

var regionValidator = regexp.MustCompile("^[a-zA-Z]{2,3}$")

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

func (s *Session) Validate() error {
	if !regionValidator.MatchString(s.Region) {
		return InvalidRegionErr
	}

	if s.SessionID == "" || s.RemoteIP == "" || s.SessionType == "" {
		return RequiredFieldErr
	}

	return nil
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

func (u *User) Validate() error {
	if !regionValidator.MatchString(u.Region) {
		return InvalidRegionErr
	}

	if u.ProfileID == 0 {
		return RequiredFieldErr
	}

	return nil
}

type Item struct {
	ProfileID   int
	Region      string
	ItemName    string
	ItemType    string
	IsUGC       bool
	IsRented    bool
	PriceGold   int
	PriceSilver int
	Created     time.Time `schema:"-"`
}

func (i *Item) Table() string {
	return "ItemEvent"
}

func (i *Item) Columns() []string {
	return []string{"Region", "ProfileID", "ItemName", "ItemType", "IsUGC", "IsRented", "PriceGold", "PriceSilver", "Created"}
}

func (i *Item) Values() []interface{} {
	return []interface{}{i.Region, i.ProfileID, i.ItemName, i.ItemType, i.IsUGC, i.IsRented, i.PriceGold, i.PriceSilver, i.Created}
}

func (i *Item) Validate() error {
	if !regionValidator.MatchString(i.Region) {
		return InvalidRegionErr
	}

	if i.ProfileID == 0 || i.ItemName == "" || i.ItemType == "" {
		return RequiredFieldErr
	}

	return nil
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

func (p *Purchase) Validate() error {
	if !regionValidator.MatchString(p.Region) {
		return InvalidRegionErr
	}

	if p.ProfileID == 0 || p.Currency == "" || p.GrossAmount == 0 || p.NetAmount == 0 || p.PaymentProvider == "" || p.Product == "" {
		return RequiredFieldErr
	}

	return nil
}
