package models

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"ledgerbolt.systems/internal/queries"
)

type Service struct {
	ID          int       `json:"ID"`
	Name        string    `json:"Name" validate:"required,min=3,max=50"`
	Description string    `json:"Description" validate:"required,min=3,max=50"`
	UnitPrice   float64   `json:"UnitPrice"`
	HourlyPrice float64   `json:"HourlyPrice"`
	Tax         int       `json:"Tax"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

func GetServices(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	userID string,
) ([]Service, error) {
	var services []Service
	var service Service

	rows, err := conn.Query(ctx, queries.GetServices, userID)
	if err != nil {
		return services, err
	}
	defer rows.Close()

	_, err = pgx.ForEachRow(rows,
		[]any{
			&service.ID,
			&service.Name,
			&service.Description,
			&service.UnitPrice,
			&service.HourlyPrice,
			&service.Tax,
			&service.CreatedAt,
			&service.UpdatedAt,
		}, func() error {
			services = append(services, service)
			return nil
		})

	if err != nil {
		return services, err
	}

	return services, nil
}

func GetService(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	userID string,
	sID string,
) (Service, error) {
	var service Service
	err := conn.QueryRow(ctx, queries.GetSingleService, userID, sID).Scan(
		&service.ID,
		&service.Name,
		&service.Description,
		&service.UnitPrice,
		&service.HourlyPrice,
		&service.Tax,
		&service.CreatedAt,
		&service.UpdatedAt,
	)
	if err != nil {
		return service, err
	}

	return service, nil
}

func CreateService(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	service Service,
	userID string,
) error {
	cmd, err := conn.Exec(
		ctx,
		queries.CreateService,
		service.Name,
		service.Description,
		service.UnitPrice,
		service.HourlyPrice,
		service.Tax,
		userID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("Something went wrong, unable to create service")
	}

	return nil
}

func UpdateService(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	service Service,
	userID string,
	sID string,
) error {
	cmd, err := conn.Exec(
		ctx,
		queries.UpdateService,
		service.Name,
		service.Description,
		service.UnitPrice,
		service.HourlyPrice,
		service.Tax,
		userID,
		sID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("Something went wrong, unable to update service")
	}

	return nil
}

func DestroyService(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	userID string,
	sID string,
) error {
	cmd, err := conn.Exec(ctx, queries.DestroyService, userID, sID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("Something went wrong, unable to remove service")
	}

	return nil
}
