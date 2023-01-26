package scene

import "github.com/hajimehoshi/ebiten/v2"

type TitleScene struct {
	count int
}

func (s *TitleScene) Update(state *GameState) error {
	return nil
}

func (s *TitleScene) Draw(r *ebiten.Image) {

}
