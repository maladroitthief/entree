package application

import (
	"image/color"

	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/common/theme"
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/scene"
)

const (
	transitionCountMax = 5
)

type SceneService struct {
	log             logs.Logger
	currentScene    scene.Scene
	nextScene       scene.Scene
	transitionCount int
	theme           theme.Colors

	settingsSvc *SettingsService
}

func NewSceneService(
	logger logs.Logger,
	settingsSvc *SettingsService,
) *SceneService {
	if logger == nil {
		panic("nil scene logger")
	}

	if settingsSvc == nil {
		panic("nil settings service")
	}

	return &SceneService{
		log:         logger,
		settingsSvc: settingsSvc,
		theme:       &theme.TokyoNight{},
	}
}

func (svc *SceneService) Update(args Inputs) error {
	// Update Settings
	err := svc.settingsSvc.Update(args)
	if err != nil {
		return err
	}

	if svc.currentScene == nil {
		err = svc.GoTo(scene.NewTitleScene(
			&scene.GameState{
				Log:      svc.log,
				SceneSvc: svc,
				InputSvc: svc.settingsSvc,
				Theme:    svc.theme,
			},
		))
	}

	if err != nil {
		return err
	}

	if svc.transitionCount <= 0 {
		return svc.currentScene.Update(
			&scene.GameState{
				Log:      svc.log,
				SceneSvc: svc,
				InputSvc: svc.settingsSvc,
				Theme:    svc.theme,
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

func (svc *SceneService) GetEntities() []canvas.Entity {
	return svc.currentScene.GetEntities()
}

func (svc *SceneService) GetBackgroundColor() color.Color {
	return svc.currentScene.GetBackgroundColor()
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
