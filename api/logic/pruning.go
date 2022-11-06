package logic

import (
	"math"

	"github.com/trixky/gomoku/models"
)

func DepthPruning(local_beta int, context *models.Context) bool {
	// return false
	// Get the beta value of the previous layer
	beta, first := context.Bests[context.State.Depth-1].Max(local_beta)

	if !first {
		// If it's not the first best comparaison of the current layer

		// Alpha is the negative beta of the previous layer
		alpha := -beta

		if local_beta*100 < alpha*100-(int(math.Abs(float64(alpha)))*(context.Options.PreComputedOptions.ReversedDepthPruningPercentage)) {
			// If beta is too weak
			return true
		}
	}

	return false
}

func WidthPruning(local_beta int, context *models.Context) bool {
	// Get the beta value of the current layer

	beta, first := context.Bests[context.State.Depth].Max(local_beta)

	if !first && local_beta*100 < beta*100-(int(math.Abs(float64(beta)))*(context.Options.PreComputedOptions.ReversedWidthPruningPercentage)) {
		// If it's not the first best comparaison of the current layer
		// and beta is too weak
		return true
	}

	return false
}
