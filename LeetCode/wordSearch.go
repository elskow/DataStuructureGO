package main

func exist(board [][]byte, word string) bool {
	var res bool
	var visited = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			res = dfs(board, word, visited, i, j, 0)
			if res {
				return res
			}
		}
	}
	return res
}

func dfs(board [][]byte, word string, visited [][]bool, i, j, k int) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || visited[i][j] || board[i][j] != word[k] {
		return false
	}
	visited[i][j] = true
	var res bool
	res = dfs(board, word, visited, i+1, j, k+1) || dfs(board, word, visited, i-1, j, k+1) || dfs(board, word, visited, i, j+1, k+1) || dfs(board, word, visited, i, j-1, k+1)
	visited[i][j] = false
	return res
}

func main() {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	println(exist(board, word))
}
