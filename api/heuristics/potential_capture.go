package heuristics

import (
	"math/rand"

	"github.com/trixky/gomoku/models"
)

// PotentialCapture computes the potential capture heuristic of a given context
func PotentialCapture(context *models.Context) int {
	if context.Options.HeuristicRandomWeight > 0 {
		// If the weight of the heuristic is not null
		return (int(rand.Int31n(101)) - 50) * context.Options.HeuristicPotentialCaptureWeight
	}

	return 0
}
