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
	Notes       string    `json:"Notes"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
	Clients     []Client  `json:"Clients"`
}

func GetProjects(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	userID string,
	sortBy string,
) ([]Project, error) {

	var projects []Project

	key := "GetProjects" + sortBy
	query, exists := queries.ProjectTemplates[key]
	if !exists {
		log.Println("Invalid query")
	}

	rows, err := conn.Query(ctx, query, userID)
	if err != nil {
		log.Println(err)
		return projects, err
	}

	defer rows.Close()

	for rows.Next() {
		var p Project
        var c Client
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Notes,
			&p.CreatedAt,
			&p.UpdatedAt,
			&c.ID,
			&c.FirstName,
			&c.LastName,
			&c.Description,
			&c.Email,
			&c.Phone,
			&c.Address,
			&c.Country,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.Starred,
		)
		if err != nil {
			return projects, err
		}

		if len(projects) == 0 || p.ID != projects[len(projects)-1].ID {
			projects = append(projects, p)
		}

		// Append the item to the last invoice
		projects[len(projects)-1].Clients = append(projects[len(projects)-1].Clients, c)

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

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	var projectID string

	err = tx.QueryRow(
		ctx,
		queries.CreateProject,
		p.Name,
		p.Description,
		userID,
		p.Notes,
	).Scan(&projectID)

	if err != nil {
        log.Println(err)
		return err
	}

	for _, client := range p.Clients {
		_, err = tx.Exec(
			ctx,
			queries.CreatePivotClientsProjects,
			client.ID,
			projectID,
			userID,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
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
