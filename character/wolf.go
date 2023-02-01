package character

import (
	"fmt"
	"werewolf/game/helper"
)

type Werewolf struct {
	RoleBase
}

func WerewolfGenerator(helper helper.AbilityHelper) Role {
	return &Werewolf{}
}

func (w *Werewolf) UseAbilityAtNight() {
	println("Werewolves, who do you want to kill? (input PlayerID)")
	var i int

	for {
		_, err := fmt.Scanf("%v", &i)
		if err != nil {
			continue
		}
		// check if can kill or not

		break
	}
}
