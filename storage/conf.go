package storage

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/eaigner/jet"
)

var (
	Db *jet.Db
)

var ErrNil = errors.New("db: nil returned")

func setupDb(dsn string) (*jet.Db, error) {
	db, err := jet.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	return db, err
}
