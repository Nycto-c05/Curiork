CREATE TABLE IF NOT EXISTS posts (
    id BIGSERIAL PRIMARY KEY,
    title TEXT,
    user_id INT,
    content TEXT,
    tags TEXT[],
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
