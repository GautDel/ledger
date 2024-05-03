package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ledgerbolt.systems/internal/pdf"
)

func homeHandler(ctx *gin.Context) {
    pdf.New(ctx, "22add092-b7fa-45ea-8e50-06feadc6981c", "3bc9b2d6-2678-4a84-a18d-b755ebb2a788")
    ctx.JSON(http.StatusOK, gin.H{
        "message": `Welcome to the home route of ledgerbolt API. Docs will be here in the future. Yours truly, Gauthier - The Dev`})
}
