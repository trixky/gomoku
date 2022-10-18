package debug

import(
	"fmt"
	"strings"
	g "github.com/trixky/gomoku/game"
	s "github.com/trixky/gomoku/structures"
)

func Game() {
	var input string
	for true {
		fmt.Print("Please tell who starts (me/ai)\n>")
		fmt.Scanln(&input)
		if input != "me" && input != "ai" {
			fmt.Print("Wrong input\n\n")
			continue
		}
		break
	}
	game := s.InitGame()
	if input == "ai" {
		g.PlayFirst()
	}
	for true {
		for true {
			fmt.Print("Please input coordinates\n>")
			fmt.Scanln(&input)
			if len(strings.Fields(input)) != 2 {
				fmt.Print("Wrong input\n\n")
				continue
			}
			break
		}
		// update Goban
		// AI play w/ debug bool arg
	}
}