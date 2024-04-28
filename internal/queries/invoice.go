package queries

const GetInvoices = `
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
   invoices.client_id,
   invoices.project_id,
   invoices.created_at,
   invoices.updated_at,
   invoices.status,
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
WHERE invoices.user_id = $1 AND invoice_items.id IS NOT NULL`

const GetInvoice = `
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
   invoices.client_id,
   invoices.project_id,
   invoices.created_at,
   invoices.updated_at,
   invoices.status,
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
WHERE invoices.user_id = $1 AND invoices.id = $2 AND invoice_items.id IS NOT NULL`

const CreateInvoice = `
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
    project_id,
    status
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)
RETURNING id`

const UpdateInvoice = `
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
    project_id = $15,
    status = $16
WHERE user_id = $17 AND id = $18`

const DestroyInvoice = `
DELETE FROM invoices
WHERE user_id = $1 AND id = $2`

const CreateInvoiceItem = `
INSERT INTO invoice_items (
    id,
    invoice_id, 
    qty, 
    name, 
    description,
    unit_price,
    hourly_price,
    total_price,
    user_id
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

const UpdateInvoiceItem = `
UPDATE invoice_items SET
    qty = $1,
    name = $2,
    description = $3,
    unit_price = $4,
    hourly_price = $5,
    total_price = $6
WHERE user_id = $7 AND invoice_id = $8`

const UpsertInvoiceItem = `
INSERT INTO invoice_items (
    id,
    invoice_id,
    qty,
    name,
    description,
    unit_price,
    hourly_price,
    total_price,
    user_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT (id)
DO UPDATE 
SET 
    qty = excluded.qty, 
    name = excluded.name, 
    description = excluded.description, 
    unit_price = excluded.unit_price, 
    hourly_price = excluded.hourly_price, 
    total_price = excluded.total_price
WHERE excluded.user_id = $9`

const DestroyInvoiceItem = `
DELETE FROM invoice_items
WHERE user_id = $1 AND id = $2`

const GetInvoiceID = `
SELECT id FROM invoice_ids 
WHERE invoice_id = $1`

const CreateInvoiceID = `
INSERT INTO invoice_ids (
    id,
    invoice_id
) VALUES ($1, $2)`

const CountInvoiceIDs = `
SELECT COUNT(*) FROM invoice_ids`
