CREATE TABLE IF NOT EXISTS match(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    type INT NOT NULL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS match_player(
    match_id UUID NOT NULL,
    player_id BIGINT,
    CONSTRAINT pk_match_player PRIMARY KEY (match_id, player_id)
);
