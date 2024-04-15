package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentStatus struct {
	ID     int
	Status string
	Color  string
}

func GetPaymentStatus(conn *pgxpool.Pool, ctx *gin.Context) ([]PaymentStatus, error) {
	var psSlice []PaymentStatus
	query := `
    SELECT * FROM payment_status`

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return psSlice, err
	}

	defer rows.Close()

	for rows.Next() {
		var ps PaymentStatus
		err := rows.Scan(
			&ps.ID,
			&ps.Status,
			&ps.Color,
		)
		if err != nil {
			return psSlice, err
		}
		psSlice = append(psSlice, ps)
	}

	if err := rows.Err(); err != nil {
		return psSlice, err
	}

	return psSlice, nil
}

func GetSinglePaymentStatus(conn *pgxpool.Pool, ctx *gin.Context, psID string) (PaymentStatus, error) {
	var ps PaymentStatus
	query := `
    SELECT * FROM payment_status WHERE id = $1`

	row := conn.QueryRow(ctx, query, psID)

	err := row.Scan(
		&ps.ID,
		&ps.Status,
		&ps.Color,
	)
	if err != nil {
		return ps, err
	}

	return ps, nil
}

func CreatePaymentStatus(conn *pgxpool.Pool, ps PaymentStatus, ctx *gin.Context) error {
	query := `
    INSERT INTO payment_status (
        status,
        color
    ) VALUES ($1, $2)`

	_, err := conn.Exec(ctx, query, ps.Status, ps.Color)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePaymentStatus(conn *pgxpool.Pool, ps PaymentStatus, ctx *gin.Context, psID string) error {
	query := `
    UPDATE payment_status SET
    status = $1,
    color = $2
    WHERE id = $3`

	_, err := conn.Exec(
		ctx,
		query,
		ps.Status,
		ps.Color,
		psID,
	)
	if err != nil {
		return err
	}

	return nil
}

func DestroyPaymentStatus(conn *pgxpool.Pool, ctx *gin.Context, psID string) error {
	query := `
    DELETE FROM payment_status
    WHERE id = $1`

	_, err := conn.Exec(
		ctx,
		query,
		psID,
	)
	if err != nil {
		return err
	}

	return nil
}
