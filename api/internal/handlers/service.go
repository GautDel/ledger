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

func getServices(ctx *gin.Context) {
	conn := db.GetPool()

	services, err := models.GetServices(conn, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get services", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, services)
}

func getService(ctx *gin.Context) {
	conn := db.GetPool()
	sID := ctx.Param("id")

	service, err := models.GetService(conn, ctx, auth.GetUser(ctx), sID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get service", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, service)
}

func createService(ctx *gin.Context) {
	conn := db.GetPool()
	var reqBody models.Service
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create service", "error": err.Error()})
		return
	}

	err = models.CreateService(conn, ctx, reqBody, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create service", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully created service!"})
}

func updateService(ctx *gin.Context) {
	conn := db.GetPool()
	sID := ctx.Param("id")

	var reqBody models.Service
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create service", "error": err.Error()})
		return
	}

	err = models.UpdateService(conn, ctx, reqBody, auth.GetUser(ctx), sID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update service", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated service!"})
}

func destroyService(ctx *gin.Context) {
	conn := db.GetPool()
    sID := ctx.Param("id")

    err := models.DestroyService(conn, ctx, auth.GetUser(ctx), sID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to remove service", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully removed service!"})
}
