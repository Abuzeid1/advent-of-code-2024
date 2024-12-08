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

	totalSum := 0       //1298300076754
	totalSecondSum := 0 //248427118972289

	for _, line := range strings.Split(string(input), "\n") {
		parts := strings.Split(line, ":")
		sum, _ := strconv.Atoi(parts[0])

		numbersPart := strings.Split(strings.TrimSpace(parts[1]), " ")

		for _, num := range operate(numbersPart, 0) {
			if num == sum {
				totalSum += sum
				break
			}
		}
		for _, num := range operateContact(numbersPart, 0) {
			if num == sum {
				totalSecondSum += sum
				break
			}
		}

	}

	// second Challenge Sums

	fmt.Println("Day 7 first challenge: ", totalSum)
	fmt.Println("Day 7 second challenge: ", totalSecondSum)
}

// First Challenge
func operate(arr []string, sum int) []int {
	// fmt.Println(sum, index)
	if len(arr) <= 0 {
		return []int{sum}
	} else {

		num, _ := strconv.Atoi(arr[0])

		arr = arr[1:]

		if sum == 0 {
			return operate(arr, sum+num)

		} else {
			return append(operate(arr, sum+num), operate(arr, sum*num)...)
		}
	}
}

func operateContact(arr []string, sum int) []int {

	if len(arr) <= 0 {

		return []int{sum}

	} else {
		num, _ := strconv.Atoi(arr[0])
		string1 := strconv.Itoa(sum) + arr[0]
		arr = arr[1:]
		if sum == 0 {
			return operateContact(arr, num)

		} else {
			getConcatenation := func() int {
				concatenation, _ := strconv.Atoi(string1)
				return concatenation
			}
			return append(append(operateContact(arr, sum+num), operateContact(arr, sum*num)...), operateContact(arr, getConcatenation())...)

		}
	}
}
