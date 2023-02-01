package character

import "werewolf/game/helper"

type Seer struct {
	RoleBase
}

func SeerGenerator(helper helper.AbilityHelper) Role {
	return &Seer{RoleBase{
		abilityHelper:        helper,
		nightAbilityPriority: NightAbilityPriorityMap["S"],
	}}
}
