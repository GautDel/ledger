package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v5"
	"ledgerbolt.systems/internal/db"
)

type Test struct {
	Name        string
	Description string
}

func homeHandler(c *gin.Context) {
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

    c.JSON(http.StatusOK, tests)
}
