package zone

// "fmt"

type error interface {
	Error() string
}

type CoordsOutsideError struct{}

func (m *CoordsOutsideError) Error() string {
	return "Coords outside goban"
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a byte, b byte) byte {
	if a > b {
		return a
	}
	return b
}

func MakeZone(board [19][19]byte, radius int, threshold int) [19][19]byte {
	var zone [19][19]byte
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			if board[i][j] != 0 {
				for k := -radius; k < radius+1; k++ {
					for l := -radius; l < radius+1; l++ {
						if i+k >= 0 && i+k <= 18 && j+l >= 0 && j+l <= 18 {
							if zone[i+k][j+l] > byte(threshold) {
								continue
							}
							zone[i+k][j+l] += byte(radius) - max(byte(abs(k)), byte(abs(l))) + 1
							if zone[i+k][j+l] > byte(threshold) {
								zone[i+k][j+l] = byte(threshold)
							}
						}
					}
				}
				zone[i][j] = board[i][j]
			}
		}
	}
	// for i := 0; i < 19; i++ {
	// 	fmt.Println(zone[i])
	// }
	// fmt.Println()
	return zone
}

func UpdateZone(zone *[19][19]byte, x int, y int, player bool, radius int, threshold int) error {
	if x < 0 && x > 18 && y < 0 && y > 18 {
		return &CoordsOutsideError{}
	}
	zone[x][y] = 5
	if player {
		zone[x][y]++
	}
	for k := -radius; k < radius+1; k++ {
		for l := -radius; l < radius+1; l++ {
			if x+k >= 0 && x+k <= 18 && y+l >= 0 && y+l <= 18 {
				if zone[x+k][y+l] > byte(threshold) {
					continue
				}
				zone[x+k][y+l] += byte(radius) - max(byte(abs(k)), byte(abs(l))) + 1
				if zone[x+k][y+l] > byte(threshold) {
					zone[x+k][y+l] = byte(threshold)
				}
			}
		}
	}
	return nil
}
