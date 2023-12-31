package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

func Education(pdf *gopdf.GoPdf) {
	pdf.SetXY(MainContentMarginLeft, StartOfEducationSection)
	pdf.SetFontSize(constants.HeaderTwo)
	// pdf.SetTextColor(uint8(DarkModePrimaryTextColor.R), uint8(DarkModePrimaryTextColor.G), uint8(DarkModePrimaryTextColor.B))
	pdf.SetTextColor(uint8(LightModePrimaryTextColor.R), uint8(LightModePrimaryTextColor.G), uint8(LightModePrimaryTextColor.B))
	pdf.Text("Education")

	AndelaBootcamp(pdf)
	FUTMinna(pdf)
}

func AndelaBootcamp(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(uint8(LightModeBoxBackgroundColor.R), uint8(LightModeBoxBackgroundColor.G), uint8(LightModeBoxBackgroundColor.B))
	// pdf.SetFillColor(uint8(DarkModePageBackgroundColor.R), uint8(DarkModePageBackgroundColor.G), uint8(DarkModePageBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft, StartOfEducationSection*1.02, MainContentMarginLeft*2.25, StartOfEducationSection*1.2, "DF", 50, 10)
}

func FUTMinna(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(uint8(LightModeBoxBackgroundColor.R), uint8(LightModeBoxBackgroundColor.G), uint8(LightModeBoxBackgroundColor.B))
	// pdf.SetFillColor(uint8(DarkModePageBackgroundColor.R), uint8(DarkModePageBackgroundColor.G), uint8(DarkModePageBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft*2.3, StartOfEducationSection*1.02, MainContentMarginLeft*3.55, StartOfEducationSection*1.2, "DF", 50, 10)
}
