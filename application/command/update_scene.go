package command

import (
	"github.com/maladroitthief/entree/common/decorator"
	"github.com/sirupsen/logrus"
)

type UpdateScene struct {
}

type UpdateSceneHandler decorator.CommandHandler[UpdateScene]

type updateSceneHandler struct {
}

func NewUpdateSceneHandler(
	logger *logrus.Entry,
) decorator.CommandHandler[UpdateScene]{
	return decorator.ApplyCommandDecorators[UpdateScene](
		updateSceneHandler{},
		logger,
	)
}

func (h updateSceneHandler) Handle(cmd UpdateScene) (err error){
  // TODO: DO SHIT HERE
  return nil
}
