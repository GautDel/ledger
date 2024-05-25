package models

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
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
	Projects    []Project
	Starred     bool    `json:"Starred"`
}

type SearchClient struct {
	Search string `json:"Search" validate:"required,max=50"`
	Sort string `json:"Sort" validate:"required,max=15"`
}

type StarClient struct {
    Starred bool `json:"Starred"`
}

func GetClients(conn *pgxpool.Pool, ctx *gin.Context, userID string, sortBy string) ([]Client, error) {
	var clients []Client
    key := "GetClients"+sortBy
    query, exists := queries.QueryTemplates[key]
    if !exists {
        log.Println("Invalid query")
    }

	rows, err := conn.Query(ctx, query, userID)
	if err != nil {
        log.Println(err)
		return clients, err
	}

	defer rows.Close()

	for rows.Next() {
		var c Client
		var p Project
		err := rows.Scan(
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
			&p.ID,
			&p.Name,
			&p.Description,
			&p.ClientID,
			&p.Notes,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return clients, err
		}

		// Check if the current invoice is different from the last one
		if len(clients) == 0 || c.ID != clients[len(clients)-1].ID {
			clients = append(clients, c)
		}


        if p.ID != 0 {
		// Append the item to the last invoice
		clients[len(clients)-1].Projects = append(clients[len(clients)-1].Projects, p)
        }
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
		&client.Starred,
	)
	if err != nil {
		log.Println(err)
		return client, err
	}

	return client, nil
}

func SearchClients(conn *pgxpool.Pool, req SearchClient, ctx *gin.Context, userID string) ([]Client, error) {
	var clients []Client
    key := "SearchClients"+req.Sort
    query, exists := queries.QueryTemplates[key]
    if !exists {
        log.Println("Invalid query")
    }

	rows, err := conn.Query(ctx, query, req.Search, userID)
	if err != nil {
		return clients, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Client
		var p Project
		err := rows.Scan(
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
			&p.ID,
			&p.Name,
			&p.Description,
			&p.ClientID,
			&p.Notes,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return clients, err
		}

		// Check if the current invoice is different from the last one
		if len(clients) == 0 || c.ID != clients[len(clients)-1].ID {
			clients = append(clients, c)
		}

        if p.ID != 0 {
		// Append the item to the last invoice
		clients[len(clients)-1].Projects = append(clients[len(clients)-1].Projects, p)
        }
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
	    client.Starred,
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

func UpdateStarClient(conn *pgxpool.Pool, client StarClient, ctx *gin.Context, clientID string, userID string) error {

	cmd, err := conn.Exec(
		ctx,
		queries.UpdateStarClient,
	    client.Starred,
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
