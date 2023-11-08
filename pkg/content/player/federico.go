package player

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func NewFederico(e *core.ECS) core.Entity {
	ai := attribute.NewAI(attribute.Input)
	state := attribute.NewState()

	position := attribute.NewPosition(100, 100, 1.6)
	movement := attribute.NewMovement()
	dimension := attribute.NewDimension(
		data.Vector{X: position.X, Y: position.Y},
		data.Vector{X: 16, Y: 16},
	)
	dimension.Offset = data.Vector{X: 0, Y: -6}
	collider := attribute.NewCollider()

	animation := attribute.NewAnimation("federico", "idle_front_1")
	animation.VariantMax = 6
	animation.Speed = 50

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
