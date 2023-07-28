package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

func Language(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*9.5))
	pdf.Text("Language")
	UnitedKingdom(pdf)
	Nigeria(pdf)
	Germany(pdf)
}

func UnitedKingdom(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*10))
	LoadIcons(pdf, "images/languages/britain.png", constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*10), 80, 70)
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsFirstValueRowHeight+10*constants.HundredAndFifty)
	pdf.SetFontSize(35)
	pdf.SetTextColor(204, 204, 204)

	pdf.Text("English")

	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsSecondValueRowHeight+10*constants.HundredAndFifty)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("Proficient")
}

func Nigeria(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*11))
	LoadIcons(pdf, "images/languages/nigeria.png", constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*11), 80, 70)
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsFirstValueRowHeight+11*constants.HundredAndFifty)
	pdf.SetFontSize(35)
	pdf.SetTextColor(204, 204, 204)

	pdf.Text("Igbo")

	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsSecondValueRowHeight+11*constants.HundredAndFifty)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("Native")
}

func Germany(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*12))
	LoadIcons(pdf, "images/languages/germany.png", constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*12), 80, 70)
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsFirstValueRowHeight+12*constants.HundredAndFifty)
	pdf.SetFontSize(35)
	pdf.SetTextColor(204, 204, 204)

	pdf.Text("German")

	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsSecondValueRowHeight+12*constants.HundredAndFifty)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("Beginner")
}
