package pdf

import (
	"fmt"
	"reflect"

	"github.com/go-pdf/fpdf"
)

type company struct {
	name    string
	address string
	tel     string
	email   string
	compNum string
}

type page struct {
	width  float64
	height float64
	margin float64
}

func New() {

	// SETUP
	pdf := fpdf.New("P", "mm", "A4", "") // 210mm x 297mm
	page := page{
		width:  210.00,
		height: 297.00,
		margin: 10.00,
	}

	pdf.AddPage()

	// SET HEADER
	pdf.SetFont("Arial", "B", 16)
	header := "Identification de l'entrepreneur"

	c := company{
		name:    "Aoife Anne McNamara, EI",
		address: "Bâtiment A1, Résidence Eden Green, 935 Chemin du Golf, 30900 Nîmes, France",
		tel:     "+33 7 89 34 23 41",
		email:   "aoife.amcnamara@gmail.com",
		compNum: "SIRET: 982 544 652 00018",
	}

	// Get the type of the struct
	structType := reflect.TypeOf(c)

	// Loop over the fields of the struct

	companyTable := func() {

		left := page.margin

		pdf.SetX(left)
		tr := pdf.UnicodeTranslatorFromDescriptor("")
		pdf.MultiCell(
			(page.width / 3),
			(12/2)+4,
			tr(header),
			"LTRB",
			"M",
			false,
		)

		next := page.margin + (page.width/3) + page.margin

		pdf.SetXY(next, 0 + page.margin)
		tr = pdf.UnicodeTranslatorFromDescriptor("")
		pdf.MultiCell(
			(page.width / 3),
			(12/2)+4,
			tr(header),
			"LTRB",
			"M",
			false,
		)



		for i := 0; i < structType.NumField(); i++ {
			// Get the field name and value
			fieldName := structType.Field(i).Name
			fieldValue := reflect.ValueOf(c).Field(i).String()

			// Print the field name and value
			fmt.Printf("%s: %s\n", fieldName, fieldValue)

			pdf.SetX(left)
			pdf.SetFont("Helvetica", "", 12)
			pdf.SetCellMargin(2.0)

			tr := pdf.UnicodeTranslatorFromDescriptor("")
			pdf.MultiCell(
				(page.width / 3),
				(12/2)+2,
				tr(fieldValue),
				"LTRB",
				"M",
				false,
			)
		}

	}

	companyTable()

	// Create file
	err := pdf.OutputFileAndClose("output.pdf")
	if err != nil {
		panic(err)
	}
}
