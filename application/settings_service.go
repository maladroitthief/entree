package application

import (
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/settings"
)

type SettingsService struct {
	repo settings.Repository
	log  logs.Logger

	inputSettings  *settings.InputSettings
	windowSettings *settings.WindowSettings

	currentKeys    []string
	currentCursorX int
	currentCursorY int
	currentInputs  []settings.Input
	inputStates    map[settings.Input]int
}

func NewSettingsService(
	logger logs.Logger,
	repo settings.Repository,
) *SettingsService {
	if logger == nil {
		panic("nil game logger")
	}

	if repo == nil {
		panic("nil settings repo")
	}

	return &SettingsService{
		repo: repo,
		log:  logger,
	}
}

func (svc *SettingsService) Update(args Inputs) error {
	if svc.inputStates == nil {
		svc.inputStates = map[settings.Input]int{}
	}

	if svc.windowSettings == nil {
		err := svc.getWindowSettings()

		if err != nil {
			return err
		}
	}

	if svc.inputSettings == nil {
		err := svc.getInputSettings()

		if err != nil {
			return err
		}
	}

	svc.currentKeys = args.Inputs
	svc.currentInputs = make([]settings.Input, len(svc.currentKeys))
	svc.currentCursorX = args.CursorX
	svc.currentCursorY = args.CursorY

	for i, k := range svc.inputSettings.Keyboard {
		for _, arg := range args.Inputs {
			if k == arg {
				svc.currentInputs = append(svc.currentInputs, i)
				svc.inputStates[i]++
				continue
			}

			svc.inputStates[i] = 0
		}
	}

	return nil
}

func (svc *SettingsService) getWindowSettings() error {
	ws, err := svc.repo.GetWindowSettings()
	if err != nil {
		return err
	}

	svc.windowSettings = &ws

	return nil
}

func (svc *SettingsService) getInputSettings() error {
	is, err := svc.repo.GetInputSettings()
	if err != nil {
		return err
	}

	svc.inputSettings = &is
	return nil
}

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

func (svc *SettingsService) GetWindowHeight() int {
	return svc.windowSettings.Height
}

func (svc *SettingsService) GetWindowWidth() int {
	return svc.windowSettings.Width
}

func (svc *SettingsService) GetWindowTitle() string {
	return svc.windowSettings.Title
}

func (svc *SettingsService) GetScale() float64 {
	return svc.windowSettings.Scale
}
