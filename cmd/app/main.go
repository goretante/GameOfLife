package main

import (
	"time"

	"github.com/goretante/GameOfLife/internal/grid"
	"github.com/goretante/GameOfLife/internal/util"
)

func main() {
	rows := util.GetIntFromUser("Enter number of rows: ")
	cols := util.GetIntFromUser("Enter number of columns: ")

	grid.SeedRandom()

	g := grid.MakeGrid(rows, cols)

	for {
		grid.PrintGrid(g)
		time.Sleep(200 * time.Millisecond)
		g = grid.Regenerate(g)
	}
}
