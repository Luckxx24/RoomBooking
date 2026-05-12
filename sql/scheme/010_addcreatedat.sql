-- +goose up
alter table users add column updated_at timestamptz not null default now();

alter table rooms add column created_at timestamptz not null default now();
alter table rooms add column updated_at timestamptz not null default now();

alter table fasilitas_ruangan add column created_at timestamptz not null default now();
alter table fasilitas_ruangan add column updated_at timestamptz not null default now();

alter table fasilitas add column created_at timestamptz not null default now();
alter table fasilitas add column updated_at timestamptz not null default now();

alter table booking add column created_at timestamptz not null default now();
alter table booking add column updated_at timestamptz not null default now();

-- +goose down 
alter table users drop column updated_at;

alter table rooms drop  column created_at ;
alter table rooms drop  column updated_at ;

alter table fasilitas_ruangan drop  column created_at ;
alter table fasilitas_ruangan drop  column updated_at ;

alter table fasilitas drop column created_at ;
alter table fasilitas drop  column updated_at ;

alter table booking drop  column created_at ;
alter table booking drop column updated_at ;

