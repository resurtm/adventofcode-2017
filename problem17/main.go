package problem17

import (
	"fmt"
)

type testCase struct {
	input  int
	result int
}

var testCasesPartOne = []testCase{
	{input: 3, result: 638},
	{input: 354, result: -1},
}

var testCasesPartTwo = []testCase{
}

func RunPartOne() {
	fmt.Printf(">>> PART ONE <<<\n\n")
	for k, v := range testCasesPartOne {
		fmt.Printf("test case : %d\n", k)
		fmt.Printf("input     : %d\n", v.input)
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
		fmt.Printf("input     : %d\n", v.input)
		fmt.Printf("expected  : %d\n", v.result)
		fmt.Print("result    : ")
		fmt.Print(v.solvePartTwo())
		fmt.Print("\n\n")
	}
}

func (tc *testCase) solvePartOne() int {
	vals := []int{0}
	pos := 0

	for i := 0; i < 2017; i++ {
		//fmt.Println(vals)

		for step := 0; step < tc.input; step++ {
			pos++
			if pos >= len(vals) {
				pos = 0
			}
		}

		pos++
		vals = append(vals[:pos], append([]int{i + 1}, vals[pos:]...)...)
	}

	j := 0
	for ; j < len(vals); j++ {
		if vals[j] == 2017 {
			break
		}
	}

	return vals[j+1]
}

func (tc *testCase) solvePartTwo() int {
	return 0
}
