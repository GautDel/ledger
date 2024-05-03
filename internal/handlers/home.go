package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ledgerbolt.systems/internal/pdf"
)

func homeHandler(ctx *gin.Context) {
    pdf.New(ctx, "8505bb57-cd55-4d6f-a7a1-632c3932fd56")
    ctx.JSON(http.StatusOK, gin.H{
        "message": `Welcome to the home route of ledgerbolt API. Docs will be here in the future. Yours truly, Gauthier - The Dev`})
}
