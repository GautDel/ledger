package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ledgerbolt.systems/internal/db"
	"ledgerbolt.systems/internal/models"
	"ledgerbolt.systems/internal/validator"
)

func getPaymentStatus(ctx *gin.Context) {
	conn := db.GetPool()

	paymentStatus, err := models.GetPaymentStatus(conn, ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get payment status", "error": err.Error()})
	}

	ctx.JSON(http.StatusOK, paymentStatus)
}

func getSinglePaymentStatus(ctx *gin.Context) {
	conn := db.GetPool()
	psID := ctx.Param("id")

	paymentStatus, err := models.GetSinglePaymentStatus(conn, ctx, psID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get payment status", "error": err.Error()})
	}

	ctx.JSON(http.StatusOK, paymentStatus)
}

func createPaymentStatus(ctx *gin.Context) {
	conn := db.GetPool()
	var reqBody models.PaymentStatus
	err := ctx.ShouldBindJSON(&reqBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update payment status", "error": err.Error()})
		return
	}

	err = models.CreatePaymentStatus(conn, reqBody, ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create payment status", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment status successfully created!"})
}

func updatePaymentStatus(ctx *gin.Context) {
	conn := db.GetPool()
	psID := ctx.Param("id")
	var reqBody models.PaymentStatus
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update payment status", "error": err.Error()})
		return
	}

	err = models.UpdatePaymentStatus(conn, reqBody, ctx, psID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update payment status", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment status successfully updated!"})
}

func destroyPaymentStatus(ctx *gin.Context) {
	conn := db.GetPool()
	psID := ctx.Param("id")

	err := models.DestroyPaymentStatus(conn, ctx, psID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to remove payment status", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment status successfully removed!"})
}
