package main

import (
	"fmt"
	"os"
	"strings"
)

func day_6() {

	input, err := os.ReadFile("./inputs/6.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	length := strings.Index(string(input), "\n") + 1

	guardPosition := strings.Index(string(input), "^")
	direction := -length
	changeDirection := func(dir int) int {
		if dir == -length {
			return 1
		} else if dir == 1 {
			return length
		} else if dir == length {
			return -1
		} else {
			return -length
		}
	}
	replaceCharachter := func(ch string) {
		input = append(append(input[0:guardPosition], []byte(ch)...), input[guardPosition+1:]...)
	}
	distinctiveMoves := 0
	loops := 0

	for {
		if guardPosition >= len(input) || guardPosition < 0 || string(input[guardPosition]) == "\n" {
			break
		}
		if string(input[guardPosition]) == "#" {
			guardPosition -= direction
			replaceCharachter("+")
			direction = changeDirection(direction)
		} else {
			if string(input[guardPosition]) != "^" {
				dir := changeDirection(direction)
				gPosition := guardPosition
				obstacles := 0
				for {

					gPosition += dir
					if gPosition >= len(input) || gPosition < 0 || string(input[gPosition]) == "\n" {
						break
					}
					if string(input[gPosition]) == "#" {
						gPosition -= dir
						if string(input[gPosition]) == "+" {
							loops++
							break
						}
						dir = changeDirection(dir)
						obstacles++
					}
					if obstacles > 10 {
						break
					}
				}
			}
			if direction == length || direction == -length {

				if string(input[guardPosition]) == "." || string(input[guardPosition]) == "^" {
					replaceCharachter("|")
					distinctiveMoves++
				} else if string(input[guardPosition]) == "_" {
					replaceCharachter("+")
				}
			} else if direction == 1 || direction == -1 {
				if string(input[guardPosition]) == "." || string(input[guardPosition]) == "^" {
					replaceCharachter("_")
					distinctiveMoves++

				} else if string(input[guardPosition]) == "|" {
					replaceCharachter("+")
				}
			}
		}
		guardPosition += direction

	}

	fmt.Println("day 6 first challenge ouput: ", distinctiveMoves)
	fmt.Println("day 6 second Challenge output: ", loops)
}
