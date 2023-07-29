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
	pdf.SetTextColor(uint8(DarkModePrimaryTextColor.R), uint8(DarkModePrimaryTextColor.G), uint8(DarkModePrimaryTextColor.B))
	pdf.Text("Latest Projects")

	// pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.SetFillColor(uint8(DarkModeBoxBackgroundColor.R), uint8(DarkModeBoxBackgroundColor.G), uint8(DarkModeBoxBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft, StartOfProject, MainContentMarginLeft*2.25, StartOfProject*1.5, "DF", 50, 10)

	// pdf.SetFillColor(uint8(BackgroundColor.R), uint8(BackgroundColor.G), uint8(BackgroundColor.B))
	pdf.SetFillColor(uint8(DarkModeBoxBackgroundColor.R), uint8(DarkModeBoxBackgroundColor.G), uint8(DarkModeBoxBackgroundColor.B))
	pdf.Rectangle(MainContentMarginLeft*2.3, StartOfProject, MainContentMarginLeft*3.55, StartOfProject*1.5, "DF", 50, 10)

}
