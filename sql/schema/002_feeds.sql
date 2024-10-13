-- +goose Up
CREATE TABLE feeds (
    id UUID primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    name text not null,
    url text unique not null,
    user_id UUID not null,
    constraint fk_user
    foreign key (user_id)
    references users(id)
    on DELETE cascade
);

-- +goose Down
DROP TABLE feeds;

--goose postgres postgres://postgres:postgres@localhost:5432/gator up
--goose postgres postgres://postgres:postgres@localhost:5432/gator down