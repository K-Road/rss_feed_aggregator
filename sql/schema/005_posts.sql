-- +goose Up
CREATE TABLE posts (
    id UUID primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    title text not null,
    url text unique not null,
    description text,
    published_at TIMESTAMP,
    feed_id UUID not null REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;