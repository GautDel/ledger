package models

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	ID          string
	FirstName   string
	LastName    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ClientRequest struct {
	FirstName   string
	LastName    string
	Description string
}

func GetClients(conn *pgxpool.Pool, ctx *gin.Context) ([]Client, error) {
	var clients []Client
	var client Client
	query := "SELECT * from clients"

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return clients, err
	}

	_, err = pgx.ForEachRow(rows,
		[]any{
			&client.ID,
			&client.FirstName,
			&client.LastName,
			&client.Description,
			&client.CreatedAt,
			&client.UpdatedAt,
		}, func() error {
			clients = append(clients, client)

			return nil
		})

	if err != nil {
		return clients, err
	}

	return clients, nil
}

func GetClient(conn *pgxpool.Pool, clientID string, ctx *gin.Context) (Client, error) {
	var client Client
	query := "SELECT * FROM clients WHERE id = $1"

	rows := conn.QueryRow(ctx, query, clientID)

	err := rows.Scan(
		&client.ID,
		&client.FirstName,
		&client.LastName,
		&client.Description,
		&client.CreatedAt,
		&client.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return client, err
	}

	return client, nil
}

func SearchClients(conn *pgxpool.Pool, searchStr string, ctx *gin.Context) ([]Client, error) {
	var clients []Client
	var client Client
	query := `
        SELECT * FROM clients
        WHERE first_name ILIKE '%' || $1 || '%'
        OR last_name ILIKE '%' || $1 || '%'`

	rows, err := conn.Query(ctx, query, searchStr)
	if err != nil {
		return clients, err
	}
	_, err = pgx.ForEachRow(rows,
		[]any{
			&client.ID,
			&client.FirstName,
			&client.LastName,
			&client.Description,
			&client.CreatedAt,
			&client.UpdatedAt,
		}, func() error {
			clients = append(clients, client)

			return nil
		})

	if err != nil {
		return clients, err
	}

	return clients, nil
}

func NewClient(conn *pgxpool.Pool, client ClientRequest, ctx *gin.Context) error {
	query := `
        INSERT INTO clients(
            first_name,
            last_name,
            description
        ) VALUES ($1, $2, $3)`
	_, err := conn.Exec(ctx, query, client.FirstName, client.LastName, client.Description)

	return err
}

func UpdateClient(conn *pgxpool.Pool, client ClientRequest, ctx *gin.Context, clientID string) error {
	query := `
        UPDATE clients SET 
        first_name = $1,
        last_name = $2,
        description = $3
        WHERE id = $4`

	_, err := conn.Exec(
		ctx,
		query,
		client.FirstName,
		client.LastName,
		client.Description,
		clientID,
	)

	return err
}

func DestroyClient(conn *pgxpool.Pool, clientID string, ctx *gin.Context) error {
	query := `
        DELETE FROM clients 
        WHERE id = $1`

	_, err := conn.Exec(ctx, query, clientID)

	return err
}
