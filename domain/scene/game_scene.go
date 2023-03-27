package scene

import "github.com/maladroitthief/entree/domain/canvas"

type GameScene struct {
	canvas *canvas.Canvas
}

func NewGameScene() *GameScene {
	c := canvas.NewCanvas()
	gs := &GameScene{
		canvas: c,
	}

	return gs
}

func (s *GameScene) Update(state *GameState) error {
	state.Log.Info("Game Scene", nil)
	return nil
}

func (s *GameScene) GetEntities() []*canvas.Entity {
	return nil
}
