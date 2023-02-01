package character

import "werewolf/game/helper"

type Hunter struct {
	RoleBase
}

func HunterGenerator(helper helper.AbilityHelper) Role {
	return &Hunter{}
}
