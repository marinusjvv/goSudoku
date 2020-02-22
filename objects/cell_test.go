package objects

import "testing"

func TestSetDefaultValue(t *testing.T) {
	Cell := Cell{}
	Cell.SetDefaultValue(5)

	if Cell.GetNumber() != 5 {
		t.Errorf("setDefault(5) expected %d, got %d", 5, Cell.GetNumber())
	}
	if Cell.IsDefault() != true {
		t.Errorf("setDefault(5) expected %t, got %t", true, Cell.IsDefault())
	}

	Cell.Clear()

	if Cell.GetNumber() != 0 {
		t.Errorf("Clear() expected %d, got %d", 0, Cell.GetNumber())
	}
	if Cell.IsDefault() != false {
		t.Errorf("Clear() expected %t, got %t", false, Cell.IsDefault())
	}
}

func TestSetGeneratedValue(t *testing.T) {
	Cell := Cell{}
	Cell.SetGeneratedValue(6)

	if Cell.GetNumber() != 6 {
		t.Errorf("SetGeneratedValue(6) expected %d, got %d", 6, Cell.GetNumber())
	}
	if Cell.IsDefault() != false {
		t.Errorf("SetGeneratedValue(6) expected %t, got %t", false, Cell.IsDefault())
	}
}
