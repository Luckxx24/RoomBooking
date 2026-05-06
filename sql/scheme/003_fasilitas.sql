-- +goose up 

create table fasilitas(
    id uuid primary key not null,
    fasilitas varchar not null,
);

-- +goose down

Drop table if exists fasilitas