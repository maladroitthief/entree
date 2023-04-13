package application

import (
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/settings"
)

type SettingsService struct {
	repo settings.Repository
	log  logs.Logger

	currentKeys    []string
	currentCursorX int
	currentCursorY int
	currentInputs  []settings.Input
	inputStates    map[settings.Input]int
	inputSettings  *settings.InputSettings
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

	if svc.inputSettings == nil {
		is, err := svc.repo.GetInputSettings()
		if err != nil {
			return err
		}

		svc.inputSettings = &is
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

func (svc *SettingsService) GetWindowSettings() (settings.WindowSettings, error) {
	ws, err := svc.repo.GetWindowSettings()
	if err != nil {
		return settings.WindowSettings{}, err
	}

	return ws, nil
}
