package doubleThree

import (
	"fmt"
	"testing"
)

func TestCoordOnGoban(t *testing.T) {
	tests := []struct {
		x     int
		y     int
		truth bool
	}{
		{
			x:     -1,
			y:     0,
			truth: false,
		}, {
			x:     -1,
			y:     -1,
			truth: false,
		}, {
			x:     0,
			y:     -1,
			truth: false,
		}, {
			x:     19,
			y:     0,
			truth: false,
		}, {
			x:     0,
			y:     19,
			truth: false,
		}, {
			x:     19,
			y:     19,
			truth: false,
		}, {
			x:     0,
			y:     18,
			truth: true,
		}, {
			x:     18,
			y:     0,
			truth: true,
		}, {
			x:     18,
			y:     18,
			truth: true,
		},
	}

	for _, test := range tests {
		testingTruth := coordOnGoban(test.x, test.y)
		if testingTruth != test.truth {
			t.Fatalf("doubleThree: coordUtils: coordOnGoban: wrong truth")
		}
	}
}

func TestCoordPlayer(t *testing.T) {
	tests := []struct {
		goban  [19][19]byte
		x      int
		y      int
		player uint8
		truth  bool
	}{
		{
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:      -1,
			y:      0,
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:      0,
			y:      -1,
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:      -1,
			y:      -1,
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:      19,
			y:      0,
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:      0,
			y:      19,
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:      19,
			y:      19,
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:      18,
			y:      0,
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:      18,
			y:      0,
			player: 255,
			truth:  true,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:      0,
			y:      0,
			player: 254,
			truth:  true,
		},
	}

	for _, test := range tests {
		testingTruth := coordPlayer(test.goban, test.x, test.y, test.player)
		if testingTruth != test.truth {
			t.Fatalf("doubleThree: coordUtils: coordPlayer: wrong truth")
		}
	}
}

func TestCoordUnoccupied(t *testing.T) {
	tests := []struct {
		goban  [19][19]byte
		x int
		y int
		player uint8
		truth  bool
	}{
		{
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:     -1,
			y:     0,
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:     -1,
			y:     61,
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:     0,
			y:     -1,
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:     19,
			y:     0,
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:     0,
			y:     19,
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:     19,
			y:     19,
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:     18,
			y:     0,
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:     1,
			y:     0,
			truth: true,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:     3,
			y:     0,
			truth: true,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			x:     0,
			y:     0,
			truth: false,
		},
	}

	for i, test := range tests {
		testingTruth := coordUnoccupied(test.goban, test.x, test.y)
		if testingTruth != test.truth {
			t.Fatalf("doubleThree: coordUtils: coordUnoccupied: "+ fmt.Sprint(i) + ": wrong truth")
		}
	}
}
