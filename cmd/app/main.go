package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter number of rows: ")
	rowInput, _ := reader.ReadString('\n')
	rows, _ := strconv.Atoi(strings.TrimSpace(rowInput))

	fmt.Print("Enter number of columns: ")
	colInput, _ := reader.ReadString('\n')
	cols, _ := strconv.Atoi(strings.TrimSpace(colInput))
	grid := MakeGrid(rows, cols)

	for {
		PrintGrid(grid)
		time.Sleep(200 * time.Millisecond)
		grid = GridRegeneration(grid, rows, cols)
	}
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

func GridRegeneration(grid [][]int, rows, cols int) [][]int {
	newGrid := make([][]int, rows)
	for i := range newGrid {
		newGrid[i] = make([]int, cols)
		for j := range newGrid[i] {
			liveNeighbours := CountNeighbours(grid, i, j, rows, cols)
			if grid[i][j] == 1 && (liveNeighbours == 2 || liveNeighbours == 3) {
				newGrid[i][j] = 1
			} else if grid[i][j] == 0 && liveNeighbours == 3 {
				newGrid[i][j] = 1
			}
		}
	}
	return newGrid
}

func CountNeighbours(grid [][]int, x, y, rows, cols int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			} else {
				r := (x + i + rows) % rows
				c := (y + j + cols) % cols
				count += grid[r][c]
			}
		}
	}
	return count
}
