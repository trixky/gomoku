package doubleThree

import (
	"testing"
	m  "github.com/thrixky/gomoku/models"
)

func TestCoordOnGoban(t *testing.T) {
	tests := []struct {
		pos m.Position
		truth bool
	}{
		{
			pos: m.Position{
				-1,0,
			},
			truth: false,
		}, {
			pos: m.Position{
				-1,-1,
			},
			truth: false,
		}, {
			pos: m.Position{
				0,-1,
			},
			truth: false,
		}, {
			pos: m.Position{
				19,0,
			},
			truth: false,
		}, {
			pos: m.Position{
				0,19,
			},
			truth: false,
		}, {
			pos: m.Position{
				19,19,
			},
			truth: false,
		}, {
			pos: m.Position{
				0,18,
			},
			truth: true,
		}, {
			pos: m.Position{
				18,0,
			},
			truth: true,
		}, {
			pos: m.Position{
				18,18,
			},
			truth: true,
		},
	}

	for _, test := range tests {
		testingTruth := coordOnGoban(test.Position)
		if testingTruth != test.truth {
			t.Fatalf("doubleThree: coordUtils: coordOnGoban: wrong truth")
		}
	}
}

func TestCoordPlayer(t *testing.T) {
	tests := []struct {
		goban  [19][19]byte
		pos		m.Position
		player int
		truth  bool
	}{
		{
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				-1,0,
			},
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				0,-1,
			},
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				-1,-1,
			},
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				19,0,
			},
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				0,19,
			},
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				19,19,
			},
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				0,18,
			},
			player: 254,
			truth:  false,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				0,18,
			},
			player: 255,
			truth:  true,
		}, {
			goban:  [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				0,0,
			},
			player: 254,
			truth:  true,
		},
	}

	for _, test := range tests {
		testingTruth := coordPlayer(test.goban, test.Position, test.player)
		if testingTruth != test.truth {
			t.Fatalf("doubleThree: coordUtils: coordPlayer: wrong truth")
		}
	}
}

func TestCoordUnoccupied(t *testing.T) {
	tests := []struct {
		goban  [19][19]byte
		pos m.Position
		player int
		truth  bool
	}{
		{
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				-1,0,
			},
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				-1,-1,
			},
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				0,-1,
			},
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				19,0,
			},
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				0,19,
			},
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				19,19,
			},
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				0,18,
			},
			truth: false,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				0,1,
			},
			truth: true,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				0,3,
			},
			truth: true,
		}, {
			goban: [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 255}},
			pos: m.Position{
				0,0,
			},
			truth: false,
		},
	}

	for _, test := range tests {
		testingTruth := coordUnoccupied(test.goban, test.Position)
		if testingTruth != test.truth {
			t.Fatalf("doubleThree: coordUtils: coordUnoccupied: wrong truth")
		}
	}
}
