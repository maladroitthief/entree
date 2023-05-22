package scene

import (
	"image/color"

	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/canvas/background"
	"github.com/maladroitthief/entree/domain/canvas/environment"
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

	pilot := player.NewPilot(player.NewPlayerInputComponent(state.InputSvc))
	gs.middleground.AddEntity(pilot)

	grass := background.Grass(100, 100)
	gs.background.AddEntity(grass)

	for i := 0; i < 8; i++ {
		wall := environment.Wall(200+(float64(i)*environment.WallSize), 200)
		gs.middleground.AddEntity(wall)
	}

	return gs
}

func (s *GameScene) Update(state *GameState) error {
	// Update the canvas
	s.middleground.Update()

	// Get the current scene actions
	for _, entity := range s.middleground.Entities() {
		entity.Update(s.middleground)
	}

	return nil
}

func (s *GameScene) GetEntities() []canvas.Entity {
	entities := s.background.Entities()
	entities = append(entities, s.middleground.Entities()...)
	return entities
}

func (s *GameScene) GetBackgroundColor() color.Color {
	return s.backgroundColor
}
