package server

type Session struct {
	ProfileID   int    // 100
	Region      string // BR
	SessionID   string // 123ABCDFG
	RemoteIP    string // 127.0.0.1
	SessionType string // Web
	Message     string // PageView
}

type User struct {
	ProfileID int
	Region    string
	Referrer  string
	Message   string
}

type Item struct {
	ProfileID   int
	Region      string
	ItemName    string
	ItemType    string
	IsUGC       bool
	PriceGold   int
	PriceSilver int
}

type Purchase struct {
	ProfileID       int
	Region          string
	Currency        string
	GrossAmount     int
	NetAmount       int
	PaymentProvider string
	Product         string
}
