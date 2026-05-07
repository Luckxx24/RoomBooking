-- +goose up

create type stats as enum('pending','approved','done');

create table booking(
    id uuid primary key not null,
    id_user uuid references users(id) not null,
    id_rooms uuid references rooms(id) not null,
    start_time timestamptz not null,
    end_time timestamptz not null,
    total_price numeric(10,2) not null,
    status stats not null
);

-- +goose down

drop type if exists stats;
drop table if exists booking;