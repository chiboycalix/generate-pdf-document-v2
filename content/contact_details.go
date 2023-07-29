package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/chiboycalix/html-to-pdf/utils"
	"github.com/signintech/gopdf"
)

var DarkModeSecondaryTextColor = constants.DarkModeSecondaryTextColor()
var DarkModePrimaryTextColor = constants.DarkModePrimaryTextColor()

func ContactDetails(pdf *gopdf.GoPdf) {
	Email(pdf)
	Website(pdf)
	PhoneNumber(pdf)
	CurrentLocation(pdf)
}

func Email(pdf *gopdf.GoPdf) {
	// utils.LoadImage(pdf, "images/logos/mail.png", constants.LeftMargin, StartOfContactSection, IconWidth, IconHeight)
	utils.LoadImage(pdf, "images/dark-mode-icons/mail-dark.png", constants.LeftMargin, StartOfContactSection, IconWidth, IconHeight)

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, StartOfEmailAddress+50)
	pdf.SetFontSize(constants.Paragraph)
	pdf.SetTextColor(uint8(DarkModeSecondaryTextColor.R), uint8(DarkModeSecondaryTextColor.G), uint8(DarkModeSecondaryTextColor.B))
	pdf.Text("Email")

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, EndOfEmailAddress+constants.HorizontalSpace/3)
	pdf.SetTextColor(uint8(DarkModePrimaryTextColor.R), uint8(DarkModePrimaryTextColor.G), uint8(DarkModePrimaryTextColor.B))
	pdf.Text("Igwechinonso77@gmail.com")

}

func Website(pdf *gopdf.GoPdf) {
	// utils.LoadImage(pdf, "images/logos/links.png", constants.LeftMargin, StartOfWebsite, IconWidth, IconHeight)
	utils.LoadImage(pdf, "images/dark-mode-icons/link-dark.png", constants.LeftMargin, StartOfWebsite, IconWidth, IconHeight)

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, StartOfWebsite+50)
	pdf.SetTextColor(uint8(DarkModeSecondaryTextColor.R), uint8(DarkModeSecondaryTextColor.G), uint8(DarkModeSecondaryTextColor.B))
	pdf.Text("Website")

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, EndOfWebsite+constants.HorizontalSpace/3)
	pdf.SetTextColor(uint8(DarkModePrimaryTextColor.R), uint8(DarkModePrimaryTextColor.G), uint8(DarkModePrimaryTextColor.B))
	pdf.Text("https://www.google.com")
}

func PhoneNumber(pdf *gopdf.GoPdf) {
	// utils.LoadImage(pdf, "images/logos/phone.png", constants.LeftMargin, StartOfPhoneNumber, IconWidth, IconHeight)
	utils.LoadImage(pdf, "images/dark-mode-icons/phone-dark.png", constants.LeftMargin, StartOfPhoneNumber, IconWidth, IconHeight)

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, StartOfPhoneNumber+50)
	pdf.SetFontSize(constants.Paragraph)
	pdf.SetTextColor(uint8(DarkModeSecondaryTextColor.R), uint8(DarkModeSecondaryTextColor.G), uint8(DarkModeSecondaryTextColor.B))
	pdf.Text("Phone")

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, EndOfPhoneNumber+constants.HorizontalSpace/3)
	pdf.SetTextColor(uint8(DarkModePrimaryTextColor.R), uint8(DarkModePrimaryTextColor.G), uint8(DarkModePrimaryTextColor.B))
	pdf.Text("(+234) 816 584 2442")
}

func CurrentLocation(pdf *gopdf.GoPdf) {
	// utils.LoadImage(pdf, "images/dark-mode-icons/location-dark.png", constants.LeftMargin, StartOfCurrentLocation, IconWidth, IconHeight)
	utils.LoadImage(pdf, "images/dark-mode-icons/location-dark.png", constants.LeftMargin, StartOfCurrentLocation, IconWidth, IconHeight)

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, StartOfCurrentLocation+50)
	pdf.SetFontSize(constants.Paragraph)
	pdf.SetTextColor(uint8(DarkModeSecondaryTextColor.R), uint8(DarkModeSecondaryTextColor.G), uint8(DarkModeSecondaryTextColor.B))
	pdf.Text("Location")

	pdf.SetXY(constants.LeftMargin+constants.HorizontalSpace*1.3, EndOfCurrentLocation+constants.HorizontalSpace/3)
	pdf.SetTextColor(uint8(DarkModePrimaryTextColor.R), uint8(DarkModePrimaryTextColor.G), uint8(DarkModePrimaryTextColor.B))
	pdf.Text("Lagos, Nigeria")
}
