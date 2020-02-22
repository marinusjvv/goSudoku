package solver

import (
	"errors"

	"github.com/marinus/soduku/objects"
)

func SolvePuzzle(grid *objects.Grid) error {
	return checkColumns(grid, 0, 0, 1)
}

func checkColumns(grid *objects.Grid, startColNumber int, startCellNumber int, startNumber int) error {
	for colNumber := startColNumber; colNumber < objects.GridSize; colNumber++ {
		err := checkCells(grid, colNumber, startCellNumber, startNumber)
		startCellNumber = 0
		startNumber = 0
		if err != nil {
			if colNumber != 0 {
				previousColNumber := colNumber - 1
				previousCol := grid.GetColumn(previousColNumber)

				previousColLastCellNumber, cellErr := getEditableCellNumber(objects.GridSize, previousCol)
				if cellErr != nil {
					return err
				}
				previousColLastCellValue := grid.GetColumn(previousColNumber).GetCell(previousColLastCellNumber).GetNumber()

				return checkColumns(grid, previousColNumber, previousColLastCellNumber, (previousColLastCellValue + 1))
			} else {
				return err
			}
		}
	}
	return nil
}

func checkCells(grid *objects.Grid, colNumber int, startCellNumber int, startNumber int) error {
	col := grid.GetColumn(colNumber)
	for cellNumber := startCellNumber; cellNumber < objects.GridSize; cellNumber++ {
		startCellNumber = 0

		cell := col.GetCell(cellNumber)
		if cell.IsDefault() == true {
			continue
		}
		cell.Clear()
		row := grid.GetRow(cellNumber)
		blockNumber := objects.GetBlockNumber(colNumber, cellNumber)
		block := grid.GetBlock(blockNumber)

		err := checkNumbers(grid, row, col, block, cell, startNumber)
		startNumber = 0
		if err != nil {
			if cellNumber != 0 {
				previousCellNumber, cellErr := getEditableCellNumber(cellNumber, col)
				if cellErr != nil {
					return err
				}
				previousCellValue := col.GetCell(previousCellNumber).GetNumber()
				return checkCells(grid, colNumber, previousCellNumber, (previousCellValue + 1))
			} else {
				return err
			}
		}
	}
	return nil
}

func getEditableCellNumber(currentCellNumber int, col *objects.Column) (int, error) {
	previousCellNumber := currentCellNumber
	for true {
		if previousCellNumber == 0 {
			return 0, errors.New("Cannot get previous editable cell in column")
		}
		previousCellNumber = previousCellNumber - 1
		if col.GetCell(previousCellNumber).IsDefault() {
			continue
		}
		if col.GetCell(previousCellNumber).GetNumber() == objects.GridSize {
			col.GetCell(previousCellNumber).Clear()
			continue
		}
		return previousCellNumber, nil
	}
	return previousCellNumber, nil
}

func checkNumbers(grid *objects.Grid, row *objects.Row, col *objects.Column, block *objects.Block, cell *objects.Cell, startNumber int) error {
	for i := startNumber; i <= objects.GridSize; i++ {
		if canSet(row, col, block, cell, i) {
			cell.SetGeneratedValue(i)
			return nil
		}
	}
	return errors.New("Cannot set number")
}

func canSet(row *objects.Row, col *objects.Column, block *objects.Block, cell *objects.Cell, val int) bool {
	if cell.IsDefault() {
		return false
	}
	if row.HasValue(val) {
		return false
	}
	if col.HasValue(val) {
		return false
	}
	if block.HasValue(val) {
		return false
	}

	return true
}
