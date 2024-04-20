package queries
const GetPaymentStatus = `
SELECT * FROM payment_status`

const GetSinglePaymentStatus = `
SELECT * FROM payment_status WHERE id = $1`

const CreatePaymentStatus = `
INSERT INTO payment_status (
    status,
    color
) VALUES ($1, $2)`

const UpdatePaymentStatus = `
UPDATE payment_status SET
    status = $1,
    color = $2
WHERE id = $3`

const DestroyPaymentStatus = `
DELETE FROM payment_status
WHERE id = $1`

