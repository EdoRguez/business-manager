-- name: CreateCompany :one
INSERT INTO 
  company.company (
    name,
    image_url,
    last_payment_date
  ) 
VALUES (
  $1, $2, $3
) 
RETURNING *;

-- name: GetCompany :one
SELECT 
  id,
  name,
  image_url,
  plan_id,
  last_payment_date,
  created_at,
  modified_at
FROM 
  company.company
WHERE 
  id = $1 
LIMIT 1;

-- name: GetCompanies :many
SELECT 
  id,
  name,
  image_url,
  plan_id,
  last_payment_date,
  created_at,
  modified_at
FROM 
  company.company
ORDER BY 
  id
LIMIT 
  $1
OFFSET 
  $2;

-- name: UpdateCompany :one
UPDATE 
  company.company
SET 
  name = $2,
  image_url = $3,
  modified_at = NOW()
WHERE 
  id = $1
RETURNING *;

-- name: DeleteCompany :exec
DELETE FROM 
  company.company
WHERE 
  id = $1;
