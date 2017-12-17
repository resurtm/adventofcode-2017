package problem17

import (
	"fmt"
	"container/list"
)

type testCase struct {
	input  int
	result int
}

var testCasesPartOne = []testCase{
	{input: 3, result: 638},
	{input: 354, result: 2000},
}

var testCasesPartTwo = []testCase{
	{input: 3, result: 1222153},
	{input: 354, result: 10242889},
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
	vls := list.New()
	vls.PushBack(0)
	e := vls.Front()

	for i := 0; i < 2017; i++ {
		for step := 0; step < tc.input; step++ {
			e = e.Next()
			if e == nil {
				e = vls.Front()
			}
		}

		e = e.Next()
		if e == nil {
			e = vls.Front()
		}

		vls.InsertAfter(i+1, e)
	}

	e = e.Next()
	if e == nil {
		e = vls.Front()
	}
	e = e.Next()
	if e == nil {
		e = vls.Front()
	}

	return (e.Value).(int)
}

func (tc *testCase) solvePartTwo() int {
	vls := list.New()
	vls.PushBack(0)
	e := vls.Front()

	for i := 0; i < 50000000; i++ {
		//if i%100000 == 0 {
		//	fmt.Println(i)
		//}

		for step := 0; step < tc.input; step++ {
			e = e.Next()
			if e == nil {
				e = vls.Front()
			}
		}

		e = e.Next()
		if e == nil {
			e = vls.Front()
		}

		vls.InsertAfter(i+1, e)
	}

	xx := vls.Front()
	for ; xx != nil; xx = xx.Next() {
		//fmt.Print(xx.Value, ",")
		if xx.Value == 0 {
			break
		}
	}
	//fmt.Println()

	return (xx.Next().Value).(int)
}
