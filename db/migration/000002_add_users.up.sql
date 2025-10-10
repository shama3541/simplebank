CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_changed_at timestamptz NOT NULL DEFAULT now(),
    created_at timestamptz NOT NULL DEFAULT now()
);

ALTER TABLE accounts
ADD CONSTRAINT fk_accounts_owner FOREIGN KEY (owner) REFERENCES users (username) ON DELETE CASCADE;

CREATE UNIQUE INDEX accounts_owner_currency_idx ON accounts (owner, currency);