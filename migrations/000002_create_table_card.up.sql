CREATE TABLE IF NOT EXISTS card (
    customer_id uuid NOT NUll,
    id varchar PRIMARY KEY,
    card_holder_name varchar,
    cvv varchar,
    expity_date date,
    password varchar,
    balance integer default 0
);

ALTER TABLE customer ADD COLUMN deleted_at timestamp;
ALTER TABLE card ADD FOREIGN KEY (customer_id) REFERENCES customer (id);


CREATE OR REPLACE FUNCTION insert_expense_from_customer()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO expense (customer_id)
    SELECT NEW.id
    FROM customer
    WHERE id = NEW.id;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_InsertOrderFromcustomer
AFTER INSERT ON customer
FOR EACH ROW
EXECUTE FUNCTION insert_expense_from_customer();