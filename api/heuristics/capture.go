package heuristics

import (
	"github.com/trixky/gomoku/models"
)

// Capture computes the capture heuristic of a given context
func Capture(context *models.Context, player uint8) int {
	if context.Options.HeuristicCaptureWeight >= 0 {
		var score int

		capturesP1 := context.State.PlayersInfo.Player_1.Captures
		capturesP2 := context.State.PlayersInfo.Player_2.Captures
		if player == 255 {
			capturesP1 = context.State.PlayersInfo.Player_2.Captures
			capturesP2 = context.State.PlayersInfo.Player_1.Captures
		}

		if capturesP1 >= 10 {
			score += context.Options.HeuristicCaptureWeight * 30000
		} else if capturesP2 >= 10 {
			score -= context.Options.HeuristicCaptureWeight * 30000
		} else if capturesP1 == 8 {
			score += context.Options.HeuristicCaptureWeight * 5000
		} else if capturesP2 == 8 {
			score -= context.Options.HeuristicCaptureWeight * 5000
		} else {
			score += context.Options.HeuristicCaptureWeight * 2000 * (int(capturesP1) - int(capturesP2))
		}

		return score
	}

	return 0
}
