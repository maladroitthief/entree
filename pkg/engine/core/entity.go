package core

import "github.com/maladroitthief/entree/common/data"

type Entity struct {
	Id          data.GenerationalIndex
	AIId        data.GenerationalIndex
	CommandId   data.GenerationalIndex
	BehaviorId  data.GenerationalIndex
	StateId     data.GenerationalIndex
	MovementId  data.GenerationalIndex
	AnimationId data.GenerationalIndex
	PositionId  data.GenerationalIndex
	DimensionId data.GenerationalIndex
	ColliderId  data.GenerationalIndex
	FactionId   data.GenerationalIndex
}

func (ecs *ECS) NewEntity() Entity {
	entity := Entity{
		Id: ecs.entityAllocator.Allocate(),
	}

	ecs.SetEntity(entity)

	return entity
}

func (ecs *ECS) GetEntity(id data.GenerationalIndex) (Entity, error) {
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

	return ecs.entityAllocator.Deallocate(data.GenerationalIndex(entity.Id))
}
