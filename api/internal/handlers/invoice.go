package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ledgerbolt.systems/internal/auth"
	"ledgerbolt.systems/internal/db"
	"ledgerbolt.systems/internal/models"
	"ledgerbolt.systems/internal/validator"
)

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

func getInvoice(ctx *gin.Context) {
	conn := db.GetPool()
    iID := ctx.Param("id")

	invoice, err := models.GetInvoice(conn, ctx, iID, auth.GetUser(ctx))
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
	ctx.JSON(http.StatusOK, invoice)
}

func createInvoice(ctx *gin.Context) {
	conn := db.GetPool()
	var reqBody models.Invoice

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
				"message": "Failed to create invoice",
				"error":   err.Error(),
			},
		)
		return
	}

	err = models.CreateInvoice(conn, ctx, reqBody, auth.GetUser(ctx))
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

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully created invoice!",
	},
	)
}

func updateInvoice(ctx *gin.Context) {
	conn := db.GetPool()
	var reqBody models.Invoice
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

	err = models.UpdateInvoice(conn, ctx, reqBody, iID, auth.GetUser(ctx))
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Invoice successfully updated!"})
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
