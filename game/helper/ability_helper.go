package helper

import "werewolf/common"

type AbilityHelper interface {
	GetAllPlayers()
	GetAlivePlayers()

	GetTeam(playerId uint64) common.ConstantTeam

	KillPlayer(id uint64)
	SavePlayer(id uint64)
}
