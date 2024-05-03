package pdf

import (
	"strconv"
	"strings"

	"ledgerbolt.systems/internal/models"
	"ledgerbolt.systems/utils"
)

func checkY(prev, curr float64) float64 {
	if curr > prev {
		return curr
	}

	return prev
}

func renderPageHeader(
	x, y, prevY float64,
	val string,
	size float64,
	color rgb,
) float64 {

	page.SetTextColor(color.red, color.green, color.blue)
	page.SetXY(x, y)

	titleLen := page.GetStringWidth(val)
	h := size / 2

	tr := page.UnicodeTranslatorFromDescriptor("")
	page.CellFormat(
		titleLen,
		h,
		tr(val),
		"",
		0,
		"R",
		false,
		0,
		"",
	)

	page.Ln(-1)
	page.SetTextColor(defaultColor.red, defaultColor.green, defaultColor.blue)

	return checkY(prevY, page.GetY())
}

func renderBoxHeader(s string, x float64, w float64) {
	page.SetX(x)
	page.SetFont("Helvetica", "B", 12)
	page.SetTextColor(defaultColor.red, defaultColor.green, defaultColor.blue)
	tr := page.UnicodeTranslatorFromDescriptor("")
	page.MultiCell(
		w,
		(12/2)+2,
		tr(s),
		"",
		"M",
		false,
	)
}

func renderBox(
	x, y, w, prevY, size float64,
	s string,
	vals []string,
	color rgb,
) float64 {

	page.SetXY(x, y)
	renderBoxHeader(s, x, w)
	var boxY float64

	for i, val := range vals {
		page.SetX(x)
		page.SetFont("Helvetica", "", size)
		page.SetTextColor(color.red, color.green, color.blue)

		tr := page.UnicodeTranslatorFromDescriptor("")
		page.MultiCell(
			w,
			(size / 2),
			tr(val),
			"",
			"LT",
			false,
		)

		if i == len(vals)-1 {
			boxY = checkY(prevY, page.GetY())
		}
	}

	return boxY
}

func renderCell(
	x, y, w, prevY, size float64,
	item string,
	color rgb,
	a string,
	nextLn bool,
) (float64, float64) {

	page.SetXY(x, y)
	page.SetTextColor(color.red, color.green, color.blue)
	tr := page.UnicodeTranslatorFromDescriptor("")
	h := size / 2
	page.CellFormat(
		w,
		h,
		tr(item),
		"",
		0,
		a,
		false,
		0,
		"",
	)

    if nextLn {
        page.Ln(-1)
    }

	return checkY(prevY, page.GetY()), page.GetX()
}

func renderHeaderItem(w float64, s, a string, size float64) {
	h := size / 2
	tr := page.UnicodeTranslatorFromDescriptor("")
	page.CellFormat(
		w,
		h,
		tr(s),
		"",
		0,
		a,
		false,
		0,
		"",
	)
}

func renderItemsHeader(x, y, prevY float64, color rgb, items []string) float64 {

	page.SetXY(x, y)
	page.SetTextColor(color.red, color.green, color.blue)

	renderHeaderItem(bw.w7_12, items[0], "L", ts.md)
	renderHeaderItem(bw.w1_12, items[1], "C", ts.md)
	renderHeaderItem(bw.w1_12, items[2], "C", ts.md)
	renderHeaderItem(bw.w1_12, items[3], "C", ts.md)
	renderHeaderItem(bw.w2_12, items[4], "R", ts.md)

	page.Ln(-1)

	return checkY(prevY, page.GetY())
}

func renderItems(x, y, prevY float64, items []models.InvoiceItem, color rgb) float64 {

	page.SetXY(x, y)

	for _, i := range items {
		page.SetFont("Helvetica", "B", ts.sm)
		renderCellInline(bw.w7_12, i.Name, "L")

		page.SetFont("Helvetica", "", ts.sm)
		renderCellInline(bw.w1_12, strconv.Itoa(i.Qty), "C")
		if i.HourlyPrice == 0 {

			renderCellInline(bw.w1_12, "", "C")
		} else {
			renderCellInline(bw.w1_12, "€ "+strconv.FormatFloat(i.HourlyPrice, 'f', -1, 64), "C")
		}

		if i.UnitPrice == 0 {
			renderCellInline(bw.w1_12, "", "C")
		} else {
			renderCellInline(bw.w1_12, "€ "+strconv.FormatFloat(i.UnitPrice, 'f', -1, 64), "C")

		}
		page.SetFont("Helvetica", "B", ts.sm)
		renderCellInline(bw.w2_12, "€ "+strconv.FormatFloat(i.TotalPrice, 'f', -1, 64), "R")
		page.Ln(-1)

		descSlice := utils.StringSplit(i.Description, "\n")
		for _, descItem := range descSlice {
			renderCellBlock(bw.w12_12, strings.TrimSpace(descItem))
		}

		page.Ln(-1)
		page.Line(ax.start, page.GetY(), ax.end, page.GetY())
		page.Ln(-1)
	}

	return checkY(prevY, page.GetY())
}

func renderCellInline(w float64, item, a string) {

	tr := page.UnicodeTranslatorFromDescriptor("")
	h := ts.sm / 2
	page.CellFormat(
		w,
		h,
		tr(item),
		"",
		0,
		a,
		false,
		0,
		"",
	)
}

func renderCellBlock(w float64, item string) {

	page.SetFont("Helvetica", "", ts.xs)
	tr := page.UnicodeTranslatorFromDescriptor("")
	h := ts.xs / 2
	page.MultiCell(
		w,
		h,
		tr(item),
		"",
		"L",
		false,
	)
}
