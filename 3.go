package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func day_3() {
	// input
	data, err := os.ReadFile("./inputs/3.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	content := string(data)

	// First Challenge
	r, err := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	if err != nil {
		fmt.Println("Error compiling regex experesson: ", err)
	}
	arr := r.FindAllString(content, -1)

	var sum int
	for _, mul := range arr {
		rNumbers, err := regexp.Compile(`[0-9]{1,3}`)
		if err != nil {
			fmt.Println("Error compiling numbers regex experesson: ", err)
		}
		numbersArr := rNumbers.FindAllString(mul, 2)
		firstNumber, _ := strconv.Atoi(numbersArr[0])
		secondNumber, _ := strconv.Atoi(numbersArr[1])
		sum += (firstNumber * secondNumber)
	}

	// Second Challenge
	rWithInstructions, err := regexp.Compile(`(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don't\(\))`)
	if err != nil {
		fmt.Println("Error compiling regex experesson: ", err)
	}
	arrWithInstruction := rWithInstructions.FindAllString(content, -1)

	var sumWithInstructions int
	do := true
	for _, instruction := range arrWithInstruction {
		if instruction == "do()" {
			do = true
			continue
		} else if instruction == "don't()" {
			do = false
			continue
		} else if do {
			rNumbers, err := regexp.Compile(`[0-9]{1,3}`)
			if err != nil {
				fmt.Println("Error compiling numbers regex experesson: ", err)
			}
			numbersArr := rNumbers.FindAllString(instruction, 2)
			firstNumber, _ := strconv.Atoi(numbersArr[0])
			secondNumber, _ := strconv.Atoi(numbersArr[1])
			sumWithInstructions += (firstNumber * secondNumber)
		}
	}

	fmt.Println("day 3 first challenge output: ", sum)
	fmt.Println("day 3 second challenge output: ", sumWithInstructions)

}
