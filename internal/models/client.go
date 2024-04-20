package models

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"ledgerbolt.systems/internal/queries"
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

	rows, err := conn.Query(ctx, queries.GetClients, userID)
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

	rows := conn.QueryRow(ctx, queries.GetClient, clientID, userID)

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

	rows, err := conn.Query(ctx, queries.SearchClients, searchStr)
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

	_, err := conn.Exec(
		ctx,
		queries.NewClient,
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

	_, err := conn.Exec(
		ctx,
		queries.UpdateClient,
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

	_, err := conn.Exec(ctx, queries.DestroyClient, clientID, userID)

	return err
}
