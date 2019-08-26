package main

func dfs(grid [][]int, row int, col int, count *int) {
	nRow := len(grid)
	if row < 0 || row >= nRow {
		return
	}
	nCol := len(grid[0])
	if col < 0 || col >= nCol {
		return
	}
	if grid[row][col] == 0 {
		return
	}

	if grid[row][col] == 1 {
		grid[row][col] = 0
		*count = *count + 1
		dfs(grid, row-1, col, count)
		dfs(grid, row, col+1, count)
		dfs(grid, row+1, col, count)
		dfs(grid, row, col-1, count)
	}
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxAreaOfIsland(grid [][]int) int {
	// Loop through all rows and columns
	// If current element is a 1 , do a dfs on it
	// During DFS, if out of bounds, then return 0
	// If current element is a 1 add 1 to result
	// mark current value as visited
	result := 0
	for i, row := range grid {
		for j, _ := range row {
			count := 0
			dfs(grid, i, j, &count)
			result = Max(count, result)
		}
	}

	return result
}
