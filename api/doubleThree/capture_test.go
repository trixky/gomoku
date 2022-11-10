package doubleThree

import (
	"fmt"
	"testing"

	m "github.com/trixky/gomoku/models"
)

func TestIsCapture(t *testing.T) {
	tests := []struct {
		goban    [19][19]byte
		x        int
		y        int
		player   uint8
		truth    bool
		captured []m.Position
	}{
		{ // 0
			goban:    [19][19]byte{{254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			x:        0,
			y:        3,
			player:   254,
			truth:    false,
			captured: []m.Position{},
		}, { // 1
			goban:    [19][19]byte{{254, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			x:        0,
			y:        0,
			player:   254,
			truth:    false,
			captured: []m.Position{},
		}, { // 2
			goban:    [19][19]byte{{0, 255, 255, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			x:        0,
			y:        0,
			player:   254,
			truth:    true,
			captured: []m.Position{{X: 1, Y: 0}, {X: 2, Y: 0}},
		}, { // 3
			goban: [19][19]byte{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			x:        0,
			y:        0,
			player:   255,
			truth:    true,
			captured: []m.Position{{X: 0, Y: 1}, {X: 0, Y: 2}},
		}, { // 4
			goban: [19][19]byte{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{254, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			x:        0,
			y:        0,
			player:   255,
			truth:    true,
			captured: []m.Position{{X: 1, Y: 1}, {X: 2, Y: 2}},
		}, { // 5
			goban: [19][19]byte{{0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			x:        1,
			y:        3,
			player:   254,
			truth:    true,
			captured: []m.Position{{X: 2, Y: 2}, {X: 3, Y: 1}},
		}, { // 6
			goban: [19][19]byte{{0, 0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			x:        7,
			y:        3,
			player:   254,
			truth:    true,
			captured: []m.Position{{X: 6, Y: 2}, {X: 5, Y: 1}},
		}, { // 7
			goban: [19][19]byte{{0, 255, 255, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{254, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			x:        0,
			y:        0,
			player:   254,
			truth:    true,
			captured: []m.Position{{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}},
		},
	}

	for i, test := range tests {
		truth, captured := isCapture(test.goban, test.x, test.y, test.player)
		if truth != test.truth {
			t.Fatalf("doubleThree: isCapture " + fmt.Sprintf("%d", i) + ": wrong truth")
		}
		if len(captured) != len(test.captured) {
			t.Fatalf("doubleThree: isCapture "+fmt.Sprintf("%d", i)+": wrong len captured "+fmt.Sprintf("%d ", len(captured)), fmt.Sprintf("%d", len(test.captured)))
		}
		for j := range test.captured {
			if captured[j].X != test.captured[j].X || captured[j].Y != test.captured[j].Y {
				fmt.Println(captured[j].X, captured[j].Y)
				fmt.Println(test.captured[j].X, test.captured[j].Y)
				t.Fatalf("doubleThree: isCapture " + fmt.Sprintf("%d", i) + ": wrong captured " + fmt.Sprintf("%d", j))
			}
		}
	}
}
