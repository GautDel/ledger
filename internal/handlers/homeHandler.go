package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"ledgerbolt.systems/internal/db"
)

type Test struct {
	Name        string
	Description string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	conn := db.GetPool()

	var tests []Test
    var test Test

	rows, err := conn.Query(context.Background(), "SELECT * FROM test")
	if err != nil {
		log.Fatal(err)
	}

	_, err = pgx.ForEachRow(rows, 
        []any{
            &test.Name, 
            &test.Description,
        }, func() error {

        tests = append(tests, test)

		return nil
	})

	if err != nil {
		fmt.Printf("ForEachRow error: %v", err)
		return
	}

	encoderErr := json.NewEncoder(w).Encode(tests)
	if encoderErr != nil {
		log.Println("Failed to Encode JSON", encoderErr)
	}
}
