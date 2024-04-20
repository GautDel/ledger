package models

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"ledgerbolt.systems/internal/queries"
)

type Client struct {
	ID          int       `json:"ID"`
	FirstName   string    `json:"FirstName" validate:"required,min=3,max=50"`
	LastName    string    `json:"LastName" validate:"required,min=3,max=50"`
	Description string    `json:"Description" validate:"max=2000"`
	Email       string    `json:"Email" validate:"required,min=5,max=320"`
	Phone       string    `json:"Phone" validate:"required,min=5,max=50"`
	Address     string    `json:"Address" validate:"required,min=5,max=1000"`
	Country     string    `json:"Country" validate:"required,max=255"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

type SearchClient struct {
	Search string `json:"Search" validate:"required,max=50"`
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

func NewClient(conn *pgxpool.Pool, client Client, ctx *gin.Context, userID string) error {

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
	if err != nil {
		return err
	}

	return err
}

func UpdateClient(conn *pgxpool.Pool, client Client, ctx *gin.Context, clientID string, userID string) error {

	cmd, err := conn.Exec(
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
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("Client doesn't exist. Invalid ID")
	}

	return nil
}

func DestroyClient(conn *pgxpool.Pool, clientID string, ctx *gin.Context, userID string) error {

	cmd, err := conn.Exec(ctx, queries.DestroyClient, clientID, userID)
	if cmd.RowsAffected() == 0 {
		return errors.New("Client doesn't exist. Invalid ID")
	}

	return err
}
