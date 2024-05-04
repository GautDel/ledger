package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func Connect(dbConnURL string) {

    fmt.Println("Attempting database connection...")

    poolConfig, err := pgxpool.ParseConfig(dbConnURL)
    if err != nil {
        log.Fatal(err)
        return
    }

    pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatal("Could not connect to Database", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

    dbPool = pool

    fmt.Println("Connection to database successful!")
}

func GetPool() *pgxpool.Pool {
    return dbPool
}
