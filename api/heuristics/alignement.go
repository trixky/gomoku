package heuristics

import (
	"fmt"

	"github.com/trixky/gomoku/models"
)

// aligned
const multiplicator_aligned_5 = 60000
const multiplicator_aligned_4 = 10000
const multiplicator_aligned_3 = 4000
const multiplicator_aligned_2 = 2000

// reset
const multiplicator_rest_2 = 1200

// All computes all heuristics of a given context
func Alignement(context *models.Context, player uint8) (heuristic int, aligned_five bool) {
	if context.Options.HeuristicAlignementWeight > 0 {
		var alignedFive, alignedFour, alignedThree, alignedTwo [2]int // 0 is player, 1 is opponent
		var restOfFour, restOfThree, restOfTwo [2]int
		var possibleCaptures [2]int

		for y := 0; y < 19; y++ {
			for x := 0; x < 19; x++ {
				if context.Goban[y][x] >= models.PLAYER_1 {
					nbAligned, canPlay, blocked, emptyMiddle := alignHeuristic(context.Goban, x, y, context.Goban[y][x])
					if nbAligned == 5 {
						aligned_five = true
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

		aggro := context.Options.HeuristicAggro

		// aligned
		multiplicator_player_aligned_5 := (multiplicator_aligned_5 / 100) * aggro
		multiplicator_opponent_aligned_5 := multiplicator_aligned_5

		multiplicator_player_aligned_4 := (multiplicator_aligned_4 / 100) * aggro
		multiplicator_opponent_aligned_4 := multiplicator_aligned_4

		multiplicator_player_aligned_3 := (multiplicator_aligned_3 / 100) * aggro
		multiplicator_opponent_aligned_3 := multiplicator_aligned_3

		multiplicator_player_aligned_2 := (multiplicator_aligned_2 / 100) * aggro
		multiplicator_opponent_aligned_2 := multiplicator_aligned_2

		// reset
		multiplicator_player_rest_4 := multiplicator_player_aligned_4
		multiplicator_opponent_rest_4 := multiplicator_opponent_aligned_4

		multiplicator_player_rest_3 := multiplicator_player_aligned_3
		multiplicator_opponent_rest_3 := multiplicator_opponent_aligned_3

		multiplicator_player_rest_2 := (multiplicator_rest_2 / 100) * aggro
		multiplicator_opponent_rest_2 := multiplicator_rest_2

		fmt.Println("aggro:", aggro)
		fmt.Println("multiplicator_player_aligned_5:", multiplicator_player_aligned_5)
		fmt.Println("multiplicator_opponent_aligned_5:", multiplicator_opponent_aligned_5)

		heuristic = context.Options.HeuristicAlignementWeight*multiplicator_player_aligned_5*(alignedFive[0]) -
			context.Options.HeuristicAlignementWeight*multiplicator_opponent_aligned_5*(alignedFive[1]) +
			context.Options.HeuristicAlignementWeight*multiplicator_player_aligned_4*(alignedFour[0]) -
			context.Options.HeuristicAlignementWeight*multiplicator_opponent_aligned_4*(alignedFour[1]) +
			context.Options.HeuristicAlignementWeight*multiplicator_player_aligned_3*(alignedThree[0]) -
			context.Options.HeuristicAlignementWeight*multiplicator_player_aligned_2*(alignedThree[1]) +
			context.Options.HeuristicAlignementWeight*multiplicator_player_aligned_2*(alignedTwo[0]) -
			context.Options.HeuristicAlignementWeight*multiplicator_opponent_aligned_2*(alignedTwo[1]) +

			context.Options.HeuristicAlignementWeight*multiplicator_player_rest_4*(restOfFour[0]) -
			context.Options.HeuristicAlignementWeight*multiplicator_opponent_rest_4*(restOfFour[1]) +
			context.Options.HeuristicAlignementWeight*multiplicator_player_rest_3*(restOfThree[0]) -
			context.Options.HeuristicAlignementWeight*multiplicator_opponent_rest_3*(restOfThree[1]) +
			context.Options.HeuristicAlignementWeight*multiplicator_player_rest_2*(restOfTwo[0]) -
			context.Options.HeuristicAlignementWeight*multiplicator_opponent_rest_2*(restOfTwo[1])

		heuristic += context.Options.HeuristicCaptureWeight * 2000 * (possibleCaptures[0] - possibleCaptures[1])

		return
	}
	return
}
