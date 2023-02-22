package query

import (
	"github.com/maladroitthief/entree/common/decorator"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/sirupsen/logrus"
)

type UIEntities struct {
}

type UIEntitiesHandler decorator.QueryHandler[UIEntities, []Entity]

type uiEntitiesHandler struct {
}

func NewUIEntityHandler(logger *logrus.Entry) decorator.QueryHandler[UIEntities, []Entity] {
	return decorator.ApplyQueryDecorators[UIEntities, []Entity](
		uiEntitiesHandler{},
		logger,
	)
}

func (h uiEntitiesHandler) Handle(q UIEntities) (e []Entity, err error) {
	defer func() {
		logs.LogCommand("Update", q, err)
	}()

	// TODO Update all the game shit here
	return nil, nil
}
