package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/chiboycalix/html-to-pdf/utils"
	"github.com/signintech/gopdf"
)

func Experience(pdf *gopdf.GoPdf) {
	pdf.SetTextColor(uint8(LightModePrimaryTextColor.R), uint8(LightModePrimaryTextColor.G), uint8(LightModePrimaryTextColor.B))
	// pdf.SetTextColor(uint8(DarkModePrimaryTextColor.R), uint8(DarkModePrimaryTextColor.G), uint8(DarkModePrimaryTextColor.B))
	pdf.SetXY(MainContentMarginLeft, StartOfExperienceSection)
	pdf.SetFontSize(constants.HeaderTwo)
	pdf.Text("Experience")

	SecondCompany(pdf)
	Bongalow(pdf)
	NIBSS(pdf)
}

func SecondCompany(pdf *gopdf.GoPdf) {
	// utils.LoadImage(pdf, "images/icons/circle.png", MainContentMarginLeft, StartOfExperienceSection+constants.VerticalSpace, constants.SmallImageWidth, constants.SmallImageHeight)
	utils.LoadImage(pdf, "images/dark-mode-icons/dot-dark.png", MainContentMarginLeft, StartOfExperienceSection+constants.VerticalSpace, constants.SmallImageWidth, constants.SmallImageHeight)

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace, StartOfExperienceSection+constants.VerticalSpace*1.3)
	pdf.SetTextColor(uint8(LightModePrimaryTextColor.R), uint8(LightModePrimaryTextColor.G), uint8(LightModePrimaryTextColor.B))
	// pdf.SetTextColor(uint8(DarkModeSecondaryTextColor.R), uint8(DarkModeSecondaryTextColor.G), uint8(DarkModeSecondaryTextColor.B))
	pdf.SetFontSize(constants.Paragraph)
	pdf.Text("Dec 2022 - Present")

	// utils.LoadImage(pdf, "images/logos/location.png", MainContentMarginLeft+constants.HorizontalSpace*5, StartOfExperienceSection+constants.VerticalSpace*0.95, constants.SmallImageWidth, constants.SmallImageHeight)
	utils.LoadImage(pdf, "images/dark-mode-icons/location-dark.png", MainContentMarginLeft+constants.HorizontalSpace*5, StartOfExperienceSection+constants.VerticalSpace*0.95, constants.SmallImageWidth, constants.SmallImageHeight)

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*5.5, StartOfExperienceSection+constants.VerticalSpace*1.3)
	pdf.SetFontSize(constants.Paragraph)
	pdf.Text("Remote, Netherlands")

	utils.LoadImage(pdf, "images/logos/sclogo.png", MainContentMarginLeft+constants.HorizontalSpace, StartOfExperienceSection+constants.VerticalSpace*2, 1200, 200)

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*2.7, StartOfExperienceSection+constants.VerticalSpace*2.3)
	pdf.Text("Full-Stack Engineer")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*2.7, StartOfExperienceSection+constants.VerticalSpace*3)
	pdf.SetTextColor(uint8(DarkModeSecondaryTextColor.R), uint8(DarkModeSecondaryTextColor.G), uint8(DarkModeSecondaryTextColor.B))
	pdf.Text("Second Company")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*1.3)
	pdf.SetTextColor(uint8(LightModePrimaryTextColor.R), uint8(LightModePrimaryTextColor.G), uint8(LightModePrimaryTextColor.B))
	// pdf.SetTextColor(uint8(LightModeSecondaryTextColor.R), uint8(LightModeSecondaryTextColor.G), uint8(LightModeSecondaryTextColor.B))
	pdf.Text("o. Improve website responsiveness across devices of all screens")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*2.0)
	pdf.Text("o. Improve website responsiveness across devices of all screens")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*2.7)
	pdf.Text("o. Improve website responsiveness across devices of all screens")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*3.4)
	pdf.Text("o. Improve website responsiveness across devices of all screens")
}

func Bongalow(pdf *gopdf.GoPdf) {
	// utils.LoadImage(pdf, "images/icons/circle.png", MainContentMarginLeft, StartOfExperienceSection+constants.VerticalSpace*5, constants.SmallImageWidth, constants.SmallImageHeight)
	utils.LoadImage(pdf, "images/dark-mode-icons/dot-dark.png", MainContentMarginLeft, StartOfExperienceSection+constants.VerticalSpace*5, constants.SmallImageWidth, constants.SmallImageHeight)

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace, StartOfExperienceSection+constants.VerticalSpace*5.3)
	pdf.Text("Apr 2022 - Nov 2022")

	// utils.LoadImage(pdf, "images/logos/location.png", MainContentMarginLeft+constants.HorizontalSpace*5.3, StartOfExperienceSection+constants.VerticalSpace*4.95, constants.SmallImageWidth, constants.SmallImageHeight)
	utils.LoadImage(pdf, "images/dark-mode-icons/location-dark.png", MainContentMarginLeft+constants.HorizontalSpace*5.3, StartOfExperienceSection+constants.VerticalSpace*4.95, constants.SmallImageWidth, constants.SmallImageHeight)

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*5.8, StartOfExperienceSection+constants.VerticalSpace*5.3)
	pdf.SetFontSize(constants.Paragraph)
	pdf.Text("Remote, Canada")

	utils.LoadImage(pdf, "images/logos/bongalow.png", MainContentMarginLeft+constants.HorizontalSpace, StartOfExperienceSection+constants.VerticalSpace*6, 1200, 200)

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*2.7, StartOfExperienceSection+constants.VerticalSpace*6.3)
	pdf.Text("Full-Stack Engineer")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*2.7, StartOfExperienceSection+constants.VerticalSpace*7)
	pdf.SetTextColor(uint8(DarkModeSecondaryTextColor.R), uint8(DarkModeSecondaryTextColor.G), uint8(DarkModeSecondaryTextColor.B))
	pdf.Text("Bongalow (Techstars ‘22)")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*5)
	pdf.SetTextColor(uint8(LightModePrimaryTextColor.R), uint8(LightModePrimaryTextColor.G), uint8(LightModePrimaryTextColor.B))
	// pdf.SetTextColor(uint8(LightModeSecondaryTextColor.R), uint8(LightModeSecondaryTextColor.G), uint8(LightModeSecondaryTextColor.B))
	pdf.Text("o. Improve website responsiveness across devices of all screens")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*5.7)
	pdf.Text("o. Improve website responsiveness across devices of all screens")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*6.4)
	pdf.Text("o. Improve website responsiveness across devices of all screens")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*7.1)
	pdf.Text("o. Improve website responsiveness across devices of all screens")

}

func NIBSS(pdf *gopdf.GoPdf) {
	// utils.LoadImage(pdf, "images/icons/circle.png", MainContentMarginLeft, StartOfExperienceSection+constants.VerticalSpace*9, constants.SmallImageWidth, constants.SmallImageHeight)
	utils.LoadImage(pdf, "images/dark-mode-icons/dot-dark.png", MainContentMarginLeft, StartOfExperienceSection+constants.VerticalSpace*9, constants.SmallImageWidth, constants.SmallImageHeight)

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace, StartOfExperienceSection+constants.VerticalSpace*9.3)
	pdf.Text("Dec 2019 - Mar 2022")

	// utils.LoadImage(pdf, "images/logos/location.png", MainContentMarginLeft+constants.HorizontalSpace*5.3, StartOfExperienceSection+constants.VerticalSpace*8.95, constants.SmallImageWidth, constants.SmallImageHeight)
	utils.LoadImage(pdf, "images/dark-mode-icons/location-dark.png", MainContentMarginLeft+constants.HorizontalSpace*5.3, StartOfExperienceSection+constants.VerticalSpace*8.95, constants.SmallImageWidth, constants.SmallImageHeight)

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*5.8, StartOfExperienceSection+constants.VerticalSpace*9.3)
	pdf.SetFontSize(constants.Paragraph)
	pdf.Text("Onsite, Nigeria")

	utils.LoadImage(pdf, "images/logos/nibbs_logo.png", MainContentMarginLeft+constants.HorizontalSpace, StartOfExperienceSection+constants.VerticalSpace*10, 500, 200)

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*2.7, StartOfExperienceSection+constants.VerticalSpace*10.3)
	pdf.Text("Frontend Engineer")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*2.7, StartOfExperienceSection+constants.VerticalSpace*11)
	pdf.SetTextColor(uint8(DarkModeSecondaryTextColor.R), uint8(DarkModeSecondaryTextColor.G), uint8(DarkModeSecondaryTextColor.B))
	pdf.Text("Nigeria Interbank Settlement Systems")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*9.3)
	pdf.SetTextColor(uint8(LightModePrimaryTextColor.R), uint8(LightModePrimaryTextColor.G), uint8(LightModePrimaryTextColor.B))
	// pdf.SetTextColor(uint8(LightModeSecondaryTextColor.R), uint8(LightModeSecondaryTextColor.G), uint8(LightModeSecondaryTextColor.B))
	pdf.Text("o. Improve website responsiveness across devices of all screens")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*10)
	pdf.Text("o. Improve website responsiveness across devices of all screens")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*10.7)
	pdf.Text("o. Improve website responsiveness across devices of all screens")

	pdf.SetXY(MainContentMarginLeft+constants.HorizontalSpace*13, StartOfExperienceSection+constants.VerticalSpace*11.4)
	pdf.Text("o. Improve website responsiveness across devices of all screens")

}
