package solver

import (
	"testing"

	"github.com/marinus/soduku/objects"
)

func TestSolvePuzzle(t *testing.T) {
	grid := objects.BuildGrid()
	grid.Clear()
	FillTestData(&grid)

	cases := []struct {
		col, row, val int
	}{
		{0, 0, 7},
		{0, 2, 8},
		{8, 6, 8},
		{8, 8, 4},
	}
	for _, c := range cases {
		actual := grid.GetColumn(c.col).GetCell(c.row).GetNumber()
		if c.val != actual {
			t.Errorf("setup failed, got %d, but want %d", actual, c.val)
		}
	}

	SolvePuzzle(&grid)

	cases = []struct {
		col, row, val int
	}{
		{0, 8, 3},
		{8, 0, 1},
		{4, 4, 1},
	}
	for _, c := range cases {
		actual := grid.GetColumn(c.col).GetCell(c.row).GetNumber()
		if c.val != actual {
			t.Errorf("setup failed, got %d, but want %d", actual, c.val)
		}
	}
}
