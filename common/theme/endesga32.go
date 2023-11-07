package theme

import "image/color"

type Endesga32 struct {
}

func (t *Endesga32) Black() color.Color {
  return color.RGBA{24, 20, 37, 0xff}
}

func (t *Endesga32) Blue() color.Color {
  return color.RGBA{58, 68, 102, 0xff}
}

func (t *Endesga32) Green() color.Color {
  return color.RGBA{38, 92, 66, 0xff}
}

func (t *Endesga32) Yellow() color.Color {
  return color.RGBA{254, 174, 52, 0xff}
}

func (t *Endesga32) Cyan() color.Color {
  return color.RGBA{0, 153, 219, 0xff}
}

func (t *Endesga32) White() color.Color {
  return color.RGBA{192, 203, 220, 0xff}
}

func (t *Endesga32) Magenta() color.Color {
  return color.RGBA{104, 56, 108, 0xff}
}

func (t *Endesga32) Red() color.Color {
  return color.RGBA{162, 38, 51, 0xff}
}

func (t *Endesga32) BrightBlack() color.Color {
  return color.RGBA{38, 43, 68, 0xff}
}

func (t *Endesga32) BrightBlue() color.Color {
  return color.RGBA{18, 78, 137, 0xff}
}

func (t *Endesga32) BrightGreen() color.Color {
  return color.RGBA{62, 137, 72, 0xff}
}

func (t *Endesga32) BrightYellow() color.Color {
  return color.RGBA{254, 231, 97, 0xff}
}

func (t *Endesga32) BrightCyan() color.Color {
  return color.RGBA{44, 232, 245, 0xff}
}

func (t *Endesga32) BrightWhite() color.Color {
  return color.RGBA{255, 255, 255, 0xff}
}

func (t *Endesga32) BrightMagenta() color.Color {
  return color.RGBA{181, 80, 136, 0xff}
}

func (t *Endesga32) BrightRed() color.Color {
  return color.RGBA{228, 59, 68, 0xff}
}
