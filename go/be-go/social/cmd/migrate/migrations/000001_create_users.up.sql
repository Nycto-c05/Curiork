-- CREATE EXTENSION IF NOT EXISTS citext; -- keep this in db_init, not in app level migrations

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    password bytea NOT NULL,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);