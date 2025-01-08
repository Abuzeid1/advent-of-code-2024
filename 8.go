package main

import (
	"fmt"
	"os"
)

type Position struct {
	y int
	x int
}

func day_8() {
	input, err := os.ReadFile("./inputs/8.txt")

	if err != nil {
		fmt.Println("Error readinf file: ", err)
	}
	firstChallengePositions := make(map[int]string)
	secondChallengePositions := make(map[int]string)

	rowLength := 0
	for i := 0; i < len(input); i++ {
		if input[i] == "\n"[0] {
			rowLength = i + 1
			break
		}
	}
	columnLength := len(input)/rowLength + 1

	getIndex := func(p Position) (index int, isInside bool) {
		if p.x >= 0 && p.x < rowLength && p.y >= 0 && p.y < columnLength {
			return rowLength*p.y + p.x, true
		}

		return 0, false
	}

	for firstAntennaIndex := 0; firstAntennaIndex < len(input); firstAntennaIndex++ {
		if input[firstAntennaIndex] == "."[0] || input[firstAntennaIndex] == "\n"[0] {
			continue
		} else {

			for secondAntennaIndex := firstAntennaIndex + 1; secondAntennaIndex < len(input); secondAntennaIndex++ {
				if string(input[firstAntennaIndex]) == string(input[secondAntennaIndex]) {
					getXY := func(index int) Position { return Position{x: index % rowLength, y: index / rowLength} }
					p1 := getXY(firstAntennaIndex)
					p2 := getXY(secondAntennaIndex)
					xDistance := p2.x - p1.x
					yDistance := p2.y - p1.y
					for _, item := range [2]Position{{x: p1.x - xDistance, y: p1.y - yDistance}, {x: p2.x + xDistance, y: p2.y + yDistance}} {
						index, isInside := getIndex(item)
						if isInside && input[index] != "\n"[0] {
							firstChallengePositions[index] = string(input[firstAntennaIndex])
						}
					}

					var divider int
					if xDistance < yDistance && xDistance >= 0 {
						divider = xDistance
					} else if xDistance*-1 < yDistance && xDistance < 0 {
						divider = xDistance * -1
					} else {
						divider = yDistance
					}
					xSmallestStep := xDistance
					ySmallestStep := yDistance
					for ; divider > 1; divider-- {

						if xSmallestStep%divider == 0 && ySmallestStep%divider == 0 {
							fmt.Println("divider", divider)
							xSmallestStep /= divider
							ySmallestStep /= divider
						}
					}
					firstPosition := p1
					for {
						index, isInside := getIndex(firstPosition)
						if !isInside || input[index] == "\n"[0] {
							firstPosition.x += xSmallestStep
							firstPosition.y += ySmallestStep
							break
						}
						firstPosition.x -= xSmallestStep
						firstPosition.y -= ySmallestStep
					}

					for {
						index, isInside := getIndex(firstPosition)
						if !isInside || input[index] == "\n"[0] {
							break
						} else {
							secondChallengePositions[index] = string(input[firstAntennaIndex])
						}

						firstPosition.x += xSmallestStep
						firstPosition.y += ySmallestStep
					}
				}
			}
		}
	}

	fmt.Println("day 8 first challenge out put: ", len(firstChallengePositions))
	fmt.Println("day 8 second challenge out put: ", len(secondChallengePositions))

}
