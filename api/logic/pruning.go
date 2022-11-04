package logic

import (
	"math"

	"github.com/trixky/gomoku/models"
)

func DepthPruning(beta int, context *models.Context) bool {
	// return false
	// Get the beta value of the previous layer
	beta, first := context.Bests[context.State.Depth-1].Max(beta)

	if !first {
		// If it's not the first best comparaison of the current layer

		// Alpha is the negative beta of the previous layer
		alpha := -beta

		if beta*100 < alpha*100-(int(math.Abs(float64(alpha)))*(100-context.Options.DepthPruningPercentage)) {
			// If beta is too weak

			return true
		}
	}

	return false
}

func WidthPruning(beta int, context *models.Context) bool {
	// Get the beta value of the current layer

	beta, first := context.Bests[context.State.Depth].Max(beta)

	if !first && beta*100 < beta*100-(int(math.Abs(float64(beta)))*(100-context.Options.WidthPruningPercentage)) {
		// If it's not the first best comparaison of the current layer
		// and beta is too weak
		return true
	}

	return false
}