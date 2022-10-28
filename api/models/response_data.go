package models

type ResponseOptionData struct {
	Time     uint16   `json:"time"` // ms
	Position Position `json:"position"`
}

type ResponseData struct {
	Options ResponseOptionData `json:"options"`
	Goban   string             `json:"goban"`
}

func (r *ResponseData) computeResponse(context *Context) {
	r.Options = ResponseOptionData{
		Time: 0,
		Position: Position{
			X: context.State.LastMove.Position.X,
			Y: context.State.LastMove.Position.Y,
		},
	}

	r.Goban = context.Goban.ToString()
}
