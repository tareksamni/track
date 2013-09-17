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
	Version = "0.1.0"
	router  *mux.Router
	queue   *storage.Queue
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
	queue = storage.NewQueue()
	go queue.Collect()
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
