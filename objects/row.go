package objects

type Row struct {
	cells [GridSize]*Cell
}

func (row *Row) setCell(position int, cell *Cell) {
	row.cells[position] = cell
}

func (row *Row) GetCells() [GridSize]*Cell {
	return row.cells
}

func (row *Row) HasValue(val int) bool {
	for _, cell := range row.cells {
		if cell.number == val {
			return true
		}
	}
	return false
}
