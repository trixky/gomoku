package models

import (
	"fmt"
	"math"
	"strconv"
)

type HeuristicGoban [19][19]int

func (hb *HeuristicGoban) Compute(contexts []Context) {
	for y := 0; y < 19; y++ {
		for x := 0; x < 19; x++ {
			hb[y][x] = math.MinInt
		}
	}

	for _, context := range contexts {
		hb[context.State.LastMove.Position.Y][context.State.LastMove.Position.X] = context.State.Beta
	}
}

// ToString convert the 2D heuristic goban to the string format
func (hb *HeuristicGoban) ToString() (str string) {
	for _, line := range hb {
		for _, cell := range line {
			if cell != math.MinInt {
				str += strconv.Itoa(cell)
			} else {
				str += "x"
			}

			str += ","
		}
	}

	str = str[:len(str)-1]

	return
}

// Print prints the 2D goban representation with heuristic in stdout
func (hb *HeuristicGoban) Print() {
	fmt.Println("============================= [HEURISTIC GOBAN]")

	for _, line := range hb {
		// For each line
		for _, cell := range line {
			// For each cell

			if cell != math.MinInt {
				fmt.Print(colorYellow)
				fmt.Printf("%4d", int(cell))
				fmt.Print(colorReset)
			} else {
				fmt.Printf("%4c", 'x')
			}

			break
		}
		fmt.Println()
	}
}
