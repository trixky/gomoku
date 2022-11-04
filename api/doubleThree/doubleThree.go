package doubleThree

import (
	m "github.com/trixky/gomoku/models"
)

// posArrayTranslation returns an array of positions that result of translations of 'pos' by (mulX, mulY). direction is for utility purpose
func posArrayTranslation(pos m.Position, mulX int, mulY int, direction int) []m.Position {
	translations := []m.Position{}
	for i := 1; i <= 4; i++ {
		translation := pos
		diff_x := int(translation.X) + mulX*i*direction
		diff_y := int(translation.Y) + mulY*i*direction
		translation.X = uint8(diff_x)
		translation.Y = uint8(diff_y)
		translations = append(translations, translation)
	}
	return translations
}

// checkEndThree returns boolean dep. if the position given would create a three w/ position in one of the ends
func checkEndThree(goban Goban, minus []m.Position, plus []m.Position, player uint8) bool {
	if coordUnoccupiedPos(goban, plus[0]) {
		if coordPlayerPos(goban, minus[0], player) {
			if coordPlayerPos(goban, minus[1], player) && coordUnoccupiedPos(goban, minus[2]) {
				return true // 011X0
			}
			if coordUnoccupiedPos(goban, minus[1]) && coordPlayerPos(goban, minus[2], player) && coordUnoccupiedPos(goban, minus[3]) {
				return true // 0101X0
			}
		}
		if coordUnoccupiedPos(goban, minus[0]) && coordPlayerPos(goban, minus[1], player) && coordPlayerPos(goban, minus[2], player) && coordUnoccupiedPos(goban, minus[3]) {
			return true // 0110X0
		}
	}

	// just reverse positions array with the few lines up
	if coordUnoccupiedPos(goban, minus[0]) {
		if coordPlayerPos(goban, plus[0], player) {
			if coordPlayerPos(goban, plus[1], player) && coordUnoccupiedPos(goban, plus[2]) {
				return true // 0X110
			}
			if coordUnoccupiedPos(goban, plus[1]) && coordPlayerPos(goban, plus[2], player) && coordUnoccupiedPos(goban, plus[4]) {
				return true // 0X1010
			}
		}
		if coordUnoccupiedPos(goban, plus[0]) && coordPlayerPos(goban, plus[1], player) && coordPlayerPos(goban, plus[2], player) && coordUnoccupiedPos(goban, plus[4]) {
			return true // 0X0110
		}
	}

	return false
}

// checMiddleThree returns boolean dep. if the position given would create a three w/ position in the middle
func checkMidddleThree(goban Goban, minus []m.Position, plus []m.Position, player uint8) bool {
	if coordPlayerPos(goban, plus[0], player) && coordUnoccupiedPos(goban, plus[1]) {
		if coordPlayerPos(goban, minus[0], player) && coordUnoccupiedPos(goban, minus[1]) {
			return true // 01X10
		}
		if coordUnoccupiedPos(goban, minus[0]) && coordPlayerPos(goban, minus[1], player) && coordUnoccupiedPos(goban, minus[2]) {
			return true // 010X10
		}
	}
	if coordUnoccupiedPos(goban, plus[0]) && coordPlayerPos(goban, plus[1], player) && coordUnoccupiedPos(goban, plus[2]) {
		if coordPlayerPos(goban, minus[0], player) && coordUnoccupiedPos(goban, minus[1]) {
			return true // 01X010
		}
	}
	return false
}

// checkThree returns boolean dep. if the position given would create a free three w/ given direction
func checkThree(goban Goban, pos m.Position, player uint8, mulX int, mulY int) bool {
	minus := posArrayTranslation(pos, mulX, mulY, -1)
	plus := posArrayTranslation(pos, mulX, mulY, 1)

	endThree := checkEndThree(goban, minus, plus, player)
	middleThree := checkMidddleThree(goban, minus, plus, player)

	if endThree || middleThree {
		return true
	}
	return false
}

// CheckDoubleThree returns boolean true if the position given would create a doubleThree, false otherwise
func CheckDoubleThree(goban Goban, pos m.Position, player uint8) (bool, int, Goban) {
	checkCapture, whoCaptured := isCapture(goban, int(pos.X), int(pos.Y), player)
	if checkCapture {
		for _, captured := range whoCaptured {
			goban[captured.X][captured.Y] = 0
		}
		goban[pos.X][pos.Y] = player
		return false, len(whoCaptured), goban
	}

	hor := checkThree(goban, pos, player, 1, 0)
	vert := checkThree(goban, pos, player, 0, 1)
	lDiag := checkThree(goban, pos, player, 1, 1)
	rDiag := checkThree(goban, pos, player, 1, -1)

	if (hor && (vert || lDiag || rDiag)) ||
		(vert && (lDiag || rDiag)) ||
		(lDiag && rDiag) { // check at least two true
		return true, 0, Goban{}
	}
	goban[pos.X][pos.Y] = player
	return false, 0, goban
}
