package game

import "werewolf/common"

type AbilityHelperImpl struct {
	game *Game

	playerToBeKilled map[uint64]bool
	playerToBeSaved  map[uint64]bool
}

func (a AbilityHelperImpl) GetAllPlayers() {
	//TODO implement me
	panic("implement me")
}

func (a AbilityHelperImpl) GetAlivePlayers() {
	//TODO implement me
	panic("implement me")
}

func (a AbilityHelperImpl) GetTeam(playerId uint64) common.ConstantTeam {
	//TODO implement me
	panic("implement me")
}

func (a AbilityHelperImpl) KillPlayer(id uint64) {
	//TODO implement me
	panic("implement me")
}

func (a AbilityHelperImpl) SavePlayer(id uint64) {
	//TODO implement me
	panic("implement me")
}

func (a AbilityHelperImpl) Commit() {

}

func (a AbilityHelperImpl) Reset() {

}
