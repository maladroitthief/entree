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

func (e *ECS) NewEntity() Entity {
	entity := Entity{
		Id: e.entityAllocator.Allocate(),
	}

	e.SetEntity(entity)

	return entity
}

func (e *ECS) GetEntity(id data.GenerationalIndex) (Entity, error) {
	entity := e.entities.Get(id)
	if !e.entityAllocator.IsLive(entity.Id) {
		return entity, ErrEnityNotFound
	}

	return entity, nil
}

func (e *ECS) GetAllEntities() []Entity {
	return e.entities.GetAll(e.entityAllocator)
}

func (e *ECS) SetEntity(entity Entity) {
	e.entities = e.entities.Set(entity.Id, entity)
}

func (e *ECS) DestroyEntity(entity Entity) bool {
	return e.entityAllocator.Deallocate(data.GenerationalIndex(entity.Id))
}
