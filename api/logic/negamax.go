package logic

import (
	"github.com/trixky/gomoku/heuristics"
	"github.com/trixky/gomoku/models"
)

// Negamax ...
func Negamax(context *models.Context) *models.Context {
	if context.State.Depth < context.Options.DepthMax {
		best_child := models.Context{}
		best_child.State.Init()

		for y, line := range context.Goban {
			// For each line
			for x, cell := range line {
				// For each cell
				if cell >= context.Options.ProximityThreshold && cell < models.PLAYER_1 {

					child := context.Next(models.Position{
						X: uint8(x),
						Y: uint8(y),
					})

					Negamax(&child)

					if child.State.HeuristicScore > best_child.State.HeuristicScore {
						best_child = child
					}
				}
			}
		}

		if context.State.Depth == 0 {
			return &best_child
		}

		context.State.HeuristicScore = -best_child.State.HeuristicScore
	} else {
		context.State.HeuristicScore = heuristics.All(context)
	}

	return nil
}
