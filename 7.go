package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_7() {

	input, err := os.ReadFile("./inputs/7.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	for key, line := range strings.Split(string(input), "\n") {
		parts := strings.Split(line, ":")
		sum, _ := strconv.Atoi(parts[0])

		numbersPart := strings.Split(strings.TrimSpace(parts[1]), " ")
		operate(numbersPart, 0, 0, [2]int{key, sum})
		operateContact(numbersPart, 0, 0, [2]int{key, sum})

	}
	totalSum := 0
	for key, possibleSums := range sums {
		for _, sum := range possibleSums {
			if sum == key[1] {
				totalSum += sum
				break
			}
		}
	}
	// second Challenge Sums
	totalSecondSum := 0
	for key, possibleSums := range secondSums {
		for _, sum := range possibleSums {
			if sum == key[1] {
				totalSecondSum += sum
				break
			}
		}
	}
	fmt.Println("Day 7 first challenge: ", totalSum)
	fmt.Println("Day 7 second challenge: ", totalSecondSum)
}

var sums = make(map[[2]int][]int)

// First Challenge
func operate(arr []string, sum int, index int, key [2]int) {
	if index >= len(arr) {
		if key[1] == sum {
			sums[key] = append(sums[key], sum)
		}
	} else {
		num, _ := strconv.Atoi(arr[index])
		if index == 0 {
			operate(arr, num, index+1, key)

		} else {
			operate(arr, sum+num, index+1, key)
			operate(arr, sum*num, index+1, key)
		}
	}
}

var secondSums = make(map[[2]int][]int)

func operateContact(arr []string, sum int, index int, key [2]int) {
	if index >= len(arr) {
		if key[1] == sum {

			secondSums[key] = append(secondSums[key], sum)
		}
	} else {
		num, _ := strconv.Atoi(arr[index])
		if index == 0 {
			operateContact(arr, num, index+1, key)

		} else {
			operateContact(arr, sum+num, index+1, key)
			operateContact(arr, sum*num, index+1, key)
			string1 := strconv.Itoa(sum) + arr[index]
			sum, _ = strconv.Atoi(string1)
			operateContact(arr, sum, index+1, key)
		}
	}
}

// 248427118972289
