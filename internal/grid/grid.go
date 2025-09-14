package grid

import (
	"fmt"
	"math/rand"
	"time"
)

func SeedRandom() {
	rand.New(rand.NewSource(time.Now().Unix()))
}

func MakeGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2)
		}
	}
	return grid
}

func PrintGrid(grid [][]int) {
	fmt.Print("\033[H\033[2J")
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				fmt.Print("■")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func Regenerate(grid [][]int) [][]int {
	rows := len(grid)
	cols := len(grid[0])
	newGrid := make([][]int, rows)
	for i := range grid {
		newGrid[i] = make([]int, cols)
		for j := range grid[i] {
			liveNeighbours := countNeighbours(grid, i, j, rows, cols)
			if grid[i][j] == 1 && (liveNeighbours == 2 || liveNeighbours == 3) {
				newGrid[i][j] = 1
			} else if grid[i][j] == 0 && liveNeighbours == 3 {
				newGrid[i][j] = 1
			}
		}
	}
	return newGrid
}

func countNeighbours(grid [][]int, x, y, rows, cols int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			r := (x + i + rows) % rows
			c := (y + j + cols) % cols
			count += grid[r][c]
		}
	}
	return count
}
