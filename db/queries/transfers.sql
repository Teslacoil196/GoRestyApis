-- name: CreateTransfer :one 
Insert into transfers (
    from_account_id,
    to_account_id,
    amount
) values (
    $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
select * from transfers where id = $1 limit 1; 

-- name: ListTransfers :many 
select * from transfers order by created_at desc limit $1 offset $2;

-- name: ListTransfersFromAccountToAccount :many
select * from transfers where from_account_id = $1 and to_account_id = $2 order by created_at desc limit $3 offset $4;

-- name: ListTransfersFromAccount :many
select * from transfers where from_account_id = $1 order by created_at desc limit $2 offset $3;

-- name: ListTransfersToAccount :many
select * from transfers where to_account_id = $1 order by created_at desc limit $2 offset $3;

-- name: UpdateTransfer :one
update transfers set amount = $2 where id = $1 RETURNING *;

-- name: DeleteTransfer :exec
delete from transfers where id = $1;