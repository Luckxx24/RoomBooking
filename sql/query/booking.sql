

-- name: CreateBooking :one

insert into booking(
    id,id_user,id_rooms,start_time,end_time,total_price,status
)values(
    $1,$2,$3,$4,$5,$6,$7
)
RETURNING *;

-- name: GetBooking :many

select b.start_time,b.end_time,b.total_price,b.status,r.nama as nama_ruangan ,u.nama as nama_user from booking b inner join rooms r on b.id_rooms = r.id 
inner join users u on b.id_user = u.id OFFSET $1 LIMIT $2;

-- name: UpdateBooking :one

update booking set start_time = $1,end_time = $2,status = $3, total_price = $4 where id = $5
RETURNING *;

-- name: DeleteBooking :exec

Delete from booking where id = $1;
