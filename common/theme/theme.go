package theme

import "image/color"

type Colors interface {
	Black() color.Color
	Blue() color.Color
	Green() color.Color
	Yellow() color.Color
	Cyan() color.Color
	White() color.Color
	Magenta() color.Color
	Red() color.Color
	BrightBlack() color.Color
	BrightBlue() color.Color
	BrightGreen() color.Color
	BrightYellow() color.Color
	BrightCyan() color.Color
	BrightWhite() color.Color
	BrightMagenta() color.Color
	BrightRed() color.Color
}
