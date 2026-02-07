-- name: CreateEntry :one 
Insert into entries (
    account_id,
    amount
) values (
    $1, $2
) RETURNING *;

-- name: GetEntry :one
select * from entries where id = $1 limit 1; 

-- name: ListEntries :many 
select * from entries order by created_at desc limit $1 offset $2;

-- name: UpdateEntry :one
update entries set amount = $2 where id = $1 RETURNING *;

-- name: DeleteEntry :exec
delete from entries where id = $1;