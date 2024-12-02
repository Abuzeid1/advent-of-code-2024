package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_2() {
	file, err := os.Open("./inputs/2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var safeRecordersCount int
	var falseRecorders [][]int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		levels := make([]int, len(parts))
		for i, part := range parts {
			levels[i], _ = strconv.Atoi(part)
		}
		_, is := isSafe(levels)
		if is {
			safeRecordersCount++
		} else {
			falseRecorders = append(falseRecorders, levels)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error Scanning file:", err)
	}
	fmt.Println("day 2 first challenge output:", safeRecordersCount)

	for _, levels := range falseRecorders {
		index, _ := isSafe(levels)
		if index == len(levels)-2 {
			safeRecordersCount++
		} else {

			for i := index - 1; i < index+2; i++ {

				var arr = make([]int, len(levels))
				copy(arr, levels)

				var newArr []int
				if i < 0 {
					i = 0
				}
				newArr = append(arr[0:i], arr[i+1:]...)

				_, is := isSafe(newArr)
				if is {

					safeRecordersCount++
					break
				}
			}

		}
	}
	fmt.Println("day 2 second challenge output", safeRecordersCount)
}

func isSafe(arr []int) (int, bool) {
	isIncreasing := false

	for i, num := range arr {
		if i >= len(arr)-1 {
			break
		}
		nextNum := arr[i+1]
		distance := nextNum - num
		if distance > 0 && i == 0 {
			isIncreasing = true
		}

		if distance > 3 || distance < -3 || distance == 0 || (distance > 0 && !isIncreasing) || (distance < 0 && isIncreasing) {
			return i, false
		}
	}

	return -1, true
}
