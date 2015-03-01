package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"bufio"
)

const DEBUG = false

const none = 0
const marker = 10
const taken = 11

type ChessPlayer int
const (
	white = 1
	black = -1
)
var strToPlayer = map[string]ChessPlayer {
	"w": white,
	"b": black,
	" ": none,
	"*": marker,
	"X": taken,
}
var playerToStr= map[ChessPlayer]string {
	white: "w",
	black: "b",
	none: " ",
	marker: "*",
	taken:  "X",
}

func (p ChessPlayer) String() string {
	return playerToStr[p]

}

func (g ChessPlayer) Invert() ChessPlayer {
	return -g
}

const (
	NEG = "\x1b[7m"
	BOLD = "\x1b[1m"
	RES = "\x1b[0m"
	FGBLACK = "\x1b[30m"
	FGRED = "\x1b[31m"
	FGGREEN = "\x1b[32m"
	FGYELLOW = "\x1b[33m"
	FGBLUE = "\x1b[34m"
	FGMAGENTA = "\x1b[35m"
	FGCYAN = "\x1b[36m"
	FGWHITE = "\x1b[37m"

	BGBLACK= "\x1b[40m"
	BGRED = "\x1b[41m"
	BGGREEN = "\x1b[42m"
	BGYELLOW = "\x1b[43m"
	BGBLUE = "\x1b[44m"
	BGMAGENTA = "\x1b[45m"
	BGCYAN = "\x1b[46m"
	BGWHITE = "\x1b[47m"
)

type ChessPieceType int
const (
	king = 1
	queen = 2
	rook = 3
	bishop = 4
	knight = 5
	pawn = 6
)
var strToPieceType = map[string]ChessPieceType {
	"K": king,
	"Q": queen,
	"R": rook,
	"B": bishop,
	"N": knight,
	"P": pawn,
	" ": none,
	"*": marker,
	"X": taken,
}
var pieceTypeToStr = map[ChessPieceType]string {
	king:   "K",
	queen:  "Q",
	rook:   "R",
	bishop: "B",
	knight: "N",
	pawn:   "P",
	none:   " ",
	marker: "*",
	taken:  "X",
}



func (p ChessPieceType) String() string {
	return pieceTypeToStr[p]

}

type ChessPiece struct {
	piece_type ChessPieceType
	player ChessPlayer
}

var pieceToStr = map[ChessPiece]string {
	whiteKing:   "♔",
	whiteQueen:  "♕",
	whiteRook:   "♖",
	whiteBishop: "♗",
	whiteKnight: "♘",
	whitePawn:   "♙",
	blackKing:   "♚",
	blackQueen:  "♛",
	blackRook:   "♜",
	blackBishop: "♝",
	blackKnight: "♞",
	blackPawn:   "♟",
	empty:       " ",
	markerPiece: "*",
	takenPiece:  "x",
}

func (p ChessPiece) String() string {
	return pieceToStr[p]
	// return fmt.Sprintf(" %s ", p.piece_type.String())
}



const (
 	_A = 0
 	_B = 1
 	_C = 2
 	_D = 3
 	_E = 4
 	_F = 5
 	_G = 6
 	_H = 7
)
const (
 	_1 = 0
 	_2 = 1
 	_3 = 2
 	_4 = 3
 	_5 = 4
 	_6 = 5
 	_7 = 6
 	_8 = 7
)

var fileToStr = map[int]string {
	_A: "a",
	_B: "b",
	_C: "c",
	_D: "d",
	_E: "e",
	_F: "f",
	_G: "g",
	_H: "h",
}

var strToFile = map[string]int {
	"a": _A,
	"b": _B,
	"c": _C,
	"d": _D,
	"e": _E,
	"f": _F,
	"g": _G,
	"h": _H,
}


var rankToStr = map[int]string {
	_1: "1",
	_2: "2",
	_3: "3",
	_4: "4",
	_5: "5",
	_6: "6",
	_7: "7",
	_8: "8",
}

var strToRank = map[string]int {
	"1": _1,
	"2": _2,
	"3": _3,
	"4": _4,
	"5": _5,
	"6": _6,
	"7": _7,
	"8": _8,
}

type ChessSquare struct {
	file int
	rank int
}

func ChessSquareAlgebraic(s string) (ChessSquare, error) {
	var file string
	var rank string
	var err error

	_, err = fmt.Sscanf(s, "%1s%1s", &file, &rank)

	return ChessSquare{strToFile[file], strToRank[rank]}, err

}

func (cs ChessSquare) String() string {
	return fmt.Sprintf("%s%s", fileToStr[cs.file], rankToStr[cs.rank])
}


func (cs ChessSquare) BoundCheck() bool {
	return cs.file >= 0 && cs.file < 8 && cs.rank >= 0 && cs.rank < 8
}

func (cs ChessSquare) SlideN(fileInc int, rankInc int, n int) []ChessSquare {
	var squares []ChessSquare
	for c := 1; c <= n; c++ {
		var s ChessSquare
		s.file = cs.file + c*fileInc
		s.rank = cs.rank + c*rankInc
		if s.BoundCheck() {
			squares = append(squares, s)
		} else {
			break
		}
	}
	return squares
}

func (cs ChessSquare) Slide(fileInc int, rankInc int) []ChessSquare {
	return cs.SlideN(fileInc, rankInc, 8)
}

func (cs ChessSquare) White() bool {
	return (cs.file + cs.rank) %2 == 0
}

type ChessMove struct {
	initial ChessSquare
	final ChessSquare
	promote ChessPieceType
}

func (cm ChessMove) String() string {
	return fmt.Sprintf("%s->%s", cm.initial.String(), cm.final.String())

}

func ChessMoveAlgebraic(s string) (ChessMove, error) {
	var s1 string
	var s2 string

	var err error

	_, err = fmt.Sscanf(s, "%2s %2s", &s1, &s2)
	if err != nil {
		return ChessMove{}, err
	}

	var m1 ChessSquare
	var m2 ChessSquare

	m1, err = ChessSquareAlgebraic(s1)
	m2, err = ChessSquareAlgebraic(s2)


	return ChessMove{m1, m2, none}, err

}

type ChessGameState struct {
	turn  ChessPlayer
	board [8][8]ChessPiece
	move ChessMove
	prev *ChessGameState
}

var empty = ChessPiece{none, none}

var whiteKing = ChessPiece{king, white}
var whiteQueen = ChessPiece{queen, white}
var whiteRook = ChessPiece{rook, white}
var whiteBishop = ChessPiece{bishop, white}
var whiteKnight = ChessPiece{knight, white}
var whitePawn = ChessPiece{pawn, white}

var blackKing = ChessPiece{king, black}
var blackQueen = ChessPiece{queen, black}
var blackRook = ChessPiece{rook, black}
var blackBishop = ChessPiece{bishop, black}
var blackKnight = ChessPiece{knight, black}
var blackPawn = ChessPiece{pawn, black}

var markerPiece = ChessPiece{marker, marker}
var takenPiece = ChessPiece{taken, taken}


var startingBoard = ChessGameState{
	white,
	[8][8]ChessPiece{
		{whiteRook,   whitePawn, empty, empty, empty, empty, blackPawn, blackRook  },
		{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
		{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
		{whiteQueen,  whitePawn, empty, empty, empty, empty, blackPawn, blackQueen },
		{whiteKing,   whitePawn, empty, empty, empty, empty, blackPawn, blackKing  },
		{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
		{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
		{whiteRook,   whitePawn, empty, empty, empty, empty, blackPawn, blackRook  },
	},
	ChessMove{},
	nil,
}

var pawnBoard = ChessGameState{
	white,
	[8][8]ChessPiece{
		{empty, whitePawn, empty, empty, empty, empty, blackPawn, empty},
		{empty, whitePawn, empty, empty, empty, empty, blackPawn, empty},
		{empty, whitePawn, empty, empty, empty, empty, blackPawn, empty},
		{empty, whitePawn, empty, empty, empty, empty, blackPawn, empty},
		{empty, whitePawn, empty, empty, empty, empty, blackPawn, empty},
		{empty, whitePawn, empty, empty, empty, empty, blackPawn, empty},
		{empty, whitePawn, empty, empty, empty, empty, blackPawn, empty},
		{empty, whitePawn, empty, empty, empty, empty, blackPawn, empty},
	},
	ChessMove{},
	nil,
}

var knightBoard = ChessGameState{
	white,
	[8][8]ChessPiece{
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, whiteKnight, empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
	},
	ChessMove{},
	nil,
}

var rookBoard = ChessGameState{
	white,
	[8][8]ChessPiece{
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, whiteRook, empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
	},
	ChessMove{},
	nil,
}

var bishopBoard = ChessGameState{
	white,
	[8][8]ChessPiece{
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, whiteBishop, empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
		{empty, empty, empty, empty,       empty, empty, empty, empty},
	},
	ChessMove{},
	nil,
}

var queenBoard = ChessGameState{
	white,
	[8][8]ChessPiece{
		{empty, empty, empty, empty,      empty, empty, empty,     empty},
		{empty, empty, empty, empty,      empty, empty, empty,     empty},
		{empty, empty, empty, empty,      empty, empty, empty,     empty},
		{empty, empty, empty, whiteQueen, empty, empty, blackPawn, empty},
		{empty, empty, empty, empty,      empty, empty, empty,     empty},
		{empty, whitePawn, empty, empty,      empty, empty, empty,     empty},
		{empty, whitePawn, empty,  empty, empty, empty,     empty},
		{empty, empty, empty, empty,      empty, empty, empty,     empty},
	},
	ChessMove{},
	nil,
}

var kingBoard = ChessGameState{
	white,
	[8][8]ChessPiece{
		{empty, empty, empty, empty,     empty, empty, empty, empty},
		{empty, empty, empty, empty,     empty, empty, empty, empty},
		{empty, empty, empty, empty,     empty, empty, empty, empty},
		{empty, empty, empty, whiteKing, empty, empty, empty, empty},
		{empty, empty, empty, empty,     empty, empty, empty, empty},
		{empty, empty, empty, empty,     empty, empty, empty, empty},
		{empty, empty, empty, empty,     empty, empty, empty, empty},
		{empty, empty, empty, empty,     empty, empty, empty, empty},
	},
	ChessMove{},
	nil,
}

func DupChessGameState(state *ChessGameState) *ChessGameState{
	var dup = new(ChessGameState)
	*dup = *state
	return dup
}

func (gs *ChessGameState) Moved(cs ChessSquare) bool {
	/*
		Has the piece at the given square been moved?
	*/
	if gs.prev == nil {
		return false
	} else {
		return gs.GetPiece(cs) != gs.prev.GetPiece(cs) || gs.prev.Moved(cs)
	}
}

func (gs *ChessGameState) GetPiece(cs ChessSquare) ChessPiece {
	/*
		Get the piece at the given square.
	*/
	return gs.board[cs.file][cs.rank]
}

func (gs *ChessGameState) GetPieceAlgebraic(file int, rank int) ChessPiece {
	/*
		Get the piece at the given file and rank.
	*/
	return gs.GetPiece(ChessSquare{file, rank})
}

func (gs *ChessGameState) BoardSquares() [64]ChessSquare {
	var squares [64]ChessSquare
	for file := range gs.board {
		for rank := range gs.board[file] {
			squares[8*file + rank] = ChessSquare{file, rank}
		}
	}
	return squares
}

func (gs *ChessGameState) Occupied(cs ChessSquare) bool {
	/*
		Is the given square occupied?
	*/
	return gs.GetPiece(cs) != empty
}

func (gs *ChessGameState) OccupiedByPlayer(cs ChessSquare, p ChessPlayer) bool {
	/*
		Is the given square occipied by a piece of the given player?
	*/
	return gs.GetPiece(cs) != empty && gs.GetPiece(cs).player == p
}

var kingIncs = [8][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
                           {1, 0}, {-1, 0}, {0, 1},  {0, -1}}

func (gs *ChessGameState) CandidateKingMoves(initial ChessSquare) []ChessMove {
	var moves []ChessMove

	initial_piece := gs.GetPiece(initial)

	if initial_piece.piece_type!= king || initial_piece.player != gs.turn {
		return  moves
	}

	for _, inc := range(kingIncs) {

		for _, f := range(initial.SlideN(inc[0], inc[1], 1)) {
			if gs.OccupiedByPlayer(f, gs.turn.Invert()) {
				moves = append(moves, ChessMove{initial, f, none})
				break
			} else if gs.Occupied(f) {
				break
			} else {
				moves = append(moves, ChessMove{initial, f, none})
			}
		}
	}

	return moves
}

var queenIncs = [8][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
                           {1, 0}, {-1, 0}, {0, 1},  {0, -1}}

func (gs *ChessGameState) CandidateQueenMoves(initial ChessSquare) []ChessMove{
	var moves []ChessMove

	initial_piece := gs.GetPiece(initial)

	if initial_piece.piece_type!= queen || initial_piece.player != gs.turn {
		return  moves
	}

	for _, inc := range(queenIncs) {

		for _, f := range(initial.Slide(inc[0], inc[1])) {
			if gs.OccupiedByPlayer(f, gs.turn.Invert()) {
				moves = append(moves, ChessMove{initial, f, none})
				break
			} else if gs.Occupied(f) {
				break
			} else {
				moves = append(moves, ChessMove{initial, f, none})
			}
		}
	}

	return moves
}

var rookIncs = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func (gs *ChessGameState) CandidateRookMoves(initial ChessSquare) []ChessMove{
	var moves []ChessMove

	initial_piece := gs.GetPiece(initial)

	if initial_piece.piece_type!= rook || initial_piece.player != gs.turn {
		return  moves
	}

	for _, inc := range(rookIncs) {

		for _, f := range(initial.Slide(inc[0], inc[1])) {
			if gs.OccupiedByPlayer(f, gs.turn.Invert()) {
				moves = append(moves, ChessMove{initial, f, none})
				break
			} else if gs.Occupied(f) {
				break
			} else {
				moves = append(moves, ChessMove{initial, f, none})
			}
		}
	}

	return moves

}

var bishopIncs = [4][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func (gs *ChessGameState) CandidateBishopMoves(initial ChessSquare) []ChessMove{
	var moves []ChessMove

	initial_piece := gs.GetPiece(initial)

	if initial_piece.piece_type!= bishop || initial_piece.player != gs.turn {
		return  moves
	}

	for _, inc := range(bishopIncs) {

		for _, f := range(initial.Slide(inc[0], inc[1])) {
			if gs.OccupiedByPlayer(f, gs.turn.Invert()) {
				moves = append(moves, ChessMove{initial, f, none})
				break
			} else if gs.Occupied(f) {
				break
			} else {
				moves = append(moves, ChessMove{initial, f, none})
			}
		}
	}

	return moves
}

var knightIncs = [8][2]int{
	{1,2},
	{2,1},
	{-1,2},
	{-2,1},
	{-1,-2},
	{-2,-1},
	{1,-2},
	{2,-1},
}

func (gs *ChessGameState) CandidateKnightMoves(initial ChessSquare) []ChessMove{
	var moves []ChessMove

	initial_piece := gs.GetPiece(initial)

	if initial_piece.piece_type!= knight || initial_piece.player != gs.turn {
		return  moves
	}

	for _, inc := range(knightIncs) {

		for _, f := range(initial.SlideN(inc[0], inc[1], 1)) {
			if gs.OccupiedByPlayer(f, gs.turn.Invert()) {
				moves = append(moves, ChessMove{initial, f, none})
				break
			} else if gs.Occupied(f) {
				break
			} else {
				moves = append(moves, ChessMove{initial, f, none})
			}
		}
	}

	return moves
}

var pawnSingleInc = [2]int{0,1}
var pawnDoubleInc = [2]int{0,2}
var pawnCaptureInc = [2]int{1,1}

func (gs *ChessGameState) CandidatePawnMoves(initial ChessSquare) []ChessMove{
	var moves []ChessMove

	initial_piece := gs.GetPiece(initial)

	if initial_piece.piece_type!= pawn || initial_piece.player != gs.turn {
		return  moves
	}

	var final ChessSquare

	var dirMult = 0
	switch initial_piece.player {
		case white:
			dirMult = 1
		case black:
			dirMult = -1
	}

	final = ChessSquare{initial.file, initial.rank + 1 * dirMult}
	if final.BoundCheck() && !gs.Occupied(final) {
		moves = append(moves, ChessMove{initial, final, none})
	}

	// Capture right
	final = ChessSquare{initial.file+1, initial.rank + 1 * dirMult}
	if final.BoundCheck() && gs.OccupiedByPlayer(final, gs.turn.Invert()) {
		moves = append(moves, ChessMove{initial, final, none})
	}

	// Capture left
	final = ChessSquare{initial.file-1, initial.rank + 1 * dirMult}
	if final.BoundCheck() && gs.OccupiedByPlayer(final, gs.turn.Invert()) {
		moves = append(moves, ChessMove{initial, final, none})
	}

	final = ChessSquare{initial.file, initial.rank + 2 * dirMult}
	if !gs.Moved(initial) && final.BoundCheck() && !gs.Occupied(final) {
		moves = append(moves, ChessMove{initial, final, none})
	}



	return moves
}

func (gs *ChessGameState) Threatened(s ChessSquare) bool {
	for _, m := range(gs.CandidateMoves()) {
		if m.final == s {
			return true
		}
	}
	return false
}

func (gs *ChessGameState) KingSquare(p ChessPlayer) ChessSquare {
	var kingPiece = ChessPiece{king, p}

	for _, square := range(gs.BoardSquares()) {
		if gs.GetPiece(square) == kingPiece {
			return square
		}
	}
	return ChessSquare{-1, -1}
}
func (gs *ChessGameState) KingThreatened(p ChessPlayer) bool {
	return gs.Threatened(gs.KingSquare(p))
}

func (gs *ChessGameState) CandidateMoves() []ChessMove {

	var moves []ChessMove

	for _, initial := range(gs.BoardSquares()) {
		var tmp []ChessMove

		switch gs.GetPiece(initial).piece_type {
		case king:
			tmp = gs.CandidateKingMoves(initial)
		case queen:
			tmp = gs.CandidateQueenMoves(initial)
		case rook:
			tmp = gs.CandidateRookMoves(initial)
		case bishop:
			tmp = gs.CandidateBishopMoves(initial)
		case knight:
			tmp = gs.CandidateKnightMoves(initial)
		case pawn:
			tmp = gs.CandidatePawnMoves(initial)
		}
		moves = append(moves, tmp...)

	}
	return moves
}

func (gs *ChessGameState) LegalMoves() []ChessMove {
	var moves[]ChessMove

	for _, c := range(gs.CandidateMoves()) {
		err := gs.CheckLegalMove(c)
		if err == nil {
			moves = append(moves, c)
		}
	}
	return moves
}

func In(moves []ChessMove, m ChessMove) bool {
	for _, c := range(moves) {
		if m == c {
			return true
		}
	}
	return false
}

func (gs *ChessGameState) CheckLegalMove(m ChessMove) error {
	if !In(gs.CandidateMoves(), m) {
		return errors.New("Invalid move")
	}

	if gs.MoveUnconditional(m).KingThreatened(gs.turn) {
		return errors.New("In check")
	}

	return nil

}

func (gs *ChessGameState) Move(m ChessMove) (*ChessGameState, error) {
	var err = gs.CheckLegalMove(m)

	if err != nil {
		return nil, err
	}
	return gs.MoveUnconditional(m), nil
}

func (gs *ChessGameState) MoveUnconditional(m ChessMove) *ChessGameState {
	var next = DupChessGameState(gs)
	next.board[m.final.file][m.final.rank] = gs.board[m.initial.file][m.initial.rank]
	next.board[m.initial.file][m.initial.rank] = empty
	next.turn = -gs.turn
	next.prev = gs
	return next
}

func (gs *ChessGameState) HighlightMoves(moves []ChessMove) *ChessGameState {
	var next = DupChessGameState(gs)

	for _, m := range(moves) {
		if next.Occupied(m.final) {
			next.board[m.final.file][m.final.rank] = takenPiece
		} else {
			next.board[m.final.file][m.final.rank] = markerPiece
		}
	}

	next.turn = -gs.turn
    next.prev = gs
	return next
}


func (gs *ChessGameState) MoveAlgebraic(f1 int, r1 int, f2 int, r2 int) (*ChessGameState, error) {
	return gs.Move(ChessMove{ChessSquare{f1, r1}, ChessSquare{f2, r2}, none})

}

func (gs ChessGameState) String() string {
	var buff bytes.Buffer

	buff.WriteString("   a  b  c  d  e  f  g  h\n")

	for rank := 7; rank>=0; rank-- {
		buff.WriteString(fmt.Sprintf("%d ", rank+1))

		for file := 0; file<8; file++ {
			var fmtStr string
			if (file + rank) % 2 == 1 {
				fmtStr = BGWHITE  + FGBLACK + " %s " + RES

			} else {
				fmtStr = BGCYAN + FGBLACK + " %s " + RES
			}
			buff.WriteString( fmt.Sprintf(fmtStr, gs.board[file][rank].String()))
		}
		buff.WriteString(fmt.Sprintf(" %d\n", rank+1))
	}
	buff.WriteString("   a  b  c  d  e  f  g  h\n")
	return buff.String()
}

func main() {
	var board = DupChessGameState(&startingBoard)
	reader := bufio.NewReader(os.Stdin)

	for {

		if DEBUG {
			var moves = make(map[ChessSquare][]ChessMove)
			for _, m := range(board.LegalMoves()) {
				moves[m.initial] = append(moves[m.initial], m)
			}
			for initial, ms := range(moves) {
				fmt.Println("========= " + initial.String() + " MOVES ========\n")
				fmt.Println(board.HighlightMoves(ms))
			}

		}

		fmt.Println("======= CURRENT BOARD ======\n")
		fmt.Println(board.String())
		fmt.Println("King square: " + board.KingSquare(board.turn).String())

		for {
			fmt.Print("move? ")
			input, _ := reader.ReadString('\n')
			move, err := ChessMoveAlgebraic(input)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(move)

			nextBoard, err := board.Move(move)
			if err != nil {
				fmt.Println(err)
			} else {
				board = nextBoard
				break
			}
		}

		oppMoves := board.LegalMoves()
		board, _ = board.Move(oppMoves[rand.Intn(len(oppMoves))])

	}

}
