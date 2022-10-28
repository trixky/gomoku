package models

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	PLAYER_0 uint8 = 0
	PLAYER_1 uint8 = 254
	PLAYER_2 uint8 = 255
)

type Context struct {
	Options *Options
	State   State
	Goban   Goban
}

func (c *Context) Next(position Position) (context Context) {
	context.Options = c.Options
	context.State = c.State
	context.Goban = c.Goban

	context.State.Depth++

	context.State.LastMove.Player = !c.State.LastMove.Player // color
	context.State.Alpha = -c.State.Beta
	context.State.Beta = -c.State.Alpha

	return
}

// OldNegamax ...
func (c *Context) OldNegamax() {
	// https://en.wikipedia.org/wiki/Negamax

	if c.State.Depth < c.Options.DepthMax {
		// value := −∞
		var value int32 = math.MinInt32

		for y, line := range c.Goban {
			// For each line
			for x, cell := range line {
				// For each cell
				if cell >= c.Options.ProximityThreshold {
					// If the cell is in the
					// @@@ foreach child in childNodes do
					child_context := c.Next(Position{
						X: uint8(x),
						Y: uint8(y),
					})

					child_context.OldNegamax()

					// @@@ value := max(value, −negamax(child, depth − 1, −β, −α, −color))

					if child_context.State.Alpha > value {
						value = -child_context.State.Alpha
					}
					// @@@ α := max(α, value)
					if c.State.Alpha < value {
						c.State.Alpha = value
					}
					// @@@ if α ≥ β then
					if c.State.Alpha >= c.State.Beta {
						return
					}
				}
			}
		}
	} else {
		c.State.Alpha = rand.Int31n(100)
	}
}

// Negamax ...
func (c *Context) Negamax() {
	if c.State.Depth < c.Options.DepthMax {
		best := math.MinInt

		for y, line := range c.Goban {
			// For each line
			for x, cell := range line {
				// For each cell
				if cell >= c.Options.ProximityThreshold {

					child_context := c.Next(Position{
						X: uint8(x),
						Y: uint8(y),
					})

					child_context.Negamax()
					child_score := -child_context.State.Super

					if child_score > best {
						best = int(child_score)
					}
				}
			}
		}

		c.State.Super = best
	} else {
		super := int(rand.Int31n(100)) - 50
		c.State.Super = super
	}
}

// Print prints context attributes
func (c *Context) Print() {
	// Prints the header separator
	fmt.Println()
	fmt.Println("* * * * * * * * * * * * * * * * * * * *")
	fmt.Println(" * * * * * * C O N T E X T * * * * * *")
	fmt.Println("* * * * * * * * * * * * * * * * * * *")
	fmt.Println()

	// Prints state attributes
	c.State.Print()
	// Prints goban attributes
	c.Goban.PrintPlayers()
	c.Goban.PrintProximity(&c.Options.ProximityThreshold)
	// Prints options attributes
	c.Options.Print()

	// Prints the footer separator
	fmt.Println(" * * * * * * * * * * * * * * * * * * * *")
}