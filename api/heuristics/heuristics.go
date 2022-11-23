package heuristics

import (
	"github.com/trixky/gomoku/models"
)

// All computes all heuristics of a given context
func All(context *models.Context) (heuristic int, aligned_five bool) {

	var player uint8 = 255
	if context.State.LastMove.Player == false {
		player = 254
	}

	heuristic, aligned_five = Alignement(context, player)
	heuristic += Capture(context, player)
	heuristic += Random(context)

	if context.Options.HeuristicDepthDivisor > 0 {
		heuristic = int(float64(heuristic) * (1 - (float64(context.State.Depth) * (float64(context.Options.HeuristicDepthDivisor) / 100))))
	}

	return
}
