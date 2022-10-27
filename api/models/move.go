package models

import "fmt"

type Move struct {
	Player   bool // false: player 1 | true: player 2
	Position Position
}

// Print prints move attributes
func (m *Move) Print() {
	fmt.Println("--------------- [moves]")

	// Prints player
	fmt.Print("player:\t ")
	if m.Player {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}

	// Prints position
	fmt.Println("x:\t\t", m.Position.X)
	fmt.Println("y:\t\t", m.Position.X)
}
