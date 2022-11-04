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

type RequestSuspicionData struct {
	Radius int `json:"radius"` // 0 = disable
}

type RequestDepthData struct {
	Max               uint8 `json:"max"`
	Min               uint8 `json:"min"`
	PruningPercentage int   `json:"pruning_percentage"`
}

type RequestWidthData struct {
	Max               int  `json:"max"`
	MultiThreading    bool `json:"multi_threading"`
	PruningPercentage int  `json:"pruning_percentage"`
}

type RequestProximityData struct {
	Radius    uint8 `json:"radius"`
	Threshold uint8 `json:"threshold"`
	Shape     uint8 `json:"shape"`
	Evolution bool  `json:"evolution"`
}

type RequestHeuristicsData struct {
	PotentialAlignementWeight int `json:"potential_alignement"`
	AlignementWeight          int `json:"alignement"`
	PotentialCaptureWeight    int `json:"potential_capture"`
	CaptureWeight             int `json:"capture"`
	RandomWeight              int `json:"random"`
}

type RequestOptionData struct {
	TimeOut    int64                 `json:"time_out"` // ms
	Position   Position              `json:"position"`
	Depth      RequestDepthData      `json:"depth"`
	Width      RequestWidthData      `json:"width"`
	Proximity  RequestProximityData  `json:"proximity"`
	Heuristics RequestHeuristicsData `json:"heuristics"`
	Suspicion  RequestSuspicionData  `json:"suspicion"`
}

type RequestData struct {
	Options RequestOptionData `json:"options"`
	Goban   string            `json:"goban"`
}

// ExtractGoban extracts the goban from the request data
func (rd *RequestData) ExtractGoban() (goban Goban, err error) {
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

// ExtractOptions extracts the options from the request data
func (rd *RequestData) ExtractOptions() (options Options, err error) {
	// Constraints
	options.TimeOut = rd.Options.TimeOut

	// Depth
	options.DepthMax = rd.Options.Depth.Max
	options.DepthMin = rd.Options.Depth.Min
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
	options.HeuristicPotentialAlignementWeight = rd.Options.Heuristics.PotentialAlignementWeight
	options.HeuristicAlignementWeight = rd.Options.Heuristics.AlignementWeight
	options.HeuristicPotentialCaptureWeight = rd.Options.Heuristics.PotentialCaptureWeight
	options.HeuristicCaptureWeight = rd.Options.Heuristics.CaptureWeight
	options.HeuristicRandomWeight = rd.Options.Heuristics.RandomWeight

	// Suspicion
	options.SuspicionRadius = rd.Options.Suspicion.Radius

	return
}

// ExtractState extracts the state from the request data
func (rd *RequestData) ExtractState() (state State, err error) {
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

	return
}

// ComputeContext compute a new context from the request data
func (rd *RequestData) ComputeContext() (context Context, err error) {
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

	return
}
