package adapter

import (
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/settings"
)

type GameAdapter struct {
	log      logs.Logger
	gameSvc  *application.GameService
	sceneSvc *application.SceneService
}

type UpdateGame struct {
	CursorX int
	CursorY int
	Inputs  []string
}

func NewGameAdapter(
	log logs.Logger,
	gameSvc *application.GameService,
	sceneSvc *application.SceneService,
) *GameAdapter {
	ga := GameAdapter{
		log:     log,
		gameSvc: gameSvc,
		sceneSvc: sceneSvc,
	}

	if gameSvc == nil {
		panic("nil game service")
	}

	if sceneSvc == nil {
		panic("nil scene service")
	}

	return &ga
}

func (ga *GameAdapter) Update(args UpdateGame) error {

	return nil
}

func (ga *GameAdapter) Layout(width, height int) (screenWidth, screenHeight int) {
	ws, err := ga.gameSvc.GetWindowSettings()
	if err != nil {
		args := struct{ width, height int }{width, height}
		ga.log.Fatal("Layout", args, err)
	}
	return ws.Width, ws.Height
}

func (ga *GameAdapter) GetWindowSettings() (settings.Window, error) {
	return ga.gameSvc.GetWindowSettings()
}
