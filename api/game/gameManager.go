package game

import (
	s "github.com/trixky/gomoku/structures"
)

func PlayFirst

func playHint(game *s.Game) {

}

func playBot(game *s.Game) {

}

func Play(game *s.Game) {
	// may be same function w/ boolean
	if game.hotseat {
		playHint(game)
	} else if game.placedStones == 0{
		playFirst(game)
	} else {
		playBot(game)
	}
	return
}