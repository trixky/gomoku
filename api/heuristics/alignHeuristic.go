package heuristics

import (
	"github.com/trixky/gomoku/models"
)

// All computes all heuristics of a given context
func Alignement(context *models.Context) int {
	var player uint8 = 255
	if context.State.LastMove.Player == false {
		player = 254
	}

	var alignedFive, alignedFour, alignedThree, alignedTwo [2]int // 0 is player, 1 is opponent
	var midNoBlockFour, midNoBlockThree, midNoBlockTwo [2]int
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
					} else if emptyMiddle == true && blocked == false {
						// not strictly aligned and unblocked
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
							if context.Goban[y][x] == player {
								restOfTwo[0]++
							} else {
								restOfTwo[1]++
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
						} else if nbAligned == 2 {
							if context.Goban[y][x] == player {
								restOfTwo[0]++
								possibleCaptures[1]++
							} else {
								restOfTwo[1]++
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
						} else if nbAligned == 2 {
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

	var score int = context.Options.HeuristicAlignementWeight*10000*(alignedFive[0]-alignedFive[1]) +
		context.Options.HeuristicAlignementWeight*5000*(alignedFour[0]-alignedFour[1]) +
		context.Options.HeuristicAlignementWeight*2000*(alignedThree[0]-alignedThree[1]) +
		context.Options.HeuristicAlignementWeight*1000*(alignedTwo[0]-alignedTwo[1]) +

		context.Options.HeuristicAlignementWeight*2000*(restOfFour[0]-restOfFour[1]) +
		context.Options.HeuristicAlignementWeight*1000*(restOfThree[0]-restOfThree[1]) +
		context.Options.HeuristicAlignementWeight*500*(restOfTwo[0]-restOfTwo[1])

	// score += context.Options.HeuristicAlignementWeight*(midNoBlockFour[0]-midNoBlockFour[1]) +
	// 	context.Options.HeuristicAlignementWeight*(midNoBlockThree[0]-midNoBlockThree[1]) +
	// 	context.Options.HeuristicAlignementWeight*(midNoBlockTwo[0]-midNoBlockTwo[1])

	
	capturesP1 := context.State.PlayersInfo.Player_1.Captures
	capturesP2 := context.State.PlayersInfo.Player_2.Captures
	if player == 255 {
		capturesP1 = context.State.PlayersInfo.Player_2.Captures
		capturesP2 = context.State.PlayersInfo.Player_1.Captures
	}

	if capturesP1 >= 10 {
		score += context.Options.HeuristicCaptureWeight * 20000
	} else if capturesP2 >= 10 {
		score -= context.Options.HeuristicCaptureWeight * 20000
	} else if capturesP1 == 8 {
		score += context.Options.HeuristicCaptureWeight * 5000
	} else if capturesP2 == 8 {
		score -= context.Options.HeuristicCaptureWeight * 5000
	} else {
		score += context.Options.HeuristicCaptureWeight * 1000 * (int(capturesP1) - int(capturesP2))
	}
	score += context.Options.HeuristicCaptureWeight * ((int(capturesP1) * possibleCaptures[0]) - (int(capturesP2) * possibleCaptures[1]))

	return score
}
