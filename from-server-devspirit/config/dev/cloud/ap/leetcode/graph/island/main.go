package main

func numIslands(grid [][]byte) int {
	cnt := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '1' {
				sunkIsland(grid, r, c)
				cnt++
			}
		}
	}
	return cnt
}

var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func sunkIsland(grid [][]byte, r, c int) {
	grid[r][c] = '0'
	for _, d := range dirs {
		newR := r + d[0]
		newC := c + d[1]
		if -1 < newR && newR < len(grid) &&
			-1 < newC && newC < len(grid[newR]) &&
			grid[newR][newC] == '1' {
			sunkIsland(grid, newR, newC)
		}
	}
}

func main() {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	numIslands(grid)
}
