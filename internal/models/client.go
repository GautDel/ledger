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
	Email       string
	Phone       string
	Address     string
	Country     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ClientRequest struct {
	FirstName   string
	LastName    string
	Description string
	Email       string
	Phone       string
	Address     string
	Country     string
}

func GetClients(conn *pgxpool.Pool, ctx *gin.Context, userID string) ([]Client, error) {
	var clients []Client
	var client Client
	query := `SELECT 
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

	rows, err := conn.Query(ctx, query, userID)
	if err != nil {
		return clients, err
	}
	defer rows.Close()

	_, err = pgx.ForEachRow(rows,
		[]any{
			&client.ID,
			&client.FirstName,
			&client.LastName,
			&client.Description,
			&client.Email,
			&client.Phone,
			&client.Address,
			&client.Country,
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

func GetClient(conn *pgxpool.Pool, clientID string, ctx *gin.Context, userID string) (Client, error) {
	var client Client
	query := `SELECT 
        id,
        first_name,
        last_name,
        description,
        email,
        phone,
        address,
        country
        FROM clients 
        WHERE id = $1
        AND user_id = $2`
	rows := conn.QueryRow(ctx, query, clientID, userID)

	err := rows.Scan(
		&client.ID,
		&client.FirstName,
		&client.LastName,
		&client.Description,
		&client.Email,
		&client.Phone,
		&client.Address,
		&client.Country,
	)
	if err != nil {
		log.Println(err)
		return client, err
	}

	return client, nil
}

func SearchClients(conn *pgxpool.Pool, searchStr string, ctx *gin.Context, userID string) ([]Client, error) {
	var clients []Client
	var client Client
	query := `
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

	rows, err := conn.Query(ctx, query, searchStr)
	if err != nil {
		return clients, err
	}
	defer rows.Close()

	_, err = pgx.ForEachRow(rows,
		[]any{
			&client.ID,
			&client.FirstName,
			&client.LastName,
			&client.Description,
			&client.Email,
			&client.Phone,
			&client.Address,
			&client.Country,
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

func NewClient(conn *pgxpool.Pool, client ClientRequest, ctx *gin.Context, userID string) error {
	query := `
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
	_, err := conn.Exec(
		ctx,
		query,
		client.FirstName,
		client.LastName,
		client.Description,
		client.Email,
		client.Phone,
		client.Address,
		client.Country,
		userID,
	)

	return err
}

func UpdateClient(conn *pgxpool.Pool, client ClientRequest, ctx *gin.Context, clientID string, userID string) error {
	query := `
        UPDATE clients SET 
        first_name = $1,
        last_name = $2,
        description = $3
        email = $4
        phone = $5
        address = $6
        country = $7
        WHERE id = $8 AND user_id = $9`

	_, err := conn.Exec(
		ctx,
		query,
		client.FirstName,
		client.LastName,
		client.Description,
		client.Email,
		client.Phone,
		client.Address,
		client.Country,
		clientID,
		userID,
	)

	return err
}

func DestroyClient(conn *pgxpool.Pool, clientID string, ctx *gin.Context, userID string) error {
	query := `
        DELETE FROM clients 
        WHERE id = $1 AND user_id = $2`

	_, err := conn.Exec(ctx, query, clientID, userID)

	return err
}
