package solver

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/marinus/soduku/objects"
)

func AddDefaults(grid *objects.Grid) {
	reader := bufio.NewReaderSize(os.Stdin, 1)

	fmt.Println("Setting up default values for grid.")
	fmt.Println("Please enter values in the following format: Column Row Number")
	fmt.Println("For example Column 5, Row 7, Value 2 will be entered:")
	fmt.Println("5 7 2[Enter]")
	fmt.Println("To stop entering values, just type \"Continue\"")

	for true {
		fmt.Println("Enter values: ")

		text, _ := reader.ReadString('\n')
		values := strings.TrimSpace(text)

		col, row, val, err := processInput(values)
		if err != nil {
			if strings.ToLower(values) == "continue" {
				break
			}
			fmt.Println("Invalid value entered")
			continue
		}
		grid.SetDefaultCellValue(col, row, val)
	}
}

func processInput(text string) (int, int, int, error) {
	match, _ := regexp.MatchString("[1-9]{1} [1-9]{1} [1-9]{1}$", text)
	if match != true {
		err := errors.New("Invalid value entered")
		return 0, 0, 0, err
	}

	col, _ := strconv.Atoi(text[0:1])
	row, _ := strconv.Atoi(text[2:3])
	val, _ := strconv.Atoi(text[4:5])

	return col - 1, row - 1, val, nil
}

func FillTestData(grid *objects.Grid) {
	// https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRicgYEnCEEINRY_h2KMCvGdiJrugdkKvaWtY02xES42iPLxx2m
	grid.SetDefaultCellValue(0, 0, 7)
	grid.SetDefaultCellValue(0, 2, 8)
	grid.SetDefaultCellValue(0, 5, 4)
	grid.SetDefaultCellValue(0, 6, 6)
	grid.SetDefaultCellValue(1, 0, 9)
	grid.SetDefaultCellValue(1, 6, 1)
	grid.SetDefaultCellValue(2, 4, 5)
	grid.SetDefaultCellValue(2, 7, 2)
	grid.SetDefaultCellValue(2, 8, 9)
	grid.SetDefaultCellValue(3, 4, 4)
	grid.SetDefaultCellValue(3, 5, 7)
	grid.SetDefaultCellValue(3, 7, 3)
	grid.SetDefaultCellValue(4, 2, 3)
	grid.SetDefaultCellValue(4, 4, 1)
	grid.SetDefaultCellValue(4, 6, 9)
	grid.SetDefaultCellValue(5, 1, 6)
	grid.SetDefaultCellValue(5, 3, 5)
	grid.SetDefaultCellValue(5, 4, 8)
	grid.SetDefaultCellValue(6, 0, 3)
	grid.SetDefaultCellValue(6, 1, 9)
	grid.SetDefaultCellValue(6, 4, 7)
	grid.SetDefaultCellValue(7, 2, 7)
	grid.SetDefaultCellValue(7, 8, 5)
	grid.SetDefaultCellValue(8, 2, 6)
	grid.SetDefaultCellValue(8, 3, 2)
	grid.SetDefaultCellValue(8, 6, 8)
	grid.SetDefaultCellValue(8, 8, 4)
}
