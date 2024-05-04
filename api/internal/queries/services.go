package queries

const GetServices = `
SELECT 
    id,
    name,
    description,
    unit_price,
    hourly_price,
    tax,
    created_at,
    updated_at
FROM services WHERE user_id = $1`

const GetSingleService = `
SELECT 
    id,
    name,
    description,
    unit_price,
    hourly_price,
    tax,
    created_at,
    updated_at
FROM services WHERE user_id = $1 AND id = $2`

const CreateService = `
INSERT INTO services(
    name,
    description,
    unit_price,
    hourly_price,
    tax,
    user_id
) VALUES($1,$2,$3,$4,$5,$6)`

const UpdateService = `
UPDATE services SET
    name = $1,
    description = $2,
    unit_price = $3,
    hourly_price = $4,
    tax = $5
WHERE user_id = $6 AND id = $7`

const DestroyService = `
DELETE FROM services
WHERE user_id = $1 AND id = $2`

