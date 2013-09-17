package server

import (
	"net/http"
	"regexp"
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

var regionValidator = regexp.MustCompile("^[a-zA-Z]{2,3}$")

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

	util.Logln(ses)

	if !regionValidator.MatchString(ses.Region) {
		writeError(w, "Invalid Region", 400)
		return
	}

	if ses.SessionID == "" || ses.RemoteIP == "" || ses.SessionType == "" {
		writeError(w, "Required field was empty", 400)
		return
	}

	ses.Created = time.Now()
	queue.Session <- ses
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

	if !regionValidator.MatchString(user.Region) {
		writeError(w, "Invalid Region", 400)
		return
	}

	if user.ProfileID == 0 {
		writeError(w, "Required field was empty", 400)
		return
	}

	user.Created = time.Now()
	queue.User <- user
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

	if !regionValidator.MatchString(item.Region) {
		writeError(w, "Invalid Region", 400)
		return
	}

	if item.ProfileID == 0 || item.ItemName == "" || item.ItemType == "" {
		writeError(w, "Required field was empty", 400)
		return
	}

	item.Created = time.Now()
	queue.Item <- item
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

	util.Logln(purchase)

	if !regionValidator.MatchString(purchase.Region) {
		writeError(w, "Invalid Region", 400)
		return
	}

	if purchase.ProfileID == 0 || purchase.Currency == "" || purchase.GrossAmount == 0 || purchase.NetAmount == 0 || purchase.PaymentProvider == "" || purchase.Product == "" {
		writeError(w, "Required field was empty", 400)
		return
	}

	purchase.Created = time.Now()
	queue.Purchase <- purchase
	w.WriteHeader(201)
}
