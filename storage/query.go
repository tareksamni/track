package storage

import (
	"database/sql"
	"fmt"
	"strings"
)

var (
	InsertSessionStmt      *sql.Stmt
	InsertBulkSessionStmt  *sql.Stmt
	InsertUserStmt         *sql.Stmt
	InsertBulkUserStmt     *sql.Stmt
	InsertItemStmt         *sql.Stmt
	InsertBulkItemStmt     *sql.Stmt
	InsertPurchaseStmt     *sql.Stmt
	InsertBulkPurchaseStmt *sql.Stmt
)

/* Currently the golang db interface does not support bulk inserts
   We use a prepared bulk insert with a fixed buffer size to work around this.

   http://code.google.com/p/go/issues/detail?id=5171
*/
func setupQueries() (err error) {
	if InsertSessionStmt, err = Db.Prepare(sessionInsertQuery(1)); err != nil {
		return err
	}

	if InsertBulkSessionStmt, err = Db.Prepare(sessionInsertQuery(bufSize)); err != nil {
		return err
	}

	if InsertUserStmt, err = Db.Prepare(userInsertQuery(1)); err != nil {
		return err
	}

	if InsertBulkUserStmt, err = Db.Prepare(userInsertQuery(bufSize)); err != nil {
		return err
	}

	if InsertItemStmt, err = Db.Prepare(itemInsertQuery(1)); err != nil {
		return err
	}

	if InsertBulkItemStmt, err = Db.Prepare(itemInsertQuery(bufSize)); err != nil {
		return err
	}

	if InsertPurchaseStmt, err = Db.Prepare(purchaseInsertQuery(1)); err != nil {
		return err
	}

	if InsertBulkPurchaseStmt, err = Db.Prepare(purchaseInsertQuery(bufSize)); err != nil {
		return err
	}

	return nil
}

func sessionInsertQuery(bulkCount int) string {
	return genericInsertQuery("SessionEvent", []string{"Region", "SessionID", "ProfileID", "RemoteIP", "SessionType", "Created", "Message"}, bulkCount)
}

func userInsertQuery(bulkCount int) string {
	return genericInsertQuery("UserEvent", []string{"Region", "ProfileID", "Referrer", "Created", "Message"}, bulkCount)
}

func itemInsertQuery(bulkCount int) string {
	return genericInsertQuery("ItemEvent", []string{"Region", "ProfileID", "ItemName", "ItemType", "IsUGC", "PriceGold", "PriceSilver", "Created"}, bulkCount)
}

func purchaseInsertQuery(bulkCount int) string {
	return genericInsertQuery("PurchaseEvent", []string{"Region", "ProfileID", "Currency", "GrossAmount", "NetAmount", "PaymentProvider", "Product", "Created"}, bulkCount)
}

func genericInsertQuery(table string, columns []string, n int) string {
	columnNames := strings.Join(columns, ", ")
	columnRepl := strings.Repeat("?, ", len(columns))
	columnRepl = fmt.Sprintf("(%s), ", columnRepl[:len(columnRepl)-2])
	columnRepl = strings.Repeat(columnRepl, n)
	columnRepl = columnRepl[:len(columnRepl)-2]

	return fmt.Sprintf("INSERT INTO %s(%s) VALUES%s", table, columnNames, columnRepl)
}

func InsertSession(ses *Session) error {
	_, err := InsertSessionStmt.Exec(ses.Region, ses.SessionID, ses.ProfileID, ses.RemoteIP, ses.SessionType, ses.Created, ses.Message)
	return err
}

func InsertSessions(ses []*Session) (err error) {
	if len(ses) == 0 {
		return
	}

	args := make([]interface{}, 0, len(ses))

	for i := 0; i < len(ses); i++ {
		args = append(args, ses[i].Region, ses[i].SessionID, ses[i].ProfileID, ses[i].RemoteIP, ses[i].SessionType, ses[i].Created, ses[i].Message)
	}

	if len(ses) != bufSize {
		// If we hit this code path it is slower because we have to create the
		// prepared statement before doing the insert.
		q := sessionInsertQuery(len(ses))
		_, err = Db.Exec(q, args...)
	} else {
		_, err = InsertBulkSessionStmt.Exec(args...)
	}

	return err
}

func InsertUser(ses *User) error {
	_, err := InsertUserStmt.Exec(ses.Region, ses.ProfileID, ses.Referrer, ses.Created, ses.Message)
	return err
}

func InsertUsers(ses []*User) (err error) {
	if len(ses) == 0 {
		return
	}
	args := make([]interface{}, 0, len(ses))

	for i := 0; i < len(ses); i++ {
		args = append(args, ses[i].Region, ses[i].ProfileID, ses[i].Referrer, ses[i].Created, ses[i].Message)
	}

	if len(ses) != bufSize {
		q := userInsertQuery(len(ses))
		_, err = Db.Exec(q, args...)
	} else {
		_, err = InsertBulkUserStmt.Exec(args...)
	}

	return err
}

func InsertItem(ses *Item) error {
	_, err := InsertItemStmt.Exec(ses.Region, ses.ProfileID, ses.ItemName, ses.ItemType, ses.IsUGC, ses.PriceGold, ses.PriceSilver, ses.Created)
	return err
}

func InsertItems(ses []*Item) (err error) {
	if len(ses) == 0 {
		return
	}
	args := make([]interface{}, 0, len(ses))

	for i := 0; i < len(ses); i++ {
		args = append(args, ses[i].Region, ses[i].ProfileID, ses[i].ItemName, ses[i].ItemType, ses[i].IsUGC, ses[i].PriceGold, ses[i].PriceSilver, ses[i].Created)
	}

	if len(ses) != bufSize {
		q := itemInsertQuery(len(ses))
		_, err = Db.Exec(q, args...)
	} else {
		_, err = InsertBulkItemStmt.Exec(args...)
	}

	return err
}

func InsertPurchase(ses *Purchase) error {
	_, err := InsertPurchaseStmt.Exec(ses.Region, ses.ProfileID, ses.Currency, ses.GrossAmount, ses.NetAmount, ses.PaymentProvider, ses.Product, ses.Created)
	return err
}

func InsertPurchases(ses []*Purchase) (err error) {
	if len(ses) == 0 {
		return
	}
	args := make([]interface{}, 0, len(ses))

	for i := 0; i < len(ses); i++ {
		args = append(args, ses[i].Region, ses[i].ProfileID, ses[i].Currency, ses[i].GrossAmount, ses[i].NetAmount, ses[i].PaymentProvider, ses[i].Product, ses[i].Created)
	}

	if len(ses) != bufSize {
		q := purchaseInsertQuery(len(ses))
		_, err = Db.Exec(q, args...)
	} else {
		_, err = InsertBulkPurchaseStmt.Exec(args...)
	}

	return err
}
