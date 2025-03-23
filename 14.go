package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_14() {
	input, err := os.ReadFile("./inputs/14.txt")
	maxX := 101
	maxY := 103
	if err != nil {
		fmt.Println("Error Reading File", input)
	}
	calculateSafetyFactor := func(num int) int {
		q1 := 0
		q2 := 0
		q3 := 0
		q4 := 0
		for _, line := range strings.Split(string(input), "\n") {
			px, _ := strconv.Atoi(line[strings.Index(line, "=")+1 : strings.Index(line, ",")])
			py, _ := strconv.Atoi(line[strings.Index(line, ",")+1 : strings.Index(line, " ")])
			vx, _ := strconv.Atoi(line[strings.LastIndex(line, "=")+1 : strings.LastIndex(line, ",")])
			vy, _ := strconv.Atoi(line[strings.LastIndex(line, ",")+1:])

			px += num * vx
			py += num * vy
			px %= maxX
			py %= maxY

			if px < 0 {
				px = maxX + px
			}
			if py < 0 {
				py = maxY + py
			}

			if px < maxX/2 && py < maxY/2 {
				q1 += 1
			} else if px > maxX/2 && py < maxY/2 {
				q2 += 1
			} else if px < maxX/2 && py > maxY/2 {
				q3 += 1
			} else if px > maxX/2 && py > maxY/2 {
				q4 += 1
			}
		}
		return q1 * q2 * q3 * q4
	}

	fmt.Println("day 14 first challenge", calculateSafetyFactor(100))

	minimum := calculateSafetyFactor(100)
	num := 100
	for index := range 10_000 {
		if calculateSafetyFactor(index+1) < minimum {
			minimum = calculateSafetyFactor(index + 1)
			num = index + 1
		}
	}
	fmt.Println(minimum, num)

	var grid [103][101]int
	for _, line := range strings.Split(string(input), "\n") {
		px, _ := strconv.Atoi(line[strings.Index(line, "=")+1 : strings.Index(line, ",")])
		py, _ := strconv.Atoi(line[strings.Index(line, ",")+1 : strings.Index(line, " ")])
		vx, _ := strconv.Atoi(line[strings.LastIndex(line, "=")+1 : strings.LastIndex(line, ",")])
		vy, _ := strconv.Atoi(line[strings.LastIndex(line, ",")+1:])

		px += num * vx
		py += num * vy
		px %= maxX
		py %= maxY

		if px < 0 {
			px = maxX + px
		}
		if py < 0 {
			py = maxY + py
		}
		grid[py][px] = 1

	}
	for _, line := range grid {
		for _, num := range line {
			if num == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(string('x'))

			}
		}
		fmt.Println()
	}

}
