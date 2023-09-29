CREATE TABLE IF NOT EXISTS wallets (
    id varchar(40) NOT NULL PRIMARY KEY,
    user_id varchar(40) NOT NULL,
    currency_id varchar(40) NOT NULL,
    wallet_name varchar(30) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE default NOW(),
	created_by varchar(50),
    updated_at TIMESTAMP WITH TIME ZONE default NOW(), 
	updated_by varchar(50),
    is_deleted boolean NOT NULL DEFAULT false,
    CONSTRAINT fk_wallet_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_wallet_currency FOREIGN KEY (currency_id) REFERENCES currencies(id) ON DELETE RESTRICT ON UPDATE CASCADE
);
