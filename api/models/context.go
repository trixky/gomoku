package models

import (
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"
)

const (
	PLAYER_0 uint8 = 0
	PLAYER_1 uint8 = 254
	PLAYER_2 uint8 = 255
)

type Context struct {
	Start    time.Time
	Options  *Options
	Bests    []Best
	State    State
	Goban    Goban
	Analyzer Analyzer
	Parent   *Context
}

// InitBests initializes the bests array used for pruning
func (c *Context) InitBests() {
	if c.Options.WidthPruningPercentage > 0 || (c.Options.DepthPruningPercentage > 0 && c.Options.DepthMax > 1) {
		c.Bests = make([]Best, c.Options.DepthMax)

		for i := uint8(0); i < c.Options.DepthMax; i++ {
			c.Bests[i] = Best{
				M:    sync.RWMutex{},
				beta: math.MinInt,
			}
		}
	} else {
		c.Bests = nil
	}
}

// Next copy and update a sub context from this parent
func (c *Context) Next(position Position) (context Context) {
	context.Start = c.Start
	context.Options = c.Options
	context.State = c.State
	context.State.Beta = math.MinInt

	context.State.PlayersInfo.Player_1 = c.State.PlayersInfo.Player_1
	context.State.PlayersInfo.Player_2 = c.State.PlayersInfo.Player_2

	context.Goban = c.Goban

	context.State.Depth++

	context.State.LastMove.Player = !c.State.LastMove.Player // color
	context.State.LastMove.Position = position

	context.Bests = c.Bests

	if c.State.LastMove.Player {
		context.Goban[position.Y][position.X] = PLAYER_2
	} else {
		context.Goban[position.Y][position.X] = PLAYER_1
	}

	if c.Options.ProximityEvolution {
		context.Goban.ComputePieceProximity(&position, context.Options.ProximityThreshold, context.Options.ProximityRadius, context.Options.ProximityShape)
	}

	context.Analyzer = c.Analyzer

	context.Parent = c

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
	// Prints analyzer
	c.Analyzer.Print()

	// Prints the footer separator
	fmt.Println(" * * * * * * * * * * * * * * * * * * * *")
}

func (c *Context) ToJSON(time int64, heuristic_goban HeuristicGoban) (string, error) {
	response := ResponseNextData{}
	response.computeResponseNext(c, time, heuristic_goban)

	marshalled, err := json.Marshal(response)

	return string(marshalled), err
}
