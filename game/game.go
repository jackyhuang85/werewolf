package game

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"werewolf/character"
	"werewolf/common"
)

func NewGameInstance() *Game {
	return &Game{}
}

type Game struct {
	mState State

	players      []*Player
	numOfPlayers uint64
}

func (g *Game) Init(config Config) error {
	err := g.setup(config)
	if err != nil {
		return err
	}
	g.mState = StateDayTime{g}
	return nil
}

func findFirstAliveWerewolf(players []*Player) *Player {
	for _, player := range players {
		if player.team == common.ConstantTeamWerewolf && !player.isDead {
			return player
		}
	}
	return nil
}

func (g *Game) Start() error {
	println("Game Start!")
	currentTurn := 0
	for {
		currentTurn++
		// change state to night
		g.changeGameState()
		// werewolf decide to kill who
		wolfPlayer := findFirstAliveWerewolf(g.players)
		wolfPlayer.UseAbilityAtNight()
		// each player(villager) uses their ability (according to the priority)
		aliveVillagers := g.getAliveVillagers()
		sort.Slice(aliveVillagers, func(i, j int) bool {
			return aliveVillagers[i].role.GetNightAbilityPriority() > aliveVillagers[j].role.GetNightAbilityPriority()
		})
		for _, villager := range aliveVillagers {
			if villager.isDead || !villager.HasAbilityAtNight() {
				continue
			}
		}
		// change state to day
		// check if having winner so far

		// each player says something / uses ability (need to check if any preemptive ability is going to be used)

		// check if having winner so far

		// voting

		// eliminating, use ability if the player has abilityAtEliminating

		// check if having winner so far
		break
	}
	return nil
}

func (g *Game) getAlivePlayers() []*Player {
	alivePlayers := make([]*Player, 0)
	for _, player := range g.players {
		if !player.isDead {
			alivePlayers = append(alivePlayers, player)
		}
	}
	return alivePlayers
}

func (g *Game) getAliveWerewolves() []*Player {
	alivePlayers := make([]*Player, 0)
	for _, player := range g.players {
		if !player.isDead && player.team == common.ConstantTeamWerewolf {
			alivePlayers = append(alivePlayers, player)
		}
	}
	return alivePlayers
}

func (g *Game) getAliveVillagers() []*Player {
	alivePlayers := make([]*Player, 0)
	for _, player := range g.players {
		if !player.isDead && player.team == common.ConstantTeamVillager {
			alivePlayers = append(alivePlayers, player)
		}
	}
	return alivePlayers
}

func (g *Game) setup(config Config) error {
	abilityHelper := &AbilityHelperImpl{
		game: g,
	}
	abilityHelper.Reset()

	// players
	numOfPlayer := uint64(0)
	for _, roleID := range config.RawRoles {
		roleGenerator := character.RoleIDToGenerator[roleID]
		if roleGenerator == nil {
			println("warning: invalid roleID: ", roleID, ", skip")
			continue
		}
		numOfPlayer++
		roleInstance := roleGenerator(abilityHelper)
		player := NewPlayer(numOfPlayer, roleInstance)
		g.players = append(g.players, &player)
	}

	// shuffle the ordering
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(int(numOfPlayer), func(i, j int) {
		g.players[i], g.players[j] = g.players[j], g.players[i]
	})

	return nil
}

func (g *Game) changeGameState() {
	fmt.Println("status before change:", g.mState.GetName())
	g.mState.Handle()
	fmt.Println("status after change:", g.mState.GetName())
}

func (g *Game) SetState(newState State) {
	g.mState = newState
}

type State interface {
	Handle()
	GetName() string
}

type StateDayTime struct {
	game *Game
}

func (s StateDayTime) GetName() string {
	return "StateDayTime"
}

func (s StateDayTime) Handle() {
	fmt.Println("changing status from day time")
	s.game.SetState(StateNightTime{game: s.game})
}

type StateNightTime struct {
	game *Game
}

func (s StateNightTime) GetName() string {
	return "StateNightTime"
}

func (s StateNightTime) Handle() {
	fmt.Println("changing status from night time")
	s.game.SetState(StateDayTime{game: s.game})
}
