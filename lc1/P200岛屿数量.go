package lc1

//var res200 = 0	//这么写有问题

var res200 int
func numIslands(grid [][]byte) int {
	res200 = 0	//这样就可以，为什么呢？
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfsGrid(grid, i, j)
				res200++
			}
		}
	}
	return res200
}

func dfsGrid(grid [][]byte, i, j int) {
	if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 || grid[i][j] != '1' {
		return
	}
	grid[i][j] = 2
	dfsGrid(grid, i - 1, j)
	dfsGrid(grid, i + 1, j)
	dfsGrid(grid, i, j - 1)
	dfsGrid(grid, i, j + 1)
}
