package queries

const GetUser = `
SELECT 
    first_name,
    last_name,
    company_name,
    email,
    phone,
    address,
    company_num
FROM users 
WHERE id = $1`

const UpdateUser = `
UPDATE users SET
    first_name = $1,
    last_name = $2,
    company_name = $3,
    email = $4,
    phone = $5,
    address = $6,
    company_num = $7
WHERE id = $8`

