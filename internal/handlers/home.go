package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func homeHandler(ctx *gin.Context) {

    ctx.JSON(http.StatusOK, gin.H{
        "message": `Welcome to the home route of ledgerbolt API. Docs will be here in the future. Yours truly, Gauthier - The Dev`})
}
