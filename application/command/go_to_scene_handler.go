package command

import (
	"github.com/maladroitthief/entree/common/decorator"
	"github.com/sirupsen/logrus"
)

type GoToScene struct {
}

type GoToSceneHandler decorator.CommandHandler[GoToScene]

type goToSceneHandler struct {
}

func NewGoToSceneHandler(
	logger *logrus.Entry,
) decorator.CommandHandler[GoToScene] {
	return decorator.ApplyCommandDecorators[GoToScene](
		goToSceneHandler{},
		logger,
	)
}

func (h goToSceneHandler) Handle(cmd GoToScene) (err error) {
	// TODO: DO SHIT HERE
	return nil
}
