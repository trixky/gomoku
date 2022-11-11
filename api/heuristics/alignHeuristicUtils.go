package heuristics

import (
	"github.com/trixky/gomoku/models"
)

func horizontalHeuristic(goban models.Goban, x int, y int, player uint8) (int, bool, bool, bool) {
	emptyLeft, emptyRight := 0, 0
	alignedNb := 1
	canPlay, blocked, emptyMiddle := false, false, false

	var opp uint8 = 255
	if player == 255 {
		opp = 254
	}
	for left := x - 1; left >= 0; left-- {
		if goban[y][left] == player {
			if emptyLeft != 0 {
				emptyMiddle = true
			}
			alignedNb++
		} else if goban[y][left] == opp {
			blocked = true
			break
		} else {
			emptyLeft++
		}
		if alignedNb >= 5 {
			return alignedNb, canPlay, blocked, emptyMiddle
		}
		if emptyLeft >= 5-alignedNb { // break if too much canPlay left
			canPlay = true
			break
		}
	}
	for right := x + 1; right <= 18; right++ {
		if goban[y][right] == player {
			if emptyRight != 0 {
				emptyMiddle = true
			}
			alignedNb++
		} else if goban[y][right] == opp {
			blocked = true
			break
		} else {
			emptyRight++
		}
		if alignedNb >= 5 {
			return alignedNb, canPlay, blocked, emptyMiddle
		}
		if emptyRight >= 5-alignedNb { // break if too much canPlay right
			canPlay = true
			break
		}
	}
	if emptyLeft+emptyRight >= 5-alignedNb {
		canPlay = true
	}
	return alignedNb, canPlay, blocked, emptyMiddle
}

func verticalHeuristic(goban models.Goban, x int, y int, player uint8) (int, bool, bool, bool) {
	emptyUp, emptyDown := 0, 0
	alignedNb := 1
	canPlay, blocked, emptyMiddle := false, false, false

	var opp uint8 = 255
	if player == 255 {
		opp = 254
	}
	for up := y - 1; up >= 0; up-- {
		if goban[up][x] == player {
			if emptyUp != 0 {
				emptyMiddle = true
			}
			alignedNb++
		} else if goban[up][x] == opp {
			blocked = true
			break
		} else {
			emptyUp++
		}
		if alignedNb >= 5 {
			return alignedNb, canPlay, blocked, emptyMiddle
		}
		if emptyUp >= 5-alignedNb { // break if too much canPlay up
			canPlay = true
			break
		}
	}
	for down := y + 1; down <= 18; down++ {
		if goban[down][x] == player {
			if emptyDown != 0 {
				emptyMiddle = true
			}
			alignedNb++
		} else if goban[down][x] == opp {
			blocked = true
			break
		} else {
			emptyDown++
		}
		if alignedNb >= 5 {
			return alignedNb, canPlay, blocked, emptyMiddle
		}
		if emptyDown >= 5-alignedNb { // break if too much canPlay down
			canPlay = true
			break
		}
	}
	if emptyUp+emptyDown >= 5-alignedNb {
		canPlay = true
	}
	return alignedNb, canPlay, blocked, emptyMiddle
}

func leftDiagHeuristic(goban models.Goban, x int, y int, player uint8) (int, bool, bool, bool) {
	emptyUpLeft, emptyDownRight := 0, 0
	alignedNb := 1
	canPlay, blocked, emptyMiddle := false, false, false

	var opp uint8 = 255
	if player == 255 {
		opp = 254
	}
	for up, left := y-1, x-1; up >= 0 && left >= 0; up, left = up-1, left-1 {
		if goban[up][left] == player {
			if emptyUpLeft != 0 {
				emptyMiddle = true
			}
			alignedNb++
		} else if goban[up][left] == opp {
			blocked = true
			break
		} else {
			emptyUpLeft++
		}
		if alignedNb >= 5 {
			return alignedNb, canPlay, blocked, emptyMiddle
		}
		if emptyUpLeft >= 5-alignedNb { // break if too much canPlay up left
			canPlay = true
			break
		}
	}
	for down, right := y+1, x+1; down <= 18 && right <= 18; down, right = down+1, right+1 {
		if goban[down][right] == player {
			if emptyDownRight != 0 {
				emptyMiddle = true
			}
			alignedNb++
		} else if goban[down][right] == opp {
			blocked = true
			break
		} else {
			emptyDownRight++
		}
		if alignedNb >= 5 {
			return alignedNb, canPlay, blocked, emptyMiddle
		}
		if emptyDownRight >= 5-alignedNb { // break if too much canPlay down right
			canPlay = true
			break
		}
	}
	if emptyUpLeft+emptyDownRight >= 5-alignedNb {
		canPlay = true
	}
	return alignedNb, canPlay, blocked, emptyMiddle
}

func rightDiagHeuristic(goban models.Goban, x int, y int, player uint8) (int, bool, bool, bool) {
	emptyDownLeft, emptyUpRight := 0, 0
	alignedNb := 1
	canPlay, blocked, emptyMiddle := false, false, false

	var opp uint8 = 255
	if player == 255 {
		opp = 254
	}
	for down, left := y+1, x-1; down <= 18 && left >= 0; down, left = down+1, left-1 {
		if goban[down][left] == player {
			if emptyDownLeft != 0 {
				emptyMiddle = true
			}
			alignedNb++
		} else if goban[down][left] == opp {
			blocked = true
			break
		} else {
			emptyDownLeft++
		}
		if alignedNb >= 5 {
			return alignedNb, canPlay, blocked, emptyMiddle
		}
		if emptyDownLeft >= 5-alignedNb { // break if too much canPlay down left
			canPlay = true
			break
		}
	}
	for up, right := y-1, x+1; up >= 0 && right <= 18; up, right = up-1, right+1 {
		if goban[up][right] == player {
			if emptyUpRight != 0 {
				emptyMiddle = true
			}
			alignedNb++
		} else if goban[up][right] == opp {
			blocked = true
			break
		} else {
			emptyUpRight++
		}
		if alignedNb >= 5 {
			return alignedNb, canPlay, blocked, emptyMiddle
		}
		if emptyUpRight >= 5-alignedNb { // break if too much canPlay up right
			canPlay = true
			break
		}
	}
	if emptyDownLeft+emptyUpRight >= 5-alignedNb {
		canPlay = true
	}
	return alignedNb, canPlay, blocked, emptyMiddle
}

// maxIndexOfFour returns the index of maximum of the four values
func maxIndexOfFour(h int, v int, ld int, rd int) int {
	if h >= v && h >= ld && h >= rd {
		return 0
	} else if v >= ld && v >= rd {
		return 1
	} else if ld >= rd {
		return 2
	} else {
		return 3
	}
}

// alignHeuristic computes the best alignement of the position considered
func alignHeuristic(goban models.Goban, x int, y int, player uint8) (int, bool, bool, bool) {
	nbH, canPlayH, blockedH, emptyMiddleH := horizontalHeuristic(goban, x, y, player)
	if nbH >= 5 && emptyMiddleH == false {
		return nbH, canPlayH, blockedH, emptyMiddleH
	}
	nbV, canPlayV, blockedV, emptyMiddleV := verticalHeuristic(goban, x, y, player)
	if nbV >= 5 && emptyMiddleV == false {
		return nbV, canPlayV, blockedV, emptyMiddleV
	}
	nbLD, canPlayLD, blockedLD, emptyMiddleLD := leftDiagHeuristic(goban, x, y, player)
	if nbLD >= 5 && emptyMiddleLD == false {
		return nbLD, canPlayLD, blockedLD, emptyMiddleLD
	}
	nbRD, canPlayRD, blockedRD, emptyMiddleRD := rightDiagHeuristic(goban, x, y, player)
	if nbRD >= 5 && emptyMiddleRD == false {
		return nbRD, canPlayRD, blockedRD, emptyMiddleRD
	}

	indexMax := maxIndexOfFour(nbH, nbV, nbLD, nbRD)
	if indexMax == 0 && canPlayH == true {
		return nbH, canPlayH, blockedH, emptyMiddleH
	} else if indexMax == 1 && canPlayV == true {
		return nbV, canPlayV, blockedV, emptyMiddleV
	} else if indexMax == 2 && canPlayLD == true {
		return nbLD, canPlayLD, blockedLD, emptyMiddleLD
	} else if indexMax == 3 && canPlayRD == true {
		return nbRD, canPlayRD, blockedRD, emptyMiddleRD
	} else if indexMax == 0 && blockedH == false {
		return nbH, canPlayH, blockedH, emptyMiddleH
	} else if indexMax == 1 && blockedV == false {
		return nbV, canPlayV, blockedV, emptyMiddleV
	} else if indexMax == 2 && blockedLD == false {
		return nbLD, canPlayLD, blockedLD, emptyMiddleLD
	} else if indexMax == 3 && blockedRD == false {
		return nbRD, canPlayRD, blockedRD, emptyMiddleRD
	}

	return nbH, canPlayH, blockedH, emptyMiddleH
}
