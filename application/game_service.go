package application

import (
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/game"
	"github.com/maladroitthief/entree/domain/settings"
)

type GameService struct {
	gameRepo     game.Repository
	settingsRepo settings.Repository
	log          logs.Logger
}

func NewGameService(
	logger logs.Logger,
	gameRepo game.Repository,
	settingsRepo settings.Repository,
) *GameService {

	if logger == nil {
		panic("nil game logger")
	}

	if gameRepo == nil {
		panic("nil game repo")
	}

	if settingsRepo == nil {
		panic("nil settings repo")
	}

	return &GameService{
		gameRepo:     gameRepo,
		settingsRepo: settingsRepo,
		log:          logger,
	}
}

func (svc *GameService) Update() error {
	// TODO Update all the game shit here

	return nil
}

func (svc *GameService) GetWindowSettings() (settings.Window, error) {
	s, err := svc.settingsRepo.GetSettings()
	if err != nil {
		return settings.Window{}, err
	}

	return s.Window, nil
}
