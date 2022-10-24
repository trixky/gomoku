package zone

import (
	// "fmt"
)

type error interface {
	Error() string
}

type CoordsOutsideError struct{}

func (m *CoordsOutsideError) Error() string {
	return "Coords outside goban"
}

func MakeZone(board [19][19]byte) [19][19]byte {
	var zone [19][19]byte
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			if board[i][j] != 0 {
				for k := -2; k < 3; k++ {
					for l := -2; l < 3; l++ {
						if (i + k >= 0 && i + k <= 18 && j + l >= 0 && j + l <= 18) {
							if zone[i + k][j + l] > 2 {
								continue
							}
							if (k == -2 || k == 2 || l == -2 || l == 2) {
								zone[i + k][j + l] += 1
							} else {
								zone[i + k][j + l] += 2
							}
							if zone[i + k][j + l] > 2 {
								zone[i + k][j + l] = 2
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

func UpdateZone(zone *[19][19]byte, x int, y int, player bool) (error) {
	if (x < 0 && x > 18 && y < 0 && y > 18) {
		return &CoordsOutsideError{}
	}
	zone[x][y] = 5
	if player {
		zone[x][y]++
	}
	for k := -2; k < 3; k++ {
		for l := -2; l < 3; l++ {
			if (x + k >= 0 && x + k <= 18 && y + l >= 0 && y + l <= 18) {
				if zone[x + k][y + l] > 2 {
					continue
				}
				if (k == -2 || k == 2 || l == -2 || l == 2) {
					zone[x + k][y + l] += 1
				} else {
					zone[x + k][y + l] += 2
				}
				if zone[x + k][y + l] > 2 {
					zone[x + k][y + l] = 2
				} 
			}
		}
	}
	return nil
}
