package scene

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/canvas/background"
	"github.com/maladroitthief/entree/domain/canvas/input"
	"github.com/maladroitthief/entree/domain/canvas/physics"
	"github.com/maladroitthief/entree/domain/canvas/player"
)

type GameScene struct {
	middleGround *canvas.Canvas
	backGround   *canvas.Canvas
}

func NewGameScene(state *GameState) *GameScene {
	mgc := canvas.NewCanvas()
	bgc := canvas.NewCanvas()
	gs := &GameScene{
		middleGround: mgc,
		backGround:   bgc,
	}

	pilot := player.NewPilot(
		input.NewPlayerInputComponent(state.InputSvc),
		physics.NewBasePhysicsComponent(),
	)
	gs.middleGround.AddEntity(pilot)

	grass := background.NewGrass(100, 100)
	gs.backGround.AddEntity(grass)

	return gs
}

func (s *GameScene) Update(state *GameState) error {
	// Get the current scene actions
	for _, entity := range s.middleGround.Entities() {
		entity.Update(s.middleGround)
	}

	return nil
}

func (s *GameScene) GetEntities() []*canvas.Entity {
  entities := s.backGround.Entities()
  entities = append(entities, s.middleGround.Entities()...)
	return entities
}
