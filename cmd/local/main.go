package main

import (
	"errors"
	"flag"
	"os"
	"strings"
	"werewolf/game"
)

var (
	roles      = flag.String("roles", "", "Required. Role config, to configure the scenario of the Werewolf Game.")
	printRoles = flag.Bool("roleID", false, "Print all ROLE_ID and corresponding character.`")
	help       = flag.Bool("help", false, "Print usage")
)

func printAllRoles() {
	roleTable := `
+---------+-----------+-------------+
| ROLE_ID | Character | Description |
+---------+-----------+-------------+
| S       | Seer      |             |
+---------+-----------+-------------+
| W       | Witch     |             |
+---------+-----------+-------------+
| V       | Villager  |             |
+---------+-----------+-------------+
| H       | Hunter    |             |
+---------+-----------+-------------+
| WW      | Werewolf  |             |
+---------+-----------+-------------+
| K       | Knight    |             |
+---------+-----------+-------------+`
	println(roleTable)
}

func printUsage() {
	usage := `Usage of ./werewolf
	-roles [ROLE_ID,...]
		Required. Role config, to configure the scenario of the Werewolf Game. 
		The length of ROLE_ID will be the number of players in the game.
		
		For example: to start a game with 6 players
			- 2 Werewolves (WW)
			- 1 Seer       (S)
			- 1 Witch 	   (W)
			- 2 Villagers  (V)

		Run the executable with -roles S,W,V,V,WW,WW (the ordering can be random)
		- $ ./werewolf -roles S,W,V,V,WW,WW
		
		See -roleID to check the ROLE_ID of each character.
	-roleID
		Print all ROLE_ID and corresponding character.`
	println(usage)
}

func main() {
	flag.Parse()
	if *printRoles {
		printAllRoles()
		os.Exit(0)
	}
	if *help {
		printUsage()
		os.Exit(0)
	}

	rolesInGame, err := parseRole(roles)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	g := game.NewGameInstance()

	_ = g.Init(game.Config{
		RawRoles: rolesInGame,
	})
	//err := g.Start()
	os.Exit(0)
}

func parseRole(input *string) ([]string, error) {
	if input == nil {
		return nil, errors.New("parse error: input is nil")
	}
	if len(*input) == 0 {
		return nil, errors.New("parse error: input is empty")
	}
	result := strings.Split(*input, ",")
	return result, nil
}
