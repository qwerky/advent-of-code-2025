package grid

import (
	"bufio"
	"os"
)

type Grid struct {
	Height int
	Width  int
	Data   [][]rune
}

func Load(path string) *Grid {
	grid := new(Grid)

	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid.Data = append(grid.Data, row)
		grid.Width = len(line)
		grid.Height++
	}
	return grid
}

func (grid *Grid) Neighbours(x int, y int) []rune {

	if x < 0 || x >= grid.Width {
		panic("X out of bounds")
	}
	if y < 0 || y >= grid.Width {
		panic("Y out of bounds")
	}

	l := x > 0
	r := x < grid.Width-1
	u := y > 0
	d := y < grid.Height-1

	result := []rune{}
	if l {
		if u {
			result = append(result, grid.Data[x-1][y-1])
		}
		result = append(result, grid.Data[x-1][y])
		if d {
			result = append(result, grid.Data[x-1][y+1])
		}
	}
	if r {
		if u {
			result = append(result, grid.Data[x+1][y-1])
		}
		result = append(result, grid.Data[x+1][y])
		if d {
			result = append(result, grid.Data[x+1][y+1])
		}
	}
	if u {
		result = append(result, grid.Data[x][y-1])
	}
	if d {
		result = append(result, grid.Data[x][y+1])
	}

	return result
}
