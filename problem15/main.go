package problem15

import (
	"fmt"
	"strconv"
)

type testCase struct {
	input1 int64
	input2 int64
	result int64
}

var testCasesPartOne = []testCase{
	{input1: 65, input2: 8921, result: 588},
	{input1: 783, input2: 325, result: 650},
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
	gena := tc.input1
	genb := tc.input2
	cnt := 0
	for iter := 0; iter < 40000000; iter++ {
		gena *= 16807
		genb *= 48271

		gena %= 2147483647
		genb %= 2147483647

		genas := fmt.Sprintf("%032s", strconv.FormatInt(gena, 2))
		genbs := fmt.Sprintf("%032s", strconv.FormatInt(genb, 2))

		if genas[16:] == genbs[16:] {
			cnt++
		}
	}
	return cnt
}

func (tc *testCase) solvePartTwo() int {
	return 0
}
