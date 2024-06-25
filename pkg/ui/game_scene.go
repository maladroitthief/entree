package ui

import (
	"context"
	"image/color"

	"github.com/maladroitthief/caravan"
	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/content/player"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/entree/pkg/engine/level"
	"github.com/maladroitthief/entree/pkg/engine/server"
	"github.com/maladroitthief/lattice"
	"github.com/maladroitthief/mosaic"
	"github.com/rs/zerolog/log"
)

type GameScene struct {
	ctx      context.Context
	gridX    int
	gridY    int
	cellSize int

	world     *content.World
	playerId  caravan.GIDX
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
		backgroundColor: state.theme.Green(),
	}

	x := gs.gridX * level.RoomWidth
	y := gs.gridY * level.RoomHeight
	gs.world = content.NewWorld(
		ctx,
		core.NewECS(),
		bt.NewManager(),
		lattice.NewSpatialGrid[core.Entity](x, y, float64(gs.cellSize)),
	)

	gs.ai = server.NewAIServer()
	gs.state = server.NewStateServer()
	gs.physics = server.NewPhysicsServer(
		gs.world,
		float64(x),
		float64(y),
		float64(gs.cellSize),
	)
	gs.animation = server.NewAnimationServer()

	player := player.NewFederico(gs.world, 0, 0)
	gs.playerId = player.Id
	gs.cameraFocus = player

	level := level.NewLevel(
		level.NewRoomFactory(),
		level.NewBlockFactory(gs.world),
		player,
		gs.gridX,
		gs.gridY,
		gs.cellSize,
	)
	level.GenerateRooms()
	level.Render(gs.world.ECS)

	gs.physics.ResetGrid()

	focus, err := gs.world.ECS.GetPosition(gs.cameraFocus)
	if err != nil {
		log.Warn().Err(err).Any("focus", focus)
	}
	gs.camera = NewCamera(
		focus.X,
		focus.Y,
		mosaic.Vector{X: 200, Y: 200},
	)

	return gs
}

func (gs *GameScene) Update(state *SceneState) error {
	inputs := state.input.CurrentInputs()

	for _, input := range inputs {
		switch input {
		case core.InputMenu:
			return Termination
		}
	}

	gs.state.Update(gs.world.ECS)
	player, err := gs.world.ECS.GetEntity(gs.playerId)
	if err != nil {
		panic("")
	}
	ProcessPlayerGameInputs(gs.world.ECS, player, inputs)
	// s.ai.Update(s.world, inputs)
	gs.physics.Update(gs.world.ECS)
	gs.animation.Update(gs.world.ECS)

	return nil
}

func (gs *GameScene) Size() (width, height int) {
	width = gs.gridX * level.RoomWidth * gs.cellSize
	height = gs.gridY * level.RoomHeight * gs.cellSize

	return width, height
}

func (gs *GameScene) GetCanvasGrid() (rows, columns int) {
	return gs.gridY, gs.gridX
}

func (gs *GameScene) CellSize() int {
	return gs.cellSize
}

func (gs *GameScene) GetWorld() *content.World {
	return gs.world
}

func (gs *GameScene) BackgroundColor() color.Color {
	return gs.backgroundColor
}

func (gs *GameScene) GetCamera() *Camera {
	focus, err := gs.world.ECS.GetPosition(gs.cameraFocus)
	if err != nil {
		log.Warn().Err(err).Any("focus", focus)
	}

	gs.camera.Update(focus.X, focus.Y)

	return gs.camera
}
