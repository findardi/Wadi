-- name: InsertRole :one
insert into workspace_roles
    (workspace_id, name, permissions, is_system)
values
    ($1, $2, $3, $4)
returning *;