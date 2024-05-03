package pdf

import (
	"log"
	"strconv"

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

func New(ctx *gin.Context, iID string) {

	data, err := loadData(ctx, iID)
	if err != nil {
		log.Println(err)
		return
	}

	// INIT
	page.AddPage()
	page.SetFont("Helvetica", "", ts.base)
	page.SetTextColor(defaultColor.red, defaultColor.green, defaultColor.blue)

	var currY, currX float64

	// START RENDER
	page.SetFont("Helvetica", "B", ts.lg)
	currY = renderPageHeader(pp.m, pp.m, currY, "Amagine Media", ts.lg, blue500)

	page.SetFont("Helvetica", "B", ts.xxl)
	currY = renderPageHeader(
		ax.getEnd(page.GetStringWidth(pp.h2)),
		pp.m,
		currY,
		pp.h2,
		ts.xxl,
		blue500,
	)

	// Margin after header
	currY += ms.lg

	_ = renderBox(
		ax.start,
		currY,
		bw.w5_12,
		currY,
		10,
		"Company Details",
		[]string{
			data.CompName,
			data.CompAddress,
			data.CompPhone,
			data.CompEmail,
		},
		neutral950,
	)

	currY = renderBox(
		ax.getEnd(bw.w4_12),
		currY,
		bw.w4_12,
		currY,
		10,
		"Client Details",
		[]string{
			data.ClientName,
			data.ClientEmail,
			data.ClientAddress,
			data.ClientPhone,
		},
		neutral950,
	)

	// Margin after Comp/Client details
	currY += ms.lg

	page.SetFont("Helvetica", "B", ts.base)
	currY, currX = renderCell(
		ax.start,
		currY,
		page.GetStringWidth("Invoice No. "),
		currY,
		ts.sm,
		"Invoice No.",
		neutral950,
        "L",
	)

	currY, currX = renderCell(
		currX,
		currY,
		page.GetStringWidth(data.InvoiceID),
		currY,
		ts.sm,
		data.InvoiceID,
		blue500,
        "L",
	)

	currY += ms.md

	page.SetFont("Helvetica", "B", ts.base)
	currY, currX = renderCell(
		ax.start,
		currY,
		page.GetStringWidth("Date: "),
		currY,
		ts.sm,
		"Date: ",
		neutral950,
        "L",
	)

	currY, currX = renderCell(
		currX,
		currY,
		page.GetStringWidth(data.InvoiceDate),
		currY,
		ts.sm,
		data.InvoiceDate,
		neutral950,
        "L",
	)

	currY += ms.lg

	page.SetFont("Helvetica", "B", ts.sm)
	page.Line(ax.start, currY, ax.end, currY)
	currY = renderItemsHeader(ax.start, currY, currY, defaultColor)
	page.Line(ax.start, currY, ax.end, currY)

	currY += ms.sm

	currY = renderItems(ax.start, currY, currY, data.InvoiceItems, defaultColor)

	currY += ms.lg

	page.SetFont("Helvetica", "B", ts.lg)
	totStr := "Total: €" + strconv.FormatFloat(data.Total, 'f', -1, 64)

	currY, currX = renderCell(
		ax.getEnd(page.GetStringWidth(totStr)),
		currY,
		page.GetStringWidth(totStr),
		currY,
		ts.lg,
		totStr,
		defaultColor,
        "R",
	)

	currY += ms.xl
	page.SetFont("Helvetica", "", ts.sm)
	currY, currX = renderCell(
		ax.getEnd(page.GetStringWidth("TVA non applicable, art. 293 B du CGI")),
		currY,
		page.GetStringWidth("TVA non applicable, art. 293 B du CGI"),
		currY,
		ts.sm,
		"TVA non applicable, art. 293 B du CGI",
		defaultColor,
        "R",
	)

	currY += ms.xxxl
	page.SetFont("Helvetica", "", ts.lg)
	currY, currX = renderCell(
		ax.getMid(page.GetStringWidth("Merci d'avance!")),
		currY,
		page.GetStringWidth("Merci d'avance!"),
		currY,
		ts.sm,
       "Merci d'avance!", 
		blue500,
        "R",
	)

	// Create file
	err = page.OutputFileAndClose(data.InvoiceID + ".pdf")
	if err != nil {
		panic(err)
	}
}
