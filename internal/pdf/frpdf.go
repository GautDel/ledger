package pdf

import (
	"strconv"

	"ledgerbolt.systems/internal/models"
)

func newFr(data models.Invoice, bank models.Bank) {
	initParams()

	var currY, currX float64
	y := &currY

	// Render Comp Name
	page.SetFont("Helvetica", "B", ts.lg)
	currY = renderPageHeader(pp.m, pp.m, currY, dtfr.h1, ts.lg, blue500)

	// Render "Facture"
	page.SetFont("Helvetica", "B", ts.xxl)
	currY = renderPageHeader(
		ax.getEnd(page.GetStringWidth(dtfr.h2)),
		pp.m,
		currY,
		dtfr.h2,
		ts.xxl,
		blue500,
	)

	ms.botMargin(y, ms.sm)

	_ = renderBox(
		ax.start,
		currY,
		bw.w5_12,
		currY,
		10,
		dtfr.companyHeader,
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
		dtfr.clientHeader,
		[]string{
			data.ClientName,
			data.ClientAddress,
			data.ClientPhone,
			data.ClientEmail,
		},
		neutral950,
	)

	ms.botMargin(y, ms.sm)
	page.SetFont("Helvetica", "B", ts.base)
	currY, _ = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dtfr.paymentHeader),
		currY,
		ts.sm,
		dtfr.paymentHeader,
		neutral950,
		"L",
		true,
	)
    
	page.SetFont("Helvetica", "", ts.sm)
	currY, _ = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dtfr.bic+bank.BIC),
		currY,
		ts.sm,
		dtfr.bic+bank.BIC,
		neutral950,
		"L",
		true,
	)


	currY, _ = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dtfr.iban+bank.IBAN),
		currY,
		ts.sm,
		dtfr.iban+bank.IBAN,
		neutral950,
		"L",
		true,
	)

	currY, _ = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dtfr.bank+bank.BankName),
		currY,
		ts.sm,
		dtfr.bank+bank.BankName,
		neutral950,
		"L",
		true,
	)

	currY, _ = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dtfr.accName+bank.AccountName),
		currY,
		ts.sm,
		dtfr.accName+bank.AccountName,
		neutral950,
		"L",
		true,
	)

	ms.botMargin(y, ms.md)

	page.SetFont("Helvetica", "B", ts.base)
	currY, currX = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dtfr.invoiceHeader),
		currY,
		ts.base,
		dtfr.invoiceHeader,
		neutral950,
		"L",
		false,
	)

	currY, _ = renderCell(
		currX - 2,
		currY,
		page.GetStringWidth(data.InvoiceID),
		currY,
		ts.base,
		data.InvoiceID,
		blue500,
		"L",
		true,
	)

	page.SetFont("Helvetica", "B", ts.base)
	currY, currX = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dtfr.dateHeader),
		currY,
		ts.sm,
		dtfr.dateHeader,
		neutral950,
		"L",
		false,
	)

	currY, _ = renderCell(
		currX,
		currY,
		page.GetStringWidth(data.InvoiceDate),
		currY,
		ts.sm,
		data.InvoiceDate,
		neutral950,
		"L",
		true,
	)

	ms.botMargin(y, ms.sm)

	page.SetFont("Helvetica", "B", ts.sm)
	page.Line(ax.start, currY, ax.end, currY)
	currY = renderItemsHeader(ax.start, currY, currY, defaultColor, dtfr.itemsHeader)
	page.Line(ax.start, currY, ax.end, currY)

	ms.botMargin(y, ms.sm)

	currY = renderItems(ax.start, currY, currY, data.InvoiceItems, defaultColor)

	ms.botMargin(y, ms.sm)

	page.SetFont("Helvetica", "B", ts.lg)
	totStr := dtfr.total + strconv.FormatFloat(data.Total, 'f', -1, 64)

	currY, _ = renderCell(
		ax.getEnd(page.GetStringWidth(totStr)),
		currY,
		page.GetStringWidth(totStr),
		currY,
		ts.lg,
		totStr,
		defaultColor,
		"R",
		true,
	)

	page.SetFont("Helvetica", "", ts.sm)
	currY, _ = renderCell(
		ax.getEnd(page.GetStringWidth(dtfr.tax)),
		currY,
		page.GetStringWidth(dtfr.tax),
		currY,
		ts.sm,
		dtfr.tax,
		defaultColor,
		"R",
		true,
	)

	ms.botMargin(y, ms.xl)

	page.SetFont("Helvetica", "I", ts.md)
	currY, _ = renderCell(
		ax.getMid(page.GetStringWidth(dtfr.message)),
		currY,
		page.GetStringWidth(dtfr.message),
		currY,
		ts.sm,
		dtfr.message,
		blue500,
		"R",
		true,
	)
}
