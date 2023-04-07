package application

import "github.com/maladroitthief/entree/domain/settings"

func (svc *SettingsService) IsAny() bool {
	return len(svc.currentKeys) > 0
}

func (svc *SettingsService) IsPressed(i settings.Input) bool {
	return svc.inputStates[i] >= 1
}

func (svc *SettingsService) IsJustPressed(i settings.Input) bool {
	return svc.inputStates[i] == 1
}

func (svc *SettingsService) GetCursor() (x, y int) {
	return svc.currentCursorX, svc.currentCursorY
}

func (svc *SettingsService) CurrentInputs() []settings.Input {
  return svc.currentInputs
}
