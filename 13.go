package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_13() {
	input, err := os.ReadFile("./inputs/13.txt")
	if err != nil {
		fmt.Println("Error Reading file ", err)
	}
	arr := strings.Split(string(input), "\n\n")
	tokens := 0
	correctedTokens := 0

	for _, machine := range arr {

		ax, _ := strconv.Atoi(machine[strings.Index(machine, "X+")+1 : strings.Index(machine, ",")])
		ay, _ := strconv.Atoi(machine[strings.Index(machine, "Y+")+1 : strings.Index(machine, "\n")])

		secondLine := machine[strings.Index(machine, "\n")+1:]
		bx, _ := strconv.Atoi(secondLine[strings.Index(secondLine, "X+")+1 : strings.Index(secondLine, ",")])
		by, _ := strconv.Atoi(secondLine[strings.Index(secondLine, "Y+")+1 : strings.Index(secondLine, "\n")])

		thirdLine := secondLine[strings.Index(secondLine, "\n")+1:]
		x, _ := strconv.Atoi(thirdLine[strings.Index(thirdLine, "X=")+2 : strings.Index(thirdLine, ",")])
		y, _ := strconv.Atoi(thirdLine[strings.Index(thirdLine, "Y=")+2:])

		na := (bx*y - by*x) / (bx*ay - by*ax)
		nb := (ay*x - ax*y) / (ay*bx - ax*by)

		if na*ax+nb*bx == x && na*ay+nb*by == y {
			tokens += na*3 + nb
		}
		// second challenge calculation
		x += 10000000000000
		y += 10000000000000
		na = (bx*y - by*x) / (bx*ay - by*ax)
		nb = (ay*x - ax*y) / (ay*bx - ax*by)

		if na*ax+nb*bx == x && na*ay+nb*by == y {
			correctedTokens += na*3 + nb
		}

	}

	fmt.Println("day 13 first challenge", tokens)
	fmt.Println("day 13 second challenge", correctedTokens)
}

// The behind the scene Math
// n1*ax + n2*bx = x
// n1*ay + n2*by = y

// n1 = (x - n2*bx) / ax
// n1 = (y - n2*by)/ay
// (x-n2*bx)/ax = (y-n2*by)/ay
// (x-n2*bx)*ay = (y-n2*by)*ax
// ay*x - ay*n2*bx = ax*y - ax*n2*by
// ay*x - ax*y = ay*n2*bx -ax*n2*by
// ay*x - ax*y = n2(ay*bx-ax*by)
// n2=(ay*x-ax*y)/(ay*bx-ax*by)

// n2 = (x-n1*ax)/bx
// n2 = (y-n1*ay)/by
// (y-n1*ay)/by = (x-n1*ax)/bx
// bx*y - bx*n1*ay = by*x - by*n1*ax
// bx*y - by*x = bx*n1*ay - by*n1*ax
// bx*y - by*x = n1(bx*ay - by*ax)
// n1 = (bx*y - by*x)/(bx*ay - by*ax)
