package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func day_11() {

	input, err := os.ReadFile("./inputs/11.txt")
	if err != nil {
		fmt.Println("Error Reading File", err)
	}

	numberOfStonesAfterBlinkingNtimes := func(depth int) int {
		var count int
		for _, stone := range strings.Split(string(input), " ") {
			num, _ := strconv.Atoi(stone)
			count += blink(num, depth)
		}

		return count
	}
	fmt.Println("day_11 first challenge output: ", numberOfStonesAfterBlinkingNtimes(25))  //213625
	fmt.Println("day_11 second challenge output: ", numberOfStonesAfterBlinkingNtimes(75)) //252442982856820

}

var cachedBlinks = make(map[[2]int]int)

func blink(stone int, depth int) int {
	if depth == 0 {
		return 1
	}
	if cachedBlinks[[2]int{stone, depth}] > 0 {
		return cachedBlinks[[2]int{stone, depth}]
	}
	numberOfDigits := int(math.Log10(float64(stone))) + 1
	var count int
	if len(strconv.Itoa(stone))%2 == 0 {
		firstNum := stone / int(math.Pow(10, float64(numberOfDigits)/2))
		secondNum := stone % int(math.Pow(10, float64(numberOfDigits)/2))
		count = blink(firstNum, depth-1) + blink(secondNum, depth-1)

	} else if stone == 0 {
		count = blink(1, depth-1)
	} else {
		count = blink(stone*2024, depth-1)
	}
	cachedBlinks[[2]int{stone, depth}] = count
	return count

}
