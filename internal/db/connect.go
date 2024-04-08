package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(dbConnURL string) {

    fmt.Println("Attempting database connection...")

	pool, err := pgxpool.New(context.Background(), dbConnURL)
	if err != nil {
		log.Fatal("Could not connect to Database", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("Connection to database successful!")
}
