package application

import (
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/settings"
)

type GameService struct {
	repo         settings.Repository
	log          logs.Logger
	sceneService *SceneService
}

func NewGameService(
	logger logs.Logger,
	repo settings.Repository,
	scene *SceneService,
) *GameService {

	if logger == nil {
		panic("nil game logger")
	}

	if repo == nil {
		panic("nil game repo")
	}

	if scene == nil {
		panic("nil scene service")
	}

	return &GameService{
		repo:         repo,
		log:          logger,
		sceneService: scene,
	}
}

func (svc *GameService) Update() error {
	// TODO Update all the game shit here

	return nil
}

func (svc *GameService) GetWindowSettings() (settings.Window, error) {
	s, err := svc.repo.GetSettings()
	if err != nil {
		return settings.Window{}, err
	}

	return s.Window, nil
}
