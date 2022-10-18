package structures

import (
	"fmt"
	"strings"
)

type Position struct {
	x uint16
	y uint16
}

type Game struct {
	// Received data
	captureCond		bool
	alignCond		bool
	hotseat			bool

	timeToSearch	float64 // 1 or 2
	depthToSearch	uint8 // min 1 max 10 (?)

	goban			[19][19]byte
	whoPlayed		byte // 1 or 2
	wherePlayed		Position
	scoreP1			uint8
	scoreP2			uint8

	// Send data
	time			string
	botStone		Position
	hints			[3]Position
}

func (g *Game) Print() {
	fmt.Println("Goban:")
	fmt.Println("align:", g.alignCond)
	fmt.Println("capture:", g.captureCond)
	fmt.Println("Last played:", g.whoPlayed)
	fmt.Println("Score P1:", g.scoreP1)
	fmt.Println("Score P2:", g.scoreP2)
	fmt.Println(strings.Repeat("-", 40))
	for _, arr := range g.goban {
		for _, val := range arr {
			fmt.Print(string(val) + " ")
		}
		fmt.Print("\n")
	}
	fmt.Println(strings.Repeat("-", 40))

}

func InitGame() *Game {
	game := &Game{}

	return game
}
