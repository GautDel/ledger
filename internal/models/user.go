package models

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	FirstName   string
	LastName    string
	CompanyName string
	Email       string
	Phone       string
	Address     string
	CompanyNum  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserRequest struct {
	FirstName   string
	LastName    string
	CompanyName string
	Email       string
	Phone       string
	Address     string
	CompanyNum  string
}

func GetUser(conn *pgxpool.Pool, ctx *gin.Context, userID string) (User, error) {
	var user User
	query := `SELECT 
        first_name,
        last_name,
        company_name,
        email,
        phone,
        address,
        company_num
        FROM users 
        WHERE id = $1`

	row := conn.QueryRow(ctx, query, userID)
	err := row.Scan(
		&user.FirstName,
		&user.LastName,
		&user.CompanyName,
		&user.Email,
		&user.Phone,
		&user.Address,
	)
	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func UpdateUser(conn *pgxpool.Pool, user UserRequest, ctx *gin.Context, userID string) error {
	query := `
        UPDATE users SET
        first_name = $1,
        last_name = $2,
        company_name = $3,
        email = $4,
        phone = $5,
        address = $6,
        company_num = $7
        WHERE id = $7`

	_, err := conn.Exec(
		ctx,
		query,
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

	return nil
}
