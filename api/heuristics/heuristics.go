package heuristics

import (
	"github.com/trixky/gomoku/models"
)

// All computes all heuristics of a given context
func All(context *models.Context) int {
	var score int = Alignement(context)
	// print(score)

	score *= (1 - (int(context.State.Depth)/(int(context.Options.DepthMax) * 10)))
	return score
}
