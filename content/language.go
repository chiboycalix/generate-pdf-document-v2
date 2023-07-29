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
	pdf.SetTextColor(uint8(DarkModeSecondaryTextColor.R), uint8(DarkModeSecondaryTextColor.G), uint8(DarkModeSecondaryTextColor.B))
	pdf.Text("English")

	pdf.SetXY(EndOfEnglishX, EndOfEnglishY)
	pdf.SetTextColor(uint8(LightModePrimaryTextColor.R), uint8(LightModePrimaryTextColor.G), uint8(LightModePrimaryTextColor.B))
	// pdf.SetTextColor(uint8(DarkModePrimaryTextColor.R), uint8(DarkModePrimaryTextColor.G), uint8(DarkModePrimaryTextColor.B))
	pdf.Text("Proficient")
}

func Nigeria(pdf *gopdf.GoPdf) {
	utils.LoadImage(pdf, "images/languages/nigeria.png", constants.LeftMargin, StartOfIgboY, IconWidth, IconHeight)

	pdf.SetXY(StartOfIgboX, StartOfIgboY+50)
	pdf.SetFontSize(constants.Paragraph)
	pdf.SetTextColor(uint8(DarkModeSecondaryTextColor.R), uint8(DarkModeSecondaryTextColor.G), uint8(DarkModeSecondaryTextColor.B))
	pdf.Text("Igbo")

	pdf.SetXY(EndOfIgboX, EndOfIgboY)
	pdf.SetTextColor(uint8(LightModePrimaryTextColor.R), uint8(LightModePrimaryTextColor.G), uint8(LightModePrimaryTextColor.B))
	// pdf.SetTextColor(uint8(DarkModePrimaryTextColor.R), uint8(DarkModePrimaryTextColor.G), uint8(DarkModePrimaryTextColor.B))
	pdf.Text("Native")
}

func Germany(pdf *gopdf.GoPdf) {
	utils.LoadImage(pdf, "images/languages/germany.png", constants.LeftMargin, StartOfGermanY, IconWidth, IconHeight)

	pdf.SetXY(StartOfGermanX, StartOfGermanY+50)
	pdf.SetFontSize(constants.Paragraph)
	pdf.SetTextColor(uint8(DarkModeSecondaryTextColor.R), uint8(DarkModeSecondaryTextColor.G), uint8(DarkModeSecondaryTextColor.B))
	pdf.Text("German")

	pdf.SetXY(EndOfGermanX, EndOfGermanY)
	pdf.SetTextColor(uint8(LightModePrimaryTextColor.R), uint8(LightModePrimaryTextColor.G), uint8(LightModePrimaryTextColor.B))
	// pdf.SetTextColor(uint8(DarkModePrimaryTextColor.R), uint8(DarkModePrimaryTextColor.G), uint8(DarkModePrimaryTextColor.B))
	pdf.Text("Beginner")
}
