package models

type ResponseNextOptionData struct {
	Time     int64    `json:"time"` // ms
	Position Position `json:"position"`
}

type ResponseNextData struct {
	Options        ResponseNextOptionData `json:"options"`
	Goban          string             `json:"goban"`
	HeuristicGoban string             `json:"heuristic_goban"`
	Analyzer       Analyzer           `json:"analyzer"`
}

func (r *ResponseNextData) computeResponseNext(context *Context, time int64, heuristic_goban HeuristicGoban) {
	r.Options = ResponseNextOptionData{
		Time: time,
		Position: Position{
			X: context.State.LastMove.Position.X,
			Y: context.State.LastMove.Position.Y,
		},
	}

	r.Goban = context.Goban.ToString()
	r.HeuristicGoban = heuristic_goban.ToString()
	r.Analyzer = context.Analyzer
}
