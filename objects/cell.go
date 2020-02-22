package objects

type Cell struct {
	number    int
	isDefault bool
}

func (cell *Cell) SetDefaultValue(number int) {
	cell.number = number
	cell.isDefault = true
}

func (cell *Cell) SetGeneratedValue(number int) {
	cell.number = number
}

func (cell *Cell) Clear() {
	cell.number = 0
	cell.isDefault = false
}

func (cell *Cell) GetNumber() int {
	return cell.number
}

func (cell *Cell) IsDefault() bool {
	return cell.isDefault
}
