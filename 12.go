package main

import (
	"fmt"
	"os"
	"strings"
)

func day_12() {

	input, err := os.ReadFile("inputs/12.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	var inputArr [][]string
	for _, line := range strings.Split(string(input), "\n") {
		inputArr = append(inputArr, strings.Split(line, ""))
	}

	totalPrice := 0
	secondTotalPrice := 0

	var current string
	for y, _ := range inputArr {
		for x, _ := range inputArr[y] {
			if inputArr[y][x] != current {
				current = inputArr[y][x]
				perimeters, areas, corners := calculate(inputArr, inputArr[y][x], y, x)

				totalPrice += perimeters * areas
				secondTotalPrice += corners * areas
			}
		}
	}

	fmt.Println("day 12 first challenge output:", totalPrice)
	fmt.Println("day 12 second challenge output:", secondTotalPrice)
}

var visitedAreas = make(map[[2]int]bool)

func calculate(inputArr [][]string, prev string, y int, x int) (perimeters int, areas int, corners int) {

	if y < 0 || y >= len(inputArr) || x < 0 || x >= len(inputArr[0]) {
		return 0, 0, 0
	} else if prev != inputArr[y][x] {
		return 0, 0, 0
	} else if visitedAreas[[2]int{y, x}] {
		return 0, 0, 0
	}
	visitedAreas[[2]int{y, x}] = true
	areas = 1
	perimeters = 0
	corners = 0
	yAfter := false
	yBefore := false
	xAfter := false
	xBefore := false
	if y-1 < 0 || inputArr[y-1][x] != inputArr[y][x] {
		perimeters += 1
		yBefore = true
	}
	if y+1 >= len(inputArr) || inputArr[y+1][x] != inputArr[y][x] {
		perimeters += 1
		yAfter = true
	}
	if x-1 < 0 || inputArr[y][x-1] != inputArr[y][x] {
		perimeters += 1
		xBefore = true
	}
	if x+1 >= len(inputArr[0]) || inputArr[y][x+1] != inputArr[y][x] {
		perimeters += 1
		xAfter = true
	}

	// top left corner
	if (yBefore && xBefore) || (!yBefore && xBefore && x-1 >= 0 && y-1 >= 0 && inputArr[y-1][x-1] == prev) {
		corners++

	}
	// top right corner
	if (yBefore && xAfter) || (!yBefore && xAfter && x+1 < len(inputArr[0]) && y-1 >= 0 && inputArr[y-1][x+1] == prev) {
		corners++
	}
	// bottom left corner
	if (yAfter && xBefore) || (!yAfter && xBefore && x-1 >= 0 && y+1 < len(inputArr) && inputArr[y+1][x-1] == prev) {
		corners += 1
	}
	// bottom right corner
	if (yAfter && xAfter) || (!yAfter && xAfter && x+1 < len(inputArr[0]) && inputArr[y+1][x+1] == prev) {
		corners++
	}

	p1, a1, c1 := calculate(inputArr, inputArr[y][x], y+1, x)
	p2, a2, c2 := calculate(inputArr, inputArr[y][x], y-1, x)
	p3, a3, c3 := calculate(inputArr, inputArr[y][x], y, x+1)
	p4, a4, c4 := calculate(inputArr, inputArr[y][x], y, x-1)

	return perimeters + p1 + p2 + p3 + p4, areas + a1 + a2 + a3 + a4, corners + c1 + c2 + c3 + c4

}
