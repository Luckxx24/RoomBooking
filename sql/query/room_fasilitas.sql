

-- name: CreateFasilitas_Ruangan :one

insert into fasilitas_ruangan(
    id,id_room,id_fasilitas 
) values (
    $1,$2,$3
)
RETURNING *;



