-- +goose up

create table rooms(
    id uuid primary key not null,
    nama varcha not null,

)