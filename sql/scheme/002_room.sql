-- +goose up

create table rooms(
    id uuid primary key not null,
    nama varchar not null,
    kapasitas int not null,
    price_per_hour numeric(10,2)not null,
    description text not null
);

-- +goose down

Drop table if exists rooms;