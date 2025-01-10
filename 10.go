package main

import (
	"fmt"
	"os"
	"strconv"
)

func day_10() {
	input, err := os.ReadFile("./inputs/10.txt")
	if err != nil {
		fmt.Println("Error Reading File")
	}
	var inputArr [][]int
	line := 0
	for index, position := range input {
		num, err := strconv.Atoi(string(position))
		if index == 0 || err != nil {
			inputArr = append(inputArr, []int{})
		}
		if err != nil {
			line += 1
		} else {

			inputArr[line] = append(inputArr[line], num)
		}
	}

	var firstScore int
	var secondScore int
	for y := 0; y < len(inputArr); y++ {
		for x := 0; x < len(inputArr[0]); x++ {
			if inputArr[y][x] == 0 {
				mapArr := make(map[[2]int]bool)
				arr := step(y, x, inputArr, -1, [][2]int{})
				for _, item := range arr {
					mapArr[[2]int{item[0], item[1]}] = true
				}
				firstScore += len(mapArr)
				secondScore += len(arr)

			}
		}
	}
	fmt.Println("day_10 first challenge output: ", firstScore)   //646
	fmt.Println("day_10 second challenge output: ", secondScore) //646
}

func step(y int, x int, inputArr [][]int, prev int, arr [][2]int) [][2]int {

	if y < 0 || y >= len(inputArr) || x < 0 || x >= len(inputArr[0]) {
		return arr
	} else if inputArr[y][x] == 9 && prev == 8 {

		return append(arr, [2]int{y, x})
	} else if prev+1 == inputArr[y][x] {

		return append(append(append(step(y-1, x, inputArr, inputArr[y][x], arr), step(y+1, x, inputArr, inputArr[y][x], arr)...), step(y, x-1, inputArr, inputArr[y][x], arr)...), step(y, x+1, inputArr, inputArr[y][x], arr)...)
	}
	return arr

}
