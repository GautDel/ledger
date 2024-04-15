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

type UserRB struct {
	FirstName   string `json:"FirstName" validate:"required,min=3,max=50"`
	LastName    string `json:"LastName" validate:"required,min=3,max=50"`
	CompanyName string `json:"CompanyName" validate:"min=3,max=100"`
	Email       string `json:"Email" validate:"required,min=5,max=320"`
	Phone       string `json:"Phone" validate:"required,min=5,max=50"`
	Address     string `json:"Address" validate:"required,min=5,max=1000"`
}

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
	var reqBody UserRB

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
		return
	}

    err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update user", "error": err.Error()})
        return
	}

	user := models.UserRequest{
		FirstName:   reqBody.FirstName,
		LastName:    reqBody.LastName,
		CompanyName: reqBody.CompanyName,
		Email: reqBody.Email,
		Phone: reqBody.Phone,
		Address: reqBody.Address,
	}

	err = models.UpdateUser(conn, user, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated user", "error": err.Error()})
}
