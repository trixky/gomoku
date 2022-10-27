package models

import "fmt"

type Goban [19][19]uint8

// Print prints a 2D goban in stdout
func (g *Goban) Print() {
	fmt.Println("============================= [GOBAN]")

	for _, line := range g {
		// For each line
		for _, cell := range line {
			// For each cell
			switch cell {
			case 254:
				fmt.Print("1 ")
				break
			case 255:
				fmt.Print("2 ")
				break
			default:
				fmt.Print(". ")
				break
			}
		}
		fmt.Println()
	}
}
