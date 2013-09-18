package server

import (
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/simonz05/track/storage"
	"github.com/simonz05/track/util"
)

var dataDecoder = schema.NewDecoder()

func writeError(w http.ResponseWriter, err string, statusCode int) {
	util.Logf("err: %v", err)
	w.WriteHeader(statusCode)
	w.Write([]byte(err))
}

func sessionHandle(w http.ResponseWriter, r *http.Request) {
	util.Logf("Session Handle")
	ses := new(storage.Session)

	if err := r.ParseForm(); err != nil {
		writeError(w, err.Error(), 501)
		return
	}

	if err := dataDecoder.Decode(ses, r.PostForm); err != nil {
		writeError(w, err.Error(), 400)
		return
	}

	if err := ses.Validate(); err != nil {
		writeError(w, err.Error(), 400)
		return
	}

	ses.Created = time.Now()
	sessionQueue.Chan <- ses
	w.WriteHeader(201)
}

func userHandle(w http.ResponseWriter, r *http.Request) {
	util.Logf("User Handle")
	user := new(storage.User)

	if err := r.ParseForm(); err != nil {
		writeError(w, err.Error(), 501)
		return
	}

	if err := dataDecoder.Decode(user, r.PostForm); err != nil {
		writeError(w, err.Error(), 400)
		return
	}

	util.Logln(user)

	if err := user.Validate(); err != nil {
		writeError(w, err.Error(), 400)
		return
	}

	user.Created = time.Now()
	userQueue.Chan <- user
	w.WriteHeader(201)
}

func itemHandle(w http.ResponseWriter, r *http.Request) {
	util.Logf("Item Handle")
	item := new(storage.Item)

	if err := r.ParseForm(); err != nil {
		writeError(w, err.Error(), 501)
		return
	}

	if err := dataDecoder.Decode(item, r.PostForm); err != nil {
		writeError(w, err.Error(), 400)
		return
	}

	util.Logln(item)

	if err := item.Validate(); err != nil {
		writeError(w, err.Error(), 400)
		return
	}

	item.Created = time.Now()
	itemQueue.Chan <- item
	w.WriteHeader(201)
}

func purchaseHandle(w http.ResponseWriter, r *http.Request) {
	util.Logf("Purchase Handle")
	purchase := new(storage.Purchase)

	if err := r.ParseForm(); err != nil {
		writeError(w, err.Error(), 501)
		return
	}

	if err := dataDecoder.Decode(purchase, r.PostForm); err != nil {
		writeError(w, err.Error(), 400)
		return
	}

	if err := purchase.Validate(); err != nil {
		writeError(w, err.Error(), 400)
		return
	}

	purchase.Created = time.Now()
	util.Logln(purchase)
	purchaseQueue.Chan <- purchase
	w.WriteHeader(201)
}
