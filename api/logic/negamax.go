package logic

import (
	"time"

	"github.com/trixky/gomoku/heuristics"
	"github.com/trixky/gomoku/models"
)

// Negamax ...
func Negamax(context *models.Context, parent_channel chan<- *models.Context) (childs []models.Context) {
	analyzed_layer := &context.Analyzer.Layers[context.State.Depth]
	best_child := &models.Context{}
	best_child.State.Init()

	child_channel := make(chan *models.Context, 361)
	childs_to_wait := 0

	depth_pruning := context.Options.DepthPruningPercentage > 0 && context.State.Depth > 1
	width_pruning := context.Options.WidthPruningPercentage > 0

	min_depth_protection := context.State.Depth < context.Options.DepthMin

	// ************************************** pruning
	context.State.SetBeta(heuristics.Random(context))

	if depth_pruning || width_pruning {
		// var analyzed_parent_layer *models.LayerInfo

		if depth_pruning && DepthPruning(context) {
			// If depth pruning is active and beta is too weak

			if min_depth_protection {
				if !context.State.Saved {
					analyzed_layer.IncrementSavedByMinDepth()
					context.State.Saved = true
				}
			} else {
				analyzed_layer.IncrementPrunedInDepth()

				if parent_channel != nil {
					parent_channel <- nil
				}
				return
			}

		}

		if width_pruning && WidthPruning(context) {
			// If width pruning is active and beta is too weak

			if min_depth_protection {
				if !context.State.Saved {
					analyzed_layer.IncrementSavedByMinDepth()
					context.State.Saved = true
				}
			} else {
				analyzed_layer.IncrementPrunedInWidth()

				if parent_channel != nil {
					parent_channel <- nil
				}
				return
			}

		}
	}

	if context.State.Depth != context.Options.DepthMax-1 {
		analyzed_child_layer := &context.Analyzer.Layers[context.State.Depth+1]

		cutted_by_max_width := false
		time_out := false

		// ************************************** childs
		for y, line := range context.Goban {
			// For each line
			for x, cell := range line {
				// For each cell
				if cell >= context.Options.ProximityThreshold && cell < models.PLAYER_1 {
					analyzed_child_layer.IncrementTotal()

					if !cutted_by_max_width && !time_out {
						// If the max width and time out are not reached

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

						childs_to_wait++

						if childs_to_wait >= context.Options.WidthMax {
							// If the max width is reached
							if min_depth_protection {
								if !context.State.Saved {
									analyzed_layer.IncrementSavedByMinDepth()
									context.State.Saved = true
								}
							} else {
								cutted_by_max_width = true
							}
						} else if elapsed_time >= context.Options.TimeOut {
							// Else if the time out is reached
							time_out = true
						}
					} else if cutted_by_max_width {
						// Else if the max width is reached

						analyzed_child_layer.IncrementCuttedByMaxWidth()
					} else {
						// Else the time out is reached

						analyzed_child_layer.IncrementCuttedByTimeOut()
					}
				}
			}
		}

		// ************************************** jugement
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

		context.State.SetBeta(-best_child.State.Beta)
	}

	if parent_channel != nil {
		analyzed_layer.IncrementSelected()
		parent_channel <- context
	}

	return
}
