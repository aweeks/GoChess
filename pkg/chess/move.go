package chess

type Move struct {
	From      Square
	To        Square
	PromoteTo PieceType
}

func FromSquares(moves []Move) map[Square]bool {
	froms := make(map[Square]bool)

	for _, m := range moves {
		froms[m.From] = true
	}

	return froms
}

func ToSquares(moves []Move) map[Square]bool {
	tos := make(map[Square]bool)

	for _, m := range moves {
		tos[m.To] = true
	}

	return tos
}

// var WhiteOOMove = Move{Square{E_, _1, NilPieceType}, Square{G_, _1}, NilPieceType}
// // var WhiteOOOMove = Move{Square{E_, _1}, Square{C_, _1}}
// // var BlackOOMove = Move{Square{E_, _8}, Square{G_, _8}}
// // var BlackOOOMove = Move{Square{E_, _8}, Square{C_, _8}}

var WhitePawnMove = Square{0, 1}
var WhitePawnDoubleMove = Square{0, 2}
var WhitePawnCaptureMoves []Square = []Square{
	{-1, 1},
	{1, 1},
}

var BlackPawnMove = Square{0, -1}
var BlackPawnDoubleMove = Square{0, -2}
var BlackPawnCaptureMoves []Square = []Square{
	{-1, -1},
	{-1, 1},
}

var KnightMoves []Square = []Square{
	{-2, -1},
	{-2, 1},
	{-1, -2},
	{-1, 2},
	{1, -2},
	{1, 2},
	{2, -1},
	{2, 1},
}

var BishopMoves []Square = []Square{
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

var RookMoves []Square = []Square{
	{-1, 0},
	{0, -1},
	{0, 1},
	{1, 0},
}

var QueenMoves []Square = []Square{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

var KingMoves []Square = []Square{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func (s Square) OnBoard() bool {
	return s.File >= A_ && s.File <= H_ && s.Rank >= _1 && s.Rank <= _8
}
