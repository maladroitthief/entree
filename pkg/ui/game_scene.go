package ui

import (
	"context"
	"image/color"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/content/player"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/entree/pkg/engine/level"
	"github.com/maladroitthief/entree/pkg/engine/server"
	"github.com/rs/zerolog/log"
)

type GameScene struct {
	ctx      context.Context
	gridX    int
	gridY    int
	cellSize int

	ecs       *core.ECS
	playerId  data.GenerationalIndex
	ai        *server.AIServer
	state     *server.StateServer
	physics   *server.PhysicsServer
	animation *server.AnimationServer

	camera          *Camera
	cameraFocus     core.Entity
	backgroundColor color.Color
}

func NewGameScene(ctx context.Context, state *SceneState) *GameScene {
	gs := &GameScene{
		ctx:             ctx,
		gridX:           2,
		gridY:           2,
		cellSize:        32,
		ecs:             core.NewECS(ctx),
		backgroundColor: state.theme.Green(),
	}

	gs.ai = server.NewAIServer()
	gs.state = server.NewStateServer()
	gs.physics = server.NewPhysicsServer(
		gs.ecs,
		float64(gs.gridX*level.RoomWidth),
		float64(gs.gridY*level.RoomHeight),
		float64(gs.cellSize),
	)
	gs.animation = server.NewAnimationServer()

	player := player.NewFederico(gs.ecs, 0, 0)
	gs.playerId = player.Id
	gs.cameraFocus = player

	level := level.NewLevel(
		level.NewRoomFactory(),
		level.NewBlockFactory(gs.ecs),
		player,
		gs.gridX,
		gs.gridY,
		gs.cellSize,
	)
	level.GenerateRooms()
	level.Render(gs.ecs)

	gs.physics.Load(gs.ecs)
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
		case core.InputMenu:
			return Termination
		}
	}

	s.state.Update(s.ecs)
	ProcessPlayerGameInputs(s.ecs, s.playerId, inputs)
	// s.ai.Update(s.world, inputs)
	s.physics.Update(s.ecs)
	s.animation.Update(s.ecs)

	return nil
}

func (s *GameScene) Size() (width, height int) {
	width = s.gridX * level.RoomWidth * s.cellSize
	height = s.gridY * level.RoomHeight * s.cellSize

	return width, height
}

func (s *GameScene) GetCanvasGrid() (rows, columns int) {
	return s.gridY, s.gridX
}

func (s *GameScene) CellSize() int {
	return s.cellSize
}

func (s *GameScene) GetState() *core.ECS {
	return s.ecs
}

func (s *GameScene) BackgroundColor() color.Color {
	return s.backgroundColor
}

func (s *GameScene) GetCamera() *Camera {
	cameraPosition, err := s.ecs.GetPosition(s.cameraFocus.Id)
	if err != nil {
		log.Warn().Err(err).Any("cameraPosition", cameraPosition)
	}

	s.camera.X = cameraPosition.X
	s.camera.Y = cameraPosition.Y
	return s.camera
}
