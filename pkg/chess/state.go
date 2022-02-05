package chess

import (
	"fmt"
	"strings"
)

type GameState struct {
	Board          Board
	Turn           Player
	WhiteCanOO     bool
	WhiteCanOOO    bool
	BlackCanOO     bool
	BlackCanOOO    bool
	EnPassanteFile File
	Result         GameResult
	ResultComment  GameResultComment
}

type GameResult int8

const (
	NilResult GameResult = -1

	WhiteWin = 1
	BlackWin = 2
	Draw     = 3
)

type GameResultComment int8

const (
	NilResultComment = -1

	WinByCheckmate = 1
	WinByTimeout   = 2

	DrawByStalemate            = 3
	DrawByInsufficientMaterial = 4
	DrawByAgreement            = 5
	DrawByRepititon            = 6
	DrawBy50MoveRule           = 7
)

func (gs GameState) String() string {
	var buff strings.Builder

	buff.WriteString(fmt.Sprintf("Turn: %s\n\n", gs.Turn))

	buff.WriteString(gs.Board.String())
	buff.WriteString("\n")
	buff.WriteString(fmt.Sprintf("White Can OO:  %t\n", gs.WhiteCanOO))
	buff.WriteString(fmt.Sprintf("White Can OOO: %t\n", gs.WhiteCanOOO))

	buff.WriteString(fmt.Sprintf("Black Can OO:  %t\n", gs.BlackCanOO))
	buff.WriteString(fmt.Sprintf("Black Can OOO: %t\n", gs.BlackCanOOO))

	buff.WriteString(fmt.Sprintf("En Passante File: %s\n", gs.EnPassanteFile))

	return buff.String()
}

func (gs GameState) Move(from Square, to Square) GameState {
	new := gs

	fromPiece := new.Board[from.File][from.Rank]
	toPiece := new.Board[to.File][to.Rank]

	new.Board[to.File][to.Rank] = fromPiece
	new.Board[from.File][from.Rank] = NilPiece
	new.Turn = -new.Turn
	new.EnPassanteFile = NilFile // En passant option lasts only one half move

	/* Move rook due to castling.  OO/OOO eligibility updated later. */
	if (fromPiece == WhiteKing && from == Square{E_, _1}) {
		if (to == Square{C_, _1}) {
			// White OOO
			new.Board[D_][_1] = WhiteRook
			new.Board[A_][_1] = NilPiece
		} else if (to == Square{G_, _1}) {
			// White OO
			new.Board[F_][_1] = WhiteRook
			new.Board[H_][_1] = NilPiece
		}
	} else if (fromPiece == BlackKing && from == Square{E_, _8}) {
		if (to == Square{C_, _8}) {
			// Black OOO
			new.Board[D_][_8] = WhiteRook
			new.Board[A_][_8] = NilPiece
		} else if (to == Square{G_, _8}) {
			// Black OO
			new.Board[F_][_8] = WhiteRook
			new.Board[H_][_8] = NilPiece
		}
	}

	switch fromPiece {

	/* Allow en passant next half move.  Pawn moving forwards two squares implies from.File == to.File per rules of chess so no case for this. */
	case WhitePawn:
		if from.Rank == _2 && to.Rank == _4 {
			new.EnPassanteFile = from.File
		}
	case BlackPawn:
		if from.Rank == _7 && to.Rank == _5 {
			new.EnPassanteFile = from.File
		}

	/* Disallow OO and OOO after king moves */
	case WhiteKing:
		new.WhiteCanOO = false
		new.WhiteCanOOO = false
	case BlackKing:
		new.BlackCanOO = false
		new.BlackCanOOO = false

	/* Disallow OO or OOO after rook moves */
	case WhiteRook:
		if (from == Square{A_, _1}) {
			new.WhiteCanOOO = false
		} else if (from == Square{H_, _1}) {
			new.WhiteCanOO = false
		}
	case BlackRook:
		if (from == Square{A_, _8}) {
			new.BlackCanOOO = false
		} else if (from == Square{H_, _8}) {
			new.BlackCanOO = false
		}
	}

	/* Disallow OO or OOO after rook capture.  King capture should never happen per rules of chess so no case for this. */
	switch toPiece {
	case WhiteRook:
		if (to == Square{A_, _1}) {
			new.WhiteCanOOO = false
		} else if (to == Square{H_, _1}) {
			new.WhiteCanOO = false
		}
	case BlackRook:
		if (to == Square{A_, _8}) {
			new.BlackCanOOO = false
		} else if (to == Square{H_, _8}) {
			new.BlackCanOO = false
		}
	}
	return new
}

func (gs GameState) LegalMoves() []Move {
	moves := []Move{}
	for _, s := range gs.Board.PlayerSquares(gs.Turn) {
		moves = append(moves, gs.MovesFrom(s)...)
	}
	return moves
}

func (gs GameState) MovesFrom(from Square) []Move {
	moves := []Move{}
	piece := gs.Board[from.File][from.Rank]

	switch piece.PieceType() {
	case Pawn:
		moves = append(moves, gs.Board.PawnMovesFrom(from, piece.Player(), gs.EnPassanteFile)...)
	case Knight:
		moves = append(moves, gs.Board.KnightMovesFrom(from, piece.Player())...)
	case Bishop:
		moves = append(moves, gs.Board.BishopMovesFrom(from, piece.Player())...)
	case Rook:
		moves = append(moves, gs.Board.RookMovesFrom(from, piece.Player())...)
	case Queen:
		moves = append(moves, gs.Board.QueenMovesFrom(from, piece.Player())...)
	case King:
		moves = append(moves, gs.Board.KingMovesFrom(from, piece.Player())...)
	}

	return moves
}
