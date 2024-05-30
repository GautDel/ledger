package queries

const CreatePivotClientsProjects = `
INSERT INTO clients_projects (
    client_id,
    project_id,
    user_id
) VALUES ($1, $2, $3);`

const DestroyPivotClientsProjects = `
DELETE FROM clients_projects 
WHERE user_id = $1 AND id = $2;`


const UpsertPivotClientsProjects = `
INSERT INTO (
    client_id,
    project_id,
    user_id
) VALUES ($1, $2, $3)
ON CONFLICT (id) 
DO NOTHING
WHERE user_id = $4;`
