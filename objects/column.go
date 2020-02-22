package objects

type Column struct {
	cells [GridSize]*Cell
}

func (column *Column) setCell(position int, cell *Cell) {
	column.cells[position] = cell
}

func (column *Column) GetCells() [GridSize]*Cell {
	return column.cells
}

func (column *Column) GetCell(position int) *Cell {
	return column.cells[position]
}

func (column *Column) HasValue(val int) bool {
	for _, cell := range column.cells {
		if cell.number == val {
			return true
		}
	}
	return false
}
