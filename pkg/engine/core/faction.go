package core

import "github.com/maladroitthief/entree/common/data"

const (
	Human Archetype = 1 << iota
	Vegetable
)

type (
	Archetype byte

	Faction struct {
		Id       data.GenerationalIndex
		EntityId data.GenerationalIndex

		Archetype Archetype
	}
)

func (e *ECS) NewFaction(a Archetype) Faction {
	faction := Faction{
		Id:        e.factionAllocator.Allocate(),
		Archetype: a,
	}
	e.factions.Set(faction.Id, faction)

	return faction
}

func (e *ECS) BindFaction(entity Entity, faction Faction) Entity {
	faction.EntityId = entity.Id
	entity.FactionId = faction.Id

	e.factions = e.factions.Set(faction.Id, faction)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetFaction(entity Entity) (Faction, error) {
	return e.GetFactionById(entity.FactionId)
}
func (e *ECS) GetFactionById(id data.GenerationalIndex) (Faction, error) {
	faction := e.factions.Get(id)
	if !e.factionAllocator.IsLive(faction.Id) {
		return faction, ErrAttributeNotFound
	}

	return faction, nil
}

func (e *ECS) GetAllFactions() []Faction {
	return e.factions.GetAll(e.factionAllocator)
}

func (e *ECS) SetFaction(faction Faction) {
	e.factions = e.factions.Set(faction.Id, faction)
}

func (e *ECS) SetArchetype(faction Faction, archetype Archetype) {
	faction.Archetype |= archetype
	e.SetFaction(faction)
}

func (e *ECS) UnsetArchetype(faction Faction, archetype Archetype) {
	faction.Archetype &= ^archetype
	e.SetFaction(faction)
}

func (faction Faction) IsArchetype(archetype Archetype) bool {
	return faction.Archetype&archetype != 0
}
