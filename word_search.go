package main

import "fmt"

func dfs(board [][]byte, r int, c int, word string, wordPos int) bool {

	if wordPos == len(word) {
		return true
	}

	if r < 0 || r == len(board) || c < 0 || c == len(board[r]) || board[r][c] == '0' {
		return false
	}

	if board[r][c] == word[wordPos] {
		fmt.Printf("%c%d\n", board[r][c], wordPos)
		board[r][c] = '0'
		status := dfs(board, r-1, c, word, wordPos+1) ||
			dfs(board, r, c-1, word, wordPos+1) ||
			dfs(board, r, c+1, word, wordPos+1) ||
			dfs(board, r+1, c, word, wordPos+1)
		board[r][c] = word[wordPos]
		return status
	} else {
		return false
	}
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	// Loop over the board , at each point do a DFS from the current character
	// In DFS Function : check if the index is valid and unvisited
	// If so , check if this patches the next position of the word , if so mark visiteda nd keep looking
	// return true if position == word.end()
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func main() {

}
