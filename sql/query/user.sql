

-- name: CreateUsers :one

insert into users(
    id,nama,email,hash_password,role,created_at
) values(
    $1,$2,$3,$4,$5,$6
)

RETURNING *;

-- name: GetUsers :one

select nama,email,hash_password from users where id = $1;

-- name: GetUserslist :many

select nama,email,hash_password from users where id = $1 order by created_at desc LIMIT $2 OFFSET $3;

-- name: UpdateUser :one

update users set nama = $1, email = $2 , hash_password = $3 where id = $4
RETURNING *;

-- name: DeleteUser :exec

Delete from users where id = $1;