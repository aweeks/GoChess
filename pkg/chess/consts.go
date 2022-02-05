package chess

import (
	"strconv"
)

type Player int8

const (
	NilPlayer Player = 0

	White = 1
	Black = -1
)

func (p Player) String() string {
	switch p {
	case White:
		return "w"
	case Black:
		return "b"
	default:
		return " "
	}
}

type PieceType int8

const (
	NilPieceType = 0

	Pawn   = 1
	Knight = 2
	Bishop = 3
	Rook   = 4
	Queen  = 5
	King   = 6
)

func (pt PieceType) String() string {
	switch pt {
	case Pawn:
		return "P"
	case Knight:
		return "N"
	case Bishop:
		return "B"
	case Rook:
		return "R"
	case Queen:
		return "Q"
	case King:
		return "K"
	default:
		return " "
	}
}

type Piece int8

const (
	NilPiece Piece = 0

	WhitePawn   = Pawn
	WhiteKnight = Knight
	WhiteBishop = Bishop
	WhiteRook   = Rook
	WhiteQueen  = Queen
	WhiteKing   = King

	BlackPawn   = -Pawn
	BlackKnight = -Knight
	BlackBishop = -Bishop
	BlackRook   = -Rook
	BlackQueen  = -Queen
	BlackKing   = -King
)

func (p Piece) Player() Player {
	if p > 0 {
		return White
	} else if p < 0 {
		return Black
	} else {
		return NilPlayer
	}
}

func (p Piece) PieceType() PieceType {
	if p < 0 {
		return PieceType(-p)
	} else {
		return PieceType(p)
	}
}

func (p Piece) ShortString() string {
	switch p {
	case WhitePawn:
		return "wP"
	case WhiteKnight:
		return "wN"
	case WhiteBishop:
		return "wB"
	case WhiteRook:
		return "wR"
	case WhiteQueen:
		return "wQ"
	case WhiteKing:
		return "wK"
	case BlackPawn:
		return "bP"
	case BlackKnight:
		return "bN"
	case BlackBishop:
		return "bB"
	case BlackRook:
		return "bR"
	case BlackQueen:
		return "bQ"
	case BlackKing:
		return "bK"
	default:
		return "--"
	}
}

type Rank int

const (
	NilRank Rank = -1
	_1      Rank = 0
	_2      Rank = 1
	_3      Rank = 2
	_4      Rank = 3
	_5      Rank = 4
	_6      Rank = 5
	_7      Rank = 6
	_8      Rank = 7
)

func (r Rank) String() string {
	if r == NilRank {
		return "Nil"
	} else {
		return strconv.Itoa(int(r) + 1)
	}
}

type File int

const (
	NilFile File = -1
	A_      File = 0
	B_      File = 1
	C_      File = 2
	D_      File = 3
	E_      File = 4
	F_      File = 5
	G_      File = 6
	H_      File = 7
)

func (f File) String() string {
	if f == NilFile {
		return "Nil"
	} else {
		return string(rune((int(f) + int('A'))))
	}
}

type Square struct {
	File File
	Rank Rank
}

var NilSquare Square = Square{NilFile, NilRank}

func (s Square) String() string {
	return s.File.String() + s.Rank.String()
}

func AddSquares(a Square, b Square) Square {
	return Square{a.File + b.File, a.Rank + b.Rank}
}

func (s Square) AddTo(o Square) {
	s.File += o.File
	s.Rank += o.Rank
}
