package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

const StartOfProfilePicture = constants.TopMargin
const StartOfFullName = StartOfProfilePicture + constants.ProfilePictureHeight + constants.VerticalSpace
const StartOfJobTitle = StartOfFullName + constants.VerticalSpace
const StartOfFirstHorizontalDivider = StartOfJobTitle + constants.VerticalSpace
const StartOfContactDetails = StartOfFirstHorizontalDivider + constants.VerticalSpace
const StartOfEmailAddress = StartOfContactDetails + 20
const EndOfEmailAddress = StartOfEmailAddress + 120
const StartOfWebsite = EndOfEmailAddress + constants.VerticalSpace*2
const EndOfWebsite = StartOfWebsite + 120
const StartOfPhoneNumber = EndOfWebsite + constants.VerticalSpace*2
const EndOfPhoneNumber = StartOfPhoneNumber + 120
const StartOfCurrentLocation = EndOfPhoneNumber + constants.VerticalSpace*2
const EndOfCurrentLocation = StartOfCurrentLocation + 120

const StartOfSocialsSection = EndOfCurrentLocation + constants.VerticalSpace*3
const StartOfGithub = StartOfSocialsSection + 180
const EndOfGithub = StartOfGithub + 120
const StartOfLinkedin = EndOfGithub + constants.VerticalSpace*2
const EndOfLinkedin = StartOfLinkedin + 120
const StartOfTwitter = EndOfLinkedin + constants.VerticalSpace*2
const EndOfTwitter = StartOfTwitter + 120

const StartOfLanguageSection = EndOfTwitter + constants.VerticalSpace*3

const StartOfEnglishY = StartOfLanguageSection + 180
const StartOfEnglishX = constants.LeftMargin + constants.HorizontalSpace*1.3
const EndOfEnglishX = constants.LeftMargin + constants.HorizontalSpace*1.3
const EndOfEnglishY = StartOfEnglishY + constants.HorizontalSpace/1.1

const StartOfIgboX = constants.LeftMargin + constants.HorizontalSpace*1.3
const StartOfIgboY = EndOfEnglishY + constants.VerticalSpace*2
const EndOfIgboX = constants.LeftMargin + constants.HorizontalSpace*1.3
const EndOfIgboY = StartOfIgboY + constants.HorizontalSpace/1.1

// const StartOfGerman = EndOfIgboY + constants.VerticalSpace*2
// const EndOfGerman = StartOfGerman + constants.HorizontalSpace/1.1

const StartOfGermanX = constants.LeftMargin + constants.HorizontalSpace*1.3
const StartOfGermanY = EndOfIgboY + constants.VerticalSpace*2
const EndOfGermanX = constants.LeftMargin + constants.HorizontalSpace*1.3
const EndOfGermanY = StartOfGermanY + constants.HorizontalSpace/1.1

const IconWidth = 180
const IconHeight = 170
const ContactRowHeight = constants.Hundred * 1.5

func LeftContent(pdf *gopdf.GoPdf) {
	ProfilePicture(pdf)
	FullNameAndJobDescription(pdf)
	HorizontalDivider(pdf, StartOfFirstHorizontalDivider)
	ContactDetails(pdf)
	HorizontalDivider(pdf, (StartOfFirstHorizontalDivider * 2.7))
	Socials(pdf)
	HorizontalDivider(pdf, (StartOfFirstHorizontalDivider * 4.2))
	Language(pdf)
}

func HorizontalDivider(pdf *gopdf.GoPdf, top float64) {
	pdf.SetStrokeColor(204, 204, 204)
	pdf.SetLineWidth(1)
	pdf.Line(constants.LeftMargin, top, constants.FirstColumnWidth-constants.RightMargin, top)
}
func LoadIcons(pdf *gopdf.GoPdf, iconPath string, leftMargin float64, topMargin float64, width float64, height float64) {
	pdf.Image(iconPath, leftMargin, topMargin, &gopdf.Rect{
		W: width,
		H: height,
	})
}
