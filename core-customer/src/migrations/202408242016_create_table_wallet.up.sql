CREATE TABLE wallets (
    id TEXT PRIMARY KEY,
    customer_id TEXT NOT NULL,
    balance TEXT NOT NULL DEFAULT '0',
    balance_invested TEXT NOT NULL DEFAULT '0',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers (id)
);