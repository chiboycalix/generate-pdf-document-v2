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

func LightModePageBackgroundColor() ColorStruct {
	c := ColorStruct{
		R: 255,
		G: 255,
		B: 255,
	}
	return c
}

func LightModeBoxBackgroundColor() ColorStruct {
	c := ColorStruct{
		R: 246,
		G: 249,
		B: 252,
	}
	return c
}

func LightModePrimaryTextColor() ColorStruct {
	c := ColorStruct{
		R: 0,
		G: 0,
		B: 0,
	}
	return c
}

func LightModeSecondaryTextColor() ColorStruct {
	c := ColorStruct{
		R: 204,
		G: 204,
		B: 204,
	}
	return c
}

func DarkModePageBackgroundColor() ColorStruct {
	c := ColorStruct{
		R: 35,
		G: 35,
		B: 37,
	}
	return c
}

func DarkModePrimaryTextColor() ColorStruct {
	c := ColorStruct{
		R: 255,
		G: 255,
		B: 255,
	}
	return c
}

func DarkModeSecondaryTextColor() ColorStruct {
	c := ColorStruct{
		R: 172,
		G: 177,
		B: 195,
	}
	return c
}
