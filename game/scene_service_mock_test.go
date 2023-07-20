package game_test

import (
	"errors"
	"image/color"

	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics"
	"github.com/maladroitthief/entree/domain/scene"
	"github.com/maladroitthief/entree/service"
)

type sceneService struct {
}

func (svc *sceneService) Update(args service.Inputs) error {
	if args.CursorX < 0 {
		return errors.New("bad cursor")
	}

	return nil
}

type focalPoint struct {
}

func (fp *focalPoint) Position() (x, y float64) {
	return x, y
}

func (svc *sceneService) GetCamera() scene.Camera {
	camera := scene.NewCamera(
		&focalPoint{},
		physics.Vector{X: 800, Y: 800},
	)

	return camera
}

func (svc *sceneService) GetCanvasSize() (int, int) {
	return 0, 0
}

func (svc *sceneService) GetCanvasCellSize() int {
	return 0
}

func (svc *sceneService) GetEntities() []canvas.Entity {
	return []canvas.Entity{}
}

func (svc *sceneService) GetBackgroundColor() color.Color {
	return color.Black
}

func (svc *sceneService) GoTo(s scene.Scene) error {
	return nil
}
