package query

import (
	"github.com/maladroitthief/entree/common/decorator"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/sirupsen/logrus"
)

type WindowSettings struct {
}

type WindowSettingsHandler decorator.QueryHandler[WindowSettings, Window]

type windowSettingsHandler struct {
  readModel WindowSettingsReadModel
}

func NewWindowSettingsHandler(
	logger *logrus.Entry,
) decorator.QueryHandler[WindowSettings, Window] {
	return decorator.ApplyQueryDecorators[WindowSettings, Window](
		windowSettingsHandler{},
		logger,
	)
}

type WindowSettingsReadModel interface {

}

func (h windowSettingsHandler) Handle(q WindowSettings) (w Window, err error) {
	defer func() {
		logs.LogCommand("Update", q, err)
	}()

	w = Window{
		Width:  1500,
		Height: 1500,
		Title:  "Entree",
	}

	return w, nil
}
