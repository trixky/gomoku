package models

import (
	"fmt"
	"math"
)

type State struct {
	// Scores
	Beta int

	LastMove Move
	Depth    uint8
}

// Init Inits the state
func (s *State) Init() {
	s.Beta = math.MinInt32
}

// Print prints state attributes
func (s *State) Print() {
	fmt.Println("============================= [STATE]")
	// Prints Depth
	fmt.Println("depth:\t", s.Depth)

	// Prints scores
	fmt.Println("--------------- [heuristic]")
	fmt.Println("score:\t", s.Beta)

	// Prints last move
	s.LastMove.Print()
}
