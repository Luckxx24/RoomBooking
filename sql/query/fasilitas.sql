-- name: CreateFasilitas :one

insert into fasilitas(
    id,
    nama
) values(
    $1,$2
)
RETURNING *;

-- name: GetFasilitas :one

select * from fasilitas ;

-- name: Updatefasilitas :one

update fasilitas set nama = $1 where id = $2
RETURNING *;

-- name: DeleteFasilitas :exec

Delete from fasilitas where id = $1;

-- name: GetFasilitasByID :one

select id from fasilitas where id = 1$;
