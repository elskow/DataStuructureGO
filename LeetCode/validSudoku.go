package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValidRow(board, i) {
			return false
		}
		if !isValidCol(board, i) {
			return false
		}
		if !isValidBox(board, i) {
			return false
		}
	}
	return true
}

func isValidRow(board [][]byte, row int) bool {
	m := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		if board[row][i] != '.' {
			if m[board[row][i]] {
				return false
			}
			m[board[row][i]] = true
		}
	}
	return true
}

func isValidCol(board [][]byte, col int) bool {
	m := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		if board[i][col] != '.' {
			if m[board[i][col]] {
				return false
			}
			m[board[i][col]] = true
		}
	}
	return true
}

func isValidBox(board [][]byte, box int) bool {
	m := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		row := box/3*3 + i/3
		col := box%3*3 + i%3
		if board[row][col] != '.' {
			if m[board[row][col]] {
				return false
			}
			m[board[row][col]] = true
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	println(isValidSudoku(board))
}
