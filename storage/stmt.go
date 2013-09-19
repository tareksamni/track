package storage

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"
	//"github.com/simonz05/track/util"
)

var (
	defaultStmtCache *stmtCache
)

/* Currently the golang db interface does not support bulk inserts
   We use a prepared bulk insert with a fixed buffer size to work around this.

   http://code.google.com/p/go/issues/detail?id=5171
*/
type stmtCache struct {
	cache map[string]*sql.Stmt
	mu    sync.RWMutex
}

func newStmtCache() *stmtCache {
	return &stmtCache{
		cache: make(map[string]*sql.Stmt),
	}
}

func (s *stmtCache) GetInsert(table Table, n int) (*sql.Stmt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	stmt, ok := s.cache[table.Table()]

	if !ok {
		var err error
		stmt, err = createInsertStmt(table, n)

		if err != nil {
			return nil, err
		}

		s.cache[table.Table()] = stmt
	}

	return stmt, nil
}

func createInsertStmt(table Table, n int) (*sql.Stmt, error) {
	return Db.Prepare(createInsertQuery(table, n))
}

func createInsertQuery(table Table, n int) string {
	columns := table.Columns()
	columnNames := strings.Join(columns, ", ")
	columnRepl := strings.Repeat("?, ", len(columns))
	columnRepl = fmt.Sprintf("(%s), ", columnRepl[:len(columnRepl)-2])
	columnRepl = strings.Repeat(columnRepl, n)
	columnRepl = columnRepl[:len(columnRepl)-2]

	return fmt.Sprintf("INSERT INTO %s(%s) VALUES%s", table.Table(), columnNames, columnRepl)
}
