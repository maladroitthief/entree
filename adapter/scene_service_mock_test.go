package adapter_test

import (
	"errors"
	"image/color"

	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics"
	"github.com/maladroitthief/entree/domain/scene"
)

type sceneService struct {
}

func (svc *sceneService) Update(args application.Inputs) error {
	if args.CursorX < 0 {
		return errors.New("bad cursor")
	}

	return nil
}

type focalPoint struct {
	position physics.Vector
}

func (fp *focalPoint) Position() physics.Vector {
	return fp.position
}

func (svc *sceneService) GetCamera() scene.Camera {
	camera := scene.NewCamera(
		&focalPoint{
			position: physics.Vector{X: 0, Y: 0},
		},
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
