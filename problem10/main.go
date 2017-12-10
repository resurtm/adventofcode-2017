package problem10

import (
	"fmt"
)

type testCase struct {
	input  string
	result int64
}

var commonTestCase = testCase{
	input:  ``,
	result: -1,
}

var testCasesPartOne = []testCase{
	{input: ``, result: -1},
	commonTestCase,
}

var testCasesPartTwo = []testCase{
	{input: ``, result: -1},
	commonTestCase,
}

func RunPartOne() {
	fmt.Printf(">>> PART ONE <<<\n\n")
	for k, v := range testCasesPartOne {
		fmt.Printf("test case : %d\n", k)
		if len(v.input) > 30 {
			fmt.Printf("input     : [TOO LONG VALUE]\n")
		} else {
			fmt.Printf("input     : %#v\n", v.input)
		}
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
		if len(v.input) > 30 {
			fmt.Printf("input     : [TOO LONG VALUE]\n")
		} else {
			fmt.Printf("input     : %#v\n", v.input)
		}
		fmt.Printf("expected  : %d\n", v.result)
		fmt.Print("result    : ")
		fmt.Print(v.solvePartTwo())
		fmt.Print("\n\n")
	}
}

func (tc *testCase) solvePartOne() int64 {
	return 0
}

func (tc *testCase) solvePartTwo() int64 {
	return 0
}
