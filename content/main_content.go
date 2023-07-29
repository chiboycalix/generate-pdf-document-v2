package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

const StartOfExperienceSection = constants.TopMargin * 1.5
const MainContentMarginLeft = constants.LeftContentWidth * 1.1

const StartOfLatestProjectsSection = StartOfExperienceSection * 10
const StartOfSkillSection = StartOfLatestProjectsSection * 1.73
const StartOfEducationSection = StartOfSkillSection * 1.26
const DotIconHeight = 80
const DotIconWidth = 80

var BackgroundColor = constants.SecondaryBackground()
var SecondaryTextColor = constants.SecondaryText()

func MainContent(pdf *gopdf.GoPdf) {
	Experience(pdf)
	LatestProjects(pdf)
	Skills(pdf)
	Education(pdf)
}
