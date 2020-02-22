package visual

import (
	"fmt"

	"github.com/marinus/soduku/objects"
)

//   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |
// 1 |   |   |   |   |   |   |   |   |   |
// 2 |   |   |   |   |   |   |   |   |   |
// 3 |   |   |   |   |   |   |   |   |   |
// 4 |   |   |   |   |   |   |   |   |   |
// 5 |   |   |   |   |   |   |   |   |   |
// 6 |   |   |   |   |   |   |   |   |   |
// 7 |   |   |   |   |   |   |   |   |   |
// 8 |   |   |   |   |   |   |   |   |   |
// 9 |   |   |   |   |   |   |   |   |   |

func DisplayGrid(grid *objects.Grid) {
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Println("-------------")
	for rowNum, row := range grid.GetRows() {
		if rowNum == 3 || rowNum == 6 {
			fmt.Println("-------------")
		}
		for cellNum, cell := range row.GetCells() {
			if cellNum == 0 || cellNum == 3 || cellNum == 6 {
				fmt.Print("|")
			}
			val := cell.GetNumber()
			if val == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(cell.GetNumber())
			}
			if cellNum == 8 {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("-------------")
	// time.Sleep(1 * time.Second)
}
