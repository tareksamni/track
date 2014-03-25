package storage

import (
	"errors"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
)

var ErrNil = errors.New("db: nil returned")

func SetupDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(32)
	//db.SetMaxOpenConns(64)

	Db = db

	if err := createUserTable(); err != nil {
		return nil, err
	}
	if err := createSessionTable(); err != nil {
		return nil, err
	}
	if err := createItemTable(); err != nil {
		return nil, err
	}
	if err := createPurchaseTable(); err != nil {
		return nil, err
	}
	defaultStmtCache = newStmtCache()
	return db, err
}

func createUserTable() (err error) {
	q := `CREATE TABLE IF NOT EXISTS UserEvent (
	  EventID int(10) unsigned NOT NULL AUTO_INCREMENT,
	  Region varchar(16) COLLATE utf8_unicode_ci NOT NULL,
	  ProfileID int(11) NOT NULL,
	  Referrer varchar(64) COLLATE utf8_unicode_ci DEFAULT NULL,
	  Message varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
	  Created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	  PRIMARY KEY (EventID)
	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci`
	_, err = Db.Exec(q)
	return err
}

func createSessionTable() (err error) {
	q := `CREATE TABLE IF NOT EXISTS SessionEvent (
	  EventID int(10) unsigned NOT NULL AUTO_INCREMENT,
	  Region varchar(16) COLLATE utf8_unicode_ci NOT NULL,
	  SessionID varchar(32) COLLATE utf8_unicode_ci NOT NULL,
	  ProfileID int(11) DEFAULT NULL,
	  RemoteIP varchar(45) COLLATE utf8_unicode_ci NOT NULL,
	  SessionType varchar(32) COLLATE utf8_unicode_ci NOT NULL,
	  Message varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
	  Created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	  PRIMARY KEY (EventID)
	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci`
	_, err = Db.Exec(q)
	return err
}

func createItemTable() (err error) {
	q := `CREATE TABLE IF NOT EXISTS ItemEvent (
	  EventID int(10) unsigned NOT NULL AUTO_INCREMENT,
	  Region varchar(16) COLLATE utf8_unicode_ci NOT NULL,
	  ProfileID int(11) NOT NULL,
	  ItemName varchar(64) COLLATE utf8_unicode_ci NOT NULL,
	  ItemType varchar(45) COLLATE utf8_unicode_ci NOT NULL,
	  IsUGC tinyint(1) NOT NULL,
	  IsRented tinyint(1) NOT NULL DEFAULT '0',
	  PriceGold int(11) DEFAULT NULL,
	  PriceSilver int(11) DEFAULT NULL,
	  Created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	  PRIMARY KEY (EventID)
	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci`
	_, err = Db.Exec(q)
	return err
}

func createPurchaseTable() (err error) {
	q := `CREATE TABLE IF NOT EXISTS PurchaseEvent (
	  EventID int(10) unsigned NOT NULL AUTO_INCREMENT,
	  Region varchar(16) COLLATE utf8_unicode_ci NOT NULL,
	  ProfileID int(11) NOT NULL,
	  Currency varchar(3) COLLATE utf8_unicode_ci NOT NULL,
	  GrossAmount decimal(19,4) NOT NULL,
	  NetAmount decimal(19,4) NOT NULL,
	  PaymentProvider varchar(45) COLLATE utf8_unicode_ci NOT NULL,
	  Product varchar(45) COLLATE utf8_unicode_ci NOT NULL,
	  Created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	  PRIMARY KEY (EventID)
	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci`
	_, err = Db.Exec(q)
	return err
}
