-- +goose up
create type roles as enum('Admin','user');

alter table users add column role roles not null;

-- +goose down 
drop type if exists roles;
alter table users drop column role;