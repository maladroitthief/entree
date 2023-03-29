package scene

import (
	"github.com/maladroitthief/entree/domain/action"
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/canvas/player"
)

type GameScene struct {
	canvas *canvas.Canvas
}

func NewGameScene() *GameScene {
	c := canvas.NewCanvas()
	gs := &GameScene{
		canvas: c,
	}

	pilot := player.NewPilot()
	gs.canvas.AddEntity(pilot)

	return gs
}

func (s *GameScene) Update(state *GameState) error {
	for _, entity := range s.canvas.Entities() {
		entity.Update([]action.Input{}, s.canvas)
	}

	return nil
}

func (s *GameScene) GetEntities() []*canvas.Entity {
	return s.canvas.Entities()
}
