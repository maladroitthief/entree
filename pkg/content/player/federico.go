package player

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func NewFederico(e *core.ECS) core.Entity {
	ai := attribute.NewAI(attribute.Input)
	state := attribute.NewState()

	position := attribute.NewPosition(data.Vector{X: 100, Y: 100})
	movement := attribute.NewMovement()
	dimension := attribute.NewDimension(position.Position, data.Vector{X: 12, Y: 18})
	dimension.Offset = data.Vector{X: 0, Y: -6}
	collider := attribute.NewCollider()

	animation := attribute.NewAnimation("federico", "idle_front_1", 0.8)
	animation.VariantMax = 6

	entity := e.NewEntity()
	entity = e.AddAI(entity, ai)
	entity = e.AddState(entity, state)
	entity = e.AddPosition(entity, position)
	entity = e.AddMovement(entity, movement)
	entity = e.AddDimension(entity, dimension)
	entity = e.AddCollider(entity, collider)
	entity = e.AddAnimation(entity, animation)

	return entity
}
