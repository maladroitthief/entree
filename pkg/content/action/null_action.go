package action

import "github.com/maladroitthief/entree/pkg/engine/core"

type Null struct {

}

func (n *Null) Execute(a core.Actor) {
  return
}
