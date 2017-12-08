package problem8

import (
	"fmt"
)

type testCase struct {
	input  string
	result string
}

var testCasesPartOne = []testCase{
}

var testCasesPartTwo = []testCase{
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
		fmt.Printf("expected  : %s\n", v.result)
		fmt.Printf("result    : %s\n\n", v.solvePartOne())
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
		fmt.Printf("expected  : %s\n", v.result)
		fmt.Printf("result    : %s\n\n", v.solvePartTwo())
	}
}

func (tc *testCase) solvePartOne() string {
	return ""
}

func (tc *testCase) solvePartTwo() string {
	return ""
}
