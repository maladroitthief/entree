package core

import (
	"github.com/maladroitthief/caravan"
)

const (
	DefaultPrimaryComboWindow   = 5
	DefaultSecondaryComboWindow = 5
	DefaultSpecialComboWindow   = 5
)

type (
	Skill    func() error
	SkillSet struct {
		Id       caravan.GIDX
		EntityId caravan.GIDX

		Primary              []Skill
		PrimaryIndex         int
		PrimaryComboWindow   int
		Secondary            []Skill
		SecondaryIndex       int
		SecondaryComboWindow int
		Special              []Skill
		SpecialIndex         int
		SpecialComboWindow   int
	}
)

func (ecs *ECS) NewSkillSet(primary, secondary, special []Skill) SkillSet {
	skillset := SkillSet{
		Id:                   ecs.skillsets.Allocate(),
		Primary:              primary,
		PrimaryIndex:         0,
		PrimaryComboWindow:   DefaultPrimaryComboWindow,
		Secondary:            secondary,
		SecondaryIndex:       0,
		SecondaryComboWindow: DefaultSecondaryComboWindow,
		Special:              special,
		SpecialIndex:         0,
		SpecialComboWindow:   DefaultSpecialComboWindow,
	}
	ecs.skillsets.Set(skillset.Id, skillset)

	return skillset
}

func (ecs *ECS) BindSkillSet(entity Entity, skillset SkillSet) Entity {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()
	ecs.skillsetMu.Lock()
	defer ecs.skillsetMu.Unlock()

	skillset.EntityId = entity.Id
	entity.SkillSetId = skillset.Id

	ecs.skillsets.Set(skillset.Id, skillset)
	ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetSkillSet(entity Entity) (SkillSet, error) {
	return ecs.GetSkillSetById(entity.SkillSetId)
}

func (ecs *ECS) GetSkillSetById(id caravan.GIDX) (SkillSet, error) {
	ecs.skillsetMu.RLock()
	defer ecs.skillsetMu.RUnlock()

	skillset := ecs.skillsets.Get(id)
	if !ecs.skillsets.IsLive(skillset.Id) {
		return skillset, ErrAttributeNotFound
	}

	return skillset, nil
}

func (ecs *ECS) GetAllSkillSets() []SkillSet {
	ecs.skillsetMu.RLock()
	defer ecs.skillsetMu.RUnlock()

	return ecs.skillsets.GetAll()
}

func (ecs *ECS) SetSkillSet(skillset SkillSet) {
	ecs.skillsetMu.Lock()
	defer ecs.skillsetMu.Unlock()

	ecs.skillsets.Set(skillset.Id, skillset)
}

func (ecs *ECS) SkillSetActive(skillset SkillSet) bool {
	return ecs.skillsets.IsLive(skillset.Id)
}
