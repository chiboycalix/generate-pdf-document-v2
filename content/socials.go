package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/chiboycalix/html-to-pdf/utils"
	"github.com/signintech/gopdf"
)

func Socials(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfSocialsSection)
	pdf.Text("Socials")
	Github(pdf)
	Linkedin(pdf)
	Twitter(pdf)
}

func Github(pdf *gopdf.GoPdf) {
	utils.LoadImage(pdf, "images/socials/github.png", constants.LeftMargin, StartOfGithub, IconWidth, IconHeight)

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, StartOfGithub+50)
	pdf.SetFontSize(constants.Paragraph)
	pdf.SetTextColor(uint8(SecondaryTextColor.R), uint8(SecondaryTextColor.G), uint8(SecondaryTextColor.B))

	pdf.Text("GitHub")

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, EndOfGithub+constants.HorizontalSpace/3)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("@chiboycalix")
}

func Linkedin(pdf *gopdf.GoPdf) {
	utils.LoadImage(pdf, "images/socials/linkedin.png", constants.LeftMargin, StartOfLinkedin, IconWidth, IconHeight)

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, StartOfLinkedin+50)
	pdf.SetFontSize(constants.Paragraph)
	pdf.SetTextColor(uint8(SecondaryTextColor.R), uint8(SecondaryTextColor.G), uint8(SecondaryTextColor.B))
	pdf.Text("Linkedin")

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, EndOfLinkedin+constants.HorizontalSpace/3)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("igwe-chinonso/")
}

func Twitter(pdf *gopdf.GoPdf) {
	utils.LoadImage(pdf, "images/socials/twitter.png", constants.LeftMargin, StartOfTwitter, IconWidth, IconHeight)

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, StartOfTwitter+50)
	pdf.SetFontSize(constants.Paragraph)
	pdf.SetTextColor(uint8(SecondaryTextColor.R), uint8(SecondaryTextColor.G), uint8(SecondaryTextColor.B))
	pdf.Text("Twitter")

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, EndOfTwitter+constants.HorizontalSpace/3)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("@thorsgardian_")
}
