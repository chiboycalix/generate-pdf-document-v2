package main

import (
	"fmt"

	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/chiboycalix/html-to-pdf/content"
	"github.com/chiboycalix/html-to-pdf/utils"
	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{
		W: constants.PageWidth,
		H: constants.PageHeight,
	}})
	pdf.AddPage()
	utils.LoadFonts(&pdf)
	DarkModePageBackgroundColor := constants.DarkModePageBackgroundColor()
	LightModePageBackgroundColor := constants.LightModePageBackgroundColor()
	fmt.Println(DarkModePageBackgroundColor)
	fmt.Println(LightModePageBackgroundColor)
	// pdf.SetFillColor(35, 35, 57)
	pdf.SetFillColor(uint8(DarkModePageBackgroundColor.R), uint8(DarkModePageBackgroundColor.G), uint8(DarkModePageBackgroundColor.B))
	pdf.SetStrokeColor(255, 255, 255)
	pdf.Rectangle(0, 0, 8000, 8000, "DF", 0, 10)

	content.LeftContent(&pdf)
	content.MainContent(&pdf)

	pdf.SetStrokeColor(204, 204, 204)
	pdf.SetLineWidth(1)
	pdf.Line(constants.LeftContentWidth, constants.TopMargin, constants.LeftContentWidth, constants.VerticalLineHeight)

	pdf.WritePdf("output/resume.pdf")

}
