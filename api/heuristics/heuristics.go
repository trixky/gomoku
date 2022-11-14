package heuristics

import (
	"github.com/trixky/gomoku/models"
)

// All computes all heuristics of a given context
func All(context *models.Context) int {

	var player uint8 = 255
	if context.State.LastMove.Player == false {
		player = 254
	}

	var score int = Alignement(context, player)
	score += Capture(context, player)
	score += Random(context)

	score *= (1 - (int(context.State.Depth) / (int(context.Options.DepthMax) * context.Options.HeuristicDepthDivisor)))
	return score
}
