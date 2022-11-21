package heuristics

import (
	"github.com/trixky/gomoku/models"
)

// All computes all heuristics of a given context
func Alignement(context *models.Context, player uint8) int {
	if context.Options.HeuristicAlignementWeight > 0 {
		var alignedFive, alignedFour, alignedThree, alignedTwo [2]int // 0 is player, 1 is opponent
		var restOfFour, restOfThree, restOfTwo [2]int
		var possibleCaptures [2]int

		for y := 0; y < 19; y++ {
			for x := 0; x < 19; x++ {
				if context.Goban[y][x] >= models.PLAYER_1 {
					nbAligned, canPlay, blocked, emptyMiddle := alignHeuristic(context.Goban, x, y, context.Goban[y][x])
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
							} else if nbAligned == 2 {
								if context.Goban[y][x] == player {
									alignedTwo[0]++
								} else {
									alignedTwo[1]++
								}
							}
						} else {
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
							} else if nbAligned == 2 {
								if emptyMiddle == false && blocked == true {
									if context.Goban[y][x] == player {
										restOfTwo[0]++
										possibleCaptures[1]++
									} else {
										restOfTwo[1]++
										possibleCaptures[0]++
									}
								} else {
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
		}

		var score int = context.Options.HeuristicAlignementWeight*50000*(alignedFive[0]-alignedFive[1]) +
			context.Options.HeuristicAlignementWeight*5000*(alignedFour[0]-alignedFour[1]) +
			context.Options.HeuristicAlignementWeight*2000*(alignedThree[0]-alignedThree[1]) +
			context.Options.HeuristicAlignementWeight*1000*(alignedTwo[0]-alignedTwo[1]) +

			context.Options.HeuristicAlignementWeight*2000*(restOfFour[0]-restOfFour[1]) +
			context.Options.HeuristicAlignementWeight*1000*(restOfThree[0]-restOfThree[1]) +
			context.Options.HeuristicAlignementWeight*500*(restOfTwo[0]-restOfTwo[1])

		score += context.Options.HeuristicCaptureWeight * 2000 * (possibleCaptures[0] - possibleCaptures[1])

		return score
	}
	return 0
}
