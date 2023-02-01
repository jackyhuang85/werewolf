package game

import (
	"werewolf/character"
	"werewolf/common"
)

type Player struct {
	playerId   uint64
	playerName string

	isDead bool
	team   common.ConstantTeam

	role character.Role
}

func NewPlayer(id uint64, role character.Role) Player {
	player := Player{
		playerId: id,
		isDead:   false,
		team:     common.ConstantTeamWerewolf,
		role:     role,
	}

	if role.IsTeamVillager() {
		player.team = common.ConstantTeamVillager
	}

	return player
}

func (p *Player) UseAbilityAtNight() {
	if !p.role.HasAbilityAtNight() {
		return
	}
	p.role.UseAbilityAtNight()
}

func (p *Player) HasAbilityAtNight() bool {
	return p.role.HasAbilityAtNight()
}
