package main

import (
	"fmt"
	"os"
	"strings"
)

// grid search
func day_4() {
	input, err := os.ReadFile("./inputs/4.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	// conver string to grid
	var grid [140][140]string
	for i, line := range strings.Split(string(input), "\n") {
		grid[i] = [140]string(strings.Split(line, ""))
	}

	// First Challenge matches
	var matches int // 2633
	// Secondc chalenge matches
	var secondChallengeMatches int // 1963

	isXMAS := func(word string) {
		if word == "XMAS" || word == "SAMX" {
			matches++
		}
	}

	isMAS := func(word string) bool {
		if word == "MS" || word == "SM" {
			return true
		}
		return false
	}
	gridLength := len(grid)
	for y := 0; y < gridLength; y++ {

		for x := 0; x < gridLength; x++ {

			// first challenge find xmas backwards, vertical, diagaonl and normal
			// horizontal matches
			if x < gridLength-3 {
				isXMAS(grid[y][x] + grid[y][x+1] + grid[y][x+2] + grid[y][x+3])
			}
			// vertical
			if y < gridLength-3 {
				isXMAS(grid[y][x] + grid[y+1][x] + grid[y+2][x] + grid[y+3][x])

				// diagonal to right
				if x < gridLength-3 {
					isXMAS(grid[y][x] + grid[y+1][x+1] + grid[y+2][x+2] + grid[y+3][x+3])
				}
				// diagonal to left
				if x > 2 {
					isXMAS(grid[y][x] + grid[y+1][x-1] + grid[y+2][x-2] + grid[y+3][x-3])
				}
			}
			// second challenge find two mas words in x shape
			if y > 0 && x > 0 && y < gridLength-1 && x < gridLength-1 {

				if grid[y][x] == "A" && isMAS(grid[y-1][x-1]+grid[y+1][x+1]) && isMAS(grid[y-1][x+1]+grid[y+1][x-1]) {
					secondChallengeMatches++

				}
			}
		}
	}
	fmt.Println("day four first challenge output: ", matches)
	fmt.Println("day four second challenge output: ", secondChallengeMatches)

}
