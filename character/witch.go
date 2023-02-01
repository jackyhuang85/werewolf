package character

import "werewolf/game/helper"

type Witch struct {
	RoleBase
}

func WitchGenerator(helper helper.AbilityHelper) Role {
	return &Witch{RoleBase{
		abilityHelper:        helper,
		nightAbilityPriority: NightAbilityPriorityMap["W"],
	}}
}
