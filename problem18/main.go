package problem18

import (
	"fmt"
)

type testCase struct {
	input  []string
	result int
}

var testCasesPartOne = []testCase{
	{input: []string{"set a 1", "add a 2", "mul a a", "mod a 5", "snd a", "set a 0", "rcv a", "jgz a -1", "set a 1", "jgz a -2",}, result: 4},
	{input: []string{"set i 31", "set a 1", "mul p 17", "jgz p p", "mul a 2", "add i -1", "jgz i -2", "add a -1", "set i 127", "set p 680", "mul p 8505", "mod p a", "mul p 129749", "add p 12345", "mod p a", "set b p", "mod b 10000", "snd b", "add i -1", "jgz i -9", "jgz a 3", "rcv b", "jgz b -1", "set f 0", "set i 126", "rcv a", "rcv b", "set p a", "mul p -1", "add p b", "jgz p 4", "snd a", "set a b", "jgz 1 3", "snd b", "set f 1", "add i -1", "jgz i -11", "snd a", "jgz f -16", "jgz a -19",}, result: -1},
}

var testCasesPartTwo = []testCase{
}

func RunPartOne() {
	fmt.Printf(">>> PART ONE <<<\n\n")
	for k, v := range testCasesPartOne {
		fmt.Printf("test case : %d\n", k)
		fmt.Printf("input     : %#v\n", v.input)
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
		fmt.Printf("input     : %#v\n", v.input)
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
