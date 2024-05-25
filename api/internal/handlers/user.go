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

func GetUserHandler(ctx *gin.Context) {
	conn := db.GetPool()

	user, err := models.GetUser(conn, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
        errMsg := utils.ErrorHandler(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update user", "error": errMsg})
		return
	}

	err = models.UpdateUser(conn, reqBody, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated user"})
}
