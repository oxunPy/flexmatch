CREATE TABLE payments(
    id SERIAL PRIMARY KEY,
    item_id UUID NOT NULL,
    player_id UUID NOT NULL,
    wallet_id BIGINT NOT NULL,
    type INT NOT NULL,
    amount numeric(19, 2) NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)