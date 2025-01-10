-- name: GetCompanies :many
select *
from closs_company
;

-- name: GetCompanyByCode :one
select *
from closs_company
where ked_codigo = ?
;

-- name: GetExistingCompanyByCode :one
select *
from closs_company
where ked_codigo = ? and (ked_status = 1 or deleted_at is null)
;

-- name: CreateCompany :one
INSERT INTO closs_company (
    ked_codigo,
    ked_nombre,
    ked_status,
    ked_enlace,
    ked_agen,
    created_at,
    updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateCompany :one
UPDATE closs_company SET 
    ked_nombre = ?,
    ked_status = ?, 
    ked_enlace = ?,
    ked_agen = ?,
    updated_at = ?
WHERE ked_codigo = ?
RETURNING *;

-- name: SoftDeleteCompany :exec
UPDATE closs_company SET 
    ked_status = 0,
    updated_at = ?,
    deleted_at = ?
WHERE ked_codigo = ?;

-- name: DeleteCompany :exec
delete from closs_company
where ked_codigo = ?
;

