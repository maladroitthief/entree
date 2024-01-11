package enemy

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func NewOnyawn(e *core.ECS, x, y float64) core.Entity {
	entity := e.NewEntity()
	state := e.NewState()

	ai := e.NewAI(core.Computer)
	root := e.Root()
	root, seq := e.RandomSequence(root)
	ai.RootBehavior = root.Id
	ai.ActiveBehavior = seq.Id
	seq, moveU := e.MovingUp(seq)
	seq, moveD := e.MovingDown(seq)
	seq, moveL := e.MovingLeft(seq)
	seq, moveR := e.MovingRight(seq)

	entity = e.BindBehavior(entity, root)
	entity = e.BindBehavior(entity, seq)
	entity = e.BindBehavior(entity, moveU)
	entity = e.BindBehavior(entity, moveD)
	entity = e.BindBehavior(entity, moveL)
	entity = e.BindBehavior(entity, moveR)

	position := e.NewPosition(x, y, 1.6)
	movement := e.NewMovement()
	dimension := e.NewDimension(
		data.Vector{X: position.X, Y: position.Y},
		data.Vector{X: 16, Y: 16},
	)
	dimension.Offset = data.Vector{X: 0, Y: -6}
	collider := e.NewCollider()

	animation := e.NewAnimation("onyawn", "idle_front_1")
	animation.VariantMax = 6
	animation.Speed = 50
	animation.Sprites = map[string][]string{
		"idle_front":      core.SpriteArray("idle_front", 2),
		"idle_front_side": core.SpriteArray("idle_front_side", 2),
		"idle_back":       core.SpriteArray("idle_back", 2),
		"idle_back_side":  core.SpriteArray("idle_front_side", 2),
		"move_front":      core.SpriteArray("move_front", 6),
		"move_front_side": core.SpriteArray("move_front_side", 6),
		"move_back":       core.SpriteArray("move_back", 6),
		"move_back_side":  core.SpriteArray("move_front_side", 6),
	}

	entity = e.BindAI(entity, ai)
	entity = e.BindState(entity, state)
	entity = e.BindPosition(entity, position)
	entity = e.BindMovement(entity, movement)
	entity = e.BindDimension(entity, dimension)
	entity = e.BindCollider(entity, collider)
	entity = e.BindAnimation(entity, animation)

	return entity
}
