package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/chiboycalix/html-to-pdf/utils"
	"github.com/signintech/gopdf"
)

func Language(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfLanguageSection)
	pdf.Text("Language")
	UnitedKingdom(pdf)
	Nigeria(pdf)
	Germany(pdf)
}

func UnitedKingdom(pdf *gopdf.GoPdf) {
	utils.LoadImage(pdf, "images/languages/britain.png", constants.LeftMargin, StartOfEnglishY, IconWidth, IconHeight)

	pdf.SetXY(StartOfEnglishX, StartOfEnglishY+50)
	pdf.SetFontSize(constants.Paragraph)
	pdf.SetTextColor(204, 204, 204)
	pdf.Text("English")

	pdf.SetXY(EndOfEnglishX, EndOfEnglishY)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("Proficient")
}

func Nigeria(pdf *gopdf.GoPdf) {
	utils.LoadImage(pdf, "images/languages/nigeria.png", constants.LeftMargin, StartOfIgboY, IconWidth, IconHeight)

	pdf.SetXY(StartOfIgboX, StartOfIgboY+50)
	pdf.SetFontSize(constants.Paragraph)
	pdf.SetTextColor(204, 204, 204)
	pdf.Text("Igbo")

	pdf.SetXY(EndOfIgboX, EndOfIgboY)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("Native")
}

func Germany(pdf *gopdf.GoPdf) {
	utils.LoadImage(pdf, "images/languages/germany.png", constants.LeftMargin, StartOfGermanY, IconWidth, IconHeight)

	pdf.SetXY(StartOfGermanX, StartOfGermanY+50)
	pdf.SetFontSize(constants.Paragraph)
	pdf.SetTextColor(204, 204, 204)
	pdf.Text("German")

	pdf.SetXY(EndOfGermanX, EndOfGermanY)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("Beginner")
}
