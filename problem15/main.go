package problem15

import (
	"fmt"
)

type testCase struct {
	input1 uint
	input2 uint
	result uint
}

var testCasesPartOne = []testCase{
	{input1: 65, input2: 8921, result: 588},
	{input1: 783, input2: 325, result: 650},
}

var testCasesPartTwo = []testCase{
	{input1: 65, input2: 8921, result: 309},
	{input1: 783, input2: 325, result: 336},
}

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

		if gena-gena>>16<<16 == genb-genb>>16<<16 {
			cnt++
		}
	}
	return cnt
}

func (tc *testCase) solvePartTwo() int {
	gena := tc.input1
	genb := tc.input2
	cnt := 0
	iter := 0
	fa := uint(0)
	fb := uint(0)

	for {
		if fa == 0 {
			gena *= 16807
			gena %= 2147483647
		}

		if fb == 0 {
			genb *= 48271
			genb %= 2147483647
		}

		if gena%4 == 0 {
			fa = gena
		}
		if genb%8 == 0 {
			fb = genb
		}

		if fa != 0 && fb != 0 {
			if fa-fa>>16<<16 == fb-fb>>16<<16 {
				cnt++
			}

			fa = 0
			fb = 0

			iter++
			if iter > 5000000 {
				break
			}
		}
	}

	return cnt
}
