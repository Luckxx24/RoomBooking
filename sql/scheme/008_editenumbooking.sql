-- +goose up
alter type stats add value 'cancelled';

-- +goose down 

alter type stats drop value 'cancelled';