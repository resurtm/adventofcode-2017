package problem3

import (
	"fmt"
)

type testCase struct {
	result int
}

var testCasesPartOne = []testCase{
	{
		result: -1,
	},
	{
		result: -1,
	},
}

var testCasesPartTwo = []testCase{
	{
		result: -1,
	},
	{
		result: -1,
	},
}

func RunPartOne() {
	fmt.Printf("PART ONE:\n\n")

	for k, v := range testCasesPartOne {
		fmt.Printf("test case %d:\n", k+1)
		fmt.Printf("result %d:\n", v.result)
		fmt.Printf("%#v\n", v)
		fmt.Printf("result: %#v\n\n", v.solvePartOne())
	}
}

func RunPartTwo() {
	fmt.Printf("PART TWO:\n\n")

	for k, v := range testCasesPartTwo {
		fmt.Printf("test case %d:\n", k+1)
		fmt.Printf("result %d:\n", v.result)
		fmt.Printf("%#v\n", v)
		fmt.Printf("result: %#v\n\n", v.solvePartTwo())
	}
}

func (tc *testCase) solvePartOne() int {
	return 0
}

func (tc *testCase) solvePartTwo() int {
	return 0
}
