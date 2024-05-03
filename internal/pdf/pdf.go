package pdf

import (
	"log"

	"github.com/gin-gonic/gin"
	"ledgerbolt.systems/internal/auth"
	"ledgerbolt.systems/internal/db"
	"ledgerbolt.systems/internal/models"
)

func loadData(ctx *gin.Context, iID string) (models.Invoice, error) {
	conn := db.GetPool()

	// Invoice returns as single object inside array
	invoice, err := models.GetInvoice(conn, ctx, iID, auth.GetUser(ctx))
	if err != nil {
		return invoice[0], err
	}

	return invoice[0], nil
}

func initParams() {
	page.AddPage()
	page.SetFont("Helvetica", "", ts.base)
	page.SetTextColor(defaultColor.red, defaultColor.green, defaultColor.blue)

}

func New(ctx *gin.Context, iID string) {

	data, err := loadData(ctx, iID)
	if err != nil {
		log.Println(err)
		return
	}

    // Create french page
    newFr(data)


	// Create file
	err = page.OutputFileAndClose(data.InvoiceID + ".pdf")
	if err != nil {
		panic(err)
	}
}
