package heuristics

import (
	"github.com/trixky/gomoku/models"
)

// All computes all heuristics of a given context
func All(context *models.Context) int {
	var player uint8 = 254
	if context.State.LastMove.Player {
		player = 255
	}

	var stones [2]int
	var alignedFive, alignedFour, alignedThree, alignedTwo [2]int // 0 is player, 1 is opponent
	var midNoBlockFour, midNoBlockThree, midNoBlockTwo [2]int
	var restOfFour, restOfThree, restOfTwo [2]int
	var possibleCaptures [2]int

	for y := 0; y < 19; y++ {
		for x := 0; x < 19; x++ {
		if context.Goban[y][x] != 0 {
				if context.Goban[y][x] == player {
					stones[0]++
				} else {
					stones[1]++
				}
				nbAligned, canPlay, blocked, emptyMiddle := alignHeuristic(context, x, y, player)
				if nbAligned == 5 {
					if context.Goban[y][x] == player {
						alignedFive[0]++
					} else {
						alignedFive[1]++
					}
					continue
				}
				if canPlay == true {
					if emptyMiddle == false && blocked == false {
						// free
						if nbAligned == 4 {
							if context.Goban[y][x] == player {
								alignedFour[0]++
							} else {
								alignedFour[1]++
							}
						} else if nbAligned == 3 {
							if context.Goban[y][x] == player {
								alignedThree[0]++
							} else {
								alignedThree[1]++
							}
						} else if nbAligned == 4 {
							if context.Goban[y][x] == player {
								alignedTwo[0]++
							} else {
								alignedTwo[1]++
							}
						}
					} else if emptyMiddle == true && blocked == false {
						// not strictly aligned and unblocked
						if nbAligned == 4 {
							if context.Goban[y][x] == player {
								midNoBlockFour[0]++
							} else {
								midNoBlockFour[1]++
							}
						} else if nbAligned == 3 {
							if context.Goban[y][x] == player {
								midNoBlockThree[0]++
							} else {
								midNoBlockThree[1]++
							}
						} else if nbAligned == 4 {
							if context.Goban[y][x] == player {
								midNoBlockTwo[0]++
							} else {
								midNoBlockTwo[1]++
							}
						}
					} else if emptyMiddle == false && blocked == true {
						// aligned but blocked one way (ABB0), two can be captured
						if nbAligned == 4 {
							if context.Goban[y][x] == player {
								restOfFour[0]++
							} else {
								restOfFour[1]++
							}
						} else if nbAligned == 3 {
							if context.Goban[y][x] == player {
								restOfThree[0]++
							} else {
								restOfThree[1]++
							}
						} else if nbAligned == 4 {
							if context.Goban[y][x] == player {
								restOfTwo[0]++ // if we do 'choose the rules'
								possibleCaptures[1]++
							} else {
								restOfTwo[1]++ // if we do 'choose the rules'
								possibleCaptures[0]++
							}
						}
					} else if emptyMiddle == true && blocked == true {
						// not aligned, blocked one way
						if nbAligned == 4 {
							if context.Goban[y][x] == player {
								restOfFour[0]++
							} else {
								restOfFour[1]++
							}
						} else if nbAligned == 3 {
							if context.Goban[y][x] == player {
								restOfThree[0]++
							} else {
								restOfThree[1]++
							}
						} else if nbAligned == 4 {
							if context.Goban[y][x] == player {
								restOfTwo[0]++
							} else {
								restOfTwo[1]++
							}
						}
					}
				}
			}
		}
	}

	var score int = context.Options.HeuristicAlignementWeight*(alignedFive[0]-alignedFive[1]) +
		context.Options.HeuristicAlignementWeight/2*(alignedFour[0]-alignedFour[1]) +
		context.Options.HeuristicAlignementWeight/10*(restOfFour[0]-restOfFour[1]) +
		context.Options.HeuristicAlignementWeight/10*(alignedThree[0]-alignedThree[1]) +
		context.Options.HeuristicAlignementWeight/30*(restOfThree[0]-restOfThree[1]) +
		context.Options.HeuristicAlignementWeight/100*(alignedTwo[0]-alignedTwo[1])

	score += context.Options.HeuristicAlignementWeight/8*(midNoBlockFour[0]-midNoBlockFour[1]) +
		context.Options.HeuristicAlignementWeight/16*(midNoBlockThree[0]-midNoBlockThree[1]) +
		context.Options.HeuristicAlignementWeight/160*(midNoBlockTwo[0]-midNoBlockTwo[1])

	// score += context.Options.HeuristicStoneWeight * (stones[0] - stones[1])

	// capturesPlayer := context.State.ScoreP1
	// capturesOpp := context.State.ScoreP2
	// if player == 255 {
	// 	capturesPlayer = context.State.ScoreP2
	// 	capturesOpp = context.State.ScoreP1
	// }
	// score += 1000 * (int(capturesPlayer) - int(capturesOpp))
	score += 2000 * (possibleCaptures[0] - possibleCaptures[1])

	return score
}
