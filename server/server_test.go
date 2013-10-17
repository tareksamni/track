package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	once       sync.Once
	serverAddr string
	server     *httptest.Server
)

func startServer() {
	err := setupServer("travis@tcp(localhost:3306)/myapp_test?charset=utf8")

	if err != nil {
		panic(err)
	}

	startCollectors()

	server = httptest.NewServer(nil)
	serverAddr = server.Listener.Addr().String()
}

type SessionTest struct {
	ProfileID   int
	Region      string
	SessionID   string
	RemoteIP    string
	SessionType string
	Message     string
}

type UserTest struct {
	ProfileID int
	Region    string
	Referrer  string
	Message   string
}

type ItemTest struct {
	ProfileID   int
	Region      string
	ItemName    string
	ItemType    string
	IsUGC       bool
	PriceGold   int
	PriceSilver int
}

type PurchaseTest struct {
	ProfileID       int
	Region          string
	Currency        string
	GrossAmount     string
	NetAmount       string
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
			Region:      "BR",
			SessionID:   "abc",
			RemoteIP:    "127.0.0.1",
			SessionType: "session type"}, 201},
		{&SessionTest{
			ProfileID:   2,
			Region:      "BR",
			SessionID:   "abc",
			RemoteIP:    "127.0.0.1",
			SessionType: "session type"}, 201},
	}

	for i, x := range tests {
		doHttp(t, i, "session", x.Data, x.StatusCode)
	}
}

func TestPurchase(t *testing.T) {
	once.Do(startServer)

	tests := []*TestCase{
		{&PurchaseTest{
			ProfileID:       1,
			Region:          "BR",
			Currency:        "USD",
			GrossAmount:     "1.99",
			NetAmount:       "1.27",
			PaymentProvider: "PayPal",
			Product:         "100 Gold",
		}, 201},
	}

	for i, x := range tests {
		doHttp(t, i, "purchase", x.Data, x.StatusCode)
	}

	time.Sleep(1e09 * 3)
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
	defer r.Body.Close()
	//fmt.Println(values.Encode())

	if err != nil {
		t.Fatalf("error posting: %s", err)
		return
	}

	if r.StatusCode != statusCode {
		//fmt.Println(r.Body)
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Printf("%s\n", body)
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
