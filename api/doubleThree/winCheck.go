package doubleThree

import (
	m "github.com/trixky/gomoku/models"
)

func isCapturesWin(context *m.Context) bool {
	if context.State.LastMove.Player == false {
		if context.State.PlayersInfo.Player_1.Captures >= 10 {
			context.State.PlayersInfo.Player_1.Win = true
			return true
		}
	} else {
		if context.State.PlayersInfo.Player_2.Captures >= 10 {
			context.State.PlayersInfo.Player_2.Win = true
			return true
		}
	}
	return false
}

func canPosBeCapturedHor(goban m.Goban, pos m.Position, player uint8, opp uint8) bool {
	if coordPlayer(goban, int(pos.X)-2, int(pos.Y), opp) && coordPlayer(goban, int(pos.X)-1, int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordUnoccupied(goban, int(pos.X)+1, int(pos.Y)) {
		return true // xoOe
	}
	if coordUnoccupied(goban, int(pos.X)-2, int(pos.Y)) && coordPlayer(goban, int(pos.X)-1, int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordPlayer(goban, int(pos.X)+1, int(pos.Y), opp) {
		return true // eoOx
	}
	if coordPlayer(goban, int(pos.X)-1, int(pos.Y), opp) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y), player) && coordUnoccupied(goban, int(pos.X)+2, int(pos.Y)) {
		return true // xOoe
	}
	if coordUnoccupied(goban, int(pos.X)-1, int(pos.Y)) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y), player) && coordPlayer(goban, int(pos.X)+2, int(pos.Y), opp) {
		return true // eOox
	}
	return false
}

func canPosBeCapturedVert(goban m.Goban, pos m.Position, player uint8, opp uint8) bool {
	if coordPlayer(goban, int(pos.X), int(pos.Y)-2, opp) && coordPlayer(goban, int(pos.X), int(pos.Y)-1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordUnoccupied(goban, int(pos.X), int(pos.Y)+1) {
		return true // xoOe
	}
	if coordUnoccupied(goban, int(pos.X), int(pos.Y)-2) && coordPlayer(goban, int(pos.X), int(pos.Y)-1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordPlayer(goban, int(pos.X), int(pos.Y)+1, opp) {
		return true // eoOx
	}
	if coordPlayer(goban, int(pos.X), int(pos.Y)-1, opp) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y)+1, player) && coordUnoccupied(goban, int(pos.X), int(pos.Y)+2) {
		return true // xOoe
	}
	if coordUnoccupied(goban, int(pos.X), int(pos.Y)-1) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y)+1, player) && coordPlayer(goban, int(pos.X), int(pos.Y)+2, opp) {
		return true // eOox
	}
	return false
}

func canPosBeCapturedLeftDiag(goban m.Goban, pos m.Position, player uint8, opp uint8) bool {
	if coordPlayer(goban, int(pos.X)-2, int(pos.Y)-2, opp) && coordPlayer(goban, int(pos.X)-1, int(pos.Y)-1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordUnoccupied(goban, int(pos.X)+1, int(pos.Y)+1) {
		return true // xoOe
	}
	if coordUnoccupied(goban, int(pos.X)-2, int(pos.Y)-2) && coordPlayer(goban, int(pos.X)-1, int(pos.Y)-1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordPlayer(goban, int(pos.X)+1, int(pos.Y)+1, opp) {
		return true // eoOx
	}
	if coordPlayer(goban, int(pos.X)-1, int(pos.Y)-1, opp) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y)+1, player) && coordUnoccupied(goban, int(pos.X)+2, int(pos.Y)+2) {
		return true // xOoe
	}
	if coordUnoccupied(goban, int(pos.X)-1, int(pos.Y)-1) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y)+1, player) && coordPlayer(goban, int(pos.X)+2, int(pos.Y)+2, opp) {
		return true // eOox
	}
	return false
}

func canPosBeCapturedRightDiag(goban m.Goban, pos m.Position, player uint8, opp uint8) bool {
	if coordPlayer(goban, int(pos.X)-2, int(pos.Y)+2, opp) && coordPlayer(goban, int(pos.X)-1, int(pos.Y)+1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordUnoccupied(goban, int(pos.X)+1, int(pos.Y)-1) {
		return true // xoOe
	}
	if coordUnoccupied(goban, int(pos.X)-2, int(pos.Y)+2) && coordPlayer(goban, int(pos.X)-1, int(pos.Y)+1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordPlayer(goban, int(pos.X)+1, int(pos.Y)-1, opp) {
		return true // eoOx
	}
	if coordPlayer(goban, int(pos.X)-1, int(pos.Y)+1, opp) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y)-1, player) && coordUnoccupied(goban, int(pos.X)+2, int(pos.Y)-2) {
		return true // xOoe
	}
	if coordUnoccupied(goban, int(pos.X)-1, int(pos.Y)+1) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y)-1, player) && coordPlayer(goban, int(pos.X)+2, int(pos.Y)-2, opp) {
		return true // eOox
	}
	return false
}

func canPositionBeCaptured(goban m.Goban, pos m.Position, player uint8, opp uint8) bool {
	if canPosBeCapturedHor(goban, pos, player, opp) {
		return true
	}
	if canPosBeCapturedVert(goban, pos, player, opp) {
		return true
	}
	if canPosBeCapturedLeftDiag(goban, pos, player, opp) {
		return true
	}
	if canPosBeCapturedRightDiag(goban, pos, player, opp) {
		return true
	}
	return false
}

func CanAlignementBecaptured(goban m.Goban, alignement []m.Position, playerBool bool) bool {
	var player uint8 = 254
	var opp uint8 = 255
	if playerBool {
		player = 255
		opp = 254
	}

	for _, pos := range alignement {
		if canPositionBeCaptured(goban, pos, player, opp) {
			return true
		}
	}
	return false
}

func getFiveAlignements(goban m.Goban, pos m.Position, playerBool bool) [][]m.Position {
	var alignements [][]m.Position

	var player uint8 = 254
	if playerBool {
		player = 255
	}

	// get horizontal
	var nbAligned uint8 = 1
	var alignement []m.Position
	for i := 1; i <= 4; i++ {
		if coordPlayer(goban, int(pos.X)-i, int(pos.Y), player) == false {
			break
		}
		alignement = append(alignement, m.Position{X: pos.X - uint8(i), Y: pos.Y})
		nbAligned++
	}
	for i := 1; i <= 4; i++ {
		if coordPlayer(goban, int(pos.X)+i, int(pos.Y), player) == false {
			break
		}
		alignement = append(alignement, m.Position{X: pos.X + uint8(i), Y: pos.Y})
		nbAligned++
	}
	if nbAligned >= 5 {
		alignement = append(alignement, m.Position{X: pos.X, Y: pos.Y})
		alignements = append(alignements, alignement)
	}

	// get vertical
	nbAligned = 1
	alignement = []m.Position{}
	for i := 1; i <= 4; i++ {
		if coordPlayer(goban, int(pos.X), int(pos.Y)-i, player) == false {
			break
		}
		alignement = append(alignement, m.Position{X: pos.X, Y: pos.Y - uint8(i)})
		nbAligned++
	}
	for i := 1; i <= 4; i++ {
		if coordPlayer(goban, int(pos.X), int(pos.Y)+i, player) == false {
			break
		}
		alignement = append(alignement, m.Position{X: pos.X, Y: pos.Y + uint8(i)})
		nbAligned++
	}
	if nbAligned >= 5 {
		alignement = append(alignement, m.Position{X: pos.X, Y: pos.Y})
		alignements = append(alignements, alignement)
	}

	// get left diag
	nbAligned = 1
	alignement = []m.Position{}
	for i := 1; i <= 4; i++ {
		if coordPlayer(goban, int(pos.X)-i, int(pos.Y)-i, player) == false {
			break
		}
		alignement = append(alignement, m.Position{X: pos.X - uint8(i), Y: pos.Y - uint8(i)})
		nbAligned++
	}
	for i := 1; i <= 4; i++ {
		if coordPlayer(goban, int(pos.X)+i, int(pos.Y)+i, player) == false {
			break
		}
		alignement = append(alignement, m.Position{X: pos.X + uint8(i), Y: pos.Y + uint8(i)})
		nbAligned++
	}
	if nbAligned >= 5 {
		alignement = append(alignement, m.Position{X: pos.X, Y: pos.Y})
		alignements = append(alignements, alignement)
	}

	// get right diag
	nbAligned = 1
	alignement = []m.Position{}
	for i := 1; i <= 4; i++ {
		if coordPlayer(goban, int(pos.X)-i, int(pos.Y)+i, player) == false {
			break
		}
		alignement = append(alignement, m.Position{X: pos.X - uint8(i), Y: pos.Y + uint8(i)})
		nbAligned++
	}
	for i := 1; i <= 4; i++ {
		if coordPlayer(goban, int(pos.X)+i, int(pos.Y)-i, player) == false {
			break
		}
		alignement = append(alignement, m.Position{X: pos.X + uint8(i), Y: pos.Y - uint8(i)})
		nbAligned++
	}
	if nbAligned >= 5 {
		alignement = append(alignement, m.Position{X: pos.X, Y: pos.Y})
		alignements = append(alignements, alignement)
	}
	return alignements
}

func IsWin(context *m.Context) bool {
	// check win by captures
	if isCapturesWin(context) {
		return true
	}

	// check win by alignement
	alignements := getFiveAlignements(context.Goban, context.State.LastMove.Position, context.State.LastMove.Player) // get five alignements, type is the direction

	for _, alignement := range alignements {
		if CanAlignementBecaptured(context.Goban, alignement, context.State.LastMove.Player) == false {
			if context.State.LastMove.Player == false {
				context.State.PlayersInfo.Player_1.Win = true
			} else {
				context.State.PlayersInfo.Player_2.Win = true
			}
			return true
		}
	}
	return false
}
