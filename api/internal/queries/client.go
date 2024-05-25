package queries

var QueryTemplates = map[string]string{
	"GetClientsASC": `
SELECT 
    clients.id,
    clients.first_name,
    clients.last_name,
    clients.description,
    clients.email,
    clients.phone,
    clients.address,
    clients.country,
    clients.created_at,
    clients.updated_at,
    clients.starred,
    COALESCE(projects.id, 0),
    COALESCE(projects.name, ''),
    COALESCE(projects.description, ''),
    COALESCE(projects.client_id, 0),
    COALESCE(projects.notes, ''),
    COALESCE(projects.created_at, '1970-01-01 00:00:00'::timestamp),
    COALESCE(projects.updated_at, '1970-01-01 00:00:00'::timestamp)
FROM clients 
LEFT JOIN projects ON clients.id = projects.client_id
WHERE clients.user_id = $1
ORDER BY clients.starred DESC, clients.first_name ASC, clients.last_name ASC`,

	"GetClientsDESC": `
SELECT 
    clients.id,
    clients.first_name,
    clients.last_name,
    clients.description,
    clients.email,
    clients.phone,
    clients.address,
    clients.country,
    clients.created_at,
    clients.updated_at,
    clients.starred,
    COALESCE(projects.id, 0),
    COALESCE(projects.name, ''),
    COALESCE(projects.description, ''),
    COALESCE(projects.client_id, 0),
    COALESCE(projects.notes, ''),
    COALESCE(projects.created_at, '1970-01-01 00:00:00'::timestamp),
    COALESCE(projects.updated_at, '1970-01-01 00:00:00'::timestamp)
FROM clients 
LEFT JOIN projects ON clients.id = projects.client_id
WHERE clients.user_id = $1
ORDER BY clients.starred DESC, clients.first_name DESC, clients.last_name DESC`,
	"GetClientsNEW": `
SELECT 
    clients.id,
    clients.first_name,
    clients.last_name,
    clients.description,
    clients.email,
    clients.phone,
    clients.address,
    clients.country,
    clients.created_at,
    clients.updated_at,
    clients.starred,
    COALESCE(projects.id, 0),
    COALESCE(projects.name, ''),
    COALESCE(projects.description, ''),
    COALESCE(projects.client_id, 0),
    COALESCE(projects.notes, ''),
    COALESCE(projects.created_at, '1970-01-01 00:00:00'::timestamp),
    COALESCE(projects.updated_at, '1970-01-01 00:00:00'::timestamp)
FROM clients 
LEFT JOIN projects ON clients.id = projects.client_id
WHERE clients.user_id = $1
ORDER BY clients.starred DESC, clients.created_at DESC;`,

	"GetClientsOLD": `
SELECT 
    clients.id,
    clients.first_name,
    clients.last_name,
    clients.description,
    clients.email,
    clients.phone,
    clients.address,
    clients.country,
    clients.created_at,
    clients.updated_at,
    clients.starred,
    COALESCE(projects.id, 0),
    COALESCE(projects.name, ''),
    COALESCE(projects.description, ''),
    COALESCE(projects.client_id, 0),
    COALESCE(projects.notes, ''),
    COALESCE(projects.created_at, '1970-01-01 00:00:00'::timestamp),
    COALESCE(projects.updated_at, '1970-01-01 00:00:00'::timestamp)
FROM clients 
LEFT JOIN projects ON clients.id = projects.client_id
WHERE clients.user_id = $1
ORDER BY clients.starred DESC, clients.created_at ASC;`,

	"SearchClientsASC": `
SELECT
    clients.id,
    clients.first_name,
    clients.last_name,
    clients.description,
    clients.email,
    clients.phone,
    clients.address,
    clients.country,
    clients.created_at,
    clients.updated_at,
    clients.starred,
    COALESCE(projects.id, 0),
    COALESCE(projects.name, ''),
    COALESCE(projects.description, ''),
    COALESCE(projects.client_id, 0),
    COALESCE(projects.notes, ''),
    COALESCE(projects.created_at, '1970-01-01 00:00:00'::timestamp),
    COALESCE(projects.updated_at, '1970-01-01 00:00:00'::timestamp)
FROM clients 
LEFT JOIN projects ON clients.id = projects.client_id
WHERE clients.first_name ILIKE '%' || $1 || '%'
OR clients.last_name ILIKE '%' || $1 || '%'
AND clients.user_id = $2
ORDER BY clients.starred DESC, clients.first_name ASC, clients.last_name ASC`,

	"SearchClientsDESC": `
SELECT
    clients.id,
    clients.first_name,
    clients.last_name,
    clients.description,
    clients.email,
    clients.phone,
    clients.address,
    clients.country,
    clients.created_at,
    clients.updated_at,
    clients.starred,
    COALESCE(projects.id, 0),
    COALESCE(projects.name, ''),
    COALESCE(projects.description, ''),
    COALESCE(projects.client_id, 0),
    COALESCE(projects.notes, ''),
    COALESCE(projects.created_at, '1970-01-01 00:00:00'::timestamp),
    COALESCE(projects.updated_at, '1970-01-01 00:00:00'::timestamp)
FROM clients 
LEFT JOIN projects ON clients.id = projects.client_id
WHERE clients.first_name ILIKE '%' || $1 || '%'
OR clients.last_name ILIKE '%' || $1 || '%'
AND clients.user_id = $2
ORDER BY clients.starred DESC, clients.first_name DESC, clients.last_name DESC`,

	"SearchClientsNEW": `
SELECT
    clients.id,
    clients.first_name,
    clients.last_name,
    clients.description,
    clients.email,
    clients.phone,
    clients.address,
    clients.country,
    clients.created_at,
    clients.updated_at,
    clients.starred,
    COALESCE(projects.id, 0),
    COALESCE(projects.name, ''),
    COALESCE(projects.description, ''),
    COALESCE(projects.client_id, 0),
    COALESCE(projects.notes, ''),
    COALESCE(projects.created_at, '1970-01-01 00:00:00'::timestamp),
    COALESCE(projects.updated_at, '1970-01-01 00:00:00'::timestamp)
FROM clients 
LEFT JOIN projects ON clients.id = projects.client_id
WHERE clients.first_name ILIKE '%' || $1 || '%'
OR clients.last_name ILIKE '%' || $1 || '%'
AND clients.user_id = $2
ORDER BY clients.starred DESC, clients.created_at DESC;`,

	"SearchClientsOLD": `
SELECT
    clients.id,
    clients.first_name,
    clients.last_name,
    clients.description,
    clients.email,
    clients.phone,
    clients.address,
    clients.country,
    clients.created_at,
    clients.updated_at,
    clients.starred,
    COALESCE(projects.id, 0),
    COALESCE(projects.name, ''),
    COALESCE(projects.description, ''),
    COALESCE(projects.client_id, 0),
    COALESCE(projects.notes, ''),
    COALESCE(projects.created_at, '1970-01-01 00:00:00'::timestamp),
    COALESCE(projects.updated_at, '1970-01-01 00:00:00'::timestamp)
FROM clients 
LEFT JOIN projects ON clients.id = projects.client_id
WHERE clients.first_name ILIKE '%' || $1 || '%'
OR clients.last_name ILIKE '%' || $1 || '%'
AND clients.user_id = $2
ORDER BY clients.starred DESC, clients.created_at ASC;`,
}

const GetClient = `
SELECT 
    id,
    first_name,
    last_name,
    description,
    email,
    phone,
    address,
    country,
    starred
    FROM clients 
WHERE id = $1 AND user_id = $2`

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
    country = $7,
    starred = $8
WHERE id = $9 AND user_id = $10`

const UpdateStarClient = `
UPDATE clients SET 
    starred = $1
WHERE id = $2 AND user_id = $3`

const DestroyClient = `
DELETE FROM clients 
WHERE id = $1 AND user_id = $2`
