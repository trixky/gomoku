package models

import (
	"fmt"
	"math"
)

type State struct {
	// Scores
	Alpha int32
	Beta  int32
	Super int

	LastMove Move
	Depth    uint8
}

// Init Inits the state
func (s *State) Init() {
	s.Alpha = math.MinInt32
	s.Beta = math.MaxInt32
}

// Print prints state attributes
func (s *State) Print() {
	fmt.Println("============================= [STATE]")
	// Prints Depth
	fmt.Println("depth:\t", s.Depth)

	// Prints scores
	fmt.Println("--------------- [scores]")
	fmt.Println("alpha:\t", s.Alpha)
	fmt.Println("beta:\t", s.Beta)

	// Prints last move
	s.LastMove.Print()
}
