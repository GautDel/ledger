package queries

const GetClients = `
SELECT 
    id,
    first_name,
    last_name,
    description,
    email,
    phone,
    address,
    country,
    created_at,
    updated_at
    FROM clients 
WHERE user_id = $1`

const GetClient = `
SELECT 
    id,
    first_name,
    last_name,
    description,
    email,
    phone,
    address,
    country
    FROM clients 
WHERE id = $1 AND user_id = $2`


const SearchClients = `
SELECT
    id,
    first_name,
    last_name,
    description,
    email,
    phone,
    address,
    country,
    created_at,
    updated_at
FROM clients 
WHERE first_name ILIKE '%' || $1 || '%'
OR last_name ILIKE '%' || $1 || '%'`

const NewClient = `
INSERT INTO clients(
    first_name,
    last_name,
    description,
    email,
    phone,
    address,
    country,
    user_id
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`


const UpdateClient = `
UPDATE clients SET 
    first_name = $1,
    last_name = $2,
    description = $3,
    email = $4,
    phone = $5,
    address = $6,
    country = $7
WHERE id = $8 AND user_id = $9`

const DestroyClient =  `
DELETE FROM clients 
WHERE id = $1 AND user_id = $2`
