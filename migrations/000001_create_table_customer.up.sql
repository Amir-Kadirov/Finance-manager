CREATE TABLE IF NOT EXISTS customer (
    id uuid PRIMARY KEY,
    phone varchar,
    password varchar NOT NUll,
    first_name varchar,
    last_name varchar,
    gmail varchar,
    history varchar,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE IF NOT EXISTS payment_info (
    customer_id uuid REFERENCES customer (id),
    card_id varchar REFERENCES "card" (id),
    history varchar,
    created_at timestamp DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS expense (
    total_expense int,
    customer_id uuid REFERENCES customer (id),
    restaurant int DEFAULT 0,
    supermarkets int DEFAULT 0,
    beauty_medecine int DEFAULT 0,
    entertaintment_sport int DEFAULT 0
);