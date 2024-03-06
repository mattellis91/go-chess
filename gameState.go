package main

type BoardState [8][8]string

type GameState struct {
	Board BoardState
	WhiteToMove bool
	MoveLog []Move
	SquareSelected Square
	PlayerClicks []Square
}

func NewGameState() *GameState {
	return &GameState{
		Board: BoardState{
			{"bR", "bN", "bB", "bQ", "bK", "bB", "bN", "bR"},
			{"bp", "bp", "bp", "bp", "bp", "bp", "bp", "bp"},
			{"--", "--", "--", "--", "--", "--", "--", "--"},
			{"--", "--", "--", "--", "--", "--", "--", "--"},
			{"--", "--", "--", "--", "--", "--", "--", "--"},
			{"--", "--", "--", "--", "--", "--", "--", "--"},
			{"wp", "wp", "wp", "wp", "wp", "wp", "wp", "wp"},
			{"wR", "wN", "wB", "wQ", "wK", "wB", "wN", "wR"},
		},
		WhiteToMove: true,
		MoveLog: []Move{},
	}
}

func (gs *GameState) MakeMove(move Move) {
	gs.Board[move.StartRow][move.StartCol] = "--"
	gs.Board[move.EndRow][move.EndCol] = move.PieceMoved
	gs.MoveLog = append(gs.MoveLog, move)
	gs.WhiteToMove = !gs.WhiteToMove
}

func (gs *GameState) UndoMove() {
	if len(gs.MoveLog) == 0 {
		return
	}
	move := gs.MoveLog[len(gs.MoveLog)-1]
	gs.Board[move.StartRow][move.StartCol] = move.PieceMoved
	gs.Board[move.EndRow][move.EndCol] = move.PieceCaptured
	gs.MoveLog = gs.MoveLog[:len(gs.MoveLog)-1]
	gs.WhiteToMove = !gs.WhiteToMove
}