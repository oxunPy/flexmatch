CREATE TABLE datafile
(
    id           UUID PRIMARY KEY     DEFAULT gen_random_uuid(),
    name         TEXT        NOT NULL,
    content_type TEXT        NOT NULL,
    size         BIGINT      NOT NULL,
    path         TEXT        NOT NULL UNIQUE,
    url          TEXT        NOT NULL,
    uploaded_by  TEXT        NOT NULL DEFAULT '',
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);

CREATE INDEX idx_files_uploaded_by ON datafile(uploaded_by);
CREATE INDEX idx_files_created_at  ON datafile(created_at DESC);