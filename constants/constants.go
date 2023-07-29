package constants

const PageWidth = 8000.00
const PageHeight = 8000.00
const LeftContentWidth = 2000
const LeftMargin = 200.00
const RightMargin = 200.00
const TopMargin = 200.00
const ProfilePictureWidth = 500.00
const ProfilePictureHeight = 500.00
const VerticalLineHeight = 8000.00
const Hundred = 100.00
const HundredAndFifty = Hundred + 50.00
const VerticalSpace = 200.00
const HorizontalSpace = 200.00
const HeaderOne = 150
const HeaderTwo = 120
const Paragraph = 90

const SmallImageWidth = 80
const SmallImageHeight = 80

type ColorStruct struct {
	R float64
	G float64
	B float64
}

func SecondaryBackground() ColorStruct {
	c := ColorStruct{
		R: 246,
		G: 249,
		B: 252,
	}
	return c
}

func SecondaryText() ColorStruct {
	c := ColorStruct{
		R: 204,
		G: 204,
		B: 204,
	}
	return c
}
