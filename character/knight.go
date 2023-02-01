package character

import "werewolf/game/helper"

type Knight struct {
	RoleBase
}

func KnightGenerator(helper helper.AbilityHelper) Role {
	return &Knight{}
}
