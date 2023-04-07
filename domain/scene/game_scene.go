package scene

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/canvas/input"
	"github.com/maladroitthief/entree/domain/canvas/physics"
	"github.com/maladroitthief/entree/domain/canvas/player"
)

type GameScene struct {
	canvas *canvas.Canvas
}

func NewGameScene(state *GameState) *GameScene {
	c := canvas.NewCanvas()
	gs := &GameScene{
		canvas: c,
	}

	pilot := player.NewPilot(
    input.NewPlayerInputComponent(state.InputSvc),
    physics.NewBasePhysicsComponent(),
  )
	gs.canvas.AddEntity(pilot)

	return gs
}

func (s *GameScene) Update(state *GameState) error {
	// Get the current scene actions
	for _, entity := range s.canvas.Entities() {
		entity.Update(s.canvas)
	}

	return nil
}

func (s *GameScene) GetEntities() []*canvas.Entity {
	return s.canvas.Entities()
}
