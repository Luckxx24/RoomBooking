-- +goose up
alter table fasilitas rename column fasilitas to nama;

-- +goose down 
alter table fasilitas rename column nama to fasilitas;

