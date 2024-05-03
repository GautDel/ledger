package pdf

import (
	"strconv"

	"ledgerbolt.systems/internal/models"
)

func newEn(data models.Invoice, bank models.Bank) {
	initParams()

	var currY, currX float64
	y := &currY

	// Render Comp Name
	page.SetFont("Helvetica", "B", ts.lg)
	currY = renderPageHeader(pp.m, pp.m, currY, dten.h1, ts.lg, blue500)

	// Render "Facture"
	page.SetFont("Helvetica", "B", ts.xxl)
	currY = renderPageHeader(
		ax.getEnd(page.GetStringWidth(dten.h2)),
		pp.m,
		currY,
		dten.h2,
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
		dten.companyHeader,
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
		dten.clientHeader,
		[]string{
			data.ClientName,
			data.ClientAddress,
			data.ClientPhone,
			data.ClientEmail,
		},
		neutral950,
	)

	ms.botMargin(y, ms.md)
	page.SetFont("Helvetica", "B", ts.base)
	currY, _ = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dten.paymentHeader),
		currY,
		ts.sm,
		dten.paymentHeader,
		neutral950,
		"L",
		true,
	)

	page.SetFont("Helvetica", "", ts.sm)
	currY, _ = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dten.bic+bank.BIC),
		currY,
		ts.sm,
		dten.bic+bank.BIC,
		neutral950,
		"L",
		true,
	)

	currY, _ = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dten.iban+bank.IBAN),
		currY,
		ts.sm,
		dten.iban+bank.IBAN,
		neutral950,
		"L",
		true,
	)

	currY, _ = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dten.bank+bank.BankName),
		currY,
		ts.sm,
		dten.bank+bank.BankName,
		neutral950,
		"L",
		true,
	)

	currY, _ = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dten.accName+bank.AccountName),
		currY,
		ts.sm,
		dten.accName+bank.AccountName,
		neutral950,
		"L",
		true,
	)

	ms.botMargin(y, ms.md)

	page.SetFont("Helvetica", "B", ts.base)
	currY, currX = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dten.invoiceHeader),
		currY,
		ts.sm,
		dten.invoiceHeader,
		neutral950,
		"L",
		false,
	)

	currY, _ = renderCell(
		currX,
		currY,
		page.GetStringWidth(data.InvoiceID),
		currY,
		ts.sm,
		data.InvoiceID,
		blue500,
		"L",
		true,
	)

	page.SetFont("Helvetica", "B", ts.base)
	currY, currX = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dten.dateHeader),
		currY,
		ts.sm,
		dten.dateHeader,
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
	currY = renderItemsHeader(ax.start, currY, currY, defaultColor, dten.itemsHeader)
	page.Line(ax.start, currY, ax.end, currY)

	ms.botMargin(y, ms.sm)

	currY = renderItems(ax.start, currY, currY, data.InvoiceItems, defaultColor)

	ms.botMargin(y, ms.sm)

	page.SetFont("Helvetica", "B", ts.lg)
	totStr := dten.total + strconv.FormatFloat(data.Total, 'f', -1, 64)

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
		ax.getEnd(page.GetStringWidth(dten.tax)),
		currY,
		page.GetStringWidth(dten.tax),
		currY,
		ts.sm,
		dten.tax,
		defaultColor,
		"R",
		true,
	)

	ms.botMargin(y, ms.xl)

	page.SetFont("Helvetica", "I", ts.md)
	currY, _ = renderCell(
		ax.getMid(page.GetStringWidth(dten.message)),
		currY,
		page.GetStringWidth(dten.message),
		currY,
		ts.sm,
		dten.message,
		blue500,
		"R",
		true,
	)
}
