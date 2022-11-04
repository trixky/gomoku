package doubleThree

import (
	"fmt"
	"testing"

	m "github.com/trixky/gomoku/models"
)

func TestCheckDoubleThree(t *testing.T) {
	tests := []struct {
		goban      [19][19]byte
		pos        m.Position
		player     uint8
		truth      bool
		nbCaptured int
		truthGoban [19][19]byte
	}{
		{
			goban: [19][19]byte{{0, 255, 255, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{254, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			pos:        m.Position{X: 0, Y: 0},
			player:     254,
			truth:      false,
			nbCaptured: 4,
			truthGoban: [19][19]byte{{254, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
	}

	for i, test := range tests {
		truth, nb, goban := CheckDoubleThree(test.goban, test.pos, test.player)
		if truth != test.truth {
			t.Fatalf("doubleThree: checkDoubleThree " + fmt.Sprintf("%d", i) + ": wrong truth")
		}
		if nb != test.nbCaptured {
			t.Fatalf("doubleThree: checkDoubleThree " + fmt.Sprintf("%d", i) + ": wrong truth")
		}
		for i := range goban {
			for j := range goban[i] {
				if goban[i][j] != test.truthGoban[i][j] {
					fmt.Println(i, j, goban[i][j], test.truthGoban[i][j])
					t.Fatalf("doubleThree: isCapture " + fmt.Sprintf("%d", i) + ": wrong goban")
				}
			}
		}
	}
}
