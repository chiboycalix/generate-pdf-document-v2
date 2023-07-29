package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

func Skills(pdf *gopdf.GoPdf) {
	pdf.SetXY(MainContentMarginLeft, StartOfSkillSection)
	pdf.SetFontSize(constants.HeaderTwo)
	pdf.SetTextColor(uint8(DarkModePrimaryTextColor.R), uint8(DarkModePrimaryTextColor.G), uint8(DarkModePrimaryTextColor.B))
	pdf.Text("Skills")

	Javascript(pdf)
	TypeScript(pdf)
	NodeJs(pdf)
	ReactJs(pdf)
	Redux(pdf)
	NextJs(pdf)
	Tailwind(pdf)
	Golang(pdf)
}

func Javascript(pdf *gopdf.GoPdf) {
	// pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.SetFillColor(uint8(DarkModePageBackgroundColor.R), uint8(DarkModePageBackgroundColor.G), uint8(DarkModePageBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft, StartOfSkillSection*1.04, MainContentMarginLeft*1.5, StartOfSkillSection*1.1, "DF", 50, 10)

	pdf.SetXY(MainContentMarginLeft*1.12, StartOfSkillSection*1.08)
	pdf.Text("Javascript.")
}
func TypeScript(pdf *gopdf.GoPdf) {
	// pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.SetFillColor(uint8(DarkModePageBackgroundColor.R), uint8(DarkModePageBackgroundColor.G), uint8(DarkModePageBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft*1.65, StartOfSkillSection*1.04, MainContentMarginLeft*2.15, StartOfSkillSection*1.1, "DF", 50, 10)

	pdf.SetXY(MainContentMarginLeft*1.77, StartOfSkillSection*1.08)
	pdf.Text("Typescript.")
}
func NodeJs(pdf *gopdf.GoPdf) {
	// pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.SetFillColor(uint8(DarkModePageBackgroundColor.R), uint8(DarkModePageBackgroundColor.G), uint8(DarkModePageBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft*2.35, StartOfSkillSection*1.04, MainContentMarginLeft*2.85, StartOfSkillSection*1.1, "DF", 50, 10)

	pdf.SetXY(MainContentMarginLeft*2.50, StartOfSkillSection*1.08)
	pdf.Text("NodeJs.")
}
func ReactJs(pdf *gopdf.GoPdf) {
	// pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.SetFillColor(uint8(DarkModePageBackgroundColor.R), uint8(DarkModePageBackgroundColor.G), uint8(DarkModePageBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft*3.05, StartOfSkillSection*1.04, MainContentMarginLeft*3.55, StartOfSkillSection*1.1, "DF", 50, 10)

	pdf.SetXY(MainContentMarginLeft*3.20, StartOfSkillSection*1.08)
	pdf.Text("ReactJs.")
}
func Redux(pdf *gopdf.GoPdf) {
	// pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.SetFillColor(uint8(DarkModePageBackgroundColor.R), uint8(DarkModePageBackgroundColor.G), uint8(DarkModePageBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft, StartOfSkillSection*1.13, MainContentMarginLeft*1.5, StartOfSkillSection*1.19, "DF", 50, 10)

	pdf.SetXY(MainContentMarginLeft*1.15, StartOfSkillSection*1.17)
	pdf.Text("Redux.")
}
func NextJs(pdf *gopdf.GoPdf) {
	// pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.SetFillColor(uint8(DarkModePageBackgroundColor.R), uint8(DarkModePageBackgroundColor.G), uint8(DarkModePageBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft*1.65, StartOfSkillSection*1.13, MainContentMarginLeft*2.15, StartOfSkillSection*1.19, "DF", 50, 10)

	pdf.SetXY(MainContentMarginLeft*1.80, StartOfSkillSection*1.17)
	pdf.Text("NextJs.")
}
func Tailwind(pdf *gopdf.GoPdf) {
	// pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.SetFillColor(uint8(DarkModePageBackgroundColor.R), uint8(DarkModePageBackgroundColor.G), uint8(DarkModePageBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft*2.35, StartOfSkillSection*1.13, MainContentMarginLeft*2.85, StartOfSkillSection*1.19, "DF", 50, 10)

	pdf.SetXY(MainContentMarginLeft*2.50, StartOfSkillSection*1.17)
	pdf.Text("Tailwind.")
}
func Golang(pdf *gopdf.GoPdf) {
	// pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.SetFillColor(uint8(DarkModePageBackgroundColor.R), uint8(DarkModePageBackgroundColor.G), uint8(DarkModePageBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft*3.05, StartOfSkillSection*1.13, MainContentMarginLeft*3.55, StartOfSkillSection*1.19, "DF", 50, 10)

	pdf.SetXY(MainContentMarginLeft*3.20, StartOfSkillSection*1.17)
	pdf.Text("Golang.")
}
