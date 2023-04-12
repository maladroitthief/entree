package theme

import "image/color"

type TokyoNight struct {
}

func (t *TokyoNight) Black() color.Color {
  return color.RGBA{0x0f, 0x0f, 0x14, 0xff}
}

func (t *TokyoNight) Blue() color.Color {
  return color.RGBA{0x34, 0x54, 0x8a, 0xff}
}

func (t *TokyoNight) Green() color.Color {
  return color.RGBA{0x33, 0x63, 0x5c, 0xff}
}

func (t *TokyoNight) Yellow() color.Color {
  return color.RGBA{0x8f, 0x5e, 0x15, 0xff}
}

func (t *TokyoNight) Cyan() color.Color {
  return color.RGBA{0x0f, 0x4b, 0x6e, 0xff}
}

func (t *TokyoNight) White() color.Color {
  return color.RGBA{0xd5, 0xd6, 0xdb, 0xff}
}

func (t *TokyoNight) Magenta() color.Color {
  return color.RGBA{0x5a, 0x4a, 0x78, 0xff}
}

func (t *TokyoNight) Red() color.Color {
  return color.RGBA{0x8c, 0x43, 0x51, 0xff}
}

func (t *TokyoNight) BrightBlack() color.Color {
  return color.RGBA{0x41, 0x48, 0x68, 0xff}
}

func (t *TokyoNight) BrightBlue() color.Color {
  return color.RGBA{0x7a, 0xa2, 0xf7, 0xff}
}

func (t *TokyoNight) BrightGreen() color.Color {
  return color.RGBA{0x73, 0xda, 0xca, 0xff}
}

func (t *TokyoNight) BrightYellow() color.Color {
  return color.RGBA{0xe0, 0xaf, 0x68, 0xff}
}

func (t *TokyoNight) BrightCyan() color.Color {
  return color.RGBA{0x7d, 0xcf, 0xff, 0xff}
}

func (t *TokyoNight) BrightWhite() color.Color {
  return color.RGBA{0xcf, 0xc9, 0xc2, 0xff}
}

func (t *TokyoNight) BrightMagenta() color.Color {
  return color.RGBA{0xbb, 0x9a, 0xf7, 0xff}
}

func (t *TokyoNight) BrightRed() color.Color {
  return color.RGBA{0xf7, 0x76, 0x8e, 0xff}
}
