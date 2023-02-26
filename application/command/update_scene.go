package command

import (
	"github.com/maladroitthief/entree/common/decorator"
	"github.com/maladroitthief/entree/domain/scene"
	"github.com/sirupsen/logrus"
)

type UpdateScene struct {
}

type UpdateSceneHandler decorator.CommandHandler[UpdateScene]

type updateSceneHandler struct {
	repo scene.Repository
}

func NewUpdateSceneHandler(
	logger *logrus.Entry,
	repo scene.Repository,
) decorator.CommandHandler[UpdateScene] {
  if repo == nil {
    panic("nil scene repo")
  }

	return decorator.ApplyCommandDecorators[UpdateScene](
		updateSceneHandler{
      repo: repo,
    },
		logger,
	)
}

func (h updateSceneHandler) Handle(cmd UpdateScene) (err error) {
	// TODO: DO SHIT HERE
	return nil
}
