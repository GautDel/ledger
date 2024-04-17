package models

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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
	query := `
   SELECT
   invoices.id,
   invoices.invoice_id,
   invoices.invoice_date,
   invoices.comp_name,
   invoices.comp_address,
   invoices.comp_email,
   invoices.comp_phone,
   invoices.sub_total,
   invoices.total,
   invoices.due_date,
   invoices.client_name,
   invoices.client_address,
   invoices.client_phone,
   invoices.client_email,
   invoices.tax_percent,
   invoices.created_at,
   invoices.updated_at,
   invoice_items.id,
   invoice_items.invoice_id,
   invoice_items.qty,
   invoice_items.name,
   invoice_items.description,
   invoice_items.unit_price,
   invoice_items.hourly_price,
   invoice_items.total_price
   FROM invoices
   LEFT JOIN invoice_items ON invoices.id = invoice_items.invoice_id
   WHERE invoices.user_id = $1
   AND invoice_items.id IS NOT NULL`

	rows, err := conn.Query(ctx, query, userID)
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
	query := `
    INSERT INTO invoices(
        invoice_id,
        invoice_date,
        comp_name,
        comp_address,
        comp_email,
        comp_phone,
        sub_total,
        total,
        due_date,
        client_name,
        client_address,
        client_phone,
        client_email,
        tax_percent,
        user_id,
        client_id,
        project_id
    ) VALUES (
        $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,
        $12,$13,$14,$15,$16,$17
    )`

	_, err := conn.Exec(
		ctx,
		query,
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
	query := `
    UPDATE invoices SET
        invoice_date = $1,
        comp_name = $2,
        comp_address = $3,
        comp_email = $4,
        comp_phone = $5,
        sub_total = $6,
        total = $7,
        due_date = $8,
        client_name = $9,
        client_address = $10,
        client_phone = $11,
        client_email = $12,
        tax_percent = $13,
        client_id = $14,
        project_id = $15
        WHERE user_id = $16 AND id = $17
    `
	_, err := conn.Exec(
		ctx,
		query,
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
	query := `
    DELETE FROM invoices
    WHERE user_id = $1 AND id = $2`

	_, err := conn.Exec(
		ctx,
		query,
		userID,
		iID,
	)

	if err != nil {
		return err
	}

	return nil
}
