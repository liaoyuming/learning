package main

import "fmt"

func numIslands(grid [][]byte) int {
	count := 0
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i int, j int)  {
	if i >= len(grid) || i < 0 || j >= len(grid[i]) || j < 0 || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '0'},
		{'0', '1', '0', '0'},
		{'1', '1', '1', '0'},
	}))

	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '0'},
		{'1', '1', '1', '1'},
		{'1', '1', '0', '0'},
		{'1', '1', '0', '1'},
		{'0', '0', '1', '0'},
	}))
}