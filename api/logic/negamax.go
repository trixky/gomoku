package logic

import (
	"time"

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

		depth_pruning := context.Options.DepthPruningPercentage > 0 && context.State.Depth > 1
		width_pruning := context.Options.WidthPruningPercentage > 0

		if depth_pruning || width_pruning {
			local_beta := heuristics.Random(context)

			if (depth_pruning && DepthPruning(local_beta, context)) || (width_pruning && WidthPruning(local_beta, context)) {
				// If depth or width pruning is active and beta is too weak

				if parent_channel != nil {
					parent_channel <- nil
				}

				return
			}
		}

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

					if context.State.Depth == 0 && context.Options.WidthMultiThreading {
						go Negamax(&child, child_channel)
					} else {
						Negamax(&child, child_channel)
					}

					elapsed_time := time.Now().Sub(context.Start).Milliseconds()

					if len(childs) >= context.Options.WidthMax || elapsed_time >= context.Options.TimeOut {
						goto childs_judgment
					}
				}
			}
		}

	childs_judgment:
		for i := 0; i < childs_to_wait; i++ {
			child := <-child_channel

			if child != nil {
				childs = append(childs, *child)

				if child.State.Beta > best_child.State.Beta {
					best_child = child
				}
			}
		}

		if len(childs) == 0 {
			if parent_channel != nil {
				parent_channel <- nil
			}

			return
		}

		context.State.Beta = -best_child.State.Beta
	} else {
		context.State.Beta = heuristics.Random(context)

		if context.State.Depth == 0 {
			childs = append(childs, *context)
		}
	}

	if parent_channel != nil {
		parent_channel <- context
	}

	return
}
