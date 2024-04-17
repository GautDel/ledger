package models

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"ledgerbolt.systems/internal/queries"
)

type InvoiceItem struct {
	ID          string  `json:"ID"`
	InvoiceID   string  `json:"InvoiceID"`
	Qty         int     `json:"Qty" validate:"required"`
	Name        string  `json:"Name" validate:"required,min=3,max=100"`
	Description string  `json:"Description" validate:"required,min=5,max=1000"`
	UnitPrice   float64 `json:"UnitPrice"`
	HourlyPrice float64 `json:"HourlyPrice"`
	TotalPrice  float64 `json:"TotalPrice" validate:"required"`
}

type Invoice struct {
	ID            string        `json:"ID"`
	InvoiceID     string        `json:"InvoiceID"`
	InvoiceDate   string        `json:"InvoiceDate" validate:"required"`
	CompName      string        `json:"CompName" validate:"min=3,max=100"`
	CompAddress   string        `json:"CompAddress" validate:"required,min=5,max=500"`
	CompEmail     string        `json:"CompEmail" validate:"required,min=5,max=320"`
	CompPhone     string        `json:"CompPhone" validate:"required,min=5,max=30"`
	SubTotal      float64       `json:"SubTotal" validate:"required"`
	Total         float64       `json:"Total" validate:"required"`
	DueDate       string        `json:"DueDate" validate:"min=5,max=25"`
	ClientName    string        `json:"ClientName" validate:"min=3,max=100"`
	ClientAddress string        `json:"ClientAddress" validate:"required,min=5,max=500"`
	ClientPhone   string        `json:"ClientPhone" validate:"required,min=5,max=50"`
	ClientEmail   string        `json:"ClientEmail" validate:"required,min=3,max=320"`
	TaxPercent    int           `json:"TaxPercent" validate:"required"`
	ClientID      int           `json:"ClientID" validate:"required"`
	ProjectID     int           `json:"ProjectID" validate:"required"`
	CreatedAt     time.Time     `json:"CreatedAt"`
	UpdatedAt     time.Time     `json:"UpdatedAt"`
	InvoiceItems  []InvoiceItem `json:"InvoiceItems" validate:"required"`
}

func GetInvoices(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	userID string,
) ([]Invoice, error) {
	var invoices []Invoice

	rows, err := conn.Query(ctx, queries.GetInvoices, userID)
	if err != nil {
		log.Println("skill issue", err)
		return invoices, err
	}

	defer rows.Close()

	for rows.Next() {
		var i Invoice
		var item InvoiceItem
		err := rows.Scan(
			&i.ID,
			&i.InvoiceID,
			&i.InvoiceDate,
			&i.CompName,
			&i.CompAddress,
			&i.CompEmail,
			&i.CompPhone,
			&i.SubTotal,
			&i.Total,
			&i.DueDate,
			&i.ClientName,
			&i.ClientAddress,
			&i.ClientPhone,
			&i.ClientEmail,
			&i.TaxPercent,
			&i.CreatedAt,
			&i.UpdatedAt,
			&item.ID,
			&item.InvoiceID,
			&item.Qty,
			&item.Name,
			&item.Description,
			&item.UnitPrice,
			&item.HourlyPrice,
			&item.TotalPrice,
		)
		if err != nil {
			return invoices, err
		}

		// Check if the current invoice is different from the last one
		if len(invoices) == 0 || i.ID != invoices[len(invoices)-1].ID {
			invoices = append(invoices, i)
		}

		// Append the item to the last invoice
		invoices[len(invoices)-1].InvoiceItems = append(invoices[len(invoices)-1].InvoiceItems, item)
	}

	if err := rows.Err(); err != nil {
		return invoices, err
	}

	return invoices, nil
}
func GetInvoice(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	iID string,
	userID string,
) ([]Invoice, error) {
	var invoices []Invoice

	rows, err := conn.Query(ctx, queries.GetInvoice, userID, iID)
	if err != nil {
		log.Println("skill issue", err)
		return invoices, err
	}

	defer rows.Close()

	for rows.Next() {
		var i Invoice
		var item InvoiceItem
		err := rows.Scan(
			&i.ID,
			&i.InvoiceID,
			&i.InvoiceDate,
			&i.CompName,
			&i.CompAddress,
			&i.CompEmail,
			&i.CompPhone,
			&i.SubTotal,
			&i.Total,
			&i.DueDate,
			&i.ClientName,
			&i.ClientAddress,
			&i.ClientPhone,
			&i.ClientEmail,
			&i.TaxPercent,
			&i.CreatedAt,
			&i.UpdatedAt,
			&item.ID,
			&item.InvoiceID,
			&item.Qty,
			&item.Name,
			&item.Description,
			&item.UnitPrice,
			&item.HourlyPrice,
			&item.TotalPrice,
		)
		if err != nil {
			return invoices, err
		}

		// Check if the current invoice is different from the last one
		if len(invoices) == 0 || i.ID != invoices[len(invoices)-1].ID {
			invoices = append(invoices, i)
		}

		// Append the item to the last invoice
		invoices[len(invoices)-1].InvoiceItems = append(invoices[len(invoices)-1].InvoiceItems, item)
	}

	if err := rows.Err(); err != nil {
		return invoices, err
	}

	return invoices, nil
}

func CreateInvoice(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	i Invoice,
	userID string,
) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	var invoiceID string
	testInvID := "20240408"

	err = tx.QueryRow(
		ctx,
		queries.CreateInvoice,
		testInvID,
		i.InvoiceDate,
		i.CompName,
		i.CompAddress,
		i.CompEmail,
		i.CompPhone,
		i.SubTotal,
		i.Total,
		i.DueDate,
		i.ClientName,
		i.ClientAddress,
		i.ClientPhone,
		i.ClientEmail,
		i.TaxPercent,
		userID,
		i.ClientID,
		i.ProjectID,
	).Scan(&invoiceID)

	if err != nil {
		return err
	}

	// Insert invoice items
	for _, item := range i.InvoiceItems {
		_, err := tx.Exec(
			ctx,
			queries.CreateInvoiceItem,
			invoiceID,
			item.Qty,
			item.Name,
			item.Description,
			item.UnitPrice,
			item.HourlyPrice,
			item.TotalPrice,
			userID,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func UpdateInvoice(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	i Invoice,
	iID string,
	userID string,
) error {

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)
	_, err = tx.Exec(
		ctx,
		queries.UpdateInvoice,
		i.InvoiceDate,
		i.CompName,
		i.CompAddress,
		i.CompEmail,
		i.CompPhone,
		i.SubTotal,
		i.Total,
		i.DueDate,
		i.ClientName,
		i.ClientAddress,
		i.ClientPhone,
		i.ClientEmail,
		i.TaxPercent,
		i.ClientID,
		i.ProjectID,
		userID,
		iID,
	)

	if err != nil {
		return err
	}

	// Update invoice items
	for _, item := range i.InvoiceItems {
		log.Println(item.Qty)
		_, err := tx.Exec(
			ctx,
			queries.UpdateInvoiceItem,
			item.Qty,
			item.Name,
			item.Description,
			item.UnitPrice,
			item.HourlyPrice,
			item.TotalPrice,
			userID,
			iID,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func DestroyInvoice(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	iID string,
	userID string,
) error {
	_, err := conn.Exec(
		ctx,
		queries.DestroyInvoice,
		userID,
		iID,
	)

	if err != nil {
		return err
	}

	return nil
}
