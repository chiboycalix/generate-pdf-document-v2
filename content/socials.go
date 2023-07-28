package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

func Socials(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*5))
	pdf.Text("Socials")
	Github(pdf)
	Linkedin(pdf)
	Twitter(pdf)
}

func Github(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*5.5))
	LoadIcons(pdf, "images/socials/github.png", constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*5.5), 80, 70)
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsFirstValueRowHeight+5.5*constants.HundredAndFifty)
	pdf.SetFontSize(35)
	pdf.SetTextColor(204, 204, 204)

	pdf.Text("GitHub")

	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsSecondValueRowHeight+5.5*constants.HundredAndFifty)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("@chiboycalix")
}

func Linkedin(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*6.5))
	LoadIcons(pdf, "images/socials/linkedin.png", constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*6.5), 80, 70)
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsFirstValueRowHeight+6.5*constants.HundredAndFifty)
	pdf.SetFontSize(35)
	pdf.SetTextColor(204, 204, 204)

	pdf.Text("Linkedin")

	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsSecondValueRowHeight+6.5*constants.HundredAndFifty)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("igwe-chinonso/")
}

func Twitter(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*7.5))
	LoadIcons(pdf, "images/socials/twitter.png", constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*7.5), 80, 70)
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsFirstValueRowHeight+7.5*constants.HundredAndFifty)
	pdf.SetFontSize(35)
	pdf.SetTextColor(204, 204, 204)

	pdf.Text("Twitter")

	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsSecondValueRowHeight+7.5*constants.HundredAndFifty)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("@thorsgardian_")
}
