package content

import (
	"github.com/chiboycalix/html-to-pdf/constants"
	"github.com/signintech/gopdf"
)

func Skills(pdf *gopdf.GoPdf) {
	pdf.SetXY(MainContentMarginLeft, StartOfSkillSection)
	pdf.SetFontSize(constants.HeaderTwo)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text("Skills")
}

func Javascript() {}
func TypeScript() {}
func NodeJs()     {}
func ReactJs()    {}
func Redux()      {}
func NextJs()     {}
func Tailwind()   {}
func Golang()     {}
func HTMLCSS()    {}
func Git()        {}
func MongoDB()    {}
func PostGres()   {}
