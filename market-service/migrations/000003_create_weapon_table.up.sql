CREATE TABLE IF NOT EXISTS weapons
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(255)   NOT NULL,
    description TEXT,
    weapon_type INT           NOT NULL DEFAULT 0,
    attr        JSONB,
    cost        NUMERIC(19, 2) NOT NULL,
    created     TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated     TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS player_weapons
(
    player_id BIGINT NOT NULL,
    weapon_id UUID   NOT NULL,
    CONSTRAINT pk_player_weapon PRIMARY KEY (player_id, weapon_id)
);
