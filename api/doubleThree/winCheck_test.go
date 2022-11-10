package doubleThree

import (
	"fmt"
	"testing"

	m "github.com/trixky/gomoku/models"
)

func TestCanPosBeCapturedVert(t *testing.T) {
	tests := []struct {
		goban  m.Goban
		pos    m.Position
		player uint8
		opp    uint8
		truth  bool
	}{
		{
			goban: m.Goban{{255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			pos:    m.Position{X: 1, Y: 0},
			player: 255,
			opp:    254,
			truth:  false,
		},
		{
			goban: m.Goban{{0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			pos:    m.Position{X: 1, Y: 1},
			player: 255,
			opp:    254,
			truth:  true,
		},
	}

	for i, test := range tests {
		truth := canPosBeCapturedVert(test.goban, test.pos, test.player, test.opp)
		if truth != test.truth {
			test.goban.PrintPlayers()
			t.Fatalf("winCheck: getFiveAlignements " + fmt.Sprintf("%d", i) + ": wrong truth")
		}
	}
}

func TestCanAlignementBecaptured(t *testing.T) {
	tests := []struct {
		goban      m.Goban
		alignement []m.Position
		player     bool
		truth      bool
	}{
		{
			goban: m.Goban{{255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			alignement: []m.Position{{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 4, Y: 0}},
			player: true,
			truth:  false,
		},
		{
			goban: m.Goban{{0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			alignement: []m.Position{{X: 0, Y: 1},
				{X: 1, Y: 1},
				{X: 2, Y: 1},
				{X: 3, Y: 1},
				{X: 4, Y: 1}},
			player: true,
			truth:  true,
		},
		{
			goban: m.Goban{{0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			alignement: []m.Position{{X: 0, Y: 1},
				{X: 1, Y: 1},
				{X: 2, Y: 1},
				{X: 3, Y: 1},
				{X: 4, Y: 1}},
			player: true,
			truth:  true,
		},
		{
			goban: m.Goban{{0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 255, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 254, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			alignement: []m.Position{{X: 1, Y: 0},
				{X: 2, Y: 1},
				{X: 3, Y: 2},
				{X: 4, Y: 3},
				{X: 5, Y: 4}},
			player: true,
			truth:  true,
		},
		{
			goban: m.Goban{{0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 254, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			alignement: []m.Position{{X: 1, Y: 4},
				{X: 2, Y: 3},
				{X: 3, Y: 2},
				{X: 4, Y: 1},
				{X: 5, Y: 0}},
			player: true,
			truth:  true,
		},
	}

	for i, test := range tests {
		truth := CanAlignementBecaptured(test.goban, test.alignement, test.player)
		if truth != test.truth {
			t.Fatalf("winCheck: CanAlignementBecaptured " + fmt.Sprintf("%d", i) + ": wrong truth")
		}
	}
}

func TestGetFiveAlignements(t *testing.T) {
	tests := []struct {
		goban  m.Goban
		pos    m.Position
		player bool
		truth  [][]m.Position
	}{
		{ // 0
			goban: m.Goban{{255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			pos:    m.Position{X: 2, Y: 0},
			player: true,
			truth: [][]m.Position{
				{{X: 1, Y: 0}, {X: 0, Y: 0}, {X: 3, Y: 0}, {X: 4, Y: 0}, {X: 2, Y: 0}}},
		},
		{ // 1
			goban: m.Goban{{0, 254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			pos:    m.Position{X: 0, Y: 0},
			player: true,
			truth: [][]m.Position{
				{{X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}, {X: 4, Y: 1}, {X: 0, Y: 1}}},
		},
		{ // 2
			goban: m.Goban{{0, 254, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			pos:    m.Position{X: 3, Y: 1},
			player: true,
			truth: [][]m.Position{
				{{X: 2, Y: 1}, {X: 1, Y: 1}, {X: 0, Y: 1}, {X: 4, Y: 1}, {X: 3, Y: 1}},
				{{X: 3, Y: 0}, {X: 3, Y: 2}, {X: 3, Y: 3}, {X: 3, Y: 4}, {X: 3, Y: 1}},
			},
		},
		{ // 3
			goban: m.Goban{{0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 255, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 254, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			pos:    m.Position{X: 5, Y: 4},
			player: true,
			truth: [][]m.Position{
				{{X: 4, Y: 3}, {X: 3, Y: 2}, {X: 2, Y: 1}, {X: 1, Y: 0}, {X: 5, Y: 4}},
			},
		},
	}

	for index, test := range tests {
		truth := getFiveAlignements(test.goban, test.pos, test.player)
		for i := range truth {
			for j := range truth[i] {
				if truth[i][j].X != test.truth[i][j].X || truth[i][j].Y != test.truth[i][j].Y {
					fmt.Printf("%d: %d %d: %d %d vs %d %d\n", index, i, j, truth[i][j].X, truth[i][j].Y, test.truth[i][j].X, test.truth[i][j].Y)
					t.Fatalf("winCheck: getFiveAlignements " + fmt.Sprintf("%d", index) + ": alignement " + fmt.Sprintf("%d", i) + ": pos " + fmt.Sprintf("%d", j) + " wrong")
				}
			}
		}
	}
}
