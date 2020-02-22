package objects

import "testing"

func TestSetCoumnCell(t *testing.T) {
	column := Column{}
	cell := Cell{}
	column.setCell(5, &cell)

	if column.cells[5] != &cell {
		t.Error("Unable to set Cell")
	}
}

func TestSetCell(t *testing.T) {
	col := Column{}
	cell := Cell{}
	col.setCell(5, &cell)

	if col.cells[5] != &cell {
		t.Error("Unable to set Cell")
	}
}

func TestGetColCells(t *testing.T) {
	col := Column{}
	for i := 0; i < GridSize; i++ {
		cell := Cell{}
		col.setCell(i, &cell)
	}
	cellReciever(col.GetCells())
}

func TestGetColCell(t *testing.T) {
	col := Column{}
	cell := Cell{}
	col.cells[5] = &cell

	if col.GetCell(5) != &cell {
		t.Error("Unable to set Cell")
	}
}
