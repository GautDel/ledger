package models

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Project struct {
	ID          int
	Name        string
	Description string
	ClientID    int
	Notes       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func GetProjects(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	userID string,
) ([]Project, error) {

	var projects []Project
	query := `
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

	rows, err := conn.Query(ctx, query, userID)
	if err != nil {
		log.Println(err)
		return projects, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Project
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.ClientID,
			&p.Notes,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return projects, err
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return projects, err
	}

	return projects, nil
}

func GetProject(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	userID string,
	projectID string,
) (Project, error) {

	var p Project
	query := `
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

	row := conn.QueryRow(ctx, query, userID, projectID)

	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.ClientID,
		&p.Notes,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return p, err
	}

	return p, nil
}

func GetProjectByClient(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	userID string,
	clientID int,
) ([]Project, error) {
    var projects []Project
	query := `
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

	rows, err := conn.Query(ctx, query, userID, clientID)
	if err != nil {
		log.Println(err)
		return projects, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Project
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.ClientID,
			&p.Notes,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return projects, err
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return projects, err
	}

	return projects, nil
}

func CreateProject(
	conn *pgxpool.Pool,
	p Project,
	ctx *gin.Context,
	userID string,
) error {

	query := `
    INSERT INTO projects(
        name,
        description,
        user_id,
        client_id,
        notes) VALUES ($1,$2,$3,$4,$5)`

	_, err := conn.Exec(
		ctx,
		query,
		p.Name,
		p.Description,
		userID,
		p.ClientID,
		p.Notes,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UpdateProject(
	conn *pgxpool.Pool,
	p Project,
	ctx *gin.Context,
	userID string,
) error {

	query := `
    UPDATE projects SET
    name = $1,
    description = $2,
    client_id = $3,
    notes = $4
    WHERE id = $5 AND user_id = $6`

	_, err := conn.Exec(
		ctx,
		query,
		p.Name,
		p.Description,
		p.ClientID,
		p.Notes,
        p.ID,
		userID,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DestroyProject(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	pID int,
	userID string,
) error {

	query := `
    DELETE FROM projects 
    WHERE id = $1 AND user_id = $2`

	_, err := conn.Exec(
		ctx,
		query,
        pID,
		userID,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
