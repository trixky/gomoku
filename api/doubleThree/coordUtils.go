package doubleThree

type Goban [19][19]byte

func coordOnGoban(x int, y int) bool {
	if x < 0 || x > 18 || y < 0 || y > 18 {
		return false
	}
	return true
}

func coordPlayer(goban Goban, x int, y int, player int) bool {
	if coordOnGoban(x, y) == true {
		if goban[x][y] == byte(player) {
			return true
		}
	}
	return false
}

func coordOpponent(goban Goban, x int, y int, player int) bool {
	if coordOnGoban(x, y) == true {
		if goban[x][y] != 0 && goban[x][y] != byte(player) {
			return true
		}
	}
	return false
}

func coordUnoccupied(goban Goban, x int, y int) bool {
	if coordOnGoban(x, y) == true {
		if goban[x][y] == 0 {
			return true
		}
	}
	return false
}
