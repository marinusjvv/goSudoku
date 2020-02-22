package objects

import "testing"

func TestBlockPosition(t *testing.T) {
	cases := []struct {
		col, row, want int
	}{
		{0, 0, 0},
		{3, 2, 1},
		{8, 1, 2},
		{2, 5, 3},
		{5, 5, 4},
		{8, 3, 5},
		{2, 8, 6},
		{3, 6, 7},
		{8, 8, 8},
	}
	for _, c := range cases {
		got := GetBlockNumber(c.col, c.row)
		if got != c.want {
			t.Errorf("blockPosition(%d, %d) == %d, want %d", c.col, c.row, got, c.want)
		}
	}
}

func TestBlockCellPosition(t *testing.T) {
	cases := []struct {
		col, row, want int
	}{
		{0, 0, 0},
		{0, 1, 3},
		{0, 2, 6},
		{2, 2, 8},
		{3, 3, 0},
		{5, 5, 8},
		{6, 6, 0},
		{8, 6, 2},
		{8, 8, 8},
	}
	for _, c := range cases {
		got := GetBlockCellPosition(c.col, c.row)
		if got != c.want {
			t.Errorf("blockCellPosition(%d, %d) == %d, want %d", c.col, c.row, got, c.want)
		}
	}
}

func TestSetBlockCell(t *testing.T) {
	block := Block{}
	cell := Cell{}
	block.setCell(5, &cell)

	if block.cells[5] != &cell {
		t.Error("Unable to set Cell")
	}
}

func TestGetCell(t *testing.T) {
	block := Block{}
	cell := Cell{}
	block.cells[5] = &cell

	if block.GetCells()[5] != &cell {
		t.Error("Unable to set Cell")
	}
}

func TestHasValue(t *testing.T) {
	block := Block{}
	for i := 0; i < GridSize; i++ {
		cell := Cell{}
		if i == 5 {
			cell.SetDefaultValue(5)
		}
		block.setCell(i, &cell)
	}

	if block.HasValue(6) != false {
		t.Error("Expected block not to have value 6")
	}
	if block.HasValue(5) != true {
		t.Error("Expected block to have value 6")
	}
}
