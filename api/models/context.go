package models

import (
	"encoding/json"
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
	context.State.Super = math.MinInt32

	context.Goban = c.Goban

	context.State.Depth++

	context.State.LastMove.Player = !c.State.LastMove.Player // color
	context.State.LastMove.Position = position
	context.State.Alpha = -c.State.Beta
	context.State.Beta = -c.State.Alpha

	return
}

// Negamax ...
func (c *Context) Negamax() *Context {
	if c.State.Depth < c.Options.DepthMax {
		best_child := Context{}
		best_child.State.Init()

		for y, line := range c.Goban {
			// For each line
			for x, cell := range line {
				// For each cell
				if cell >= c.Options.ProximityThreshold && cell < PLAYER_1 {

					child := c.Next(Position{
						X: uint8(x),
						Y: uint8(y),
					})

					child.Negamax()

					if child.State.Super > best_child.State.Super {
						best_child = child
					}
				}
			}
		}

		c.State.Super = -best_child.State.Super

		if c.State.Depth == 0 {
			return &best_child
		}
	} else {
		super := int(rand.Int31n(101)) - 50
		c.State.Super = super
	}

	return nil
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

func (c *Context) ToJSON() (string, error) {
	response := ResponseData{}
	response.computeResponse(c)

	marshalled, err := json.Marshal(response)

	return string(marshalled), err
}
