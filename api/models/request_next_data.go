package models

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const cells_nbr = 19 * 19

var (
	ERR_RD_GOBAN_LENGTH             = errors.New("goban string need 361 characters")
	ERR_RD_GOBAN_CHARACTER_ONLY_012 = errors.New("goban string accepts only [0,1,2] characters")
	ERR_RD_LAST_MOVE_NO_PLAYER      = errors.New("cell of the last move need to be taken by a player")
)

type RequestNextSuspicionData struct {
	Radius int `json:"radius"` // 0 = disable
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
	Aggro            int `json:"aggro"`
	DepthDivisor     int `json:"depth_divisor"`
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

// Sanitize sanitizes these attributes
func (rndd *RequestNextDepthData) Sanitize() error {
	if rndd.Max < rndd.Min {
		return errors.New("max depth need to be greater or equal than the min depth")
	}
	if rndd.PruningPercentage < 0 {
		return errors.New("depth pruning percentage can't be negative")
	}
	if rndd.PruningPercentage > 400 {
		return errors.New("depth pruning percentage need to be smaller or equal than 400")
	}

	return nil
}

// Sanitize sanitizes these attributes
func (rnwd *RequestNextWidthData) Sanitize() error {
	if rnwd.PruningPercentage > cells_nbr {
		return errors.New("max width cutting need to be smaller or equal than " + strconv.Itoa(cells_nbr))
	}
	if rnwd.PruningPercentage > 200 {
		return errors.New("depth pruning percentage need to be smaller or equal than 200")
	}

	return nil
}

// Sanitize sanitizes these attributes
func (rnpd *RequestNextProximityData) Sanitize() error {
	if rnpd.Radius > 9 {
		return errors.New("max proximity shape radius must be smaller or equal than 9")
	}
	if rnpd.Threshold > 36 {
		return errors.New("max proximity threshold must be smaller or equal than 36")
	}
	if rnpd.Shape != SHAPE_SQUARE && rnpd.Shape != SHAPE_STAR {
		return errors.New("proximity shape need to be <neighbour(square)/square/star>")
	}

	return nil
}

// Sanitize sanitizes these attributes
func (rnhd *RequestNextHeuristicsData) Sanitize() error {
	if rnhd.Aggro < 0 {
		return errors.New("aggro can't be negative")
	}
	if rnhd.Aggro > 300 {
		return errors.New("aggro can't be higher than 300")
	}
	if rnhd.DepthDivisor < 0 {
		return errors.New("depth divisor can't be negative")
	}
	if rnhd.DepthDivisor > 100 {
		return errors.New("depth heuristic divisor must be smaller or equal than 100")
	}
	if rnhd.AlignementWeight < 0 || rnhd.AlignementWeight > 10 {
		return errors.New("alignement heuristic weight must be between 0 and 10")
	}
	if rnhd.CaptureWeight < 0 || rnhd.CaptureWeight > 10 {
		return errors.New("capture heuristic weight must be between 0 and 10")
	}
	if rnhd.RandomWeight < 0 || rnhd.RandomWeight > 10 {
		return errors.New("random heuristic weight must be between 0 and 10")
	}

	return nil
}

// Sanitize sanitizes these attributes
func (rnsd *RequestNextSuspicionData) Sanitize() error {
	if rnsd.Radius < 0 {
		return errors.New("suspicion radius must can't be negative")
	}
	if rnsd.Radius > 19 {
		return errors.New("suspicion radius must be smaller or equal than 19")
	}
	return nil
}

// Sanitize sanitizes these attributes
func (rnod *RequestNextOptionData) Sanitize() error {
	if rnod.TimeOut < 0 || rnod.TimeOut > 60000 {
		return errors.New("time out must be between 0 and 60000")
	}
	if err := rnod.Position.Sanitize(); err != nil {
		return err
	}
	if err := rnod.Depth.Sanitize(); err != nil {
		return err
	}
	if err := rnod.Width.Sanitize(); err != nil {
		return err
	}
	if err := rnod.Proximity.Sanitize(); err != nil {
		return err
	}
	if err := rnod.Heuristics.Sanitize(); err != nil {
		return err
	}
	if err := rnod.Suspicion.Sanitize(); err != nil {
		return err
	}

	return nil
}

// Sanitize sanitizes these attributes
func (rd *RequestNextData) Sanitize() error {
	if err := rd.Options.Sanitize(); err != nil {
		return err
	}
	if err := rd.PlayersInfo.Sanitize(); err != nil {
		return err
	}
	if len(rd.Goban) != cells_nbr {
		return errors.New("goban (19x19) expected " + strconv.Itoa(cells_nbr) + " cells, got " + strconv.Itoa(len(rd.Goban)))
	}
	if strings.Count(rd.Goban, "0")+strings.Count(rd.Goban, "1")+strings.Count(rd.Goban, "2") != cells_nbr {
		return errors.New("goban corrupted, it only accepts '0', '1', '2' characters")
	}

	return nil
}

// ExtractGoban extracts the goban from the requestNext data
func (rd *RequestNextData) ExtractGoban() (goban Goban, err error) {
	if len(rd.Goban) != cells_nbr {
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
	options.HeuristicAlignementWeight = rd.Options.Heuristics.AlignementWeight
	options.HeuristicCaptureWeight = rd.Options.Heuristics.CaptureWeight
	options.HeuristicRandomWeight = rd.Options.Heuristics.RandomWeight
	options.HeuristicAggro = rd.Options.Heuristics.Aggro
	options.HeuristicDepthDivisor = rd.Options.Heuristics.DepthDivisor

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
