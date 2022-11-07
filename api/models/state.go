package models

import (
	"fmt"
	"math"
)

type State struct {
	// Scores
	Beta           int
	BetaPercentage int

	LastMove    Move
	Depth       uint8
	Saved       bool
	PlayersInfo RequestNextPlayersInfoData
}

// Init Inits the state
func (s *State) Init() {
	s.Beta = math.MinInt
	s.BetaPercentage = math.MinInt
}

// SetBeta sets beta and pre-compute his percentage
func (s *State) SetBeta(beta int) {
	s.Beta = beta
	s.BetaPercentage = beta * 100
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

	// Prints players infos
	fmt.Println("--------------- [players 1]")
	fmt.Printf("captures:\t%2d\n", s.PlayersInfo.Player_1.Captures)
	fmt.Println("--------------- [players 2]")
	fmt.Printf("captures:\t%2d\n", s.PlayersInfo.Player_2.Captures)
}
