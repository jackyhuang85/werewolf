package character

import "werewolf/game/helper"

type RoleGenerator func(abilityHelper helper.AbilityHelper) Role

var RoleIDToGenerator = map[string]RoleGenerator{
	"W":  WitchGenerator,
	"S":  SeerGenerator,
	"H":  HunterGenerator,
	"K":  KnightGenerator,
	"WW": WerewolfGenerator,
	"V":  VillagerGenerator,
}

type Role interface {
	IsTeamVillager() bool

	HasAbilityAtNight() bool
	GetNightAbilityPriority() int

	HasAbilityAtDay() bool
	IsAtDayAbilityPreemptive() bool

	HasAbilityAtDying() bool

	UseAbilityAtDay()
	UseAbilityAtNight()
	UseAbilityAtDying()
}

type RoleBase struct {
	abilityHelper helper.AbilityHelper

	nightAbilityPriority int
}

func (r RoleBase) GetNightAbilityPriority() int {
	return r.nightAbilityPriority
}

func (r RoleBase) UseAbilityAtDay() {
	//TODO implement me
	panic("implement me")
}

func (r RoleBase) UseAbilityAtNight() {
	//TODO implement me
	panic("implement me")
}

func (r RoleBase) UseAbilityAtDying() {
	//TODO implement me
	panic("implement me")
}

func (r RoleBase) IsTeamVillager() bool {
	//TODO implement me
	panic("implement me")
}

func (r RoleBase) HasAbilityAtNight() bool {
	//TODO implement me
	panic("implement me")
}

func (r RoleBase) HasAbilityAtDay() bool {
	//TODO implement me
	panic("implement me")
}

func (r RoleBase) IsAtDayAbilityPreemptive() bool {
	//TODO implement me
	panic("implement me")
}

func (r RoleBase) HasAbilityAtDying() bool {
	//TODO implement me
	panic("implement me")
}

func (r RoleBase) Ability() string {
	//TODO implement me
	panic("implement me")
}
