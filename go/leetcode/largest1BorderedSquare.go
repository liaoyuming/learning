package main

func largest1BorderedSquare(grid [][]int) int {
	maxSide := 0
	flag := false
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[i]); j++ {
			if grid[i][j] == 1 {
				for k:=maxSide; k<len(grid[i])-j; k++ {
					if isSquare(grid, i, j, k) {
						flag = true
						maxSide = k
					}
				}
			}
		}
	}
	if !flag {
		return 0
	}
	maxSide++
	return maxSide * maxSide
}

func isSquare(grid [][]int, startx int, starty int, side int) bool {
	if startx + side >= len(grid) || starty + side >= len(grid[0]) {
		return false
	}
	for i:=0; i<=side; i++ {
		if grid[startx+i][starty] != 1 || grid[startx][starty+i] != 1 || grid[startx + i][starty+side] != 1 || grid[startx + side][starty+i] != 1 {
			return false
		}
	}
	return true
}