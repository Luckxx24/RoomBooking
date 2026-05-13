-- name: Createroom :one

insert into rooms(
    id,nama,kapasitas,price_per_hour,description,created_at,updated_at
) values(
    $1,$2,$3,$4,$5,$6,$7
)

RETURNING *;

-- name: GetRoom :many

select r.nama,r.kapasitas,r.price_per_hour,b.status as booking_status from rooms r inner join booking b on b.id_rooms = r.id OFFSET $1 LIMIT $2;

-- name: GetRoomDetail :one

select r.nama,r.kapasitas,r.price_per_hour,r.kapasitas,f.fasilitas as room_fasilitas ,b.status as booking_status from rooms r 
inner join booking b on b.id_rooms = r.id 
inner join room_fasilitas rf on rf id_room = r.id 
inner join fasilitas f on f.id = rf.id where r.id = $1;

-- name: UpdateRoom :one

update rooms set nama = $1, kapasitas = $2, price_per_hour = $3, description = $4,updated_at = $5 where id = $6
RETURNING *;

-- name: DeleteRoom :exec

Delete from rooms where id = $1;