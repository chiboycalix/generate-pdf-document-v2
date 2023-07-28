package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

func ProfilePicture(pdf *gopdf.GoPdf) {
	pdf.Image("images/pic-modified.png", constants.LeftMargin, StartOfProfilePicture, &gopdf.Rect{
		W: constants.ProfilePictureWidth,
		H: constants.ProfilePictureHeight,
	})
}

func FullNameAndJobDescription(pdf *gopdf.GoPdf) {
	pdf.SetXY(constants.LeftMargin, StartOfFullName)
	pdf.SetFontSize(constants.HeaderOne)
	pdf.Text("Igwe Chinonso")

	pdf.SetXY(constants.LeftMargin, StartOfJobTitle)
	pdf.SetFontSize(constants.HeaderTwo)
	pdf.SetTextColor(164, 120, 232)
	pdf.Text("Full-Stack Engineer")
}
