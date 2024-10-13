-- +goose Up
CREATE TABLE feed_follows (
    id UUID primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    user_id UUID not null,
    feed_id UUID not null,
    constraint fk_user
    foreign key (user_id)
    references users(id)
    on DELETE cascade,
    constraint fk_feed
    foreign key (feed_id)
    references feeds(id)
    on DELETE cascade,
    constraint unique_user_feed
    unique (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;