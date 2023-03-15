package adapter

import (
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/settings"
)

type GameAdapter struct {
	log         logs.Logger
	sceneSvc    *application.SceneService
	settingsSvc *application.SettingsService
}

type UpdateArgs struct {
	CursorX int
	CursorY int
	Inputs  []string
}

func NewGameAdapter(
	log logs.Logger,
	sceneSvc *application.SceneService,
	settingsSvc *application.SettingsService,
) *GameAdapter {
	ga := GameAdapter{
		log:         log,
		settingsSvc: settingsSvc,
		sceneSvc:    sceneSvc,
	}

	if sceneSvc == nil {
		panic("nil scene service")
	}

	if settingsSvc == nil {
		panic("nil settings service")
	}

	return &ga
}

func (ga *GameAdapter) Update(args UpdateArgs) error {
	return ga.sceneSvc.Update(application.InputArgs{
    CursorX: args.CursorX,
    CursorY: args.CursorY,
    Inputs: args.Inputs,
  })
}

func (ga *GameAdapter) Layout(width, height int) (screenWidth, screenHeight int) {
	ws, err := ga.settingsSvc.GetWindowSettings()
	if err != nil {
		args := struct{ width, height int }{width, height}
		ga.log.Fatal("Layout", args, err)
	}
	return ws.Width, ws.Height
}

func (ga *GameAdapter) GetWindowSettings() (settings.WindowSettings, error) {
	return ga.settingsSvc.GetWindowSettings()
}
