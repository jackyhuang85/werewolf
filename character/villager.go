package character

import "werewolf/game/helper"

type Villager struct {
	RoleBase
}

func VillagerGenerator(helper helper.AbilityHelper) Role {
	return &Villager{}
}
