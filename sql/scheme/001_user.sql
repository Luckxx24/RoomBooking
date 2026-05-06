-- +goose up

create table users(
    id uuid primary key not null,
    nama varchar not null,
    email varchar not null,
    hash_password varchar not null,
);

-- +goose down

Drop table if exists users;