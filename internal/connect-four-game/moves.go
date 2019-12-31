package connect_four_game

func ValidMove(board [][]int32, column int32, player int32) bool {
	col := board[column]
	col = append(col, player)
	if len(col) <= 6 {
		board[column] = col
	}
	return len(col) <= 6
}

func wonVertical(board [][]int32) int32 {
	for col := range board {
		negCounter := 0
		posCounter := 0
		for row := range board[col] {
			num := board[col][row]
			if num == 1 {
				negCounter = 0
				posCounter++
			} else if num == -1 {
				posCounter = 0
				negCounter++
			} else {
				break
			}
			if negCounter == 4 {
				return -1
			}
			if posCounter == 4 {
				return 1
			}
		}
	}
	return 0
}

func wonHorizontal(board [][]int32) int32 {
	// getting the height, given that board is rectangle
	for row := 0; row < 6; row++ {
		for startCol := 0; startCol <= len(board) - 4; startCol++ {
			negCounter := 0
			posCounter := 0
			for col := startCol; col < startCol + 4; col++ {
				if len(board[col])  <= row {
					break
				}
				num := board[col][row]
				if num == 1 {
					negCounter = 0
					posCounter++
				} else if num == -1 {
					posCounter = 0
					negCounter++
				} else {
					break
				}
				if negCounter == 4 {
					return -1
				}
				if posCounter == 4 {
					return 1
				}
			}
		}
	}
	return 0
}

func wonDiagonalForward(board [][]int32) int32 {
	for row := 0; row <= 6 - 4; row++ {
		for startCol := 0; startCol <= len(board) - 4; startCol++ {
			negCounter := 0
			posCounter := 0
			for col := 0; col < 4; col++ {
				if len(board[startCol + col]) <= row + col {
					break
				}
				num := board[startCol + col][row + col]
				if num == 1 {
					negCounter = 0
					posCounter++
				} else if num == -1 {
					posCounter = 0
					negCounter++
				} else {
					break
				}
				if negCounter == 4 {
					return -1
				}
				if posCounter == 4 {
					return 1
				}
			}
		}
	}
	return 0
}

func wonDiagonalBackwards(board [][]int32) int32{
	for row := 0; row <= 6 - 4; row++ {
		for startCol := len(board) - 1; startCol >= 3; startCol-- {
			negCounter := 0
			posCounter := 0
			for col := 0; col < 4; col++ {
				if len(board[startCol - col]) <= row + col {
					break
				}
				num := board[startCol - col][row + col]
				if num == 1 {
					negCounter = 0
					posCounter++
				} else if num == -1 {
					posCounter = 0
					negCounter++
				} else {
					break
				}
				if negCounter == 4 {
					return -1
				}
				if posCounter == 4 {
					return 1
				}
			}
		}
	}
	return 0
}