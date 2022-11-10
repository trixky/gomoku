package heuristics

import (
	"fmt"
	"testing"
)

func TestMaxIndexOfFour(t *testing.T) {
	tests := []struct {
		nb1        int
		nb2        int
		nb3        int
		nb4        int
		truthIndex int
	}{
		{
			nb1:        0,
			nb2:        1,
			nb3:        0,
			nb4:        0,
			truthIndex: 1,
		}, {
			nb1:        0,
			nb2:        0,
			nb3:        1,
			nb4:        0,
			truthIndex: 2,
		}, {
			nb1:        0,
			nb2:        0,
			nb3:        0,
			nb4:        1,
			truthIndex: 3,
		}, {
			nb1:        0,
			nb2:        3,
			nb3:        2,
			nb4:        5,
			truthIndex: 3,
		}, {
			nb1:        0,
			nb2:        0,
			nb3:        0,
			nb4:        0,
			truthIndex: 0,
		}, {
			nb1:        4,
			nb2:        3,
			nb3:        2,
			nb4:        1,
			truthIndex: 0,
		}, {
			nb1:        1,
			nb2:        2,
			nb3:        4,
			nb4:        3,
			truthIndex: 2,
		},
	}

	for i, test := range tests {
		maxIndex := maxIndexOfFour(test.nb1, test.nb2, test.nb3, test.nb4)
		if maxIndex != test.truthIndex {
			t.Fatalf("heuristics: maxIndexOfFour " + fmt.Sprintf("%d", i) + ": wrong index")
		}
	}
}
