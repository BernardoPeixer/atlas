CREATE TABLE IF NOT EXISTS cryptos
(
    id          INTEGER                            NOT NULL primary key,
    symbol      VARCHAR(10)                        NOT NULL,
    name        VARCHAR(100)                       NOT NULL,
    status_code TINYINT  DEFAULT 0                 NOT NULL,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    modified_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user
(
    id             INTEGER AUTO_INCREMENT PRIMARY KEY,
    username       VARCHAR(100) NULL,
    wallet_address TEXT         NOT NULL
);

CREATE TABLE IF NOT EXISTS festival_cards
(
    id             INTEGER AUTO_INCREMENT PRIMARY KEY,
    id_crypto_type INTEGER                            NOT NULL,
    id_user        INTEGER                            NOT NULL,
    balance        DECIMAL(10, 2)                     NOT NULL,
    crypto_price   DECIMAL(18, 10)                    NOT NULL,
    sold_at        DATETIME                           NULL,
    status_code    TINYINT  DEFAULT 0                 NOT NULL,
    created_at     DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    modified_at_   DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_crypto_type) REFERENCES cryptos (id),
    FOREIGN KEY (id_user) REFERENCES user (id)
);

