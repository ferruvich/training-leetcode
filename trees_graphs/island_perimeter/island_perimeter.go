package island_perimeter

func dfsCountingPerimeter(grid [][]int, visited map[int]map[int]bool, i, j int) int {
	if v, ok := visited[i][j]; ok && v {
		return 0
	}

	if _, ok := visited[i]; !ok {
		visited[i] = map[int]bool{}
	}
	visited[i][j] = true

	boxPerimeter := 4
	if i-1 >= 0 && grid[i-1][j] == 1 {
		boxPerimeter += dfsCountingPerimeter(grid, visited, i-1, j) - 1
	}
	if j-1 >= 0 && grid[i][j-1] == 1 {
		boxPerimeter += dfsCountingPerimeter(grid, visited, i, j-1) - 1
	}
	if i+1 < len(grid) && grid[i+1][j] == 1 {
		boxPerimeter += dfsCountingPerimeter(grid, visited, i+1, j) - 1
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
		boxPerimeter += dfsCountingPerimeter(grid, visited, i, j+1) - 1
	}

	return boxPerimeter
}

func islandPerimeter(grid [][]int) int {
	visited := map[int]map[int]bool{}
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return dfsCountingPerimeter(grid, visited, i, j)
			}
		}
	}

	return 0
}
