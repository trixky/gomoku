package heuristics

import (
	"github.com/trixky/gomoku/models"
)

// All computes all heuristics of a given context
func All(context *models.Context) int {
	var score int

	score += PotentialAlignement(context)
	score += Alignement(context)
	score += PotentialCapture(context)
	score += Capture(context)
	score += Random(context)

	return score
}
