package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

func MainContent(pdf *gopdf.GoPdf) {
	Experience(pdf)
}

func Experience(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.FirstColumnWidth+constants.LeftMargin, StartOfProfilePicture+constants.Hundred)
	pdf.SetFontSize(60)
	pdf.Text("Experience")

	LoadIcons(pdf, "images/logos/nibbs_logo.png", constants.FirstColumnWidth+constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*5.5), 280, 70)
	LoadIcons(pdf, "images/logos/sclogo.png", constants.FirstColumnWidth+constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*6.5), 600, 80)
	LoadIcons(pdf, "images/logos/bongalow.png", constants.FirstColumnWidth+constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*7.5), 600, 80)

	pdf.SetLineWidth(50)
	pdf.SetStrokeColor(255, 0, 0)
	pdf.SetFillColor(255, 0, 0)
	// pdf.Oval(constants.FirstColumnWidth+constants.LeftMargin, 900, 900, 950)
	pdf.Oval(constants.FirstColumnWidth+constants.LeftMargin, 250, constants.FirstColumnWidth+constants.LeftMargin+250, 300)
}
