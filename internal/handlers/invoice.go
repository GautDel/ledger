package handlers

import (
)

//type InvoiceRB struct {
//	FirstName   string `json:"FirstName" validate:"required,min=3,max=50"`
//	LastName    string `json:"LastName" validate:"required,min=3,max=50"`
//	CompanyName string `json:"CompanyName" validate:"min=3,max=100"`
//	Email       string `json:"Email" validate:"required,min=5,max=320"`
//	Phone       string `json:"Phone" validate:"required,min=5,max=50"`
//	Address     string `json:"Address" validate:"required,min=5,max=1000"`
//}
//
//func getInvoice(ctx *gin.Context) {
//	conn := db.GetPool()
//
//	invoices, err := models.GetInvoices(conn, ctx, auth.GetUser(ctx))
//	if err != nil {
//		log.Println(err)
//		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get invoices from database", "error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, invoices)
//}
