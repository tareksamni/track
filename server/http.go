package server

import (
	"math/big"
	"net/http"
	"reflect"

	"github.com/gorilla/schema"
	"github.com/simonz05/track/storage"
	"github.com/simonz05/util/log"
)

var (
	invalidValue = reflect.Value{}
	dataDecoder  = schema.NewDecoder()
)

func init() {
	dataDecoder.RegisterConverter(big.Rat{}, convertRat)
}

func convertRat(value string) reflect.Value {
	r := new(big.Rat)

	if v, ok := r.SetString(value); ok {
		return reflect.ValueOf(*v)
	}

	return invalidValue
}

func writeError(w http.ResponseWriter, err string, statusCode int) {
	log.Errorf("err: %v", err)
	w.WriteHeader(statusCode)
	w.Write([]byte(err))
}

func makeSesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventHandle(w, r, sessionQueue.Chan, storage.NewSession())
	}
}

func makeUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventHandle(w, r, userQueue.Chan, storage.NewUser())
	}
}

func makeItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventHandle(w, r, itemQueue.Chan, storage.NewItem())
	}
}

func makePurchaseHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventHandle(w, r, purchaseQueue.Chan, storage.NewPurchase())
	}
}

func eventHandle(w http.ResponseWriter, r *http.Request, q chan storage.TableRecord, t storage.TableValidator) {
	if err := r.ParseForm(); err != nil {
		writeError(w, err.Error(), 400)
		return
	}

	if err := dataDecoder.Decode(t, r.PostForm); err != nil {
		writeError(w, err.Error(), 400)
		return
	}

	if err := t.Validate(); err != nil {
		writeError(w, err.Error(), 400)
		return
	}

	q <- t
	w.WriteHeader(201)
}
