package solver

import "testing"

func TestProcessInput(t *testing.T) {
	cases := []struct {
		input         string
		col, row, val int
	}{
		{"1 1 1", 0, 0, 1},
		{"1 2 3", 0, 1, 3},
		{"9 9 9", 8, 8, 9},
	}
	for _, c := range cases {
		col, row, val, _ := processInput(c.input)
		if col != c.col || row != c.row || val != c.val {
			t.Errorf("blockCellPoprocessInputsition(%s) == %d, %d, %d, but want %d, %d, %d", c.input, col, row, val, c.col, c.row, c.val)
		}
	}

	invalidCases := []struct {
		input string
	}{
		{"1 1 10"},
		{"0 0 0"},
		{"continue"},
	}
	for _, c := range invalidCases {
		_, _, _, err := processInput(c.input)
		if err == nil {
			t.Errorf("Expected failure, got none for input %s", c.input)
		}
	}
}
