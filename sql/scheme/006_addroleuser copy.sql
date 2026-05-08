-- +goose up
create type roles as enum('Admin','user');

alter table users add column role roles not null;

-- +goose down 
alter table users drop column role;
drop type if exists roles;