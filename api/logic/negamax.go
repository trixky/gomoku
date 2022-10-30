package logic

import (
	"github.com/trixky/gomoku/heuristics"
	"github.com/trixky/gomoku/models"
)

// Negamax ...
func Negamax(context *models.Context, channel chan *models.Context) {
	if context.State.Depth < context.Options.DepthMax {
		best_child := models.Context{}
		best_child.State.Init()

		child_to_wait := 0

		for y, line := range context.Goban {
			// For each line
			for x, cell := range line {
				// For each cell
				if cell >= context.Options.ProximityThreshold && cell < models.PLAYER_1 {
					child_to_wait++

					child := context.Next(models.Position{
						X: uint8(x),
						Y: uint8(y),
					})

					if context.State.Depth == 0 {
						go Negamax(&child, channel)
					} else {
						Negamax(&child, channel)
					}
				}
			}
		}

		for i := 0; i < child_to_wait; i++ {
			child := <-channel

			if child.State.HeuristicScore > best_child.State.HeuristicScore {
				best_child = *child
			}
		}

		context.State.HeuristicScore = -best_child.State.HeuristicScore

		if channel != nil {
			channel <- &best_child
		}
	} else {
		context.State.HeuristicScore = heuristics.All(context)
		if channel != nil {
			channel <- context
		}
	}
}
