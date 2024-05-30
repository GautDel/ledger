package queries

var ProjectTemplates = map[string]string{
    "GetProjectsNEW": `SELECT
    p.id,
    p.name,
    p.description,
    p.notes,
    p.created_at,
    p.updated_at,
    c.id,
    c.first_name,
    c.last_name,
    c.description,
    c.email,
    c.phone,
    c.address,
    c.country,
    c.created_at,
    c.updated_at,
    c.starred
FROM projects p
JOIN clients_projects cp ON p.id = cp.project_id
JOIN clients c on cp.client_id = c.id
WHERE p.user_id = $1`,
}

const GetProjects = `SELECT 
    id,
    name,
    description,
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
    notes
) VALUES ($1,$2,$3,$4)
RETURNING id`

const UpdateProject = `
UPDATE projects SET
    name = $1,
    description = $2,
    notes = $3
WHERE id = $4 AND user_id = $5`

const DestroyProject = `
DELETE FROM projects 
WHERE id = $1 AND user_id = $2`
