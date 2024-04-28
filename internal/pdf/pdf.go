package pdf

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-pdf/fpdf"
)

func strDelimit(str string, sepstr string, sepcount int) string {
	pos := len(str) - sepcount
	for pos > 0 {
		str = str[:pos] + sepstr + str[pos:]
		pos = pos - sepcount
	}
	return str
}

func New() {

	pdf := fpdf.New("P", "mm", "A4", "") // 210mm x 297mm
	pdf.AddPage()

	// Create dynamic cell
	//	text := "Name:"
	//	textWidth := pdf.GetStringWidth(text)
	//	log.Println(textWidth)
	//	pdf.Cell(textWidth, 4, text)
	//
	//	// Create second dynamic cell
	//	pdf.SetFont("Arial", "", 16)
	//	text = "Gauthier"
	//	textWidth = pdf.GetStringWidth(text)
	//	log.Println(textWidth)
	//	pdf.Cell(textWidth, 4, text)

	pdf.SetFont("Arial", "B", 12)
	type countryType struct {
		nameStr, capitalStr, areaStr, popStr string
	}


	countryList := make([]countryType, 0, 8)

	header := []string{"Country", "Capital", "Area (sq km)", "Pop. (thousands)"}

	loadData := func(fileStr string) {
		fl, err := os.Open(fileStr)
		if err == nil {
			scanner := bufio.NewScanner(fl)
			var c countryType
			for scanner.Scan() {
				// Austria;Vienna;83859;8075
				lineStr := scanner.Text()
				list := strings.Split(lineStr, ";")
				if len(list) == 4 {
					c.nameStr = list[0]
					c.capitalStr = list[1]
					c.areaStr = list[2]
					c.popStr = list[3]
					countryList = append(countryList, c)
				} else {
					err = fmt.Errorf("error tokenizing %s", lineStr)
				}
			}
			fl.Close()
			if len(countryList) == 0 {
				err = fmt.Errorf("error loading data from %s", fileStr)
			}
		}
		if err != nil {
			pdf.SetError(err)
		}
	}

	// Simple table
	basicTable := func() {
		// Page width - num cells * cell width / 2 -> 2 = center
		left := (210.0 - 4*40) / 2
		pdf.SetX(left)
		for _, str := range header {
			pdf.CellFormat(40, 7, str, "1", 0, "", false, 0, "")
		}

		pdf.Ln(-1)

		for _, c := range countryList {
			pdf.SetX(left)
			pdf.CellFormat(40, 6, c.nameStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.capitalStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.areaStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.popStr, "1", 0, "", false, 0, "")
			pdf.Ln(-1)
		}
	}

	// Better table
	improvedTable := func() {
		// Column widths
		w := []float64{40.0, 35.0, 40.0, 45.0}
		wSum := 0.0
		for _, v := range w {
			wSum += v
		}
		left := (210 - wSum) / 2
		// 	Header
		pdf.SetX(left)
		for j, str := range header {
			pdf.CellFormat(w[j], 7, str, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)
		// Data
		for _, c := range countryList {
			pdf.SetX(left)
			pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", false, 0, "")
			pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", false, 0, "")
			pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3),
				"LR", 0, "R", false, 0, "")
			pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3),
				"LR", 0, "R", false, 0, "")
			pdf.Ln(-1)
		}
		pdf.SetX(left)
		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	}

	loadData("examplepdf.txt")
	basicTable()
	improvedTable()

	// Create file
	err := pdf.OutputFileAndClose("output.pdf")
	if err != nil {
		panic(err)
	}
}
