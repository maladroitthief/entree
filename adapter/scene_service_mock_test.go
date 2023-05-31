package adapter_test

import (
	"errors"
	"image/color"

	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/domain/canvas"
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

func (svc *sceneService) GetEntities() []canvas.Entity {
	return []canvas.Entity{}
}

func (svc *sceneService) GetBackgroundColor() color.Color {
	return color.Black
}

func (svc *sceneService) GoTo(s scene.Scene) error {
	return nil
}
