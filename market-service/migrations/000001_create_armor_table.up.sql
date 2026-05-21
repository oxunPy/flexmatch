CREATE TABLE IF NOT EXISTS armors
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    cost        NUMERIC(19, 2),
    attr        JSONB,
    created     TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated     TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS player_armors
(
    player_id BIGINT NOT NULL,
    armor_id  UUID   NOT NULL,
    CONSTRAINT pk_player_armor PRIMARY KEY (player_id, armor_id)
);
