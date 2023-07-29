package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

func Education(pdf *gopdf.GoPdf) {
	pdf.SetXY(MainContentMarginLeft, StartOfEducationSection)
	pdf.SetFontSize(constants.HeaderTwo)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("Education")

	AndelaBootcamp(pdf)
	FUTMinna(pdf)
}

func AndelaBootcamp(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft, StartOfEducationSection*1.02, MainContentMarginLeft*2.25, StartOfEducationSection*1.2, "DF", 50, 10)
}

func FUTMinna(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft*2.3, StartOfEducationSection*1.02, MainContentMarginLeft*3.55, StartOfEducationSection*1.2, "DF", 50, 10)
}
