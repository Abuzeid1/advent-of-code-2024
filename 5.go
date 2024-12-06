package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_5() {

	input, err := os.ReadFile("./inputs/5.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	data := [2]string(strings.SplitN(string(input), "\n\n", 2))
	rules, instructions := data[0], data[1]

	var rulesArr [][2]string

	for i := 0; i < len(rules); i += 6 {

		before := string([]byte{rules[i], rules[i+1]})
		after := string([]byte{rules[i+3], rules[i+4]})
		rulesArr = append(rulesArr, [2]string{before, after})

	}
	isBefore := func(before string, after string) bool {
		for _, rule := range rulesArr {
			if before == rule[0] && after == rule[1] {
				return true
			} else if before == rule[1] && after == rule[0] {
				return false
			}
		}
		return true
	}
	// sum the middle number in each update
	sortedUpdatesSum := 0   // first challenge
	unSortedUpdatesSum := 0 //second challenge

	for _, line := range strings.Split(instructions, "\n") {
		// first challenge sum sorted updates middle number
		isSorted := true
		lineArr := strings.Split(line, ",")
		for i := 0; i < len(lineArr)-1; i++ {

			if !isBefore(lineArr[i], lineArr[i+1]) {
				isSorted = false

				break
			}
		}
		if isSorted {
			num, _ := strconv.Atoi(lineArr[len(lineArr)/2])
			sortedUpdatesSum += num
		} else {
			// second challenge sort unsorted updates & sum their middle numbers
			for x := 0; x < len(lineArr)-1; x++ {

				if !isBefore(lineArr[x], lineArr[x+1]) {
					lineArr[x], lineArr[x+1] = lineArr[x+1], lineArr[x]
					if x > 0 {
						x -= 2
					}
				}

			}
			num, _ := strconv.Atoi(lineArr[len(lineArr)/2])
			unSortedUpdatesSum += num
			isSorted = false
		}

	}
	fmt.Println("day 5 first challenge output: ", sortedUpdatesSum)
	fmt.Println("day 5 second challenge output: ", unSortedUpdatesSum)
}
