package storage

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/simonz05/track/util"
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

func InsertEvents(any []interface{}) (err error) {
	if len(any) > 0 {
		switch any[0].(type) {
		case *Session:
			err = InsertSessions(any)
		case *User:
			err = InsertUsers(any)
		case *Item:
			err = InsertItems(any)
		case *Purchase:
			err = InsertPurchases(any)
		default:
			err = typeErr
		}
	}

	if err != nil {
		util.Errln(err)
	}
	return err
}

func InsertSession(ses *Session) error {
	_, err := InsertSessionStmt.Exec(ses.Region, ses.SessionID, ses.ProfileID, ses.RemoteIP, ses.SessionType, ses.Created, ses.Message)
	return err
}

func InsertSessions(any []interface{}) (err error) {
	args, err := flattenSessions(any)

	if err != nil {
		return err
	}

	if len(any) != bufSize {
		// If we hit this code path it is slower because we have to create the
		// prepared statement before doing the insert.
		q := sessionInsertQuery(len(any))
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

func InsertUsers(any []interface{}) (err error) {
	args, err := flattenUsers(any)

	if err != nil {
		return err
	}

	if len(any) != bufSize {
		q := userInsertQuery(len(any))
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

func InsertItems(any []interface{}) (err error) {
	args, err := flattenItems(any)

	if err != nil {
		return err
	}

	if len(any) != bufSize {
		q := itemInsertQuery(len(any))
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

func InsertPurchases(any []interface{}) (err error) {
	args, err := flattenPurchases(any)

	if err != nil {
		return err
	}

	if len(any) != bufSize {
		q := purchaseInsertQuery(len(any))
		_, err = Db.Exec(q, args...)
	} else {
		_, err = InsertBulkPurchaseStmt.Exec(args...)
	}

	return err
}

func flattenSessions(any []interface{}) (args []interface{}, err error) {
	columns := 7
	args = make([]interface{}, 0, len(any)*columns)

	for i := 0; i < len(any); i++ {
		ses, ok := any[i].(*Session)

		if !ok {
			return nil, err
		}

		args = append(args, ses.Region, ses.SessionID, ses.ProfileID, ses.RemoteIP, ses.SessionType, ses.Created, ses.Message)
	}

	return
}

func flattenUsers(any []interface{}) (args []interface{}, err error) {
	columns := 5
	args = make([]interface{}, 0, len(any)*columns)

	for i := 0; i < len(any); i++ {
		user, ok := any[i].(*User)

		if !ok {
			return nil, err
		}

		args = append(args, user.Region, user.ProfileID, user.Referrer, user.Created, user.Message)
	}

	return
}

func flattenItems(any []interface{}) (args []interface{}, err error) {
	columns := 8
	args = make([]interface{}, 0, len(any)*columns)

	for i := 0; i < len(any); i++ {
		item, ok := any[i].(*Item)

		if !ok {
			return nil, err
		}

		args = append(args, item.Region, item.ProfileID, item.ItemName, item.ItemType, item.IsUGC, item.PriceGold, item.PriceSilver, item.Created)
	}

	return
}

func flattenPurchases(any []interface{}) (args []interface{}, err error) {
	columns := 8
	args = make([]interface{}, 0, len(any)*columns)

	for i := 0; i < len(any); i++ {
		purchase, ok := any[i].(*Purchase)

		if !ok {
			return nil, err
		}

		args = append(args, purchase.Region, purchase.ProfileID, purchase.Currency, purchase.GrossAmount, purchase.NetAmount, purchase.PaymentProvider, purchase.Product, purchase.Created)
	}

	return
}
