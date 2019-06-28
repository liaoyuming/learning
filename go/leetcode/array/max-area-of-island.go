package main

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) < 1 {
		return 0
	}
	max := 0;
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[i]); j++ {
			area := 0
			getArea(&grid, i, j, &area)
			if max < area {
				max = area
			}
		}
	}
	return max
}

func getArea(grid *[][]int, i int, j int, area *int){
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[i]) || (*grid)[i][j] != 1 {
		return
	}
	*area++
	(*grid)[i][j] = 2
	getArea(grid, i-1, j, area)
	getArea(grid, i, j-1, area)
	getArea(grid, i+1, j, area)
	getArea(grid, i, j+1, area)
}

func main() {
	println(maxAreaOfIsland([][]int{
		{1,1,0,1,1},
		{1,0,0,0,0},
		{0,0,0,0,1},
		{1,1,0,1,1},
	}))
}