

-- name: CreateFasilitas_Ruangan :one

insert into fasilitas_ruangan(
    id,id_room,id_fasilitas 
) values (
    $1,$2,$3
)
RETURNING *;

-- name: GetFasilitas_Ruangan :one

Select * from fasilitas_ruangan;

-- name: DeleteFasilitas_Ruangan :exec

Delete from fasilitas_ruangan where id_room = $1;





