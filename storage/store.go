package storage

import (
	"errors"
	//"github.com/simonz05/track/util"
)

// Expects rows of same Table
func InsertMulti(rows []TableRecord) (err error) {
	if len(rows) == 0 {
		return
	}

	args, err := recordArgs(rows)

	if err != nil {
		return err
	}

	table := rows[0]
	n := len(rows)

	if n != insertBufSize {
		// If we hit this code path it is slower because we have to create the
		// prepared statement before doing the insert.
		q := createInsertQuery(table, n)
		_, err = Db.Exec(q, args...)
	} else {
		stmt, err := defaultStmtCache.GetInsert(table, n)

		if err != nil {
			return err
		}

		_, err = stmt.Exec(args...)
	}

	return
}

func recordArgs(rows []TableRecord) ([]interface{}, error) {
	if len(rows) == 0 {
		return nil, errors.New("Expected at least one row")
	}

	table := rows[0]
	name := table.Table()
	args := make([]interface{}, 0, len(rows)*len(table.Columns()))

	for i := 0; i < len(rows); i++ {
		if name != rows[i].Table() {
			return nil, typeErr
		}
		args = append(args, rows[i].Values()...)
	}

	return args, nil
}
