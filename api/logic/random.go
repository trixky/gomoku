package logic

import (
	"math/rand"

	"github.com/trixky/gomoku/models"
)

// Random chooses a random child
func Random(context *models.Context, proximity bool) models.Context {
	for {
		for y, line := range context.Goban {
			// For each line
			for x, cell := range line {
				// For each cell
				if (!proximity || cell >= context.Options.ProximityThreshold) && cell < models.PLAYER_1 && rand.Intn(361) == 42 {
					return context.Next(models.Position{
						X: uint8(x),
						Y: uint8(y),
					})
				}
			}
		}
	}
}
