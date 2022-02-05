package chess

import (
	"strings"
)

type Board [8][8]Piece

func (b Board) String() string {
	return b.PrettyBoard(make(map[Square]bool))
}

func (b Board) PrettyBoard(highlight map[Square]bool) string {
	var buff strings.Builder

	for rank := _8; rank >= _1; rank-- {
		/* Header */
		if rank == _8 {
			buff.WriteString("  ")
			for file := A_; file <= H_; file++ {
				buff.WriteString(file.String() + "  ")
			}
			buff.WriteString("\n")
		}

		buff.WriteString(rank.String() + " ")

		/* Pieces */
		for file := A_; file <= H_; file++ {

			pieceString := b[file][rank].ShortString()

			if h, ok := highlight[Square{file, rank}]; ok && h {
				pieceString = "**"
			}

			buff.WriteString(pieceString + " ")
		}

		buff.WriteString(rank.String())
		buff.WriteString("\n")

		/* Footer */
		if rank == _1 {
			buff.WriteString("  ")
			for file := A_; file <= H_; file++ {
				buff.WriteString(file.String() + "  ")
			}
			buff.WriteString("\n")
		}
	}
	return buff.String()
}

// PlayerSquares() returns a list of Squares occupied by the given player
func (b Board) PlayerSquares(p Player) []Square {
	squares := []Square{}
	for rank := _1; rank <= _8; rank++ {
		for file := A_; file <= H_; file++ {
			if b[file][rank].Player() == p {
				squares = append(squares, Square{file, rank})
			}
		}
	}
	return squares
}

func (b Board) shapedMoveFrom(from Square, p Player, maxHops int, shape Square) []Move {
	moves := []Move{}

	for to := AddSquares(from, shape); to.OnBoard(); to = AddSquares(to, shape) {

		if b[to.File][to.Rank] == NilPiece {
			moves = append(moves, Move{from, to, NilPieceType})
		} else {
			if b[to.File][to.Rank].Player() == -p {
				moves = append(moves, Move{from, to, NilPieceType}) // Capture enemy piece
			}
			break // Stop after hitting first piece
		}

		maxHops--
		if maxHops == 0 {
			break
		}
	}
	return moves
}

func (b Board) MovesFrom(from Square, p Player, maxHops int, shapes ...Square) []Move {
	moves := []Move{}

	for _, s := range shapes {
		moves = append(moves, b.shapedMoveFrom(from, p, maxHops, s)...)
	}

	return moves
}

func (b Board) PawnMovesFrom(from Square, p Player, enPassanteFile File) []Move {
	pawnMove := WhitePawnMove
	pawnDoubleMove := WhitePawnDoubleMove
	pawnCaptureMoves := WhitePawnCaptureMoves
	if p == Black {
		pawnMove = BlackPawnMove
		pawnDoubleMove = BlackPawnDoubleMove
		pawnCaptureMoves = BlackPawnCaptureMoves
	}

	moves := []Move{}

	/* Single pawn move */
	moves = append(moves, b.MovesFrom(from, p, 1, pawnMove)...)

	/* Double pawn move */
	if from.Rank == _2 && p == White || from.Rank == _7 && p == Black {
		moves = append(moves, b.MovesFrom(from, p, 1, pawnDoubleMove)...)
	}

	/* Capture moves */
	for _, s := range pawnCaptureMoves {
		captureSquare := AddSquares(from, s)

		if !captureSquare.OnBoard() {
			continue
		}

		/* Normal captures */
		if b[captureSquare.File][captureSquare.Rank].Player() == -p {
			moves = append(moves, Move{from, captureSquare, NilPieceType})
		}

		/* En Passante captures */
		if captureSquare.File == enPassanteFile && (from.Rank == _4 && p == White || from.Rank == _5 && p == Black) {
			moves = append(moves, Move{from, captureSquare, NilPieceType})
		}
	}

	return moves
}

func (b Board) KnightMovesFrom(from Square, p Player) []Move {
	return b.MovesFrom(from, p, 1, KnightMoves...)
}

func (b Board) BishopMovesFrom(from Square, p Player) []Move {
	return b.MovesFrom(from, p, -1, BishopMoves...)
}

func (b Board) RookMovesFrom(from Square, p Player) []Move {
	return b.MovesFrom(from, p, -1, RookMoves...)
}

func (b Board) QueenMovesFrom(from Square, p Player) []Move {
	return b.MovesFrom(from, p, -1, QueenMoves...)
}

func (b Board) KingMovesFrom(from Square, p Player) []Move {
	return b.MovesFrom(from, p, 1, KingMoves...)
}

func (b Board) OOMove(from Square, p Player) []Move {
	return b.MovesFrom(from, p, 1, KingMoves...)
}

func (b Board) OOOMove(from Square, p Player) []Move {
	return b.MovesFrom(from, p, 1, KingMoves...)
}
