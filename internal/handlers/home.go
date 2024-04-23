package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ledgerbolt.systems/utils"
)

func homeHandler(ctx *gin.Context) {

    utils.GenInvoiceID(1)
    ctx.JSON(http.StatusOK, gin.H{
        "message": `Welcome to the home route of ledgerbolt API. Docs will be here in the future. Yours truly, Gauthier - The Dev`})
}
