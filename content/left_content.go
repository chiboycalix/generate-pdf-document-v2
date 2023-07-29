package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

// ProfilePicture Section
const StartOfProfilePicture = constants.TopMargin
const StartOfFullName = StartOfProfilePicture + constants.ProfilePictureHeight + constants.VerticalSpace
const StartOfJobTitle = StartOfFullName + constants.VerticalSpace
const StartOfFirstHorizontalDivider = StartOfJobTitle + constants.VerticalSpace

// Contact Section
const StartOfContactSection = StartOfFirstHorizontalDivider + constants.VerticalSpace
const StartOfEmailAddress = StartOfContactSection + 20
const EndOfEmailAddress = StartOfEmailAddress + 120
const StartOfWebsite = EndOfEmailAddress + constants.VerticalSpace*2
const EndOfWebsite = StartOfWebsite + 120
const StartOfPhoneNumber = EndOfWebsite + constants.VerticalSpace*2
const EndOfPhoneNumber = StartOfPhoneNumber + 120
const StartOfCurrentLocation = EndOfPhoneNumber + constants.VerticalSpace*2
const EndOfCurrentLocation = StartOfCurrentLocation + 120

// Socials Section
const StartOfSocialsSection = EndOfCurrentLocation + constants.VerticalSpace*3
const StartOfGithub = StartOfSocialsSection + 180
const EndOfGithub = StartOfGithub + 120
const StartOfLinkedin = EndOfGithub + constants.VerticalSpace*2
const EndOfLinkedin = StartOfLinkedin + 120
const StartOfTwitter = EndOfLinkedin + constants.VerticalSpace*2
const EndOfTwitter = StartOfTwitter + 120

// Language Section
const StartOfLanguageSection = EndOfTwitter + constants.VerticalSpace*3

const StartOfEnglishY = StartOfLanguageSection + 180
const StartOfEnglishX = constants.LeftMargin + constants.HorizontalSpace*1.3
const EndOfEnglishX = constants.LeftMargin + constants.HorizontalSpace*1.3
const EndOfEnglishY = StartOfEnglishY + constants.HorizontalSpace/1.1

const StartOfIgboX = constants.LeftMargin + constants.HorizontalSpace*1.3
const StartOfIgboY = EndOfEnglishY + constants.VerticalSpace*2
const EndOfIgboX = constants.LeftMargin + constants.HorizontalSpace*1.3
const EndOfIgboY = StartOfIgboY + constants.HorizontalSpace/1.1

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
	pdf.SetStrokeColor(uint8(SecondaryTextColor.R), uint8(SecondaryTextColor.G), uint8(SecondaryTextColor.B))
	pdf.SetLineWidth(1)
	pdf.Line(constants.LeftMargin, top, constants.LeftContentWidth-constants.RightMargin, top)
}
