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
	CurrentPlayerInCheck bool
	Stalemate		 bool
	CheckState		 bool
	Pins		     []AttactedSquare
	Checks		    []AttactedSquare
}

type PieceDelta struct {
	row int
	col int
}

type AttactedSquare struct {
	row int
	col int
	direction PieceDelta
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

	//TODO: ADD option select piece to promote to
	if move.IsPawnPromotion {
		gs.Board[move.EndRow][move.EndCol] = move.PieceMoved[:1] + "Q"
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
	moves := []Move{}

	gs.CurrentPlayerInCheck, gs.Pins, gs.Checks = gs.CheckForPinsAndChecks() 

	var kingRow, kingCol int

	if gs.WhiteToMove {
		kingRow = gs.WhiteKingSquare.row
		kingCol = gs.WhiteKingSquare.col
	} else {
		kingRow = gs.BlackKingSquare.row
		kingCol = gs.BlackKingSquare.col
	}

	if gs.CurrentPlayerInCheck {
		if len(gs.Checks) == 1 { // only one check, block check or move king
			moves = gs.GetAllPossibleMoves()
			// to block a check you must move a piece into one of the squares between the enemy piece and the king
			check := gs.Checks[0]
			checkRow := check.row
			checkCol := check.col
			pieceAttacking := gs.Board[checkRow][checkCol]
			validSquares := []Square{}
			if pieceAttacking[1] == 'N' { // if knight, must capture knight or move king
				validSquares = append(validSquares, Square{checkRow, checkCol})
			} else { // if rook, bishop, or queen, you can block the check by moving a piece in between the king and the enemy piece
				for i := 1; i < 8; i++ {
					endRow := kingRow + check.direction.row * i
					endCol := kingCol + check.direction.col * i
					validSquares = append(validSquares, Square{endRow, endCol})
					if endRow == checkRow && endCol == checkCol {
						break
					}
				}
			}
			for i := len(moves) - 1; i >= 0; i-- { // remove moves that don't block check or move king
				if moves[i].PieceMoved[1] != 'K' { 
					moveSquareInValidSquares := false
					for _, validSquare := range validSquares {
						if moves[i].EndRow == validSquare.row && moves[i].EndCol == validSquare.col { 
							moveSquareInValidSquares = true
							break
						}
					}
					if !moveSquareInValidSquares { //remove move that does not block check
						moves = append(moves[:i], moves[i+1:]...)
					}  
				}
			}
		} else { // double check, king has to move
			moves = append(moves, gs.GetKingMoves(kingRow, kingCol)...)	
		}
	} else {
		moves = gs.GetAllPossibleMoves()
	}

	// if len(moves) == 0 {
	// 	if gs.InCheck() {
	// 		gs.Checkmate = true
	// 	} else {
	// 		gs.Stalemate = true
	// 	}
	// } else {
	// 	gs.Checkmate = false
	// 	gs.Stalemate = false
	// }
	
	return moves
}

func (gs * GameState) CheckForPinsAndChecks() (bool, []AttactedSquare, []AttactedSquare) {
	pins := []AttactedSquare{}
	checks := []AttactedSquare{}
	inCheck := false

	var enemyColor, allyColor byte
	var startRow, startCol int

	if gs.WhiteToMove {
		enemyColor = 'b'
		allyColor = 'w'
		startRow = gs.WhiteKingSquare.row
		startCol = gs.WhiteKingSquare.col	
	} else {
		enemyColor = 'w'
		allyColor = 'b'
		startRow = gs.BlackKingSquare.row
		startCol = gs.BlackKingSquare.col
	}

	// directions 0 to 3 are orthogonal, 4 to 7 are diagonal
	directions := []PieceDelta{{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {-1, 1}, {1, -1}, {1, 1}}
	
	for j := 0; j < len(directions); j++ {
		d := directions[j]
		possiblePin := AttactedSquare{-1, -1, PieceDelta{}}
		for i := 1; i < 8; i++ {
			endRow := startRow + d.row * i
			endCol := startCol + d.col * i
			if 0 <= endRow && endRow < 8 && 0 <= endCol && endCol < 8 {
				endPiece := gs.Board[endRow][endCol]
				if endPiece[0] == allyColor && endPiece[1] != 'K' {
					if possiblePin.row == -1 {
						possiblePin = AttactedSquare{endRow, endCol, d}
					} else {
						break
					} 
				} else if endPiece[0] == enemyColor {
					pieceType := endPiece[1]
					// depending on direction that is being checked, only certain pieces can attack the king
					if  pieceType == 'R' && 0 <= j && j <= 3 || pieceType == 'B' && 4 <= j && j <= 7 ||
						i == 1 && pieceType == 'p' && ((enemyColor == 'w' && 6 <= j && j <= 7) || (enemyColor == 'b' && 4 <= j && j <= 5)) || 
						pieceType == 'Q' || (i == 1 && pieceType == 'K') {
							if possiblePin.row == -1 {
								inCheck = true
								checks = append(checks, AttactedSquare{endRow, endCol, d})
								break
							} else { // there is a piece blocking so pin
								pins = append(pins, possiblePin)
								break
							}
					} else { // enemy piece is not attacking the king
						break
					}
				}
			} else {
				break
			}
		}
	}

	// knight checks
	knightMoves := []PieceDelta{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	for _, m := range knightMoves {
		endRow := startRow + m.row
		endCol := startCol + m.col
		if 0 <= endRow && endRow < 8 && 0 <= endCol && endCol < 8 {
			endPiece := gs.Board[endRow][endCol]
			if endPiece[0] == enemyColor && endPiece[1] == 'N' {
				inCheck = true
				checks = append(checks, AttactedSquare{endRow, endCol, m})
			}
		}
	}

	return inCheck, pins, checks

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
	piecePinned := false
	pinDirection := PieceDelta{}

	for i := len(gs.Pins) - 1; i >= 0; i-- {
		if gs.Pins[i].row == r && gs.Pins[i].col == c {
			piecePinned = true
			pinDirection = gs.Pins[i].direction
			gs.Pins = append(gs.Pins[:i], gs.Pins[i+1:]...)
			break
		}
	}


	if gs.WhiteToMove {
		if r-1 >= 0 && gs.Board[r-1][c] == "--" { //move one square
			if !piecePinned || pinDirection == (PieceDelta{-1, 0}) {
				moves = append(moves, NewMove(Square{r, c}, Square{r - 1, c}, gs.Board))
				if r == 6 && gs.Board[r-2][c] == "--" { //move two squares
					moves = append(moves, NewMove(Square{r, c}, Square{r - 2, c}, gs.Board))
				}
			}
		}
		if r-1 >= 0 && c-1 >= 0 && gs.Board[r-1][c-1][0] == 'b' { //capture to the left
			if !piecePinned || pinDirection == (PieceDelta{-1, -1}) {
				moves = append(moves, NewMove(Square{r, c}, Square{r - 1, c - 1}, gs.Board))
			}
		}
		if r-1 >= 0 && c+1 < 8 && gs.Board[r-1][c+1][0] == 'b' { //capture to the right
			if !piecePinned || pinDirection == (PieceDelta{-1, 1}) {
				moves = append(moves, NewMove(Square{r, c}, Square{r - 1, c + 1}, gs.Board))
			}
		}
	} else {
		if r+1 < 8 && gs.Board[r+1][c] == "--" { //move one square
			if !piecePinned || pinDirection == (PieceDelta{1, 0}) {
				moves = append(moves, NewMove(Square{r, c}, Square{r + 1, c}, gs.Board))
				if r == 1 && gs.Board[r+2][c] == "--" { //move two squares
					moves = append(moves, NewMove(Square{r, c}, Square{r + 2, c}, gs.Board))
				}
			}
		}
		if r+1 < 8 && c-1 >= 0 && gs.Board[r+1][c-1][0] == 'w' { //capture to the left
			if !piecePinned || pinDirection == (PieceDelta{1, -1}) {
				moves = append(moves, NewMove(Square{r, c}, Square{r + 1, c - 1}, gs.Board))
			}
		}
		if r+1 < 8 && c+1 < 8 && gs.Board[r+1][c+1][0] == 'w' { //capture to the right
			if !piecePinned || pinDirection == (PieceDelta{1, 1}) {
				moves = append(moves, NewMove(Square{r, c}, Square{r + 1, c + 1}, gs.Board))
			}
		}
	}
	return moves
}

func (gs *GameState) GetRookMoves(r int, c int) []Move {
	moves := []Move{}
	directions := []PieceDelta{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	piecePinned := false
	pinDirection := PieceDelta{}
	for i := len(gs.Pins) - 1; i >= 0; i-- {
		if gs.Pins[i].row == r && gs.Pins[i].col == c {
			piecePinned = true
			pinDirection = gs.Pins[i].direction
			if gs.Board[r][c][1] != 'Q' {
				gs.Pins = append(gs.Pins[:i], gs.Pins[i+1:]...)
			}
			break
		}
	}

	for _, direction := range directions {
		for i := 1; i < 8; i++ {
			endRow := r + direction.row*i
			endCol := c + direction.col*i
			if endRow < 0 || endRow >= 8 || endCol < 0 || endCol >= 8 {
				break
			}
			if !piecePinned || pinDirection == direction || pinDirection == (PieceDelta{-direction.row, -direction.col}) {
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
	}
	return moves
}

func (gs *GameState) GetKnightMoves(r int, c int) []Move {
	moves := []Move{}
	directions := []PieceDelta{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}

	piecePinned := false

	for i := len(gs.Pins) - 1; i >= 0; i-- {
		if gs.Pins[i].row == r && gs.Pins[i].col == c {
			piecePinned = true
			gs.Pins = append(gs.Pins[:i], gs.Pins[i+1:]...)
			break
		}
	}

	for _, direction := range directions {
		endRow := r + direction.row
		endCol := c + direction.col
		if endRow >= 0 && endRow < 8 && endCol >= 0 && endCol < 8 {
			if !piecePinned {
				if gs.Board[endRow][endCol] == "--" || gs.Board[endRow][endCol][0] != gs.Board[r][c][0] { //empty square or enemy piece
					moves = append(moves, NewMove(Square{r, c}, Square{endRow, endCol}, gs.Board))
				}
			}
		}

	}
	return moves
}

func (gs *GameState) GetBishopMoves(r int, c int) []Move {
	moves := []Move{}
	directions := []PieceDelta{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	piecePinned := false
	pinDirection := PieceDelta{}

	for i := len(gs.Pins) - 1; i >= 0; i-- {
		if gs.Pins[i].row == r && gs.Pins[i].col == c {
			piecePinned = true
			pinDirection = gs.Pins[i].direction
			if gs.Board[r][c][1] != 'Q' {
				gs.Pins = append(gs.Pins[:i], gs.Pins[i+1:]...)
			}
			break
		}
	}

	for _, direction := range directions {
		for i := 1; i < 8; i++ {
			endRow := r + direction.row*i
			endCol := c + direction.col*i
			if endRow < 0 || endRow >= 8 || endCol < 0 || endCol >= 8 {
				break
			}
			if !piecePinned || pinDirection == direction || pinDirection == (PieceDelta{-direction.row, -direction.col}) {
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

	rowMoves := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	colMoves := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	var allyColor byte
	if gs.WhiteToMove {
		allyColor = 'w'
	} else {
		allyColor = 'b'
	}

	for i := 0; i < 8; i++ {
		endRow := r + rowMoves[i]
		endCol := c + colMoves[i]
		if 0 <= endRow && endRow < 8 && 0 <= endCol && endCol < 8 {
			endPiece := gs.Board[endRow][endCol]
			if endPiece[0] != allyColor {
				if allyColor == 'w' {
					gs.WhiteKingSquare = Square{endRow, endCol}
				} else {
					gs.BlackKingSquare = Square{endRow, endCol}
				}
				
				inCheck, _, _ := gs.CheckForPinsAndChecks()
				
				if !inCheck {
					mmoves = append(mmoves, NewMove(Square{r, c}, Square{endRow, endCol}, gs.Board))
				}
				
				if allyColor == 'w' {
					gs.WhiteKingSquare = Square{r, c}
				} else {
					gs.BlackKingSquare = Square{r, c}
				}
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
