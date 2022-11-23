package logic

import (
	"time"

	"github.com/trixky/gomoku/doubleThree"
	"github.com/trixky/gomoku/heuristics"
	"github.com/trixky/gomoku/models"
)

// Negamax is the main logic function of the program.
// Using the negamax algorithm, boosted by width and depth alpha-beta pruning and tranlsated tables.
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
	local_beta, _ := heuristics.All(context)
	context.State.SetBeta(local_beta)

	if depth_pruning || width_pruning {
		// If depth or width pruning is active

		if depth_pruning && DepthPruning(context) {
			// If depth pruning is active and beta is too weak

			if min_depth_protection {
				// If the min depth protection is active
				// Save the childs
				if !context.State.Saved {
					analyzed_layer.IncrementSavedByMinDepth()
					context.State.Saved = true
				}
			} else {
				// Else the min depth protection is not active
				analyzed_layer.IncrementPrunedInDepth()

				if parent_channel != nil {
					// If the parent channel exists
					// Return himself without start its childs
					parent_channel <- context
				}

				return
			}

		}

		if width_pruning && WidthPruning(context) {
			// If width pruning is active and beta is too weak

			if min_depth_protection {
				// If the min depth protection is active
				// Save the childs
				if !context.State.Saved {
					analyzed_layer.IncrementSavedByMinDepth()
					context.State.Saved = true
				}
			} else {
				// Else the min depth protection is not active

				analyzed_layer.IncrementPrunedInWidth()

				if parent_channel != nil {
					// If the parent channel exists
					// Return himself without start its childs
					parent_channel <- context
				}

				return
			}

		}
	}

	// ************************************** childs
	if context.State.Depth != context.Options.DepthMax-1 {
		// If the context layer is not the last.
		// Start the childs processus

		analyzed_child_layer := &context.Analyzer.Layers[context.State.Depth+1]

		cutted_by_max_width := false
		time_out := false

		for y, line := range context.Goban {
			// For each line
			for x, cell := range line {
				// For each cell
				if cell >= context.Options.ProximityThreshold && cell < models.PLAYER_1 {

					analyzed_child_layer.IncrementTotal()

					if !cutted_by_max_width && !time_out {

						// If the max width and time out are not reached

						// Create a child for the current cell
						child := context.Next(models.Position{
							X: uint8(x),
							Y: uint8(y),
						})

						isDoubleThree, lenCaptured, newGoban := doubleThree.CheckDoubleThree(child.Goban,
							child.State.LastMove.Position,
							child.State.LastMove.Player)

						if isDoubleThree {
							// If the child is a double three
							// Skip it
							continue
						}

						child.Goban = newGoban

						if lenCaptured > 0 {
							// Else if at least one capture occured
							// Update the goban is these captures

							// Increment the captures in the state of the concerned player
							if child.State.LastMove.Player == false {
								child.State.PlayersInfo.Player_1.Captures += uint8(lenCaptured)
							} else {
								child.State.PlayersInfo.Player_2.Captures += uint8(lenCaptured)
							}
						}

						if context.State.Depth == 0 && context.Options.WidthMultiThreading {
							// If it's the first layer and the multi-threading option is active
							// Start the negamax of the child in a go-routine
							go Negamax(&child, child_channel)
						} else {
							// Start the negamax of the child
							Negamax(&child, child_channel)
						}

						// Compute the elapsed time from the beginning
						elapsed_time := time.Now().Sub(context.Start).Milliseconds()

						// Increment the number of childs to wait
						childs_to_wait++

						if childs_to_wait >= context.Options.WidthMax {
							// If the max width is reached
							if min_depth_protection {
								// If the min depth protection is active
								// Save the child
								if !context.State.Saved {
									analyzed_layer.IncrementSavedByMinDepth()
									context.State.Saved = true
								}
							} else {
								// Else the min depth protection is not active
								// Don't save the child

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

		// ************************************** jugement of childs
		for i := 0; i < childs_to_wait; i++ {
			// For each started child
			// Wait the child result
			child := <-child_channel

			if child != nil {
				// If the child have result
				// Save it
				childs = append(childs, *child)

				if child.State.Beta > best_child.State.Beta {
					// If the child is the best one
					// Save it as the best one
					best_child = child
				}
			}
		}

		if len(childs) == 0 {
			// If no childs
			if parent_channel != nil {
				// If the parent channel exists
				// Return himself
				parent_channel <- context
			}

			return
		}

		context.State.SetBeta(-best_child.State.Beta)
	}

	if parent_channel != nil {
		// If the parent channel exists
		analyzed_layer.IncrementSelected()
		// Return himself
		parent_channel <- context
	}

	return
}
