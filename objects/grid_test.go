package objects

import "testing"

func TestBuild(t *testing.T) {
	Grid := BuildGrid()

	if len(Grid.rows) != GridSize {
		t.Errorf("BuildGrid() expected rows %d, got %d", GridSize, len(Grid.rows))
	}
	if len(Grid.columns) != GridSize {
		t.Errorf("BuildGrid() expected columns %d, got %d", GridSize, len(Grid.columns))
	}
	if len(Grid.blocks) != GridSize {
		t.Errorf("BuildGrid() expected blocks %d, got %d", GridSize, len(Grid.blocks))
	}
	if len(Grid.rows[8].cells) != GridSize {
		t.Errorf("BuildGrid() expected cells %d, wgotant %d", GridSize, len(Grid.rows[8].cells))
	}
	if len(Grid.columns[8].cells) != GridSize {
		t.Errorf("BuildGrid() expected cells %d, got %d", GridSize, len(Grid.columns[8].cells))
	}
	if len(Grid.blocks[8].cells) != GridSize {
		t.Errorf("BuildGrid() expected cells %d, got %d", GridSize, len(Grid.blocks[8].cells))
	}

	rowsReciever(Grid.GetRows())
	colsReciever(Grid.GetColumns())
	blocksReciever(Grid.GetBlocks())
}

func TestSetDefaultCellValue(t *testing.T) {
	Grid := BuildGrid()

	if Grid.columns[8].HasValue(5) != false {
		t.Errorf("HasValue(5) expected %t, got %t", false, Grid.columns[8].HasValue(5))
	}

	Grid.SetDefaultCellValue(8, 6, 5)

	if Grid.columns[8].cells[6].number != 5 {
		t.Errorf("setDefault(8, 6, 5) expected %d, got %d", 5, Grid.columns[8].cells[6].number)
	}
	if Grid.columns[8].cells[6].isDefault != true {
		t.Errorf("setDefault(8, 6, 5) expected %t, got %t", true, Grid.columns[8].cells[6].isDefault)
	}
	if Grid.columns[8].HasValue(5) != true {
		t.Errorf("HasValue(5) expected %t, got %t", true, Grid.columns[8].HasValue(5))
	}

	if Grid.rows[6].cells[8].number != 5 {
		t.Errorf("setDefault(8, 6, 5) expected %d, got %d", 5, Grid.rows[6].cells[8].number)
	}
	if Grid.rows[6].cells[8].isDefault != true {
		t.Errorf("setDefault(8, 6, 5) expected %t, got %t", true, Grid.rows[6].cells[8].isDefault)
	}

	if Grid.blocks[8].cells[2].number != 5 {
		t.Errorf("setDefault(8, 6, 5) expected %d, got %d", 5, Grid.blocks[6].cells[2].number)
	}
	if Grid.blocks[8].cells[2].isDefault != true {
		t.Errorf("setDefault(8, 6, 5) expected %t, got %t", true, Grid.blocks[6].cells[2].isDefault)
	}

	Grid.Clear()

	if Grid.columns[8].HasValue(5) != false {
		t.Errorf("HasValue(5) expected %t, got %t", false, Grid.columns[8].HasValue(5))
	}
	if Grid.columns[8].cells[6].number != 0 {
		t.Errorf("setDefault(8, 6, 5) expected %d, got %d", 0, Grid.columns[8].cells[6].number)
	}
	if Grid.columns[8].cells[6].isDefault != false {
		t.Errorf("setDefault(8, 6, 5) expected %t, got %t", false, Grid.columns[8].cells[6].isDefault)
	}
}

func TestSetRow(t *testing.T) {
	Grid := BuildGrid()
	row := Row{}
	Grid.setRow(8, &row)

	if Grid.rows[8] != &row {
		t.Error("Unable to set Row")
	}
}

func TestSetColumn(t *testing.T) {
	Grid := BuildGrid()
	column := Column{}
	Grid.setColumn(8, &column)

	if Grid.columns[8] != &column {
		t.Error("Unable to set Column")
	}
}

func TestSetBlock(t *testing.T) {
	Grid := BuildGrid()
	block := Block{}
	Grid.setBlock(8, &block)

	if Grid.blocks[8] != &block {
		t.Error("Unable to set Block")
	}
}

func TestGoAssociation(t *testing.T) {
	ce := Cell{}
	slice1 := [1]*Cell{&ce}
	slice2 := [1]*Cell{&ce}

	ce.SetDefaultValue(3)

	if slice1[0].number != 3 {
		t.Errorf("expected %d, got %d", 3, slice1[0].number)
	}
	if slice2[0].number != 3 {
		t.Errorf("expected %d, got %d", 3, slice1[0].number)
	}
}

func rowsReciever([9]*Row)     {}
func colsReciever([9]*Column)  {}
func blocksReciever([9]*Block) {}
func cellReciever([9]*Cell)    {}
