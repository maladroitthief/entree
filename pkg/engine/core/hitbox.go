package core

import "github.com/maladroitthief/caravan"

type (
	Hitbox struct {
		Id         caravan.GIDX
		HitboxesId caravan.GIDX

		PositionId  caravan.GIDX
		DimensionId caravan.GIDX
		ColliderId  caravan.GIDX

		OffsetX float64
		OffsetY float64
		OffsetZ float64
	}

	Hitboxes struct {
		Id       caravan.GIDX
		EntityId caravan.GIDX

		hitboxes []caravan.GIDX
	}
)
