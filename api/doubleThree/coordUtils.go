package doubleThree

type Goban [19][19]byte

// coordOnGoban test if the coordinates are on the goban
func coordOnGoban(x int, y int) bool {
	if x < 0 || x > 18 || y < 0 || y > 18 {
		return false
	}
	return true
}

// coordPlayer test if the value of the coordinates on the goban has player number
func coordPlayer(goban Goban, x int, y int, player int) bool {
	if coordOnGoban(x, y) == true {
		if goban[x][y] == byte(player) {
			return true
		}
	}
	return false
}

// coordUnoccupied test if the value of the coordinates on the goban is zero (empty)
func coordUnoccupied(goban Goban, x int, y int) bool {
	if coordOnGoban(x, y) == true {
		if goban[x][y] == 0 {
			return true
		}
	}
	return false
}
