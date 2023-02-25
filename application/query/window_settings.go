package query

import (
	"github.com/maladroitthief/entree/common/decorator"
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
	w = Window{
		Width:  1920,
		Height: 1080,
		Title:  "Entree",
	}

	return w, nil
}
