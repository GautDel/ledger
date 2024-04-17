package models

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"ledgerbolt.systems/internal/queries"
)

type Invoice struct {
	ID            string
	InvoiceID     string
	InvoiceDate   time.Time
	CompName      string
	CompAddress   *string
	CompEmail     *string
	CompPhone     *string
	SubTotal      float64
	Total         float64
	DueDate       *time.Time
	ClientName    string
	ClientAddress *string
	ClientPhone   *string
	ClientEmail   *string
	TaxPercent    *int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	InvoiceItems  []InvoiceItem
}

type InvoiceItem struct {
	ID          string
	InvoiceID   string
	Qty         int
	Name        string
	Description string
	UnitPrice   *float64
	HourlyPrice *float64
	TotalPrice  float64
}

type InvoiceRequest struct {
	InvoiceDate   time.Time
	CompName      string
	CompAddress   string
	CompEmail     string
	CompPhone     string
	SubTotal      float64
	Total         float64
	DueDate       time.Time
	ClientName    string
	ClientAddress string
	ClientPhone   string
	ClientEmail   string
	TaxPercent    int
	ClientID      int
	ProjectID     int
	CreatedAt     time.Time
	UpdatedAT     time.Time
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

func CreateInvoice(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	i InvoiceRequest,
	userID string,
) error {
	_, err := conn.Exec(
		ctx,
		queries.CreateInvoice,
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
	)
	if err != nil {
		return err
	}

	return nil
}

func UpdateInvoice(
	conn *pgxpool.Pool,
	ctx *gin.Context,
	i InvoiceRequest,
	iID string,
	userID string,
) error {

	_, err := conn.Exec(
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
