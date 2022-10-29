package models

import (
	"fmt"
	"math"
)

type State struct {
	// Scores
	HeuristicScore int

	LastMove Move
	Depth    uint8
}

// Init Inits the state
func (s *State) Init() {
	s.HeuristicScore = math.MinInt32
}

// Print prints state attributes
func (s *State) Print() {
	fmt.Println("============================= [STATE]")
	// Prints Depth
	fmt.Println("depth:\t", s.Depth)

	// Prints scores
	fmt.Println("--------------- [heuristic]")
	fmt.Println("score:\t", s.HeuristicScore)

	// Prints last move
	s.LastMove.Print()
}
