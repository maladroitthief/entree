package player

import (
	"fmt"
	"strings"

	"github.com/maladroitthief/entree/domain/canvas"
)

var (
	spriteVariants = map[string]int{
		"idle_front":      6,
		"idle_front_side": 4,
		"idle_back":       6,
		"idle_back_side":  4,
		"move_front":      6,
		"move_front_side": 6,
		"move_back":       6,
		"move_back_side":  6,
	}
)

func NewPilot(
	input canvas.InputComponent,
	physics canvas.PhysicsComponent,
) *canvas.Entity {
	return &canvas.Entity{
		Width:             200,
		Height:            200,
		X:                 100,
		Y:                 100,
		DeltaX:            0,
		DeltaY:            0,
		Acceleration:      canvas.DefaultAcceleration,
		Deceleration:      canvas.DefaultDeceleration,
		MaxVelocity:       canvas.DefaultMaxVelocity,
		Mass:              canvas.DefaultMass,
		Sheet:             "pilot",
		Sprite:            "idle_front_1",
		SpriteSpeed:       5,
		SpriteVariant:     1,
		SpriteMaxVariants: 6,
		State:             "idle",
		StateCounter:      0,
		OrientationX:      canvas.Neutral,
		OrientationY:      canvas.South,
		Input:             input,
		Physics:           physics,
		Graphics:          &pilotGraphics{},
	}
}

type pilotGraphics struct {
}

func (g *pilotGraphics) Update(e *canvas.Entity) {
	spriteName := []string{e.State}
	if e.OrientationY == canvas.South {
		spriteName = append(spriteName, "front")
	} else {
		spriteName = append(spriteName, "front")
	}

	if e.OrientationX != canvas.Neutral {
		spriteName = append(spriteName, "side")
	}

	sprite := strings.Join(spriteName, "_")
	e.VariantUpdate()
	e.Sprite = fmt.Sprintf("%s_%d", sprite, e.SpriteVariant)
}
