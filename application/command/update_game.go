package command

import (
	"github.com/maladroitthief/entree/common/decorator"
	"github.com/sirupsen/logrus"
)

type UpdateGame struct {
  CursorX int
	CursorY int
	Inputs  []string
}

type UpdateGameHandler decorator.CommandHandler[UpdateGame]

type updateGameHandler struct {
	sceneService SceneService
}

func NewUpdateGameHandler(
  logger *logrus.Entry,
  sceneService SceneService,
) decorator.CommandHandler[UpdateGame] {
	return decorator.ApplyCommandDecorators[UpdateGame](
		updateGameHandler{
      sceneService: sceneService,
    },
		logger,
	)
}

func (h updateGameHandler) Handle(cmd UpdateGame) (err error) {
	// TODO Update all the game shit here
	return nil
}
