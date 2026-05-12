-- +goose up

create table fasilitas_ruangan(
    id uuid primary key not null,
    id_room uuid not null references rooms(id)  on Delete cascade,
    id_fasilitas uuid references fasilitas(id) not null
);

-- +goose down
Drop table if exists fasilitas_ruangan;