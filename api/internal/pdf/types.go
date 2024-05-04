package pdf

import (
	"github.com/go-pdf/fpdf"
)

var page *fpdf.Fpdf = fpdf.New("P", "mm", "A4", "") // 210mm x 297mm

type pageParams struct {
	w float64
	h float64
	m float64
}

var pp = pageParams{
	w: 210.00,
	h: 297.00,
	m: 10.00,
}

type defaultText struct {
	h1            string
	h2            string
	companyHeader string
	clientHeader  string
	paymentHeader string
	bic           string
	iban          string
	bank          string
	accName       string
	invoiceHeader string
	dateHeader    string
	itemsHeader   []string
	total         string
	tax           string
	message       string
}

var dtfr = defaultText{
	h1:            "Amagine Media",
	h2:            "FACTURE",
	companyHeader: "Identification De l'Entrepreneur",
	clientHeader:  "Identification Du Client",
	paymentHeader: "Détails De Paiement",
	bic:           "BIC: ",
	iban:          "IBAN: ",
	bank:          "Banque: ",
	accName:       "Titulaire Du Compte: ",
	invoiceHeader: "Numéro Facture:",
	dateHeader:    "Date: ",
	itemsHeader:   []string{"Désignation Des Prestations", "Qty/Hrs", "€/Heure", "€/Unité", "Total"},
	total:         "Total: €",
	tax:           "TVA non applicable, art. 293 B du CGI",
	message:       "Merci d'avance!",
}

var dten = defaultText{
	h1:            "Amagine Media",
	h2:            "INVOICE",
	companyHeader: "Contractor Details",
	clientHeader:  "Client Details",
	paymentHeader: "Payment Details",
	bic:           "BIC: ",
	iban:          "IBAN: ",
	bank:          "Bank: ",
    accName:       "Account Holder: ",
	invoiceHeader: "Invoice No. ",
	dateHeader:    "Date: ",
	itemsHeader:   []string{"Item Details", "Qty/Hrs", "€/hour", "€/Unit", "Total"},
	total:         "Total: €",
	tax:           "TVA not applicable, art. 293 B du CGI",
	message:       "Thank you in advance!",
}

type boxWidth struct {
	w1_12  float64
	w2_12  float64
	w3_12  float64
	w4_12  float64
	w5_12  float64
	w6_12  float64
	w7_12  float64
	w8_12  float64
	w9_12  float64
	w10_12 float64
	w11_12 float64
	w12_12 float64
}

var bw = boxWidth{
	// page width - (page margin * 2 (for either side)) / fractional unit * column num)
	w1_12:  ((pp.w - (pp.m * 2)) / 12),
	w2_12:  ((pp.w - (pp.m * 2)) / 12) * 2,
	w3_12:  ((pp.w - (pp.m * 2)) / 12) * 3,
	w4_12:  ((pp.w - (pp.m * 2)) / 12) * 4,
	w5_12:  ((pp.w - (pp.m * 2)) / 12) * 5,
	w6_12:  ((pp.w - (pp.m * 2)) / 12) * 6,
	w7_12:  ((pp.w - (pp.m * 2)) / 12) * 7,
	w8_12:  ((pp.w - (pp.m * 2)) / 12) * 8,
	w9_12:  ((pp.w - (pp.m * 2)) / 12) * 9,
	w10_12: ((pp.w - (pp.m * 2)) / 12) * 10,
	w11_12: ((pp.w - (pp.m * 2)) / 12) * 11,
	w12_12: ((pp.w - (pp.m * 2)) / 12) * 12,
}

type textSizing struct {
	xs   float64
	sm   float64
	base float64
	md   float64
	lg   float64
	xl   float64
	xxl  float64
}

var ts = textSizing{
	xs:   8,
	sm:   10,
	base: 12,
	md:   16,
	lg:   24,
	xl:   32,
	xxl:  46,
}

type marginSizing struct {
	xs   float64
	sm   float64
	md   float64
	lg   float64
	xl   float64
	xxl  float64
	xxxl float64
}

var ms = marginSizing{
	xs:   2,
	sm:   4,
	md:   6,
	lg:   10,
	xl:   12,
	xxl:  14,
	xxxl: 32,
}

func (ms *marginSizing) botMargin(currY *float64, m float64) {
	*currY += m
}

type alignX struct {
	start  float64
	end    float64
	middle float64
}

var ax = alignX{
	start:  pp.m,
	end:    pp.w - pp.m,
	middle: (pp.w / 2),
}

func (a *alignX) getEnd(bw float64) float64 {
	return a.end - bw
}

func (a *alignX) getMid(bw float64) float64 {
	return a.middle - (bw / 2)
}

type rgb struct {
	red   int
	green int
	blue  int
}

var (
	defaultColor = rgb{red: 38, green: 38, blue: 38}
	neutral500   = rgb{red: 82, green: 82, blue: 82}
	neutral950   = rgb{red: 10, green: 10, blue: 10}
	blue500      = rgb{red: 59, green: 130, blue: 246}
)
