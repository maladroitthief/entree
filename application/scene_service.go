package application

import (
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/scene"
)

const (
	transitionCountMax = 20
)

type SceneService struct {
	log             logs.Logger
	currentScene    scene.Scene
	nextScene       scene.Scene
	transitionCount int

	settingsSvc *SettingsService
	graphicsSvc *GraphicsService
}

func NewSceneService(
	logger logs.Logger,
	settingsSvc *SettingsService,
  graphicsSvc *GraphicsService,
) *SceneService {
	if logger == nil {
		panic("nil scene logger")
	}

	if settingsSvc == nil {
		panic("nil settings service")
	}

	if graphicsSvc == nil {
		panic("nil graphics service")
	}

	return &SceneService{
		log:         logger,
		settingsSvc: settingsSvc,
    graphicsSvc: graphicsSvc,
	}
}

func (svc *SceneService) Update(args InputArgs) error {
	// Update Settings
	err := svc.settingsSvc.Update(args)
	if err != nil {
		return err
	}

	if svc.currentScene == nil {
		svc.GoTo(&scene.TitleScene{})
	}

	if svc.transitionCount <= 0 {
		return svc.currentScene.Update(
			&scene.GameState{
				Log:      svc.log,
				SceneSvc: svc,
				InputSvc: svc.settingsSvc,
			},
		)
	}

	svc.transitionCount--

	if svc.transitionCount > 0 {
		return nil
	}

	svc.currentScene = svc.nextScene
	svc.nextScene = nil

	return nil
}

func (svc *SceneService) GoTo(s scene.Scene) error {
	if svc.currentScene == nil {
		svc.currentScene = s
	} else {
		svc.nextScene = s
		svc.transitionCount = transitionCountMax
	}

	return nil
}
