package models

import "fmt"

type PlayerInfo struct {
	Captures   uint8 `json:"captures"`
	Alignement bool  `json:"alignement"`
	Win        bool  `json:"win"`
}

type PlayersInfo struct {
	Player_1 PlayerInfo `json:"player_1"`
	Player_2 PlayerInfo `json:"player_2"`
}

// Print prints players info
func (pi *PlayersInfo) Print() {
	// Prints player 1 infos
	fmt.Println("--------------- [players 1]")
	fmt.Printf("captures:\t\t%d\n", pi.Player_1.Captures)
	fmt.Printf("alignements:\t%t\n", pi.Player_1.Alignement)
	fmt.Printf("win:\t\t%t\n", pi.Player_1.Win)

	// Prints player 1 infos
	fmt.Println("--------------- [players 2]")
	fmt.Printf("captures:\t\t%d\n", pi.Player_2.Captures)
	fmt.Printf("alignements:\t%t\n", pi.Player_2.Alignement)
	fmt.Printf("win:\t\t%t\n", pi.Player_2.Win)
}
