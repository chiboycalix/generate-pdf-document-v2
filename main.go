package main

import (
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

	content.LeftContent(&pdf)
	content.MainContent(&pdf)

	pdf.SetStrokeColor(204, 204, 204)
	pdf.SetLineWidth(1)
	pdf.Line(constants.FirstColumnWidth, constants.TopMargin, constants.FirstColumnWidth, constants.VerticalLineHeight)

	pdf.WritePdf("output/resume.pdf")

}
