package models

import (
	"fmt"
	"strings"
)

type Game struct {
	// Received data
	captureCond bool
	alignCond   bool
	hotseat     bool

	timeToSearch  int   // in ms
	depthToSearch uint8 // min 2 max 10

	goban       [19][19]byte
	whoPlayed   byte // 5 or 6
	wherePlayed Position
	scoreP1     uint8
	scoreP2     uint8

	// Send data
	time   string
	eval   int
	depth  int
	stones [3]Position
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
		fmt.Print(arr)
	}
	fmt.Println(strings.Repeat("-", 40))

}
