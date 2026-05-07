-- +goose up
alter table users add column created_at timestamptz not null;

-- +goose down 

alter table users drop column created_at;