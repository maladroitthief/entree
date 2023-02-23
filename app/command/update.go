package command

import (
	"github.com/maladroitthief/entree/common/decorator"
	"github.com/sirupsen/logrus"
)

type Update struct {
	CursorX int
	CursorY int
	Inputs  []string
}

type UpdateHandler decorator.CommandHandler[Update]

type updateHandler struct {
}

func NewUpdateHandler(logger *logrus.Entry) decorator.CommandHandler[Update] {
	return decorator.ApplyCommandDecorators[Update](
		updateHandler{},
		logger,
	)
}

func (h updateHandler) Handle(cmd Update) (err error) {
	// TODO Update all the game shit here
	return nil
}
