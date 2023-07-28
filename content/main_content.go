package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

const StartOfExperienceSection = constants.TopMargin * 1.5
const MainContentMarginLeft = constants.LeftContentWidth * 1.1

const StartOfLatestProjectsSection = constants.TopMargin * 15
const DotIconHeight = 80
const DotIconWidth = 80

func MainContent(pdf *gopdf.GoPdf) {
	Experience(pdf)
	LatestProjects(pdf)
}
