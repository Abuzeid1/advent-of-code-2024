package main

import (
	"fmt"
	"os"
	"strconv"
)

func day_9() {
	input, err := os.ReadFile("./inputs/9.txt")
	if err != nil {
		fmt.Println("Error Reading File")
	}

	var line []int
	var fileId int
	for index, part := range input {
		blockCount, _ := strconv.Atoi(string(part))

		for n := 0; n < blockCount; n++ {
			if index%2 == 0 {
				line = append(line, fileId)
				if n == blockCount-1 {
					fileId += 1
				}
			} else {
				line = append(line, -1)
			}
		}
	}
	var firstLine []int
	firstLine = append(firstLine, line...)
	lastIndex := len(firstLine) - 1
	var checkSum int
	for index := 0; index <= lastIndex; index++ {
		if firstLine[index] == -1 {
			for i := lastIndex; i > index; i-- {
				if firstLine[i] != -1 {
					firstLine[i], firstLine[index] = firstLine[index], firstLine[i]
					lastIndex = i - 1
					break
				}
			}
		}
		if firstLine[index] != -1 {
			checkSum += firstLine[index] * index
		}
	}

	lastIndex = len(line) - 1
	firstIndex := 0

	for index := len(line) - 1; index > firstIndex; index-- {

		if line[index] != -1 {
			fileBlockCount := 1
			for i := index; i > firstIndex && line[i] == line[i-1]; i-- {
				fileBlockCount++
				index -= 1
			}

			first := false

			for n := firstIndex; n < index; n++ {
				if line[n] == -1 {
					if !first {
						firstIndex = n
						first = true
					}
					emptyBlockCount := 1

					for x := n; line[x] == -1 && line[x+1] == -1; x++ {
						emptyBlockCount++
						n++
					}

					if emptyBlockCount >= fileBlockCount {

						for i := 0; i < fileBlockCount; i++ {
							line[index+i], line[(n-emptyBlockCount)+i+1] = line[(n-emptyBlockCount)+i+1], line[index+i]

						}

						break
					}
				}

			}
		}
	}
	secondCheckSum := 0
	for index, block := range line {
		if block != -1 {
			secondCheckSum += block * index
		}
	}

	fmt.Println("day 9 first Challenge output", checkSum)        // 6283404590840
	fmt.Println("day 9 second Challenge output", secondCheckSum) // 6304576012713

}
