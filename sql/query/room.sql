-- name: Createroom :one

insert into rooms(
    id,nama,kapasitas,price_per_hour,description 
) values(
    $1,$2,$3,$4,$5
)

RETURNING *;

-- name: GetRoom :many

select r.nama,r.kapasitas,r.price_per_hour,b.status as booking_status from rooms r inner join booking b on b.id_rooms = r.id OFFSET $1 LIMIT $2;

-- name: GetRoomDetail :one

select r.nama,r.kapasitas,r.price_per_hour,r.kapasitas,b.status as booking_status from rooms r inner join booking b on b.id_rooms = r.id where r.id = $1 OFFSET $2 LIMIT $3;

-- name: UpdateRoom :one

update rooms set nama = $1, kapasitas = $2, price_per_hour = $3, description = $4 where id = $5
RETURNING *;

-- name: DeleteRoom :exec

Delete from rooms where id = $1;