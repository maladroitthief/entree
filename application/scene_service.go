package application

import (
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/scene"
)

type SceneService struct {
	repo scene.Repository
	log  logs.Logger
}

func NewSceneService(logger logs.Logger, repo scene.Repository) *SceneService {
	if logger == nil {
		panic("nil scene logger")
	}

	if repo == nil {
		panic("nil scene repo")
	}

	return &SceneService{
		repo: repo,
		log:  logger,
	}
}

func (svc *SceneService) Update() error {
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
