package models

import (
	"errors"
	"time"
)

var (
	ERR_RD_GOBAN_LENGTH             = errors.New("goban string need 361 characters")
	ERR_RD_GOBAN_CHARACTER_ONLY_012 = errors.New("goban string accepts only [0,1,2] characters")
	ERR_RD_LAST_MOVE_NO_PLAYER      = errors.New("cell of the last move need to be taken by a player")
)

type RequestNextSuspicionData struct {
	Radius uint8 `json:"radius"` // 0 = disable
}

type RequestNextDepthData struct {
	Max               uint8 `json:"max"`
	Min               uint8 `json:"min"`
	PruningPercentage int   `json:"pruning_percentage"`
}

type RequestNextWidthData struct {
	Max               int  `json:"max"`
	MultiThreading    bool `json:"multi_threading"`
	PruningPercentage int  `json:"pruning_percentage"`
}

type RequestNextProximityData struct {
	Radius    uint8 `json:"radius"`
	Threshold uint8 `json:"threshold"`
	Shape     uint8 `json:"shape"`
	Evolution bool  `json:"evolution"`
}

type RequestNextHeuristicsData struct {
	AlignementWeight int `json:"alignement"`
	CaptureWeight    int `json:"capture"`
	RandomWeight     int `json:"random"`
}

type RequestNextOptionData struct {
	TimeOut    int64                     `json:"time_out"` // ms
	Position   Position                  `json:"position"`
	Depth      RequestNextDepthData      `json:"depth"`
	Width      RequestNextWidthData      `json:"width"`
	Proximity  RequestNextProximityData  `json:"proximity"`
	Heuristics RequestNextHeuristicsData `json:"heuristics"`
	Suspicion  RequestNextSuspicionData  `json:"suspicion"`
}
type RequestNextData struct {
	Options     RequestNextOptionData `json:"options"`
	PlayersInfo PlayersInfo           `json:"players_info"`
	Goban       string                `json:"goban"`
}

// ExtractGoban extracts the goban from the requestNext data
func (rd *RequestNextData) ExtractGoban() (goban Goban, err error) {
	if len(rd.Goban) != 361 {
		// If the goban is not exactly the right length
		return goban, ERR_RD_GOBAN_LENGTH
	}

	i := 0
	for y := 0; y < 19; y++ {
		for x := 0; x < 19; x++ {
			switch rd.Goban[i] {
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

// ExtractOptions extracts the options from the requestNext data
func (rd *RequestNextData) ExtractOptions() (options Options, err error) {
	// Constraints
	options.TimeOut = rd.Options.TimeOut

	// Depth
	options.DepthMax = rd.Options.Depth.Max + 1
	options.DepthMin = rd.Options.Depth.Min + 1
	options.DepthPruningPercentage = rd.Options.Depth.PruningPercentage

	// Width
	options.WidthMax = rd.Options.Width.Max
	options.WidthMultiThreading = rd.Options.Width.MultiThreading
	options.WidthPruningPercentage = rd.Options.Width.PruningPercentage

	// Proximity
	options.ProximityRadius = rd.Options.Proximity.Radius
	options.ProximityThreshold = rd.Options.Proximity.Threshold
	options.ProximityShape = rd.Options.Proximity.Shape
	options.ProximityEvolution = rd.Options.Proximity.Evolution

	// Heuristics
	options.HeuristicAlignementWeight = 1000 *rd.Options.Heuristics.AlignementWeight
	options.HeuristicCaptureWeight = 1000 * rd.Options.Heuristics.CaptureWeight
	options.HeuristicRandomWeight = rd.Options.Heuristics.RandomWeight

	// Suspicion
	options.SuspicionRadius = rd.Options.Suspicion.Radius

	// Pre computed options
	options.PreCompute()

	return
}

// ExtractState extracts the state from the requestNext data
func (rd *RequestNextData) ExtractState() (state State, err error) {
	var player bool

	// Find the player of the last move
	switch rd.Goban[int32(rd.Options.Position.Y)*19+int32(rd.Options.Position.X)] {
	case '0':
		// If no player on the last position move
		return state, ERR_RD_LAST_MOVE_NO_PLAYER
	case '1':
		// If the player 1 is on the last position move
		player = false
		break
	case '2':
		// If the player 2 is on the last position move
		player = true
		break
	default:
		// If an unknown character is on the last position move
		return state, ERR_RD_GOBAN_CHARACTER_ONLY_012
	}

	// Init the Alpha and Beta
	state.Init()
	// Set the last move
	state.LastMove = Move{
		Player: player,
		Position: Position{
			X: rd.Options.Position.X,
			Y: rd.Options.Position.Y,
		},
	}

	state.PlayersInfo = rd.PlayersInfo

	return
}

// ComputeContext compute a new context from the requestNext data
func (rd *RequestNextData) ComputeContext() (context Context, err error) {
	// Start the time
	context.Start = time.Now()

	// Extract the goban
	if goban, err := rd.ExtractGoban(); err != nil {
		return context, err
	} else {
		context.Goban = goban
	}

	// Extract the options
	if options, err := rd.ExtractOptions(); err != nil {
		return context, err
	} else {
		context.Options = &options
	}

	// Extract the state
	if state, err := rd.ExtractState(); err != nil {
		return context, err
	} else {
		context.State = state
	}

	context.InitBests()

	context.Analyzer.Init(context.Options.DepthMax)

	context.Parent = nil

	return
}
