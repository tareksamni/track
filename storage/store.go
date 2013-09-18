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

// Package storage implements the track storage abstraction
// on top of MySQL.
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
