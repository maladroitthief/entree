package scene

import (
	"image/color"

	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/canvas/background"
	"github.com/maladroitthief/entree/domain/canvas/components/input"
	"github.com/maladroitthief/entree/domain/canvas/components/physics"
	"github.com/maladroitthief/entree/domain/canvas/player"
)

type GameScene struct {
	middleground    *canvas.Canvas
	background      *canvas.Canvas
	backgroundColor color.Color
}

func NewGameScene(state *GameState) *GameScene {
	mgc := canvas.NewCanvas(8, 8, 16)
	bgc := canvas.NewCanvas(8, 8, 16)
	gs := &GameScene{
		middleground:    mgc,
		background:      bgc,
		backgroundColor: state.Theme.Green(),
	}

	pilot := player.NewPilot(
		input.NewPlayerInputComponent(state.InputSvc),
		physics.NewBasePhysicsComponent(),
	)
	gs.middleground.AddEntity(pilot)

	grass := background.Grass(100, 100)
	gs.background.AddEntity(grass)

	return gs
}

func (s *GameScene) Update(state *GameState) error {
	// Get the current scene actions
	for _, entity := range s.middleground.Entities() {
		entity.Update(s.middleground)
	}

	return nil
}

func (s *GameScene) GetEntities() []*canvas.Entity {
	entities := s.background.Entities()
	entities = append(entities, s.middleground.Entities()...)
	return entities
}

func (s *GameScene) GetBackgroundColor() color.Color {
	return s.backgroundColor
}
