package objects

const GridSize = 9

type Grid struct {
	rows    [GridSize]*Row
	columns [GridSize]*Column
	blocks  [GridSize]*Block
}

type cellContainer interface {
	setCell(position int, cell *Cell)
	GetCells() [GridSize]*Cell
	HasValue(val int) bool
}

func (grid Grid) GetRows() [9]*Row {
	return grid.rows
}

func (grid Grid) GetRow(position int) *Row {
	return grid.rows[position]
}

func (grid Grid) GetColumns() [9]*Column {
	return grid.columns
}

func (grid Grid) GetColumn(position int) *Column {
	return grid.columns[position]
}

func (grid Grid) GetBlocks() [9]*Block {
	return grid.blocks
}

func (grid Grid) GetBlock(position int) *Block {
	return grid.blocks[position]
}

func (grid *Grid) setRow(position int, row *Row) {
	grid.rows[position] = row
}

func (grid *Grid) setColumn(position int, column *Column) {
	grid.columns[position] = column
}

func (grid *Grid) setBlock(position int, block *Block) {
	grid.blocks[position] = block
}

func BuildGrid() Grid {
	Grid := Grid{}

	for i := 0; i < GridSize; i++ {
		row := Row{}
		Grid.setRow(i, &row)
		column := Column{}
		Grid.setColumn(i, &column)
		block := Block{}
		Grid.setBlock(i, &block)
	}

	for col := 0; col < GridSize; col++ {
		for row := 0; row < GridSize; row++ {
			Cell := Cell{}
			Grid.columns[col].setCell(row, &Cell)
			Grid.rows[row].setCell(col, &Cell)
			blockPosition := GetBlockNumber(col, row)
			cellPosition := GetBlockCellPosition(col, row)
			Grid.blocks[blockPosition].setCell(cellPosition, &Cell)
		}
	}

	return Grid
}

func (grid *Grid) SetDefaultCellValue(col int, row int, val int) {
	grid.columns[col].cells[row].SetDefaultValue(val)
}

func (grid *Grid) Clear() {
	for _, col := range grid.columns {
		for _, cell := range col.cells {
			cell.Clear()
		}
	}
}
