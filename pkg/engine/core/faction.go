package core

import "github.com/maladroitthief/entree/common/data"

const (
	Human Archetype = 1 << iota
	Vegetable
	Fruit
)

type (
	Archetype byte

	Faction struct {
		Id       data.GenerationalIndex
		EntityId data.GenerationalIndex

		Archetype Archetype
	}
)

func (ecs *ECS) NewFaction(a Archetype) Faction {
	faction := Faction{
		Id:        ecs.factionAllocator.Allocate(),
		Archetype: a,
	}
	ecs.factions.Set(faction.Id, faction)

	return faction
}

func (ecs *ECS) BindFaction(entity Entity, faction Faction) Entity {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()
	ecs.factionMu.Lock()
	defer ecs.factionMu.Unlock()

	faction.EntityId = entity.Id
	entity.FactionId = faction.Id

	ecs.factions = ecs.factions.Set(faction.Id, faction)
	ecs.entities = ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetFaction(entity Entity) (Faction, error) {
	return ecs.GetFactionById(entity.FactionId)
}
func (ecs *ECS) GetFactionById(id data.GenerationalIndex) (Faction, error) {
	ecs.factionMu.RLock()
	defer ecs.factionMu.RUnlock()

	faction := ecs.factions.Get(id)
	if !ecs.factionAllocator.IsLive(faction.Id) {
		return faction, ErrAttributeNotFound
	}

	return faction, nil
}

func (ecs *ECS) GetAllFactions() []Faction {
	ecs.factionMu.RLock()
	defer ecs.factionMu.RUnlock()

	return ecs.factions.GetAll(ecs.factionAllocator)
}

func (ecs *ECS) SetFaction(faction Faction) {
	ecs.factionMu.Lock()
	defer ecs.factionMu.Unlock()

	ecs.factions = ecs.factions.Set(faction.Id, faction)
}

func (a Archetype) Set(archetype Archetype) Archetype {
	a |= archetype
	return a
}

func (a Archetype) Unset(archetype Archetype) Archetype {
	a &= ^archetype
	return a
}

func (a Archetype) Check(archetype Archetype) bool {
	return a&archetype != 0
}

func (ecs *ECS) SetArchetype(faction Faction, archetype Archetype) {
	ecs.factionMu.Lock()
	defer ecs.factionMu.Unlock()

	faction.Archetype = faction.Archetype.Set(archetype)
	ecs.SetFaction(faction)
}

func (ecs *ECS) UnsetArchetype(faction Faction, archetype Archetype) {
	ecs.factionMu.Lock()
	defer ecs.factionMu.Unlock()

	faction.Archetype = faction.Archetype.Unset(archetype)
	ecs.SetFaction(faction)
}

func (faction Faction) IsArchetype(archetype Archetype) bool {
	return faction.Archetype.Check(archetype)
}
