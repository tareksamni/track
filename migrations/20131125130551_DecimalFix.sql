-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE PurchaseEvent MODIFY COLUMN GrossAmount DECIMAL(19,4) NOT NULL;
ALTER TABLE PurchaseEvent MODIFY COLUMN NetAmount DECIMAL(19,4) NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE PurchaseEvent MODIFY COLUMN GrossAmount DECIMAL(10,2) NOT NULL;
ALTER TABLE PurchaseEvent MODIFY COLUMN NetAmount DECIMAL(10,2) NOT NULL;
