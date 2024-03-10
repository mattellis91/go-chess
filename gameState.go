package main

type BoardState [8][8]string

type GameState struct {
	Board             BoardState
	WhiteToMove       bool
	MoveLog           []Move
	SquareSelected    Square
	PlayerClicks      []Square
	ValidMoves        []Move
	MoveMade          bool
	HiglightedSquares []Square
	BlackKingSquare   Square
	WhiteKingSquare   Square
	Checkmate		 bool
	Stalemate		 bool
}

type PieceDelta struct {
	row int
	col int
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
		WhiteToMove:    true,
		MoveMade:       false,
		SquareSelected: GetNullSquare(),
	}
}

func GetNullSquare() Square {
	return Square{-1, -1}
}
 
func (gs *GameState) MakeMove(move Move) {
	gs.Board[move.StartRow][move.StartCol] = "--"
	gs.Board[move.EndRow][move.EndCol] = move.PieceMoved
	gs.MoveLog = append(gs.MoveLog, move)

	if move.PieceMoved == "wK" {
		gs.WhiteKingSquare = Square{move.EndRow, move.EndCol}
	} else if move.PieceMoved == "bK" {
		gs.BlackKingSquare = Square{move.EndRow, move.EndCol}
	}

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

	if move.PieceMoved == "wK" {
		gs.WhiteKingSquare = Square{move.StartRow, move.StartCol}
	} else if move.PieceMoved == "bK" {
		gs.BlackKingSquare = Square{move.StartRow, move.StartCol}
	}

	gs.WhiteToMove = !gs.WhiteToMove
}

func (gs *GameState) GetValidMoves() []Move {
	moves := gs.GetAllPossibleMoves()
	for i := len(moves) - 1; i >= 0; i-- {
		gs.MakeMove(moves[i])
		gs.WhiteToMove = !gs.WhiteToMove
		if gs.InCheck() {
			moves = append(moves[:i], moves[i+1:]...)
		}
		gs.WhiteToMove = !gs.WhiteToMove
		gs.UndoMove()
	}

	if len(moves) == 0 {
		if gs.InCheck() {
			gs.Checkmate = true
		} else {
			gs.Stalemate = true
		}
	} else {
		gs.Checkmate = false
		gs.Stalemate = false
	}
	
	return moves
}

func (gs *GameState) IsValidMove(move Move) bool {
	for _, validMove := range gs.ValidMoves {
		if move.MoveId == validMove.MoveId {
			return true		
		}
	}
	return false
}

func (gs *GameState) SquareAlreadyHighlighted(square Square) bool {
	for _, currentSquare := range gs.HiglightedSquares {
		if square == currentSquare {
			return true
		}
	}
	return false
}

func (gs *GameState) RemoveSquareFromSlice(squares []Square, square Square) []Square {
	for i, currentSquare := range squares {
		if currentSquare == square {
			return append(squares[:i], squares[i+1:]...)
		}
	}
	return squares
}

func (gs *GameState) GetAllPossibleMoves() []Move {
	moves := []Move{}
	for r := 0; r < len(gs.Board); r++ {
		for c := 0; c < len(gs.Board[r]); c++ {
			turn := gs.Board[r][c][0]
			if (turn == 'w' && gs.WhiteToMove) || (turn == 'b' && !gs.WhiteToMove) {
				piece := gs.Board[r][c][1]
				switch piece {
				case 'p':
					moves = append(moves, gs.GetPawnMoves(r, c)...)
				case 'R':
					moves = append(moves, gs.GetRookMoves(r, c)...)
				case 'N':
					moves = append(moves, gs.GetKnightMoves(r, c)...)
				case 'B':
					moves = append(moves, gs.GetBishopMoves(r, c)...)
				case 'Q':
					moves = append(moves, gs.GetQueenMoves(r, c)...)
				case 'K':
					moves = append(moves, gs.GetKingMoves(r, c)...)
				}
			}
		}
	}
	return moves
}

func (gs *GameState) GetPawnMoves(r int, c int) []Move {
	moves := []Move{}
	if gs.WhiteToMove {
		if r-1 >= 0 && gs.Board[r-1][c] == "--" { //move one square
			moves = append(moves, NewMove(Square{r, c}, Square{r - 1, c}, gs.Board))
			if r == 6 && gs.Board[r-2][c] == "--" { //move two squares
				moves = append(moves, NewMove(Square{r, c}, Square{r - 2, c}, gs.Board))
			}
		}
		if r-1 >= 0 && c-1 >= 0 && gs.Board[r-1][c-1][0] == 'b' { //capture to the left
			moves = append(moves, NewMove(Square{r, c}, Square{r - 1, c - 1}, gs.Board))
		}
		if r-1 >= 0 && c+1 < 8 && gs.Board[r-1][c+1][0] == 'b' { //capture to the right
			moves = append(moves, NewMove(Square{r, c}, Square{r - 1, c + 1}, gs.Board))
		}
	} else {
		if r+1 < 8 && gs.Board[r+1][c] == "--" { //move one square
			moves = append(moves, NewMove(Square{r, c}, Square{r + 1, c}, gs.Board))
			if r == 1 && gs.Board[r+2][c] == "--" { //move two squares
				moves = append(moves, NewMove(Square{r, c}, Square{r + 2, c}, gs.Board))
			}
		}
		if r+1 < 8 && c-1 >= 0 && gs.Board[r+1][c-1][0] == 'w' { //capture to the left
			moves = append(moves, NewMove(Square{r, c}, Square{r + 1, c - 1}, gs.Board))
		}
		if r+1 < 8 && c+1 < 8 && gs.Board[r+1][c+1][0] == 'w' { //capture to the right
			moves = append(moves, NewMove(Square{r, c}, Square{r + 1, c + 1}, gs.Board))
		}
	}
	return moves
}

func (gs *GameState) GetRookMoves(r int, c int) []Move {
	moves := []Move{}
	directions := []PieceDelta{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, direction := range directions {
		for i := 1; i < 8; i++ {
			endRow := r + direction.row*i
			endCol := c + direction.col*i
			if endRow < 0 || endRow >= 8 || endCol < 0 || endCol >= 8 {
				break
			}
			if gs.Board[endRow][endCol] == "--" {
				moves = append(moves, NewMove(Square{r, c}, Square{endRow, endCol}, gs.Board))
			} else {
				if gs.Board[endRow][endCol][0] != gs.Board[r][c][0] { //enemy piece
					moves = append(moves, NewMove(Square{r, c}, Square{endRow, endCol}, gs.Board))
				}
				break
			}
		}
	}
	return moves
}

func (gs *GameState) GetKnightMoves(r int, c int) []Move {
	moves := []Move{}
	directions := []PieceDelta{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	for _, direction := range directions {
		endRow := r + direction.row
		endCol := c + direction.col
		if endRow >= 0 && endRow < 8 && endCol >= 0 && endCol < 8 {
			if gs.Board[endRow][endCol] == "--" || gs.Board[endRow][endCol][0] != gs.Board[r][c][0] { //empty square or enemy piece
				moves = append(moves, NewMove(Square{r, c}, Square{endRow, endCol}, gs.Board))
			}
		}

	}
	return moves
}

func (gs *GameState) GetBishopMoves(r int, c int) []Move {
	moves := []Move{}
	directions := []PieceDelta{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, direction := range directions {
		for i := 1; i < 8; i++ {
			endRow := r + direction.row*i
			endCol := c + direction.col*i
			if endRow < 0 || endRow >= 8 || endCol < 0 || endCol >= 8 {
				break
			}
			if gs.Board[endRow][endCol] == "--" {
				moves = append(moves, NewMove(Square{r, c}, Square{endRow, endCol}, gs.Board))
			} else {
				if gs.Board[endRow][endCol][0] != gs.Board[r][c][0] { //enemy piece
					moves = append(moves, NewMove(Square{r, c}, Square{endRow, endCol}, gs.Board))
				}
				break
			}
		}

	}
	return moves
}

func (gs *GameState) GetQueenMoves(r int, c int) []Move {
	moves := []Move{}
	rookMoves := gs.GetRookMoves(r, c)
	bishopMoves := gs.GetBishopMoves(r, c)
	moves = append(moves, rookMoves...)
	moves = append(moves, bishopMoves...)
	return moves
}

func (gs *GameState) GetKingMoves(r int, c int) []Move {
	mmoves := []Move{}
	directions := []PieceDelta{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, direction := range directions {
		endRow := r + direction.row
		endCol := c + direction.col
		if endRow >= 0 && endRow < 8 && endCol >= 0 && endCol < 8 {
			if gs.Board[endRow][endCol] == "--" || gs.Board[endRow][endCol][0] != gs.Board[r][c][0] { //empty square or enemy piece
				mmoves = append(mmoves, NewMove(Square{r, c}, Square{endRow, endCol}, gs.Board))
			}
		}
	}
	return mmoves
}

func (gs *GameState) InCheck() bool {
	if gs.WhiteToMove {
		return gs.SquareAttacked(gs.WhiteKingSquare.row, gs.WhiteKingSquare.col)
	} else {
		return gs.SquareAttacked(gs.BlackKingSquare.row, gs.BlackKingSquare.col)
	}
}

func (gs *GameState) SquareAttacked(r int, c int) bool {
	gs.WhiteToMove = !gs.WhiteToMove
	opponentMoves := gs.GetAllPossibleMoves()
	gs.WhiteToMove = !gs.WhiteToMove
	for _, move := range opponentMoves {
		if move.EndRow == r && move.EndCol == c {
			return true
		}
	}
	return false
}
