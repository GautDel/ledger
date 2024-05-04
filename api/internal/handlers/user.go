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

func GetUserHandler(ctx *gin.Context) {
	conn := db.GetPool()

	user, err := models.GetUser(conn, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
        log.Println("hit here")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user from database", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func UpdateUserHandler(ctx *gin.Context) {
	conn := db.GetPool()
	var reqBody models.User

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
        log.Println("hit here")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
		return
	}

    err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
        log.Println("hit here 2")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update user", "error": err.Error()})
        return
	}

	err = models.UpdateUser(conn, reqBody, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
        log.Println("hit here 3")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated user"})
}
