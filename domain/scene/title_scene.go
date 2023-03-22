package scene

import "github.com/maladroitthief/entree/domain/canvas"

type TitleScene struct {
}

func (s *TitleScene) Update(state *GameState) error {
	state.Log.Info("Title Scene", nil)
	if state.InputSvc.IsAny() {
		state.SceneSvc.GoTo(NewGameScene())
	}

	return nil
}

func (s *TitleScene) Draw() []*canvas.Entity {
	return nil
}
