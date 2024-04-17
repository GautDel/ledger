package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ledgerbolt.systems/internal/auth"
	"ledgerbolt.systems/internal/db"
	"ledgerbolt.systems/internal/models"
	"ledgerbolt.systems/internal/validator"
	"ledgerbolt.systems/utils"
)

type InvoiceRB struct {
	InvoiceDate   string  `json:"InvoiceDate" validate:"required"`
	CompName      string  `json:"CompName" validate:"min=3,max=100"`
	CompAddress   string  `json:"CompAddress" validate:"required,min=5,max=500"`
	CompEmail     string  `json:"CompEmail" validate:"required,min=5,max=320"`
	CompPhone     string  `json:"CompPhone" validate:"required,min=5,max=30"`
	SubTotal      float64 `json:"SubTotal" validate:"required"`
	Total         float64 `json:"Total" validate:"required"`
	DueDate       string  `json:"DueDate" validate:"min=5,max=25"`
	ClientName    string  `json:"ClientName" validate:"min=3,max=100"`
	ClientAddress string  `json:"ClientAddress" validate:"required,min=5,max=500"`
	ClientPhone   string  `json:"ClientPhone" validate:"required,min=5,max=50"`
	ClientEmail   string  `json:"ClientEmail" validate:"required,min=3,max=320"`
	TaxPercent    int     `json:"TaxPercent" validate:"required"`
	ClientID      int     `json:"ClientID" validate:"required"`
	ProjectID     int     `json:"ProjectID" validate:"required"`
}

func getInvoices(ctx *gin.Context) {
	conn := db.GetPool()

	invoices, err := models.GetInvoices(conn, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Failed to get invoices from database",
				"error":   err.Error(),
			},
		)
		return
	}
	ctx.JSON(http.StatusOK, invoices)
}

func createInvoice(ctx *gin.Context) {
	conn := db.GetPool()
	var reqBody InvoiceRB

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Failed to read request body",
				"error":   err.Error(),
			},
		)
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Failed to create project",
				"error":   err.Error(),
			},
		)
		return
	}

	dateFormat := "2/1/2006"

	invoice := models.InvoiceRequest{
		InvoiceDate:   utils.ParseTime(dateFormat, reqBody.InvoiceDate),
		CompName:      reqBody.CompName,
		CompAddress:   reqBody.CompAddress,
		CompEmail:     reqBody.CompEmail,
		CompPhone:     reqBody.CompPhone,
		SubTotal:      reqBody.SubTotal,
		Total:         reqBody.Total,
		DueDate:       utils.ParseTime(dateFormat, reqBody.DueDate),
		ClientName:    reqBody.ClientName,
		ClientAddress: reqBody.ClientAddress,
		ClientPhone:   reqBody.ClientPhone,
		ClientEmail:   reqBody.ClientEmail,
		TaxPercent:    reqBody.TaxPercent,
		ClientID:      reqBody.ClientID,
		ProjectID:     reqBody.ProjectID,
	}

	err = models.CreateInvoice(conn, ctx, invoice, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Failed to create invoice",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, invoice)
}

func updateInvoice(ctx *gin.Context) {
	conn := db.GetPool()
	var reqBody InvoiceRB
	iID := ctx.Param("id")

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Failed to read request body",
				"error":   err.Error(),
			},
		)
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Failed to create project",
				"error":   err.Error(),
			},
		)
		return
	}

	dateFormat := "2/1/2006"

	invoice := models.InvoiceRequest{
		InvoiceDate:   utils.ParseTime(dateFormat, reqBody.InvoiceDate),
		CompName:      reqBody.CompName,
		CompAddress:   reqBody.CompAddress,
		CompEmail:     reqBody.CompEmail,
		CompPhone:     reqBody.CompPhone,
		SubTotal:      reqBody.SubTotal,
		Total:         reqBody.Total,
		DueDate:       utils.ParseTime(dateFormat, reqBody.DueDate),
		ClientName:    reqBody.ClientName,
		ClientAddress: reqBody.ClientAddress,
		ClientPhone:   reqBody.ClientPhone,
		ClientEmail:   reqBody.ClientEmail,
		TaxPercent:    reqBody.TaxPercent,
		ClientID:      reqBody.ClientID,
		ProjectID:     reqBody.ProjectID,
	}

	err = models.UpdateInvoice(conn, ctx, invoice, iID, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Failed to update invoice",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, invoice)
}

func destroyInvoice(ctx *gin.Context) {
	conn := db.GetPool()
	iID := ctx.Param("id")

	err := models.DestroyInvoice(conn, ctx, iID, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Failed to remove invoice",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Invoice successfully removed!"})
}
