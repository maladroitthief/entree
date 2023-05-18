package player

import (
	"fmt"
	"strings"

	"github.com/maladroitthief/entree/domain/canvas"
)

type PlayerGraphicsComponent struct {
	Speed         float64
	Variant       int
	VariantMax    int
	VariantMaxLUT map[string]int
}

func NewPlayerGraphicsComponent(vmLUT map[string]int) *PlayerGraphicsComponent {
	pgc := &PlayerGraphicsComponent{
		Speed:         canvas.DefaultSpriteSpeed,
		Variant:       1,
		VariantMaxLUT: vmLUT,
	}

	return pgc
}

func (g *PlayerGraphicsComponent) Update(e canvas.Entity) {
	spriteName := []string{e.State()}
	if e.OrientationY() == canvas.South {
		spriteName = append(spriteName, "front")
	} else {
		spriteName = append(spriteName, "back")
	}

	if e.OrientationX() != canvas.Neutral {
		spriteName = append(spriteName, "side")
	}

	sprite := strings.Join(spriteName, "_")
	g.VariantMax = g.VariantMaxLUT[sprite]
	g.VariantUpdate(float64(e.StateCounter()))
	e.SetSprite(fmt.Sprintf("%s_%d", sprite, g.Variant))
}

func (g *PlayerGraphicsComponent) Receive(e canvas.Entity, msg, val string) {

}

func (g *PlayerGraphicsComponent) VariantUpdate(counter float64) {
	speed := counter / (g.Speed / float64(g.VariantMax))
	g.Variant = int(speed)%g.VariantMax + 1
}
