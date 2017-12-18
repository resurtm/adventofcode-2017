package problem18

import (
	"fmt"
	"strings"
	"strconv"
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
	curr := 0

	lastSnd := 0
	resultSnd := 0

	regs := map[string]int{}
	ensureReg := func(name string) {
		if _, ok := regs[name]; !ok {
			regs[name] = 0
		}
	}

Loop:
	for {
		instr := tc.input[curr]

		switch instr[:3] {
		case "set":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				regs[parts[1]] = regs[parts[2]]
			} else {
				regs[parts[1]] = val
			}

		case "add":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				regs[parts[1]] += regs[parts[2]]
			} else {
				regs[parts[1]] += val
			}

		case "mul":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			ensureReg(parts[2])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				regs[parts[1]] *= regs[parts[2]]
			} else {
				regs[parts[1]] *= val
			}

		case "mod":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				regs[parts[1]] %= regs[parts[2]]
			} else {
				regs[parts[1]] %= val
			}

		case "snd":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			lastSnd = regs[parts[1]]

		case "rcv":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			if regs[parts[1]] != 0 {
				resultSnd = lastSnd
				break Loop
			}

		case "jgz":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, _ := strconv.Atoi(parts[2])
			if regs[parts[1]] > 0 {
				curr += val
				continue Loop
			}
		}

		curr++
	}

	return resultSnd
}

func (tc *testCase) solvePartTwo() int {
	return 0
}
