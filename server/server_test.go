package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/simonz05/track/util"
)

var (
	once       sync.Once
	serverAddr string
	server     *httptest.Server
)

func startServer() {
	util.LogLevel = 0
	err := setupServer("")

	if err != nil {
		panic(err)
	}

	server = httptest.NewServer(nil)
	serverAddr = server.Listener.Addr().String()
}

type SessionTest struct {
	ProfileID   int
	SessionID   string
	RemoteIP    string
	SessionType string
	Message     string
}

type UserTest struct {
	ProfileID int
	Referrer  string
	Message   string
}

type ItemTest struct {
	ProfileID   int
	ItemName    string
	ItemType    string
	IsUGC       bool
	PriceGold   int
	PriceSilver int
}

type PruchaseTest struct {
	ProfileID       int
	Currency        string
	GrossAmount     int
	NetAmount       int
	PaymentProvider string
	Product         string
}

type TestCase struct {
	Data       interface{}
	StatusCode int
}

func TestSession(t *testing.T) {
	once.Do(startServer)

	tests := []*TestCase{
		{&SessionTest{
			ProfileID:   1,
			SessionID:   "abc",
			RemoteIP:    "127.0.0.1",
			SessionType: "session type"}, 201},
	}

	for i, x := range tests {
		doHttp(t, i, "session", x.Data, x.StatusCode)
	}
}

func doHttp(t *testing.T, index int, endpoint string, data interface{}, statusCode int) {
	values := url.Values{}
	var uri string

	s := reflect.ValueOf(data).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		values.Add(typeOfT.Field(i).Name, fmt.Sprintf("%v", f.Interface()))
		//fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	params := strings.NewReader(values.Encode())
	uri = fmt.Sprintf("http://%s/api/1.0/track/%s/", serverAddr, endpoint)

	req, _ := http.NewRequest("POST", uri, params)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	r, err := client.Do(req)

	if err != nil {
		t.Fatalf("error posting: %s", err)
		return
	}

	if r.StatusCode != statusCode {
		t.Fatalf("expected status code %d, got %d", statusCode, r.StatusCode)
	}
}

func BenchmarkServer(b *testing.B) {
	once.Do(startServer)

	values := url.Values{}
	values.Add("ProfileID", "1")
	values.Add("SessionID", "abc")
	values.Add("RemoteIP", "127.0.0.1")
	values.Add("SessionType", "session type")

	params := strings.NewReader(values.Encode())
	uri := fmt.Sprintf("http://%s/api/1.0/track/session/", serverAddr)

	client := &http.Client{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", uri, params)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		_, err := client.Do(req)

		if err != nil {
			b.Fatalf("error posting: %s", err)
		}
	}
}
