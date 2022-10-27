package models

import (
	"fmt"
)

type Goban [19][19]uint8

const (
	// color
	// https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
)

// PrintPlayers prints the 2D goban representation with players in stdout
func (g *Goban) PrintPlayers() {
	fmt.Println("============================= [GOBAN players]")

	for _, line := range g {
		// For each line
		for _, cell := range line {
			// For each cell
			switch cell {
			case PLAYER_1:
				// If the player 1 is on the last position move
				fmt.Print(colorRed + "F " + colorReset)
				break
			case PLAYER_2:
				// If the player 2 is on the last position move
				fmt.Print(colorRed + "S " + colorReset)
				break
			default:
				fmt.Print(". ")
				break
			}
		}
		fmt.Println()
	}
}

// PrintProximity prints the 2D goban representation with proximity in stdout
func (g *Goban) PrintProximity(threshold *uint8) {
	fmt.Println("============================= [GOBAN proximity]")

	for _, line := range g {
		// For each line
		for _, cell := range line {
			// For each cell
			switch cell {
			case PLAYER_0:
				// If no player on the last position move
				fmt.Print("   .")
				break
			case PLAYER_1:
				// If the player 1 is on the last position move
				fmt.Print(colorRed + "   F" + colorReset)
				break
			case PLAYER_2:
				// If the player 2 is on the last position move
				fmt.Print(colorRed + "   S" + colorReset)
				break
			default:
				if threshold != nil && *threshold == cell {
					fmt.Print(colorYellow)
				}
				fmt.Printf("%4d", int(cell))
				fmt.Print(colorReset)

				break
			}
		}
		fmt.Println()
	}
}
