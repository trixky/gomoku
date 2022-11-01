package logic

import (
	"github.com/trixky/gomoku/heuristics"
	"github.com/trixky/gomoku/models"
)

// Negamax ...
func Negamax(context *models.Context, parent_channel chan *models.Context) (childs []models.Context) {
	if context.State.Depth < context.Options.DepthMax {
		best_child := &models.Context{}
		best_child.State.Init()

		child_channel := make(chan *models.Context, context.Options.WidthMax)
		childs_to_wait := 0

		for y, line := range context.Goban {
			// For each line
			for x, cell := range line {
				// For each cell
				if cell >= context.Options.ProximityThreshold && cell < models.PLAYER_1 {
					childs_to_wait++

					child := context.Next(models.Position{
						X: uint8(x),
						Y: uint8(y),
					})

					childs = append(childs, child)

					if context.State.Depth == 0 && context.Options.WidthMultiThreading {
						go Negamax(&childs[len(childs)-1], child_channel)
					} else {
						Negamax(&childs[len(childs)-1], child_channel)
					}

					if len(childs) >= context.Options.WidthMax {
						goto childs_judgment
					}
				}
			}
		}

	childs_judgment:
		for i := 0; i < childs_to_wait; i++ {
			child := <-child_channel
			if context.State.Depth == 0 {
				childs = append(childs, *child)
			}
			if child.State.HeuristicScore > best_child.State.HeuristicScore {
				best_child = child
			}
		}

		context.State.HeuristicScore = -best_child.State.HeuristicScore

	} else {
		context.State.HeuristicScore = heuristics.Random(context)

		if context.State.Depth == 0 {
			childs = append(childs, *context)
		}
	}

	if parent_channel != nil {
		parent_channel <- context
	}

	return
}
