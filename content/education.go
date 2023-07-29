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
}
