package logic

import (
	"github.com/trixky/gomoku/heuristics"
	"github.com/trixky/gomoku/models"
)

// Negamax ...
func Negamax(context *models.Context, parent_channel chan *models.Context) (childs []models.Context) {
	if context.State.Depth < context.Options.DepthMax {
		best_child := models.Context{}
		best_child.State.Init()

		child_to_wait := 0
		child_channel := make(chan *models.Context, context.Options.WidthMax)

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

					if context.State.Depth == 0 && context.Options.WidthMultiThreading {
						go Negamax(&child, child_channel)
					} else {
						Negamax(&child, child_channel)
					}

					if child_to_wait >= context.Options.WidthMax {
						goto childs_judgment
					}
				}
			}
		}

	childs_judgment:

		for i := 0; i < child_to_wait; i++ {
			child := <-child_channel

			if context.State.Depth == 0 {
				childs = append(childs, *child)
			}
			if child.State.HeuristicScore > best_child.State.HeuristicScore {
				best_child = *child
			}
		}

		context.State.HeuristicScore = -best_child.State.HeuristicScore

		parent_channel <- &best_child
	} else {
		context.State.HeuristicScore = heuristics.Random(context)

		if context.State.Depth == 0 {
			childs = append(childs, *context)
		}

		parent_channel <- context
	}

	return
}
