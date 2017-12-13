package problem13

import (
	"fmt"
)

type testCase struct {
	input  []string
	result int
}

var commonTestCase = testCase{input: []string{"0: 3", "1: 2", "2: 4", "4: 4", "6: 5", "8: 6", "10: 6", "12: 8", "14: 6", "16: 6", "18: 9", "20: 8", "22: 8", "24: 8", "26: 12", "28: 8", "30: 12", "32: 12", "34: 12", "36: 10", "38: 14", "40: 12", "42: 10", "44: 8", "46: 12", "48: 14", "50: 12", "52: 14", "54: 14", "56: 14", "58: 12", "62: 14", "64: 12", "66: 12", "68: 14", "70: 14", "72: 14", "74: 17", "76: 14", "78: 18", "84: 14", "90: 20", "92: 14"}, result: -1}

var testCasesPartOne = []testCase{
	{input: []string{"0: 3", "1: 2", "4: 4", "6: 4"}, result: 24},
	commonTestCase,
}

var testCasesPartTwo = []testCase{
	{input: []string{"0: 3", "1: 2", "4: 4", "6: 4"}, result: 24},
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

func (tc *testCase) solvePartOne() int {
	return 0
}

func (tc *testCase) solvePartTwo() int {
	return 0
}
