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

func getBanks(ctx *gin.Context) {
	conn := db.GetPool()

	bank, err := models.GetBanks(conn, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get bank details", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bank)
}

func getBank(ctx *gin.Context) {
	conn := db.GetPool()
	bankID := ctx.Param("id")

	bank, err := models.GetBank(conn, ctx, auth.GetUser(ctx), bankID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get bank details", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bank)
}

func createBank(ctx *gin.Context) {
	conn := db.GetPool()

	var reqBody models.Bank

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create bank", "error": err.Error()})
		return
	}

	err = models.CreateBank(conn, ctx, &reqBody, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create bank", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Bank successfully created!"})
}

func updateBank(ctx *gin.Context) {
	conn := db.GetPool()
	var reqBody models.Bank
	bankID := ctx.Param("id")

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
        return
	}

	err = models.UpdateBank(conn, ctx, reqBody, auth.GetUser(ctx), bankID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update bank details", "error": err.Error()})
        return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated Bank details"})
}

func destroyBank(ctx *gin.Context) {
	conn := db.GetPool()
	bankID := ctx.Param("id")

	err := models.DestroyBank(conn, ctx, auth.GetUser(ctx), bankID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to remove bank details", "error": err.Error()})
        return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully removed Bank details"})
}
