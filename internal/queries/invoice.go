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
WHERE invoices.user_id = $1 AND invoice_items.id IS NOT NULL`

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
    project_id
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`

const UpdateInvoice = `
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
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`

const DestroyInvoice = `
DELETE FROM invoices
WHERE user_id = $1 AND id = $2`
