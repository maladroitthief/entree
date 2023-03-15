package scene

type GameScene struct {
}

func NewGameScene() *GameScene {
	gs := &GameScene{}

	return gs
}

func (s *GameScene) Update(state *GameState) error {
	state.Log.Info("Game Scene", nil)
	return nil
}
