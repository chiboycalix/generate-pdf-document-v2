package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

func ContactDetails(pdf *gopdf.GoPdf) {
	Email(pdf)
	Website(pdf)
	PhoneNumber(pdf)
	CurrentLocation(pdf)
}

func Email(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails)
	LoadIcons(pdf, "images/logos/mail.png", constants.LeftMargin, StartOfContactDetails, 80, 70)
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsFirstValueRowHeight)
	pdf.SetFontSize(35)
	pdf.SetTextColor(204, 204, 204)

	pdf.Text("Email")

	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsSecondValueRowHeight)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("Igwechinonso77@gmail.com")
}

func Website(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails)
	LoadIcons(pdf, "images/logos/links.png", constants.LeftMargin, StartOfContactDetails+ContactRowHeight, 80, 70)
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsFirstValueRowHeight+constants.HundredAndFifty)
	pdf.SetFontSize(35)

	pdf.SetTextColor(204, 204, 204)
	pdf.Text("Website")

	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsSecondValueRowHeight+constants.HundredAndFifty)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("https://www.google.com")
}

func PhoneNumber(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails)
	LoadIcons(pdf, "images/logos/phone.png", constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*2), 80, 70)
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsFirstValueRowHeight+2*constants.HundredAndFifty)
	pdf.SetFontSize(35)
	pdf.SetTextColor(204, 204, 204)
	pdf.Text("Phone")
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsSecondValueRowHeight+2*constants.HundredAndFifty)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("(+234) 816 584 2442")
}

func CurrentLocation(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfContactDetails)
	LoadIcons(pdf, "images/logos/location.png", constants.LeftMargin, StartOfContactDetails+(ContactRowHeight*3), 80, 70)
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsFirstValueRowHeight+3*constants.HundredAndFifty)
	pdf.SetFontSize(35)
	pdf.SetTextColor(204, 204, 204)
	pdf.Text("Location")
	pdf.SetXY(constants.LeftMargin+constants.Hundred+50, FirstContactDetailsSecondValueRowHeight+3*constants.HundredAndFifty)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("Lagos, Nigeria")
}
