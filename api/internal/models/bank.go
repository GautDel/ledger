package models

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"ledgerbolt.systems/internal/encrypt"
	"ledgerbolt.systems/internal/queries"
)

type Bank struct {
	ID           uuid.UUID `json:"ID"`
	BIC          string    `json:"BIC" validate:"required,min=8,max=11"`
	IBAN         string    `json:"IBAN" validate:"required,min=14,max=34"`
	AccountName  string    `json:"AccountName" validate:"required,min=3,max=50"`
	BankName     string    `json:"BankNAme" validate:"required,min=3,max=100"`
	BankLocation string    `json:"BankLocation" validate:"required,min=3,max=25"`
}
type EncBank struct {
	ID           uuid.UUID
	BIC          []byte
	IBAN         []byte
	AccountName  []byte
	BankName     []byte
	BankLocation []byte
}

func GetBanks(conn *pgxpool.Pool, ctx *gin.Context, userID string) ([]Bank, error) {
	var encBank EncBank
	var encBanks []EncBank
	var banks []Bank

	rows, err := conn.Query(ctx, queries.GetBankDetails, userID)
	if err != nil {
		return banks, err
	}

    defer rows.Close()

	_, err = pgx.ForEachRow(rows,
		[]any{
			&encBank.ID,
			&encBank.BIC,
			&encBank.IBAN,
			&encBank.AccountName,
			&encBank.BankName,
			&encBank.BankLocation,
		}, func() error {
			encBanks = append(encBanks, encBank)

			return nil
		})

	for _, encBank := range encBanks {
		var bank Bank
		bank.ID = encBank.ID
		bank.BIC = encrypt.DecryptField(ctx, encBank.BIC)
		bank.IBAN = encrypt.DecryptField(ctx, encBank.IBAN)
		bank.AccountName = encrypt.DecryptField(ctx, encBank.AccountName)
		bank.BankName = encrypt.DecryptField(ctx, encBank.BankName)
		bank.BankLocation = encrypt.DecryptField(ctx, encBank.BankLocation)
		banks = append(banks, bank)
	}

	return banks, err
}

func GetBank(conn *pgxpool.Pool, ctx *gin.Context, userID string, bankID string) (Bank, error) {
	var encBank EncBank
	var bank Bank	

	rows := conn.QueryRow(ctx, queries.GetSingleBankDetails, userID, bankID)

	err := rows.Scan(
		&encBank.ID,
		&encBank.BIC,
		&encBank.IBAN,
		&encBank.AccountName,
		&encBank.BankName,
		&encBank.BankLocation,
	)
	if err != nil {
		log.Println(err)
		return bank, err
	}

	bank.ID = encBank.ID
	bank.BIC = encrypt.DecryptField(ctx, encBank.BIC)
	bank.IBAN = encrypt.DecryptField(ctx, encBank.IBAN)
	bank.AccountName = encrypt.DecryptField(ctx, encBank.AccountName)
	bank.BankName = encrypt.DecryptField(ctx, encBank.BankName)
	bank.BankLocation = encrypt.DecryptField(ctx, encBank.BankLocation)

	return bank, err
}

func CreateBank(conn *pgxpool.Pool, ctx *gin.Context, bank *Bank, userID string) error {
	var encBank EncBank

	uuid := uuid.New()
	encBank.BIC = encrypt.EncryptField(ctx, bank.BIC)
	encBank.IBAN = encrypt.EncryptField(ctx, bank.IBAN)
	encBank.AccountName = encrypt.EncryptField(ctx, bank.AccountName)
	encBank.BankName = encrypt.EncryptField(ctx, bank.BankName)
	encBank.BankLocation = encrypt.EncryptField(ctx, bank.BankLocation)

	_, err := conn.Exec(
		ctx,
		queries.CreateBankDetails,
		uuid,
		encBank.BIC,
		encBank.IBAN,
		encBank.AccountName,
		encBank.BankName,
		encBank.BankLocation,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}

func UpdateBank(
    conn *pgxpool.Pool, 
    ctx *gin.Context, 
    bank Bank, 
    userID string, 
    bankID string,
) error {

	var encBank EncBank

	encBank.BIC = encrypt.EncryptField(ctx, bank.BIC)
	encBank.IBAN = encrypt.EncryptField(ctx, bank.IBAN)
	encBank.AccountName = encrypt.EncryptField(ctx, bank.AccountName)
	encBank.BankName = encrypt.EncryptField(ctx, bank.BankName)
	encBank.BankLocation = encrypt.EncryptField(ctx, bank.BankLocation)

	cmd, err := conn.Exec(
		ctx,
		queries.UpdateBankDetails,
		encBank.BIC,
		encBank.IBAN,
		encBank.AccountName,
		encBank.BankName,
		encBank.BankLocation,
		userID,
		bankID,
	)
	if err != nil {
		return err
	}

    if cmd.RowsAffected() == 0 {
        return errors.New("Something went wrong, unable to update bank details")
    }

	return nil
}

func DestroyBank(
    conn *pgxpool.Pool, 
    ctx *gin.Context, 
    userID string, 
    bankID string,
) error {
    cmd, err := conn.Exec(ctx, queries.DestroyBankDetails, userID, bankID) 
	if err != nil {
		return err
	}

    if cmd.RowsAffected() == 0 {
        return errors.New("Something went wrong, unable to delete bank details")
    }

	return nil
}
