package array

func numIslands(grid [][]byte) int {
	c := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				c++
				bfs(i, j, grid)
			}
		}
	}

	return c
}

func bfs(i, j int, grid [][]byte) {
	grid[i][j] = 0

	if i > 0 && grid[i-1][j] == '1' {
		bfs(i-1, j, grid)
	}

	if i < len(grid)-1 && grid[i+1][j] == '1' {
		bfs(i+1, j, grid)
	}

	if j > 0 && grid[i][j-1] == '1' {
		bfs(i, j-1, grid)
	}

	if j < len(grid[i])-1 && grid[i][j+1] == '1' {
		bfs(i, j+1, grid)
	}
}

/* 23:55 11.02.2022 - 00:02 12.02.2022 */
