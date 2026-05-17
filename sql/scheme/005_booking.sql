-- +goose up

create type stats as enum('pending','approved','done');

create table booking(
    id uuid primary key not null,
    id_user uuid not null references users(id)  on Delete cascade,
    id_rooms uuid not null references rooms(id)  on Delete cascade,
    start_time timestamptz not null,
    end_time timestamptz not null,
    total_price numeric(10,2) not null,
    status stats not null Default 'pending'
);

-- +goose down

drop table if exists booking;
drop type if exists stats;