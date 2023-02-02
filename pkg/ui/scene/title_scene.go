package scene

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	fontBaseSize = 6
)

var (
	titleFonts map[int]font.Face
	titleChars = []string{
		`______ ______                     _____                          _____`,
		`\ .  / \  . /                     \ . |                          \  .|`,
		` | .|   |. |                       | .|                           |. |`,
		` |. |___| .|   _____  _____ _____  |. | ___     ______  ____  ___ | .|  ____`,
		` |:::___:::|   \::::\ \:::| \:::|  |::|/:::\   /::::::\ \:::|/:::\|::| /::/`,
		` |xx|   |xx|  ___ \xx| |xx|  |xx|  |xx|  \xx\ |xx|__)xx| |xx|  \x||xx|/x/`,
		` |xx|   |xx| /xxx\|xx| |xx|  |xx|  |xx|   |xx||xx|\xxxx| |xx|     |xxxxxx\`,
		` |XX|   |XX||XX(__|XX| |XX\__|XX|  |XX|__/XXX||XX|_____  |XX|     |XX| \XX\_`,
		` |XX|   |XX| \XXXX/\XX\ \XXX/|XXX\/XXX/\XXXX/  \XXXXXX/ /XXXX\   /XXXX\ \XXX\`,
		` |XX|   |XX|_________________________________________________________________`,
		` |XX|   |XX||XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\`,
		`_|XX|   |XX|_`,
		`\XXX|   |XXX/`,
		` \XX|   |XX/`,
		`  \X|   |X/`,
		`   \|   |/`,
	}
	titleColors = []string{
		`LLLLLL LLLLLL                     LLLLL                          LLLLL`,
		`ERRRRE ERRRRE                     ERRRE                          ERRRE`,
		` ERRE   ERRE                       ERRE                           ERRE`,
		` ERRELLLERRE   LLLLL  LLLLL LLLLL  ERRE LLL     LLLLLL  LLLL  LLL ERRE  LLLL`,
		` ERRREEERRRE   ERRRRL ERRRE ERRRE  ERREERRRL   LRRRRRRL ERRRLLRRRLERRE LRRE`,
		` ERRE   ERRE  LLL ERRE ERRE  ERRE  ERRE  ERRL ERRELLERRE ERRE  EREERRELRE`,
		` EOOE   EOOE LOOOEEOOE EOOE  EOOE  EOOE   EOOEEOOEEOOOOE EOOE     EOOOOOOL`,
		` EGGE   EGGEEGGELLEGGE EGGLLLEGGE  EGGELLLGGGEEGGELLLLL  EGGE     EGGE EGGLL`,
		` EYYE   EYYE EYYYYEEYYE EYYYEEYYYLLYYYEEYYYYE  EYYYYYYE LYYYYL   LYYYYL EYYYL`,
		` EYYE   EYYELLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL`,
		` EYYE   EYYEEYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYL`,
		`LEYYE   EYYEL`,
		`EYYYE   EYYYE`,
		` EYYE   EYYE`,
		`  EYE   EYE`,
		`   EE   EE`,
	}
	colorLUT = map[string]color.Color{
		"L": color.NRGBA{207, 216, 220, 255}, // light warm gray
		"E": color.NRGBA{176, 190, 197, 255}, // warm gray
		"R": color.NRGBA{183, 28, 28, 255},   // red
		"O": color.NRGBA{239, 108, 0, 255},   // carrot
		"G": color.NRGBA{255, 179, 0, 255},   // gold
		"Y": color.NRGBA{255, 234, 0, 255},   // yellow
		" ": color.NRGBA{0, 0, 0, 0},         // nothing
	}
)

type TitleScene struct {
	count int
}

func (s *TitleScene) Update(state *GameState) error {
	if state.Input.IsAnyAction() {
		state.SceneManager.GoTo(NewGameScene())
	}

	return nil
}

func (s *TitleScene) Draw(r *ebiten.Image) {
	s.drawTitle(r, 1)
}

func (s *TitleScene) drawTitle(r *ebiten.Image, scale int) {
	offset := 32
	for y, line := range titleChars {
    yPos := (y * scale * fontBaseSize * 4) + offset
		for x, char := range line {
      xPos := (x * scale * fontBaseSize * 2) + offset
			clr := colorLUT[string(titleColors[y][x])]
      text.Draw(r, string(char), getTitleFonts(4), xPos, yPos, clr)
		}
	}
}

func getTitleFonts(scale int) font.Face {
	if titleFonts == nil {
		tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
		if err != nil {
			log.Fatal(err)
		}

		titleFonts = map[int]font.Face{}
		for i := 1; i <= 4; i++ {
			const dpi = 72
			titleFonts[i], err = opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    float64(fontBaseSize * i),
				DPI:     dpi,
				Hinting: font.HintingFull,
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return titleFonts[scale]
}
