package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

const StartOfProfilePicture = constants.TopMargin
const StartOfFullName = StartOfProfilePicture + constants.ProfilePictureHeight + constants.Hundred
const StartOfFullDescription = StartOfFullName + constants.Hundred
const StartOfFirstHorizontalDivider = StartOfFullDescription + constants.Hundred
const StartOfContactDetails = StartOfFirstHorizontalDivider + constants.Hundred
const ContactRowHeight = constants.Hundred * 1.5
const FirstContactDetailsFirstValueRowHeight = StartOfContactDetails + 25
const FirstContactDetailsSecondValueRowHeight = FirstContactDetailsFirstValueRowHeight + 55

func LeftContent(pdf *gopdf.GoPdf) {
	ProfilePicture(pdf)
	FullNameAndJobDescription(pdf)
	HorizontalDivider(pdf, StartOfFirstHorizontalDivider)
	ContactDetails(pdf)
	HorizontalDivider(pdf, (StartOfFirstHorizontalDivider*2 + 150))
	Socials(pdf)
	HorizontalDivider(pdf, (StartOfFirstHorizontalDivider*3 + 250))
	Language(pdf)
}

func HorizontalDivider(pdf *gopdf.GoPdf, top float64) {
	pdf.SetStrokeColor(204, 204, 204)
	pdf.SetLineWidth(1)
	pdf.Line(constants.LeftMargin, top, constants.FirstColumnWidth-constants.RightMargin, top)
}
func LoadIcons(pdf *gopdf.GoPdf, iconPath string, leftMargin float64, topMargin float64) {
	pdf.Image(iconPath, leftMargin, topMargin, &gopdf.Rect{
		W: 80,
		H: 70,
	})
}
