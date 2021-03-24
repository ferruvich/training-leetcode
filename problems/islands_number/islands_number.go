package islands_number

const (
	water byte = '0'
	land byte = '1'
)

func dfs(grid [][]byte, x, y int) {
	grid[x][y] = water
	if y-1 >= 0 && grid[x][y-1] == land {
		dfs(grid, x, y-1)
	}
	if y+1 < len(grid[x]) && grid[x][y+1] == land {
		dfs(grid, x, y+1)
	}
	if x-1 >= 0 && grid[x-1][y] == land {
		dfs(grid, x-1, y)
	}
	if x+1 < len(grid) && grid[x+1][y]== land {
		dfs(grid, x+1, y)
	}
}

func NumIslands(grid [][]byte) int {
	var islands int
	for x, row := range grid {
		for y, el := range row {
			if el == land {
				islands++
				dfs(grid, x, y)
			}
		}
	}

	return islands
}