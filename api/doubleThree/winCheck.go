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

func canPosBeCapturedHor(goban m.Goban, pos m.Position, player uint8, opp uint8) (bool, []m.Position) {
	var mustPlays []m.Position
	if coordPlayer(goban, int(pos.X)-2, int(pos.Y), opp) && coordPlayer(goban, int(pos.X)-1, int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordUnoccupied(goban, int(pos.X)+1, int(pos.Y)) {
		mustPlays = append(mustPlays, m.Position{X: pos.X + 1, Y: pos.Y})
	}
	if coordUnoccupied(goban, int(pos.X)-2, int(pos.Y)) && coordPlayer(goban, int(pos.X)-1, int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordPlayer(goban, int(pos.X)+1, int(pos.Y), opp) {
		mustPlays = append(mustPlays, m.Position{X: pos.X - 2, Y: pos.Y})
	}
	if coordPlayer(goban, int(pos.X)-1, int(pos.Y), opp) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y), player) && coordUnoccupied(goban, int(pos.X)+2, int(pos.Y)) {
		mustPlays = append(mustPlays, m.Position{X: pos.X + 2, Y: pos.Y})
	}
	if coordUnoccupied(goban, int(pos.X)-1, int(pos.Y)) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y), player) && coordPlayer(goban, int(pos.X)+2, int(pos.Y), opp) {
		mustPlays = append(mustPlays, m.Position{X: pos.X - 1, Y: pos.Y})
	}
	if len(mustPlays) > 0 {
		return true, mustPlays
	}
	return false, nil
}

func canPosBeCapturedVert(goban m.Goban, pos m.Position, player uint8, opp uint8) (bool, []m.Position) {
	var mustPlays []m.Position
	if coordPlayer(goban, int(pos.X), int(pos.Y)-2, opp) && coordPlayer(goban, int(pos.X), int(pos.Y)-1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordUnoccupied(goban, int(pos.X), int(pos.Y)+1) {
		mustPlays = append(mustPlays, m.Position{X: pos.X, Y: pos.Y + 1})
	}
	if coordUnoccupied(goban, int(pos.X), int(pos.Y)-2) && coordPlayer(goban, int(pos.X), int(pos.Y)-1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordPlayer(goban, int(pos.X), int(pos.Y)+1, opp) {
		mustPlays = append(mustPlays, m.Position{X: pos.X, Y: pos.Y - 2})
	}
	if coordPlayer(goban, int(pos.X), int(pos.Y)-1, opp) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y)+1, player) && coordUnoccupied(goban, int(pos.X), int(pos.Y)+2) {
		mustPlays = append(mustPlays, m.Position{X: pos.X, Y: pos.Y + 2})
	}
	if coordUnoccupied(goban, int(pos.X), int(pos.Y)-1) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y)+1, player) && coordPlayer(goban, int(pos.X), int(pos.Y)+2, opp) {
		mustPlays = append(mustPlays, m.Position{X: pos.X, Y: pos.Y - 1})
	}
	if len(mustPlays) > 0 {
		return true, mustPlays
	}
	return false, nil
}

func canPosBeCapturedLeftDiag(goban m.Goban, pos m.Position, player uint8, opp uint8) (bool, []m.Position) {
	var mustPlays []m.Position
	if coordPlayer(goban, int(pos.X)-2, int(pos.Y)-2, opp) && coordPlayer(goban, int(pos.X)-1, int(pos.Y)-1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordUnoccupied(goban, int(pos.X)+1, int(pos.Y)+1) {
		mustPlays = append(mustPlays, m.Position{X: pos.X + 1, Y: pos.Y + 1})
	}
	if coordUnoccupied(goban, int(pos.X)-2, int(pos.Y)-2) && coordPlayer(goban, int(pos.X)-1, int(pos.Y)-1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordPlayer(goban, int(pos.X)+1, int(pos.Y)+1, opp) {
		mustPlays = append(mustPlays, m.Position{X: pos.X - 2, Y: pos.Y - 2})
	}
	if coordPlayer(goban, int(pos.X)-1, int(pos.Y)-1, opp) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y)+1, player) && coordUnoccupied(goban, int(pos.X)+2, int(pos.Y)+2) {
		mustPlays = append(mustPlays, m.Position{X: pos.X + 2, Y: pos.Y + 2})
	}
	if coordUnoccupied(goban, int(pos.X)-1, int(pos.Y)-1) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y)+1, player) && coordPlayer(goban, int(pos.X)+2, int(pos.Y)+2, opp) {
		mustPlays = append(mustPlays, m.Position{X: pos.X - 1, Y: pos.Y - 1})
	}
	if len(mustPlays) > 0 {
		return true, mustPlays
	}
	return false, nil
}

func canPosBeCapturedRightDiag(goban m.Goban, pos m.Position, player uint8, opp uint8) (bool, []m.Position) {
	var mustPlays []m.Position

	if coordPlayer(goban, int(pos.X)-2, int(pos.Y)+2, opp) && coordPlayer(goban, int(pos.X)-1, int(pos.Y)+1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordUnoccupied(goban, int(pos.X)+1, int(pos.Y)-1) {
		mustPlays = append(mustPlays, m.Position{X: pos.X + 1, Y: pos.Y - 1})
	}
	if coordUnoccupied(goban, int(pos.X)-2, int(pos.Y)+2) && coordPlayer(goban, int(pos.X)-1, int(pos.Y)+1, player) &&
		coordPlayer(goban, int(pos.X), int(pos.Y), player) && coordPlayer(goban, int(pos.X)+1, int(pos.Y)-1, opp) {
		mustPlays = append(mustPlays, m.Position{X: pos.X - 2, Y: pos.Y + 2})
	}
	if coordPlayer(goban, int(pos.X)-1, int(pos.Y)+1, opp) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y)-1, player) && coordUnoccupied(goban, int(pos.X)+2, int(pos.Y)-2) {
		mustPlays = append(mustPlays, m.Position{X: pos.X + 2, Y: pos.Y - 2})
	}
	if coordUnoccupied(goban, int(pos.X)-1, int(pos.Y)+1) && coordPlayer(goban, int(pos.X), int(pos.Y), player) &&
		coordPlayer(goban, int(pos.X)+1, int(pos.Y)-1, player) && coordPlayer(goban, int(pos.X)+2, int(pos.Y)-2, opp) {
		mustPlays = append(mustPlays, m.Position{X: pos.X - 1, Y: pos.Y + 1})
	}
	if len(mustPlays) > 0 {
		return true, mustPlays
	}
	return false, nil
}

func canPositionBeCaptured(goban m.Goban, pos m.Position, player uint8, opp uint8) (bool, []m.Position) {
	var mustPlays []m.Position

	hor, horPlays := canPosBeCapturedHor(goban, pos, player, opp)
	if hor {
		mustPlays = append(mustPlays, horPlays...)
	}
	vert, vertPlays := canPosBeCapturedVert(goban, pos, player, opp)
	if vert {
		mustPlays = append(mustPlays, vertPlays...)
	}
	ld, ldPlays := canPosBeCapturedLeftDiag(goban, pos, player, opp)
	if ld {
		mustPlays = append(mustPlays, ldPlays...)
	}
	rd, rdPlays := canPosBeCapturedRightDiag(goban, pos, player, opp)
	if rd {
		mustPlays = append(mustPlays, rdPlays...)
	}
	return hor || vert || ld || rd, mustPlays
}

func CanAlignementBecaptured(goban m.Goban, alignement []m.Position, playerBool bool) (bool, []m.Position) {
	var player uint8 = 254
	var opp uint8 = 255
	if playerBool {
		player = 255
		opp = 254
	}

	var rtnBool bool = false
	var mustPlays []m.Position
	for _, pos := range alignement {
		canBeCaptured, curPlays := canPositionBeCaptured(goban, pos, player, opp)
		rtnBool = canBeCaptured || rtnBool
		mustPlays = append(mustPlays, curPlays...)
	}
	return rtnBool, mustPlays
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

func IsWin(context *m.Context) (bool, []m.Position) {
	// check win by captures
	if isCapturesWin(context) {
		return true, nil
	}

	// check win by alignement
	alignements := getFiveAlignements(context.Goban, context.State.LastMove.Position, context.State.LastMove.Player) // get five alignements, type is the direction

	mustPlays := []m.Position{}
	for _, alignement := range alignements {
		canBeCaptured, whereMustPlay := CanAlignementBecaptured(context.Goban, alignement, context.State.LastMove.Player)
		if canBeCaptured == false {
			if context.State.LastMove.Player == false {
				context.State.PlayersInfo.Player_1.Win = true
			} else {
				context.State.PlayersInfo.Player_2.Win = true
			}
			return true, nil
		}
		mustPlays = append(mustPlays, whereMustPlay...)
	}
	return false, mustPlays
}
