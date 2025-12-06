package day4

import (
	"fmt"
	"qwerky/learngo/aocutil/grid"
)

type point struct {
	x int
	y int
}

func Part1() {
	grid := grid.Load("data/day4/input.txt")
	movable := getMovable(grid)
	fmt.Printf("Movable roll count is %v", len(movable))
}

// Returns a slice of points that are movable
func getMovable(grid *grid.Grid) []point {
	movable := []point{}

	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			if grid.Data[x][y] == '@' {
				n := grid.Neighbours(x, y)
				ncount := 0
				for _, c := range n {
					if c == '@' {
						ncount++
					}
				}
				if ncount < 4 {
					movable = append(movable, point{x, y})
				}
			}
		}
	}
	return movable
}

func Part2() {
	grid := grid.Load("data/day4/input.txt")

	totalRemoved := 0
	movable := getMovable(grid)
	for len(movable) > 0 {
		totalRemoved += len(movable)
		// Remove the rolls that are movable
		for _, p := range movable {
			grid.Data[p.x][p.y] = '.'
		}
		movable = getMovable(grid)
	}
	fmt.Printf("Total removed is %v", totalRemoved)
}
