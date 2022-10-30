package doubleThree

import (
	"testing"
	m "github.com/trixky/gomoku/models"
)

func TestPosArrayTranslation(t *testing.T) {
	tests := []struct {
		pos m.Position
		mulX int
		mulY int
		direction int
		trueArray []m.Position
	}{
		{
			pos: m.Position{
				10,12,
			},
			
		}, 

}

