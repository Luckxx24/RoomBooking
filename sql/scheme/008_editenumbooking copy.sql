-- +goose up
alter type stats add value 'cancelled';

-- +goose down 
