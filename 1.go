package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("1.csv")
	if err != nil {
		fmt.Println("Error opening file", err)
		return

	}
	defer file.Close()
	var inputArr1 = make([]int, 0, 1000)
	var inputArr2 = make([]int, 0, 1000)
	reader := csv.NewReader(file)

	recordes, err := reader.ReadAll()
	for _, record := range recordes {
		value, _ := strconv.Atoi(record[0])
		inputArr1 = append(inputArr1, value)
		secondValue, _ := strconv.Atoi(record[1])
		inputArr2 = append(inputArr2, secondValue)
	}

	sort.Ints(inputArr1)
	sort.Ints(inputArr2)
	DifferenceSum := 0
	for index, number := range inputArr1 {

		if number-inputArr2[index] > 0 {
			DifferenceSum += number - inputArr2[index]
		} else if number-inputArr2[index] < 0 {
			DifferenceSum += (number - inputArr2[index]) * -1
		}
	}
	fmt.Println(DifferenceSum)
	// Second Challenge Find Similarities	occurrences := make(map[int]int)

	occurrences := make(map[int]int)
	binarysearch := func(number int) bool {
		high := 999
		low := 0

		for low <= high {
			mid := low + (high-low)/2

			if inputArr1[mid] == number {
				occurrences[number] += 1

				return true
			} else if inputArr1[mid] > number {
				high = mid - 1
			} else if inputArr1[mid] < number {
				low = mid + 1
			}

		}
		return false
	}

	for _, num := range inputArr2 {
		binarysearch(num)
	}
	var sum int
	for i, num := range occurrences {
		sum += i * num
	}
	fmt.Println(sum)
}
