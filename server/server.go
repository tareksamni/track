// Copyright (c) 2013 Simon Zimmermann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package server implements HTTP interface for track
// service.
package server

import (
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/simonz05/track/storage"
	"github.com/simonz05/track/util"
)

var (
	router        *mux.Router
	sessionQueue  *storage.Queue
	userQueue     *storage.Queue
	itemQueue     *storage.Queue
	purchaseQueue *storage.Queue
)

func sigTrapCloser(l net.Listener) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for _ = range c {
			l.Close()
			// TODO: Db Close
			util.Logf("Closed listener %s", l.Addr())
		}
	}()
}

func setupCollector() {
	sessionQueue = storage.NewInsertQueue()
	userQueue = storage.NewInsertQueue()
	itemQueue = storage.NewInsertQueue()
	purchaseQueue = storage.NewInsertQueue()
}

func setupServer(dsn string) error {
	if _, err := storage.SetupDb(dsn); err != nil {
		return err
	}

	setupCollector()

	// HTTP endpoints
	router = mux.NewRouter()
	router.HandleFunc("/api/1.0/track/session/", sessionHandle).Methods("GET", "POST").Name("session")
	router.HandleFunc("/api/1.0/track/user/", userHandle).Methods("POST").Name("user")
	router.HandleFunc("/api/1.0/track/item/", itemHandle).Methods("POST").Name("item")
	router.HandleFunc("/api/1.0/track/purchase/", purchaseHandle).Methods("POST").Name("purchase")
	router.StrictSlash(false)
	http.Handle("/", router)

	return nil
}

func ListenAndServe(laddr, dsn string) error {
	if err := setupServer(dsn); err != nil {
		return err
	}

	l, err := net.Listen("tcp", laddr)

	if err != nil {
		return err
	}

	util.Logf("Listen on %s", l.Addr())

	sigTrapCloser(l)
	err = http.Serve(l, nil)
	util.Logf("Shutting down ..")
	return err
}
