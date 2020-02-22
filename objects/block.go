package objects

import "math"

type Block struct {
	cells [GridSize]*Cell
}

func (block *Block) setCell(position int, cell *Cell) {
	block.cells[position] = cell
}

func (block *Block) GetCells() [9]*Cell {
	return block.cells
}

func (block *Block) HasValue(val int) bool {
	for _, cell := range block.cells {
		if cell.number == val {
			return true
		}
	}
	return false
}

func GetBlockNumber(col int, row int) int {
	colPos := math.Ceil((float64(col) + 1) / 3)
	rowPos := (math.Ceil((float64(row)+1)/3) - 1) * 3

	return int(colPos) + int(rowPos) - 1
}

func GetBlockCellPosition(col int, row int) int {
	colPos := col + 1
	for colPos > 3 {
		colPos = colPos - 3
	}
	rowPos := row + 1
	for rowPos > 3 {
		rowPos = rowPos - 3
	}
	val := colPos + ((rowPos - 1) * 3)
	return val - 1
}
