-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS UserEvent (
    EventID int(10) unsigned NOT NULL AUTO_INCREMENT,
    Region varchar(16) COLLATE utf8_unicode_ci NOT NULL,
    ProfileID int(11) NOT NULL,
    Referrer varchar(64) COLLATE utf8_unicode_ci DEFAULT NULL,
    Message varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
    Created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    Language varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
    PRIMARY KEY (EventID)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE IF NOT EXISTS SessionEvent (
    EventID int(10) unsigned NOT NULL AUTO_INCREMENT,
    Region varchar(16) COLLATE utf8_unicode_ci NOT NULL,
    SessionID varchar(32) COLLATE utf8_unicode_ci NOT NULL,
    ProfileID int(11) DEFAULT NULL,
    RemoteIP varchar(45) COLLATE utf8_unicode_ci NOT NULL,
    SessionType varchar(32) COLLATE utf8_unicode_ci NOT NULL,
    Message varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
    Created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    Language varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
    PRIMARY KEY (EventID)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE IF NOT EXISTS ItemEvent (
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
    Language varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
    PRIMARY KEY (EventID)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE IF NOT EXISTS PurchaseEvent (
    EventID int(10) unsigned NOT NULL AUTO_INCREMENT,
    Region varchar(16) COLLATE utf8_unicode_ci NOT NULL,
    ProfileID int(11) NOT NULL,
    Currency varchar(3) COLLATE utf8_unicode_ci NOT NULL,
    GrossAmount decimal(10,2) NOT NULL,
    NetAmount decimal(10,2) NOT NULL,
    PaymentProvider varchar(45) COLLATE utf8_unicode_ci NOT NULL,
    Product varchar(45) COLLATE utf8_unicode_ci NOT NULL,
    Created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    Language varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
    PRIMARY KEY (EventID)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- DROP TABLE UserEvent;
-- DROP TABLE SessionEvent;
-- DROP TABLE ItemEvent;
-- DROP TABLE PurchaseEvent;
