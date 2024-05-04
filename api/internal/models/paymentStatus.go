package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"ledgerbolt.systems/internal/queries"
)

type PaymentStatus struct {
	ID int `json:"ID"`
	Status string `json:"Status" validate:"required,min=3,max=50"`
	Color  string `json:"Color" validate:"min=3,max=100"`
}
func GetPaymentStatus(conn *pgxpool.Pool, ctx *gin.Context) ([]PaymentStatus, error) {
	var psSlice []PaymentStatus

	rows, err := conn.Query(ctx, queries.GetPaymentStatus)
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

	row := conn.QueryRow(ctx, queries.GetSinglePaymentStatus, psID)

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
	_, err := conn.Exec(ctx, queries.CreatePaymentStatus, ps.Status, ps.Color)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePaymentStatus(conn *pgxpool.Pool, ps PaymentStatus, ctx *gin.Context, psID string) error {
	
	cmd, err := conn.Exec(
		ctx,
		queries.UpdatePaymentStatus,
		ps.Status,
		ps.Color,
		psID,
	)
	if err != nil {
		return err
	}

    if cmd.RowsAffected() == 0 {
        return errors.New("Payment status does not exist, invalid ID")
    }

	return nil
}

func DestroyPaymentStatus(conn *pgxpool.Pool, ctx *gin.Context, psID string) error {

	cmd, err := conn.Exec(
		ctx,
		queries.DestroyPaymentStatus,
		psID,
	)
	if err != nil {
		return err
	}

    if cmd.RowsAffected() == 0 {
        return errors.New("Payment status does not exist, invalid ID")
    }

	return nil
}
