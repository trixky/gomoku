package models

import (
	"encoding/json"
	"fmt"
	"math"
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

// Next copy and update a sub context from this parent
func (c *Context) Next(position Position) (context Context) {
	context.Options = c.Options
	context.State = c.State
	context.State.HeuristicScore = math.MinInt32

	context.Goban = c.Goban

	context.State.Depth++

	context.State.LastMove.Player = !c.State.LastMove.Player // color
	context.State.LastMove.Position = position

	if c.State.LastMove.Player {
		context.Goban[position.Y][position.X] = PLAYER_2
	} else {
		context.Goban[position.Y][position.X] = PLAYER_1
	}

	if c.Options.ProximityEvolution {
		context.Goban.ComputePieceProximity(&position, context.Options.ProximityThreshold, context.Options.ProximityRadius, context.Options.ProximityShape)
	}

	return
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

func (c *Context) ToJSON(time int64, heuristic_goban HeuristicGoban) (string, error) {
	response := ResponseData{}
	response.computeResponse(c, time, heuristic_goban)

	marshalled, err := json.Marshal(response)

	return string(marshalled), err
}
