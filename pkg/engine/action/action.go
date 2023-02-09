package action

import "github.com/maladroitthief/entree/pkg/engine/core"

type Action interface {
	Execute(core.Actor)
}
