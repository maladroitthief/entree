package service_test

import (
	"github.com/maladroitthief/entree/domain/settings"
	"github.com/maladroitthief/entree/service"
)

type settingsService struct {
}

func (s *settingsService) Update(args service.Inputs) error {
	return nil
}

func (s *settingsService) IsAny() bool {
	return false
}

func (s *settingsService) IsPressed(i settings.Input) bool {
	return false
}

func (s *settingsService) IsJustPressed(i settings.Input) bool {
	return false
}

func (s *settingsService) GetCursor() (x, y int) {
	return 0, 0
}

func (s *settingsService) CurrentInputs() []settings.Input {
	return []settings.Input{}
}

func (s *settingsService) GetWindowHeight() int {
	return 0
}

func (s *settingsService) GetWindowWidth() int {
	return 0
}

func (s *settingsService) GetWindowTitle() string {
	return ""
}

func (s *settingsService) GetScale() float64 {
	return 0
}
