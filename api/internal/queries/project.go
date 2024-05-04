package queries

const GetProjects = `
SELECT 
    id,
    name,
    description,
    client_id,
    notes,
    created_at,
    updated_at
    FROM projects 
WHERE user_id = $1`

const GetProject = `
SELECT 
    id,
    name,
    description,
    client_id,
    notes,
    created_at,
    updated_at
    FROM projects 
WHERE user_id = $1 
AND id = $2`

const GetProjectsByClient = `
SELECT 
    id,
    name,
    description,
    client_id,
    notes,
    created_at,
    updated_at
FROM projects 
WHERE user_id = $1 
AND client_id = $2`

const CreateProject = `
INSERT INTO projects(
    name,
    description,
    user_id,
    client_id,
    notes
) VALUES ($1,$2,$3,$4,$5)`

const UpdateProject = `
UPDATE projects SET
    name = $1,
    description = $2,
    client_id = $3,
    notes = $4
WHERE id = $5 AND user_id = $6`

const DestroyProject = `
DELETE FROM projects 
WHERE id = $1 AND user_id = $2`
