package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ledgerbolt.systems/internal/db"
	"ledgerbolt.systems/internal/models"
)

type ClientRB struct {
	FirstName   string `json:"FirstName" validate:"required,min=3,max=50"`
	LastName    string `json:"LastName" validate:"required,min=3,max=50"`
	Description string `json:"Description" validate:"required,max=2000"`
}

type SearchRB struct {
    Search string `json:"Search"`
}

func getClientsHandler(ctx *gin.Context) {
	conn := db.GetPool()

	clients, err := models.GetClients(conn, ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve clients from database", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, clients)
}

func getClientHandler(ctx *gin.Context) {
	conn := db.GetPool()
	clientID := ctx.Param("id")

	client, err := models.GetClient(conn, clientID, ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve client from database", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, client)
}

func searchClientsHandler(ctx *gin.Context) {
    conn := db.GetPool()

    var reqBody SearchRB
    err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to find clients", "error": err})
		return
	}
    
    clients, err := models.SearchClients(conn, reqBody.Search, ctx) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to find clients", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, clients)
}

func newClientHandler(ctx *gin.Context) {
	conn := db.GetPool()

	var reqBody ClientRB
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create client", "error": err})
		return
	}

	validate := validator.New()
	if err := validate.Struct(reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := models.ClientRequest{
		FirstName:   reqBody.FirstName,
		LastName:    reqBody.LastName,
		Description: reqBody.Description,
	}

	err = models.NewClient(conn, client, ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create client", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Client successfully created!"})
}

func updateClientHandler(ctx *gin.Context) {
	conn := db.GetPool()
    clientID := ctx.Param("id")

	var reqBody ClientRB
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update client", "error": err})
		return
	}

	validate := validator.New()
	if err := validate.Struct(reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := models.ClientRequest{
		FirstName:   reqBody.FirstName,
		LastName:    reqBody.LastName,
		Description: reqBody.Description,
	}

	err = models.UpdateClient(conn, client, ctx, clientID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update client", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Client successfully updated!"})
}

func destroyClientHandler(ctx *gin.Context) {
    conn := db.GetPool()
    clientID := ctx.Param("id")

    err := models.DestroyClient(conn, clientID, ctx)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't delete client", "error": err})
    }


	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully removed client!"})
}
