package chess

var StartingPosition Board = [8][8]Piece{
	{WhiteRook, WhitePawn, NilPiece, NilPiece, NilPiece, NilPiece, BlackPawn, BlackRook},
	{WhiteKnight, WhitePawn, NilPiece, NilPiece, NilPiece, NilPiece, BlackPawn, BlackKnight},
	{WhiteBishop, WhitePawn, NilPiece, NilPiece, NilPiece, NilPiece, BlackPawn, BlackBishop},
	{WhiteQueen, WhitePawn, NilPiece, NilPiece, NilPiece, NilPiece, BlackPawn, BlackQueen},
	{WhiteKing, WhitePawn, NilPiece, NilPiece, NilPiece, NilPiece, BlackPawn, BlackKing},
	{WhiteBishop, WhitePawn, NilPiece, NilPiece, NilPiece, NilPiece, BlackPawn, BlackBishop},
	{WhiteKnight, WhitePawn, NilPiece, NilPiece, NilPiece, NilPiece, BlackPawn, BlackKnight},
	{WhiteRook, WhitePawn, NilPiece, NilPiece, NilPiece, NilPiece, BlackPawn, BlackRook},
}

var StartingGameState GameState = GameState{
	Board:          StartingPosition,
	Turn:           White,
	WhiteCanOO:     true,
	WhiteCanOOO:    true,
	BlackCanOO:     true,
	BlackCanOOO:    true,
	EnPassanteFile: NilFile,
	Result:         NilResult,
	ResultComment:  NilResultComment,
}

var RookKingOnlyPosition Board = [8][8]Piece{
	{WhiteRook, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, BlackRook},
	{NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece},
	{NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece},
	{NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece},
	{WhiteKing, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, BlackKing},
	{NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece},
	{NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece},
	{WhiteRook, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, BlackRook},
}

var RookKingOnlyState GameState = GameState{
	Board:          RookKingOnlyPosition,
	Turn:           White,
	WhiteCanOO:     true,
	WhiteCanOOO:    true,
	BlackCanOO:     true,
	BlackCanOOO:    true,
	EnPassanteFile: NilFile,
	Result:         NilResult,
	ResultComment:  NilResultComment,
}

var NoPawnsPosition Board = [8][8]Piece{
	{WhiteRook, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, BlackRook},
	{WhiteKnight, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, BlackKnight},
	{WhiteBishop, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, BlackBishop},
	{WhiteQueen, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, BlackQueen},
	{WhiteKing, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, BlackKing},
	{WhiteBishop, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, BlackBishop},
	{WhiteKnight, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, BlackKnight},
	{WhiteRook, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, NilPiece, BlackRook},
}

var NoNoPawnsState GameState = GameState{
	Board:          NoPawnsPosition,
	Turn:           White,
	WhiteCanOO:     true,
	WhiteCanOOO:    true,
	BlackCanOO:     true,
	BlackCanOOO:    true,
	EnPassanteFile: NilFile,
	Result:         NilResult,
	ResultComment:  NilResultComment,
}
