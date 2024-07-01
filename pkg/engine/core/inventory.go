package core

import (
	"github.com/maladroitthief/caravan"
	"github.com/maladroitthief/mosaic"
)

type Inventory struct {
	Id       caravan.GIDX
	EntityId caravan.GIDX

	Size    mosaic.Vector
	Scale   float64
	Offset  mosaic.Vector
	Polygon mosaic.Polygon
}

func (ecs *ECS) NewInventory() Inventory {
	inventory := Inventory{
		Id: ecs.inventories.Allocate(),
	}
	ecs.inventories.Set(inventory.Id, inventory)

	return inventory
}

func (ecs *ECS) BindInventory(entity Entity, inventory Inventory) Entity {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()
	ecs.inventoryMu.Lock()
	defer ecs.inventoryMu.Unlock()

	inventory.EntityId = entity.Id
	entity.InventoryId = inventory.Id

	ecs.inventories.Set(inventory.Id, inventory)
	ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetInventory(entity Entity) (Inventory, error) {
	return ecs.GetInventoryById(entity.InventoryId)
}

func (ecs *ECS) GetInventoryById(id caravan.GIDX) (Inventory, error) {
	ecs.inventoryMu.RLock()
	defer ecs.inventoryMu.RUnlock()

	inventory := ecs.inventories.Get(id)
	if !ecs.inventories.IsLive(inventory.Id) {
		return inventory, ErrAttributeNotFound
	}

	return inventory, nil
}

func (ecs *ECS) GetAllInventories() []Inventory {
	ecs.inventoryMu.RLock()
	defer ecs.inventoryMu.RUnlock()

	return ecs.inventories.GetAll()
}

func (ecs *ECS) SetInventory(inventory Inventory) {
	ecs.inventoryMu.Lock()
	defer ecs.inventoryMu.Unlock()

	ecs.inventories.Set(inventory.Id, inventory)
}

func (ecs *ECS) InventoryActive(inventory Inventory) bool {
	return ecs.inventories.IsLive(inventory.Id)
}
