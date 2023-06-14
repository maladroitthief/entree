package scene

import (
	"image/color"

	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/canvas/background"
	"github.com/maladroitthief/entree/domain/canvas/environment"
	"github.com/maladroitthief/entree/domain/canvas/player"
	"github.com/maladroitthief/entree/domain/physics"
	"github.com/maladroitthief/entree/domain/settings"
)

type GameScene struct {
	columns         int
	rows            int
	cellSize        int
	camera          Camera
	middleground    *canvas.Canvas
	background      *canvas.Canvas
	backgroundColor color.Color
}

func NewGameScene(state *GameState) *GameScene {
	gs := &GameScene{
		columns:         8,
		rows:            8,
		cellSize:        16,
		backgroundColor: state.Theme.Green(),
	}
	gs.middleground = canvas.NewCanvas(gs.columns, gs.rows, gs.cellSize)
	gs.background = canvas.NewCanvas(gs.columns, gs.rows, gs.cellSize)

	pilot := player.NewPilot(player.NewPlayerInputComponent(state.InputSvc))
	gs.middleground.AddEntity(pilot)

	grass := background.Grass(100, 100)
	gs.background.AddEntity(grass)

	for i := 0; i < 8; i++ {
		wall := environment.Wall((float64(i)*environment.WallSize + environment.WallSize/2), 120)
		gs.middleground.AddEntity(wall)
	}

	gs.camera = NewCamera(
		pilot,
		physics.Vector{X: 800, Y: 800},
	)

	return gs
}

func (s *GameScene) Update(state *GameState) error {
	for _, input := range state.InputSvc.CurrentInputs() {
		switch input {
		case settings.Menu:
			return SceneTermination
		}
	}

	// Update the canvas
	s.middleground.Update()

	// Get the current scene actions
	for _, entity := range s.middleground.Entities() {
		entity.Update(s.middleground)
	}

	return nil
}

func (s *GameScene) GetCanvasSize() (width, height int) {
	return s.columns * s.cellSize, s.rows * s.cellSize
}

func (s *GameScene) GetCanvasGrid() (rows, columns int) {
	return s.rows, s.columns
}

func (s *GameScene) GetCanvasCellSize() int {
	return s.cellSize
}

func (s *GameScene) GetEntities() []canvas.Entity {
	entities := s.background.Entities()
	entities = append(entities, s.middleground.Entities()...)
	return entities
}

func (s *GameScene) GetBackgroundColor() color.Color {
	return s.backgroundColor
}

func (s *GameScene) GetCamera() Camera {
	return s.camera
}
