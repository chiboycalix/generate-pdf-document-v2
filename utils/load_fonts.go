package utils

import (
	"log"

	"github.com/signintech/gopdf"
)

func LoadFonts(pdf *gopdf.GoPdf) {
	err := pdf.AddTTFFont("Outfit-Regular", "fonts/Outfit/Outfit-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("Outfit-Regular", "", 0)
	if err != nil {
		log.Print(err.Error())
		return
	}
}
