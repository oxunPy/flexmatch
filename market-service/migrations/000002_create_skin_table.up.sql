CREATE TABLE IF NOT EXISTS skins
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    attr JSONB,
    cost NUMERIC(19, 2) NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS player_skins
(
    player_id BIGINT NOT NULL,
    skin_id   UUID NOT NULL,
    CONSTRAINT pk_player_skin PRIMARY KEY (player_id, skin_id)
);