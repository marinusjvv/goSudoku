package objects

import "testing"

func TestSetRowCell(t *testing.T) {
	row := Row{}
	cell := Cell{}
	row.setCell(5, &cell)

	if row.cells[5] != &cell {
		t.Error("Unable to set Cell")
	}
}

func TestGetRowCells(t *testing.T) {
	row := Row{}
	for i := 0; i < GridSize; i++ {
		cell := Cell{}
		row.setCell(i, &cell)
	}
	cellReciever(row.GetCells())
}

func TestRowHasValue(t *testing.T) {
	row := Row{}
	for i := 0; i < GridSize; i++ {
		cell := Cell{}
		if i == 5 {
			cell.SetDefaultValue(5)
		}
		row.setCell(i, &cell)
	}

	if row.HasValue(6) != false {
		t.Error("Expected block not to have value 6")
	}
	if row.HasValue(5) != true {
		t.Error("Expected block to have value 6")
	}
}
