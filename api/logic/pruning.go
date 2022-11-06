package logic

import (
	"math"

	"github.com/trixky/gomoku/models"
)

func DepthPruning(context *models.Context) bool {
	// Get the best beta value of the previous layer
	best_beta, best_beta_percentage, first := context.Bests[context.State.Depth].MaxAndGetAll(context.State.Beta)

	if !first {
		// If it's not the first best comparaison of the current layer

		// Alpha is the negative beta of the previous layer
		best_alpha := -best_beta
		best_alpha_percentage := -best_beta_percentage

		if context.State.BetaPercentage < best_alpha_percentage-(int(math.Abs(float64(best_alpha)))*(context.Options.PreComputedOptions.ReversedDepthPruningPercentage)) {
			// If beta is too weak
			return true
		}
	}

	return false
}

func WidthPruning(context *models.Context) bool {
	// Get the best beta value of the current layer

	best_beta, best_beta_percentage, first := context.Bests[context.State.Depth].MaxAndGetAll(context.State.Beta)

	if !first && context.State.BetaPercentage < best_beta_percentage-(int(math.Abs(float64(best_beta)))*(context.Options.PreComputedOptions.ReversedWidthPruningPercentage)) {
		// If it's not the first best comparaison of the current layer
		// and beta is too weak
		return true
	}

	return false
}
