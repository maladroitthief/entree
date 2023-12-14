package ui

import (
	"image/color"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/pkg/content/player"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/entree/pkg/engine/level"
	"github.com/maladroitthief/entree/pkg/engine/server"
)

type GameScene struct {
	columns  int
	rows     int
	cellSize int

	world     *core.ECS
	ai        *server.AIServer
	state     *server.StateServer
	physics   *server.PhysicsServer
	animation *server.AnimationServer

	camera          *Camera
	cameraFocus     core.Entity
	log             logs.Logger
	backgroundColor color.Color
}

func NewGameScene(state *SceneState) *GameScene {
	gs := &GameScene{
		columns:         6,
		rows:            6,
		cellSize:        32,
		world:           core.NewECS(),
		log:             state.log,
		backgroundColor: state.theme.Green(),
	}

	gs.ai = server.NewAIServer()
	gs.state = server.NewStateServer()
	gs.physics = server.NewPhysicsServer(
		state.log,
		float64(gs.columns*level.RoomWidth),
		float64(gs.rows*level.RoomHeight),
		float64(gs.cellSize),
	)
	gs.animation = server.NewAnimationServer()

	player := player.NewFederico(gs.world)
	// player := enemy.NewOnyawn(gs.world)
	gs.cameraFocus = player

	level := level.NewLevel(
		level.NewRoomFactory(),
		level.NewBlockFactory(),
		player,
	)
	level.GenerateRooms()
	level.Render(gs.world)

	gs.physics.Load(gs.world)
	gs.camera = NewCamera(
		0,
		0,
		data.Vector{X: 200, Y: 200},
	)

	return gs
}

func (s *GameScene) Update(state *SceneState) error {
	inputs := state.input.CurrentInputs()

	for _, input := range inputs {
		switch input {
		case core.Menu:
			return Termination
		}
	}

	s.state.Update(s.world)
	s.ai.Update(s.world, inputs)
	s.physics.Update(s.world)
	s.animation.Update(s.world)

	return nil
}

func (s *GameScene) Size() (width, height int) {
	width = s.columns * level.RoomWidth * s.cellSize
	height = s.rows * level.RoomHeight * s.cellSize

	return width, height
}

func (s *GameScene) GetCanvasGrid() (rows, columns int) {
	return s.rows, s.columns
}

func (s *GameScene) CellSize() int {
	return s.cellSize
}

func (s *GameScene) GetState() *core.ECS {
	return s.world
}

func (s *GameScene) BackgroundColor() color.Color {
	return s.backgroundColor
}

func (s *GameScene) GetCamera() *Camera {
	cameraPosition, err := s.world.GetPosition(s.cameraFocus.Id)
	if err != nil {
		s.log.Error("GameScene.GetCamera", cameraPosition, err)
	}

	s.camera.X = cameraPosition.X
	s.camera.Y = cameraPosition.Y
	return s.camera
}
