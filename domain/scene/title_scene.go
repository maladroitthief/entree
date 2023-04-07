package scene

import "github.com/maladroitthief/entree/domain/canvas"

type TitleScene struct {
}

func (s *TitleScene) Update(state *GameState) error {
	// state.Log.Info("Title Scene", nil)
	if state.InputSvc.IsAny() {
		return state.SceneSvc.GoTo(NewGameScene(state))
	}

	return nil
}

func (s *TitleScene) GetEntities() []*canvas.Entity {
	return nil
}
