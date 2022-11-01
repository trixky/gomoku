package models

type ResponseOptionData struct {
	Time     int64    `json:"time"` // ms
	Position Position `json:"position"`
}

type ResponseData struct {
	Options ResponseOptionData `json:"options"`
	Goban   string             `json:"goban"`
}

func (r *ResponseData) computeResponse(context *Context, time int64) {
	r.Options = ResponseOptionData{
		Time: time,
		Position: Position{
			X: context.State.LastMove.Position.X,
			Y: context.State.LastMove.Position.Y,
		},
	}

	r.Goban = context.Goban.ToString()
}
