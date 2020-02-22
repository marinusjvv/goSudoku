package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/marinus/soduku/objects"
	"github.com/marinus/soduku/solver"
	"github.com/marinus/soduku/visual"
)

// Rules:
// Cannot have same number double in row vertical
// Cannot have same number double in row horizontal
// Cannot have same number double in square

func main() {
	grid := objects.BuildGrid()
	recieveInputs(&grid)
}

func recieveInputs(grid *objects.Grid) {
	fmt.Println("Type \"Help\" to see all commands")

mainloop:
	for {
		reader := bufio.NewReader(os.Stdin)

		text, _ := reader.ReadString('\n')
		cleanText := strings.ToLower(strings.TrimSpace(text))

		switch cleanText {
		case "test":
			buildInputBox()
		case "help":
			displayHelp()
		case "fill":
			solver.FillTestData(grid)
			visual.DisplayGrid(grid)
		case "add":
			solver.AddDefaults(grid)
			visual.DisplayGrid(grid)
		case "show":
			visual.DisplayGrid(grid)
		case "clear":
			grid.Clear()
			fmt.Println("Grid cleared")
		case "solve":
			start := time.Now()
			solver.SolvePuzzle(grid)
			elapsed := time.Since(start)
			fmt.Printf("Solved puzzle in %s", elapsed)
			visual.DisplayGrid(grid)
		case "exit":
			break mainloop
		default:
			fmt.Println("Invalid input")
			displayHelp()
		}
	}
}

func displayHelp() {
	fmt.Println("Available commands:")
	fmt.Println("")

	fmt.Println("\"Help\" - Displays this list of commands")
	fmt.Println("\"Fill\" - Fill the grid with an example puzzle")
	fmt.Println("\"Show\" - Show the grid with entered numbers")
	fmt.Println("\"Solve\" - Solve the sudoku puzzle")
	fmt.Println("\"Add\" - Add starting numbers to the grid")
	fmt.Println("\"Clear\" - Clear all numbers from the grid")
	fmt.Println("\"Exit\" - Close the program")
}
