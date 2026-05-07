-- +goose up

create table fasilitas_ruangan(
    id uuid primary key not null,
    id_room uuid references rooms(id) not null,
    id_fasilitas uuid references fasilitas(id) not null
);

Drop table if exists fasilitas_ruangan;