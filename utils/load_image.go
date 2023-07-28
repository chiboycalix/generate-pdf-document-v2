package utils

import "github.com/signintech/gopdf"

func LoadImage(pdf *gopdf.GoPdf, path string, left float64, top float64, width float64, height float64) {
	pdf.Image(path, left, top, &gopdf.Rect{
		W: width,
		H: height,
	})
}
