package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ledgerbolt.systems/internal/auth"
	"ledgerbolt.systems/internal/db"
	"ledgerbolt.systems/internal/models"
)

func homeHandler(ctx *gin.Context) {
	conn := db.GetPool()

	user, userErr := models.GetUser(conn, ctx, auth.GetUser(ctx))
	if userErr != nil {
		log.Println(userErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting user", "error": userErr.Error()})
	}

	ctx.JSON(http.StatusOK, user)
}
