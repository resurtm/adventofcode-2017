package problem10

import (
	"fmt"
)

type testCase struct {
	input       []int
	stdListSize int
	result      int
}

var testCasesPartOne = []testCase{
	{input: []int{3, 4, 1, 5}, stdListSize: 5, result: 12},
	{input: []int{76, 1, 88, 148, 166, 217, 130, 0, 128, 254, 16, 2, 130, 71, 255, 229}, stdListSize: 256, result: -1},
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
		fmt.Printf("expected  : %d\n", v.result)
		fmt.Print("result    : ")
		fmt.Print(v.solvePartOne())
		fmt.Print("\n\n")
	}
}

func RunPartTwo() {
	return
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
	lst := []int{}
	for i := 0; i < tc.stdListSize; i ++ {
		lst = append(lst, i)
	}

	currPos := 0
	skipSize := 0

	for _, length := range tc.input {
		//fmt.Println("=================")

		// step 1
		tmpLst := []int{}
		for i, j := 0, currPos; i < length; i, j = i+1, j+1 {
			tmpLst = append(tmpLst, lst[j])
			if j == len(lst)-1 {
				j = -1
			}
		}
		tmpLst = reverseInts(tmpLst)

		// step 2
		tmpLstIdx := currPos
		for _, tmpLstItem := range tmpLst {
			lst[tmpLstIdx] = tmpLstItem
			tmpLstIdx++
			if tmpLstIdx == len(lst) {
				tmpLstIdx = 0
			}
		}

		// step 3
		currPos += length + skipSize
		currPos %= len(lst)
		skipSize++

		//fmt.Printf("%#v\n", lst)
	}

	return lst[0] * lst[1]
}

func (tc *testCase) solvePartTwo() int {
	return 0
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
