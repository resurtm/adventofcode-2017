package problem15

import (
	"fmt"
)

type testCase struct {
	input1 int
	input2 int
	result int
}

var testCasesPartOne = []testCase{
	{input1: 65, input2: 8921, result: 588},
}

var testCasesPartTwo = []testCase{}

func RunPartOne() {
	fmt.Printf(">>> PART ONE <<<\n\n")
	for k, v := range testCasesPartOne {
		fmt.Printf("test case : %d\n", k)
		fmt.Printf("input1    : %#v\n", v.input1)
		fmt.Printf("input2    : %#v\n", v.input2)
		fmt.Printf("expected  : %d\n", v.result)
		fmt.Print("result    : ")
		fmt.Print(v.solvePartOne())
		fmt.Print("\n\n")
	}
}

func RunPartTwo() {
	fmt.Printf(">>> PART TWO <<<\n\n")
	for k, v := range testCasesPartTwo {
		fmt.Printf("test case : %d\n", k)
		fmt.Printf("input1    : %#v\n", v.input1)
		fmt.Printf("input2    : %#v\n", v.input2)
		fmt.Printf("expected  : %d\n", v.result)
		fmt.Print("result    : ")
		fmt.Print(v.solvePartTwo())
		fmt.Print("\n\n")
	}
}

func (tc *testCase) solvePartOne() int {

	return 0
}

func (tc *testCase) solvePartTwo() int {
	return 0
}
