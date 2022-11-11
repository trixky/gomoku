package models

import (
	"fmt"

	"github.com/trixky/gomoku/utils"
)

type Goban [19][19]uint8

const (
	// color
	// https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
)

// ComputeGlobalProximity computes the global proximity
func (g *Goban) ComputeGlobalProximity(threshold uint8, radius uint8, shape byte, suspicious_radius int, last_move Move) {
	position := Position{}

	for x := uint8(0); x < 19; x++ {
		// For each line
		for y := uint8(0); y < 19; y++ {
			// For each cell

			if suspicious_radius == 0 || (int(x) >= int(last_move.Position.X)-suspicious_radius && int(x) <= int(last_move.Position.X)+suspicious_radius && int(y) >= int(last_move.Position.Y)-suspicious_radius && int(y) <= int(last_move.Position.Y)+suspicious_radius) {
				// If the suspicious radius in disable or respected
				if g[y][x] >= PLAYER_1 {
					// If the cell is taken by a player
					position.X = x
					position.Y = y
					g.ComputePieceProximity(&position, threshold, radius, shape)
				}
			}
		}
	}
}

// ComputeGlobalProximity computes a local proximity
func (g *Goban) ComputePieceProximity(position *Position, threshold uint8, radius uint8, shape byte) {
	radius_plus := radius + 1
	radius_end := 19 - radius

	// For each cell
	// Compute sub xx start
	var xx_start uint8
	if position.X > radius {
		xx_start = position.X - radius
	} else {
		xx_start = 0
	}

	// Compute sub yy start
	var yy_start uint8
	if position.Y > radius {
		yy_start = position.Y - radius
	} else {
		yy_start = 0
	}

	// // Compute sub xx end
	var xx_end uint8
	if position.X < radius_end {
		xx_end = position.X + radius_plus
	} else {
		xx_end = 19
	}

	// // Compute sub yy end
	var yy_end uint8
	if position.Y < radius_end {
		yy_end = position.Y + radius_plus
	} else {
		yy_end = 19
	}

	for yy := yy_start; yy < yy_end; yy++ {
		// For each line of local radius
		for xx := xx_start; xx < xx_end; xx++ {
			// For each cell of local radius

			// Compute x and y diffs
			diff_x := utils.DiffUint8(position.X, xx)
			diff_y := utils.DiffUint8(position.Y, yy)

			if g[yy][xx] < PLAYER_1 {
				// If the cell is not taken by a player
				if shape == SHAPE_SQUARE || diff_x == diff_y || diff_x == 0 || diff_y == 0 {
					// If the cell is in the shape
					g[yy][xx] += radius + 1 - utils.MaxUint8(diff_x, diff_y)
					g[yy][xx] = utils.MinUint8(g[yy][xx], threshold)
				}
			}
		}
	}
}

// IsBlocked check if is blocked by no proximity
func (g *Goban) IsBlocked(threshold uint8) bool {
	for x := uint8(0); x < 19; x++ {
		// For each line
		for y := uint8(0); y < 19; y++ {
			// For each cell

			if g[y][x] >= threshold && g[y][x] < PLAYER_1 {
				// If the goban has at least one active cell
				// Is not blocked

				return false
			}
		}
	}

	// No active cell have be find
	return true
}

// ToString convert the 2D goban to the string format
func (g *Goban) ToString() (str string) {
	for _, line := range g {
		for _, cell := range line {
			switch cell {
			case PLAYER_1:
				str += "1"
				break
			case PLAYER_2:
				str += "2"
				break
			default:
				str += "0"
				break
			}
		}
	}

	return
}

// Extra extracts the goban from the request data
func (g *Goban) Extract(requestGoban string) (goban Goban, err error) {
	if len(requestGoban) != 361 {
		// If the goban is not exactly the right length
		return goban, ERR_RD_GOBAN_LENGTH
	}

	i := 0
	for y := 0; y < 19; y++ {
		for x := 0; x < 19; x++ {
			switch requestGoban[i] {
			case '0':
				// If no player on the last position move
				goban[y][x] = PLAYER_0
				break
			case '1':
				// If the player 1 is on the last position move
				goban[y][x] = PLAYER_1
				break
			case '2':
				// If the player 2 is on the last position move
				goban[y][x] = PLAYER_2
				break
			default:
				// If an unknown character is on the last position move
				return goban, ERR_RD_GOBAN_CHARACTER_ONLY_012
			}
			i++
		}
	}

	return
}

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
