package models

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"ledgerbolt.systems/internal/queries"
)

type User struct {
	FirstName   string `json:"FirstName" validate:"required,min=3,max=50"`
	LastName    string `json:"LastName" validate:"required,min=3,max=50"`
	CompanyName string `json:"CompanyName" validate:"min=3,max=100"`
	Email       string `json:"Email" validate:"required,min=5,max=320"`
	Phone       string `json:"Phone" validate:"required,min=5,max=50"`
	Address     string `json:"Address" validate:"required,min=5,max=1000"`
	CompanyNum     string `json:"CompanyNum" validate:"required,min=5,max=1000"`
	CreatedAt     time.Time `json:"CreatedAt"`
	UpdatedAt     time.Time `json:"UpdatedAt"`
}

func GetUser(conn *pgxpool.Pool, ctx *gin.Context, userID string) (User, error) {
	var user User

	row := conn.QueryRow(ctx, queries.GetUser, userID)
	err := row.Scan(
		&user.FirstName,
		&user.LastName,
		&user.CompanyName,
		&user.Email,
		&user.Phone,
		&user.Address,
		&user.CompanyNum,
	)
	if err != nil {
		log.Println(err)
		return user, err
	}

    log.Println(userID)
	return user, nil
}

func UpdateUser(conn *pgxpool.Pool, user User, ctx *gin.Context, userID string) error {

	cmd, err := conn.Exec(
		ctx,
		queries.UpdateUser,
		user.FirstName,
		user.LastName,
		user.CompanyName,
		user.Email,
		user.Phone,
		user.Address,
		user.CompanyNum,
		userID,
	)
	if err != nil {
		log.Println(err)
		return err
	}

    if cmd.RowsAffected() == 0 {
        return errors.New("User doesn't exist, invalid ID")
    }

	return nil
}
