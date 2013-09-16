package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/simonz05/track/util"
	"github.com/simonz05/track/storage"
)

func init() {
	util.LogLevel = 0
}

var dataDecoder = schema.NewDecoder()

func dataError(w http.ResponseWriter, err string) {
	util.Logf("err: %v", err)
	w.WriteHeader(400)
	fmt.Fprintf(w, err)
}

func appError(w http.ResponseWriter, err error) {
	util.Logf("err: %v", err)
	w.WriteHeader(501)
	fmt.Fprintln(w, err)
}

func sessionHandle(w http.ResponseWriter, r *http.Request) {
	//util.Logf("Session Handle")
	session := new(storage.Session)

	if err := r.ParseForm(); err != nil {
		appError(w, err)
		return
	}

	if err := dataDecoder.Decode(session, r.PostForm); err != nil {
		dataError(w, fmt.Sprintln(err))
		return
	}

	util.Logln(session)
	// TODO validate that required fields exists
	// TODO Store session event
	w.WriteHeader(201)
}

func userHandle(w http.ResponseWriter, r *http.Request) {
	util.Logf("User Handle")
	w.WriteHeader(201)
}

func itemHandle(w http.ResponseWriter, r *http.Request) {
	util.Logf("Item Handle")
	w.WriteHeader(201)
}

func purchaseHandle(w http.ResponseWriter, r *http.Request) {
	util.Logf("Purchase Handle")
	w.WriteHeader(201)
}
