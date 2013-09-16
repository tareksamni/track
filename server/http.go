package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/simonz05/profanity/util"
)

func dataError(w http.ResponseWriter, error string) {
	w.WriteHeader(400)
	fmt.Fprintf(w, error)
}

func sessionHandle(w http.ResponseWriter, r *http.Request) {
	util.Logf("Session Handle")
	w.WriteHeader(201)
}

func userHandle(w http.ResponseWriter, r *http.Request) {
	util.Logf("User Handle")

	profileID, err := strconv.Atoi(r.PostFormValue("ProfileID"))
	if err != nil {
		util.Logln(err)
		dataError(w, "Expected ProfileID")
	}

	util.Logf("ProfileID %d", profileID)
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
