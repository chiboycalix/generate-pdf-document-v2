package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

const ProjectOneX0 = 196.6 * 2
const ProjectOneX1 = 336.8 * 2
const StartOfProject = StartOfLatestProjectsSection + constants.VerticalSpace

func LatestProjects(pdf *gopdf.GoPdf) {

	pdf.SetXY(MainContentMarginLeft, StartOfLatestProjectsSection)
	pdf.SetFontSize(constants.HeaderTwo)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("Latest Projects")

	pdf.SetFillColor(225, 231, 254)
	pdf.Rectangle(MainContentMarginLeft, StartOfProject, MainContentMarginLeft*2.25, StartOfProject*1.5, "DF", 50, 10)

	pdf.SetFillColor(225, 231, 254)
	pdf.Rectangle(MainContentMarginLeft*2.3, StartOfProject, MainContentMarginLeft*3.55, StartOfProject*1.5, "DF", 50, 10)

}
