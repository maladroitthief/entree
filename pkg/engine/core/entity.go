package core

import (
	"fmt"

	"github.com/maladroitthief/caravan"
)

type Entity struct {
	Name        string
	Id          caravan.GIDX
	AIId        caravan.GIDX
	CommandId   caravan.GIDX
	BehaviorId  caravan.GIDX
	StateId     caravan.GIDX
	MovementId  caravan.GIDX
	AnimationId caravan.GIDX
	PositionId  caravan.GIDX
	DimensionId caravan.GIDX
	ColliderId  caravan.GIDX
	FactionId   caravan.GIDX
}

func (ecs *ECS) NewEntity(name string) Entity {
	entity := Entity{
		Id: ecs.entityAllocator.Allocate(),
	}
	entity.Name = fmt.Sprintf("%v_%v", name, entity.Id.Info())

	ecs.SetEntity(entity)

	return entity
}

func (ecs *ECS) GetEntity(id caravan.GIDX) (Entity, error) {
	ecs.entityMu.RLock()
	defer ecs.entityMu.RUnlock()

	entity := ecs.entities.Get(id)
	if !ecs.entityAllocator.IsLive(entity.Id) {
		return entity, ErrEnityNotFound
	}

	return entity, nil
}

func (ecs *ECS) GetAllEntities() []Entity {
	ecs.entityMu.RLock()
	defer ecs.entityMu.RUnlock()

	return ecs.entities.GetAll(ecs.entityAllocator)
}

func (ecs *ECS) SetEntity(entity Entity) {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()

	ecs.entities = ecs.entities.Set(entity.Id, entity)
}

func (ecs *ECS) DestroyEntity(entity Entity) bool {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()

	return ecs.entityAllocator.Deallocate(caravan.GIDX(entity.Id))
}
