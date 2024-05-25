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

func getClientsHandler(ctx *gin.Context) {
	conn := db.GetPool()
	sortBy := ctx.Param("sort")

	clients, err := models.GetClients(conn, ctx, auth.GetUser(ctx), sortBy)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve clients from database", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, clients)
}

func getClientHandler(ctx *gin.Context) {
	conn := db.GetPool()
	clientID := ctx.Param("id")

	client, err := models.GetClient(conn, clientID, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve client from database", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, client)
}

func searchClientsHandler(ctx *gin.Context) {
	conn := db.GetPool()

	var reqBody models.SearchClient
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to find client", "error": err.Error()})
		return
	}

	clients, err := models.SearchClients(conn, reqBody, ctx, auth.GetUser(ctx))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to find clients", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, clients)
}

func newClientHandler(ctx *gin.Context) {
	conn := db.GetPool()

	var reqBody models.Client
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		errMsg := utils.ErrorHandler(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create client", "error": errMsg})
		return
	}

	err = models.NewClient(conn, reqBody, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create client", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Client successfully created!"})
}

func updateClientHandler(ctx *gin.Context) {
	conn := db.GetPool()
	clientID := ctx.Param("id")

	var reqBody models.Client
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
		return
	}

	err = validator.Validate(&reqBody)

	if err != nil {
		log.Println(err)
		errMsg := utils.ErrorHandler(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update client", "error": errMsg})
		return
	}

	err = models.UpdateClient(conn, reqBody, ctx, clientID, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update client", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Client successfully updated!"})
}

func starClientHandler(ctx *gin.Context) {
	conn := db.GetPool()
	clientID := ctx.Param("id")

	var reqBody models.StarClient
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		errMsg := utils.ErrorHandler(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update client", "error": errMsg})
		return
	}

	err = models.UpdateStarClient(conn, reqBody, ctx, clientID, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update client", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Client successfully starred!"})
}

func destroyClientHandler(ctx *gin.Context) {
	conn := db.GetPool()
	clientID := ctx.Param("id")

	err := models.DestroyClient(conn, clientID, ctx, auth.GetUser(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't delete client", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully removed client!"})
}
