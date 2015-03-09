
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE UserEvent ADD 
(
	`RemoteIP` varchar(45) COLLATE utf8_unicode_ci NOT NULL
);

ALTER TABLE ItemEvent ADD 
(
	`RemoteIP` varchar(45) COLLATE utf8_unicode_ci NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE UserEvent DROP RemoteIP;
ALTER TABLE ItemEvent DROP RemoteIP;