package application

import (
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/scene"
)

type SceneService struct {
	repo scene.Repository
	log  logs.Logger

	settingsSvc *SettingsService
}

func NewSceneService(
	logger logs.Logger,
	repo scene.Repository,
	settingsSvc *SettingsService,
) *SceneService {
	if logger == nil {
		panic("nil scene logger")
	}

	if repo == nil {
		panic("nil scene repo")
	}

	if settingsSvc == nil {
		panic("nil settings service")
	}

	return &SceneService{
		repo:        repo,
		log:         logger,
		settingsSvc: settingsSvc,
	}
}

func (svc *SceneService) Update(args InputArgs) error {
	svc.log.Info("scene update", args)
  err := svc.settingsSvc.Update(args)
  if err != nil {
    return err
  }

	if svc.repo.GetTransitionCount() <= 0 {
		//return s.current.Update(
		//	&GameState{
		//		SceneManager: s,
		//		Input:        s.input,
		//	},
		//)
		return nil
	}

	svc.repo.SetTransitionCount(svc.repo.GetTransitionCount() - 1)

	if svc.repo.GetTransitionCount() > 0 {
		return nil
	}

	svc.repo.SetCurrentScene(svc.repo.GetNextScene())
	svc.repo.SetNextScene(nil)

	return nil
}

func (svc *SceneService) GoTo(s *scene.Scene) error {
	if svc.repo.GetCurrentScene() == nil {
		svc.repo.SetCurrentScene(s)
	} else {
		svc.repo.SetNextScene(s)
		svc.repo.SetTransitionCount(scene.TransitionMaxCount)
	}

	return nil
}
