package game_test

import (
	"github.com/maladroitthief/entree/domain/settings"
	"github.com/maladroitthief/entree/service"
)

type settingsService struct {
}

func (svc *settingsService) Update(args service.Inputs) error {
	return nil
}

func (svc *settingsService) IsAny() bool {
	return false
}

func (svc *settingsService) IsPressed(i settings.Input) bool {
	return false
}

func (svc *settingsService) IsJustPressed(i settings.Input) bool {
	return false
}

func (svc *settingsService) GetCursor() (x, y int) {
	return 0, 0
}

func (svc *settingsService) CurrentInputs() []settings.Input {
	return []settings.Input{}
}

func (svc *settingsService) GetWindowHeight() int {
	return 0
}

func (svc *settingsService) GetWindowWidth() int {
	return 0
}

func (svc *settingsService) GetWindowTitle() string {
	return ""
}

func (svc *settingsService) GetScale() float64 {
	return 0
}
