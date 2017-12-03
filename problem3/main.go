package problem3

import (
	"fmt"
	"math"
)

type testCase struct {
	input  int
	result int
}

var testCasesPartOne = []testCase{
	{input: 1, result: 0},
	{input: 12, result: 3},
	{input: 23, result: 2},
	{input: 1024, result: 31},
	{input: 277678, result: -1},
}

var testCasesPartTwo = []testCase{
	{result: -1},
}

func RunPartOne() {
	fmt.Printf("PART ONE:\n\n")
	for k, v := range testCasesPartOne {
		fmt.Printf("test case %d:\n", k+1)
		fmt.Printf("result: %d\n", v.result)
		fmt.Printf("%#v\n", v)
		fmt.Printf("result: %d\n\n", v.solvePartOne())
	}
}

func RunPartTwo() {
	return
	fmt.Printf("PART TWO:\n\n")
	for k, v := range testCasesPartTwo {
		fmt.Printf("test case %d:\n", k+1)
		fmt.Printf("result %d:\n", v.result)
		fmt.Printf("%#v\n", v)
		fmt.Printf("result: %#v\n\n", v.solvePartTwo())
	}
}

func (tc *testCase) solvePartOne() int {
	const (
		right    = iota
		up
		left
		down
		resetDir
	)

	directions, direction := []int{}, right
	for i := 1; i < 1000; i += 1 {
		for k := 0; k < 2; k++ {
			for j := 0; j < i; j++ {
				directions = append(directions, direction)
			}
			direction += 1
			if direction == resetDir {
				direction = right
			}
		}
	}

	const size = 600
	elems := make([][]int, size)
	for i := range elems {
		elems[i] = make([]int, size)
	}

	currX, currY := int(size/2), int(size/2)
	dirN, dir := 0, directions[0]
	number := 1

	for {
		if (number == tc.input) {
			break
		}

		elems[currX][currY] = number
		number += 1

		switch dir {
		case right:
			currX += 1
		case up:
			currY -= 1
		case left:
			currX -= 1
		case down:
			currY += 1
		}

		dirN += 1
		dir = directions[dirN]

		if (currX == size || currY == size || currX == -1 || currY == -1) {
			break
		}
	}

	//fmt.Println(number)
	//printElems(elems)

	return int(math.Abs(float64(currX-int(size/2)))) +
		int(math.Abs(float64(currY-int(size/2))))
}

func (tc *testCase) solvePartTwo() int {
	return 0
}

func printElems(elems [][]int) {
	for _, m := range elems {
		for _, n := range m {
			fmt.Printf("%3d", n)
		}
		fmt.Printf("\n")
	}
}
