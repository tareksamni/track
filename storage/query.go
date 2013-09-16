package storage

import (
	"time"
)

func CreateSession(ses *Session, created time.Time) (error) {
	q := `
	INSERT INTO SessionEvent()
	VALUES()`
	return Db.Query(q, created).Run()
}
