package pdf

import (
	"strconv"

	"ledgerbolt.systems/internal/models"
)

func newFr(data models.Invoice) {
	initParams()

	var currY, currX float64
	y := &currY

	// Render Comp Name
	page.SetFont("Helvetica", "B", ts.lg)
	currY = renderPageHeader(pp.m, pp.m, currY, dtfr.h1, ts.lg, blue500)

	// Render "Facture"
	page.SetFont("Helvetica", "B", ts.xxl)
	currY = renderPageHeader(
		ax.getEnd(page.GetStringWidth(pp.h2)),
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
			data.ClientEmail,
			data.ClientAddress,
			data.ClientPhone,
		},
		neutral950,
	)

	ms.botMargin(y, ms.md)

	page.SetFont("Helvetica", "B", ts.base)
	currY, currX = renderCell(
		ax.start,
		currY,
		page.GetStringWidth(dtfr.invoiceHeader),
		currY,
		ts.sm,
		dtfr.invoiceHeader,
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

	ms.botMargin(y, ms.sm)

	page.SetFont("Helvetica", "B", ts.sm)
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
