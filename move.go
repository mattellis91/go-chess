package main

type Move struct {
	StartRow int
	StartCol int
	EndRow int
	EndCol int	
	PieceMoved string
	PieceCaptured string
	IsPawnPromotion bool
	IsEnPassant bool
	IsCastleMove bool
	MoveId int
}

func NewMove(startSquare Square, endSquare Square, boardState BoardState, isEnPassant bool, isCastleMove bool) Move {
	pieceMoved := boardState[startSquare.row][startSquare.col]

	isPawnPromotion := (pieceMoved == "wp" && endSquare.row == 0) || (pieceMoved == "bp" && endSquare.row == 7)

	return Move{
		StartRow: startSquare.row,
		StartCol: startSquare.col,
		EndRow: endSquare.row,
		EndCol: endSquare.col,
		PieceMoved: pieceMoved,
		PieceCaptured: boardState[endSquare.row][endSquare.col],
		IsPawnPromotion: isPawnPromotion,
		IsEnPassant: isEnPassant,
		MoveId: startSquare.row * 1000 + startSquare.col * 100 + endSquare.row * 10 + endSquare.col,
		IsCastleMove: isCastleMove,
	} 
}

func (m *Move) GetChessNotation() string {
	return m.getSquareNotationFromIndexes(m.StartRow, m.StartCol) + " - " + m.getSquareNotationFromIndexes(m.EndRow, m.EndCol)
}

func (m *Move) getSquareNotationFromIndexes(row int, col int) string {
	rowsToRank := map[int]string{
		0: "8", 1: "7", 2: "6", 3: "5", 4: "4", 5: "3", 6: "2", 7: "1",
	}
	colsToFile := map[int]string{
		0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 5: "f", 6: "g", 7: "h",
	}
	return colsToFile[col] + rowsToRank[row]
}

func (m *Move) getSquareFromNotation(square string) Square {
	rankToRows := map[string]int{
		"8": 0, "7": 1, "6": 2, "5": 3, "4": 4, "3": 5, "2": 6, "1": 7,
	}
	fileToCols := map[string]int{
		"a": 0, "b": 1, "c": 2, "d": 3, "e": 4, "f": 5, "g": 6, "h": 7,
	}
	return Square{
		row: rankToRows[string(square[1])],
		col: fileToCols[string(square[0])],
	}
}
