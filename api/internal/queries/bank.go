package queries

const GetBankDetails = `
SELECT 
    id,
    bic,
    iban,
    account_name,
    bank_name,
    bank_location
FROM bank_details
WHERE user_id = $1`

const GetSingleBankDetails = `
SELECT 
    id,
    bic,
    iban,
    account_name,
    bank_name,
    bank_location
FROM bank_details
WHERE user_id = $1 AND id = $2`

const CreateBankDetails = `
INSERT INTO bank_details(
    id,
    bic,
    iban,
    account_name,
    bank_name,
    bank_location,
    user_id
) VALUES ($1,$2,$3,$4,$5,$6,$7)`

const UpdateBankDetails = `
UPDATE bank_details SET
    bic = $1,
    iban = $2,
    account_name = $3,
    bank_name = $4,
    bank_location = $5
WHERE user_id = $6 AND id = $7`

const DestroyBankDetails = `
DELETE FROM bank_details
WHERE user_id = $1 AND id = $2` 
