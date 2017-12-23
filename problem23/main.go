package problem23

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
	{input: []string{"set b 67", "set c b", "jnz a 2", "jnz 1 5", "mul b 100", "sub b -100000", "set c b", "sub c -17000", "set f 1", "set d 2", "set e 2", "set g d", "mul g e", "sub g b", "jnz g 2", "set f 0", "sub e -1", "set g e", "sub g b", "jnz g -8", "sub d -1", "set g d", "sub g b", "jnz g -13", "jnz f 2", "sub h -1", "set g b", "sub g c", "jnz g 2", "jnz 1 3", "sub b -17", "jnz 1 -23"}, result: 4225},
}

var testCasesPartTwo = []testCase{
	{input: []string{"set b 67", "set c b", "jnz a 2", "jnz 1 5", "mul b 100", "sub b -100000", "set c b", "sub c -17000", "set f 1", "set d 2", "set e 2", "set g d", "mul g e", "sub g b", "jnz g 2", "set f 0", "sub e -1", "set g e", "sub g b", "jnz g -8", "sub d -1", "set g d", "sub g b", "jnz g -13", "jnz f 2", "sub h -1", "set g b", "sub g c", "jnz g 2", "jnz 1 3", "sub b -17", "jnz 1 -23"}, result: -1},
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
	instrs := tc.input
	curr := 0

	regs := map[string]int{}
	ensureReg := func(name string) {
		if _, ok := regs[name]; !ok {
			regs[name] = 0
		}
	}

	mulTimes := 0

Loop:
	for {
		if curr < 0 || curr == len(instrs) {
			break Loop
		}
		instr := instrs[curr]

		switch instr[:3] {
		case "set":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err == nil {
				regs[parts[1]] = val
			} else {
				ensureReg(parts[2])
				regs[parts[1]] = regs[parts[2]]
			}

		case "sub":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err == nil {
				regs[parts[1]] -= val
			} else {
				ensureReg(parts[2])
				regs[parts[1]] -= regs[parts[2]]
			}

		case "mul":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err == nil {
				regs[parts[1]] *= val
			} else {
				ensureReg(parts[2])
				regs[parts[1]] *= regs[parts[2]]
			}
			mulTimes++

		case "jnz":
			parts := strings.Split(instr, " ")
			val1, err1 := strconv.Atoi(parts[1])
			if err1 != nil {
				ensureReg(parts[1])
				val1 = regs[parts[1]]
			}
			if val1 != 0 {
				val2, err2 := strconv.Atoi(parts[2])
				offset := 0
				if err2 == nil {
					offset = val2
				} else {
					ensureReg(parts[2])
					offset = regs[parts[2]]
				}
				curr += offset
				continue Loop
			}
		}

		curr++
	}

	return mulTimes
}

func (tc *testCase) solvePartTwo() int {
	instrs := tc.input
	curr := 0

	regs := map[string]int{}
	ensureReg := func(name string) {
		if _, ok := regs[name]; !ok {
			regs[name] = 0
		}
	}
	regs["a"] = 1

	prevD := -1

Loop:
	for {
		if regs["d"] != prevD {
			fmt.Println(regs["d"])
			prevD = regs["d"]
			regss := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
			for _, reg := range regss {
				fmt.Print(reg + " " + strconv.Itoa(regs[reg]) + " ")
			}
			fmt.Println()
		}

		//if curr < 0 || curr == len(instrs) {
		//	break Loop
		//}

		//if _, ok := regs["h"]; ok {
		//	fmt.Println(regs["h"])
		//}

		//regss := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
		//for _, reg := range regss {
		//	fmt.Print(reg + " " + strconv.Itoa(regs[reg]) + " ")
		//}
		//fmt.Println()
		//time.Sleep(time.Millisecond*500)

		instr := instrs[curr]

		switch instr[:3] {
		case "set":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err == nil {
				regs[parts[1]] = val
			} else {
				ensureReg(parts[2])
				regs[parts[1]] = regs[parts[2]]
			}

		case "sub":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err == nil {
				regs[parts[1]] -= val
			} else {
				ensureReg(parts[2])
				regs[parts[1]] -= regs[parts[2]]
			}

		case "mul":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err == nil {
				regs[parts[1]] *= val
			} else {
				ensureReg(parts[2])
				regs[parts[1]] *= regs[parts[2]]
			}

		case "jnz":
			parts := strings.Split(instr, " ")
			val1, err1 := strconv.Atoi(parts[1])
			if err1 != nil {
				ensureReg(parts[1])
				val1 = regs[parts[1]]
			}
			if val1 != 0 {
				val2, err2 := strconv.Atoi(parts[2])
				offset := 0
				if err2 == nil {
					offset = val2
				} else {
					ensureReg(parts[2])
					offset = regs[parts[2]]
				}
				curr += offset
				continue Loop
			}
		}

		curr++
	}

	return 0
}
