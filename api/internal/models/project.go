package models

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"ledgerbolt.systems/internal/queries"
)

type Project struct {
	ID          int       `json:"ID"`
	Name        string    `json:"Name" validate:"required,min=3,max=50"`
	Description string    `json:"Description" validate:"required,min=3,max=1000"`
	ClientID    int       `json:"ClientID" validate:"required"`
	Notes       string    `json:"Notes"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

func GetProjects(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	userID string,
) ([]Project, error) {

	var projects []Project

	rows, err := conn.Query(ctx, queries.GetProjects, userID)
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

	row := conn.QueryRow(ctx, queries.GetProject, userID, projectID)

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
	clientID string,
) ([]Project, error) {
	var projects []Project

	rows, err := conn.Query(ctx, queries.GetProjectsByClient, userID, clientID)
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

	_, err := conn.Exec(
		ctx,
		queries.CreateProject,
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
	pID string,
	userID string,
) error {

	cmd, err := conn.Exec(
		ctx,
		queries.UpdateProject,
		p.Name,
		p.Description,
		p.ClientID,
		p.Notes,
		pID,
		userID,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("Project doesn't exist. Invalid ID")
	}

	return nil
}

func DestroyProject(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	pID string,
	userID string,
) error {

	cmd, err := conn.Exec(
		ctx,
		queries.DestroyProject,
		pID,
		userID,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("Project doesn't exist. Invalid ID")
	}
	return nil
}
