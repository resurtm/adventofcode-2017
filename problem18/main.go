package problem18

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	"runtime"
)

type testCase struct {
	input  []string
	result int
}

var testCasesPartOne = []testCase{
	{input: []string{"set a 1", "add a 2", "mul a a", "mod a 5", "snd a", "set a 0", "rcv a", "jgz a -1", "set a 1", "jgz a -2"}, result: 4},
	{input: []string{"set i 31", "set a 1", "mul p 17", "jgz p p", "mul a 2", "add i -1", "jgz i -2", "add a -1", "set i 127", "set p 680", "mul p 8505", "mod p a", "mul p 129749", "add p 12345", "mod p a", "set b p", "mod b 10000", "snd b", "add i -1", "jgz i -9", "jgz a 3", "rcv b", "jgz b -1", "set f 0", "set i 126", "rcv a", "rcv b", "set p a", "mul p -1", "add p b", "jgz p 4", "snd a", "set a b", "jgz 1 3", "snd b", "set f 1", "add i -1", "jgz i -11", "snd a", "jgz f -16", "jgz a -19"}, result: 3188},
}

var testCasesPartTwo = []testCase{
	{input: []string{"snd 1", "snd 2", "snd p", "rcv a", "rcv b", "rcv c", "rcv d"}, result: 3},
	{input: []string{"set i 31", "set a 1", "mul p 17", "jgz p p", "mul a 2", "add i -1", "jgz i -2", "add a -1", "set i 127", "set p 680", "mul p 8505", "mod p a", "mul p 129749", "add p 12345", "mod p a", "set b p", "mod b 10000", "snd b", "add i -1", "jgz i -9", "jgz a 3", "rcv b", "jgz b -1", "set f 0", "set i 126", "rcv a", "rcv b", "set p a", "mul p -1", "add p b", "jgz p 4", "snd a", "set a b", "jgz 1 3", "snd b", "set f 1", "add i -1", "jgz i -11", "snd a", "jgz f -16", "jgz a -19"}, result: 7112},
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
	return runProg1(tc.input)
}

func runProg1(instrs []string) int {
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
		instr := instrs[curr]

		switch instr[:3] {
		case "set":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				ensureReg(parts[2])
				regs[parts[1]] = regs[parts[2]]
			} else {
				regs[parts[1]] = val
			}

		case "add":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				ensureReg(parts[2])
				regs[parts[1]] += regs[parts[2]]
			} else {
				regs[parts[1]] += val
			}

		case "mul":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				ensureReg(parts[2])
				regs[parts[1]] *= regs[parts[2]]
			} else {
				regs[parts[1]] *= val
			}

		case "mod":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				ensureReg(parts[2])
				regs[parts[1]] %= regs[parts[2]]
			} else {
				regs[parts[1]] %= val
			}

		case "snd":
			parts := strings.Split(instr, " ")
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				ensureReg(parts[1])
				lastSnd = regs[parts[1]]
			} else {
				lastSnd = val
			}

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

var ch1, ch2 chan int
var cnt []int

func (tc *testCase) solvePartTwo() int {
	runtime.GOMAXPROCS(4)

	const chSize = 100000000
	ch1 = make(chan int, chSize)
	ch2 = make(chan int, chSize)
	cnt = make([]int, 2)

	go runProg2(tc.input, 0)
	go runProg2(tc.input, 1)

	time.Sleep(time.Second * 5)

	return cnt[1]
}

func runProg2(instrs []string, pid int) int {
	sendTo := ch1
	if pid == 1 {
		sendTo = ch2
	}
	receiveFrom := ch2
	if pid == 1 {
		receiveFrom = ch1
	}

	curr := 0

	regs := map[string]int{}
	regs["p"] = pid

	ensureReg := func(name string) {
		if _, ok := regs[name]; !ok {
			regs[name] = 0
		}
	}

Loop:
	for {
		//time.Sleep(time.Millisecond * 50)
		instr := instrs[curr]
		//fmt.Printf("%d: inst %s\n", pid, instr)

		switch instr[:3] {
		case "set":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				ensureReg(parts[2])
				regs[parts[1]] = regs[parts[2]]
			} else {
				regs[parts[1]] = val
			}

		case "add":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				ensureReg(parts[2])
				regs[parts[1]] += regs[parts[2]]
			} else {
				regs[parts[1]] += val
			}

		case "mul":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				ensureReg(parts[2])
				regs[parts[1]] *= regs[parts[2]]
			} else {
				regs[parts[1]] *= val
			}

		case "mod":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				ensureReg(parts[2])
				regs[parts[1]] %= regs[parts[2]]
			} else {
				regs[parts[1]] %= val
			}

		case "snd":
			parts := strings.Split(instr, " ")
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				ensureReg(parts[1])
				sendTo <- regs[parts[1]]
				//fmt.Printf("%d: send %d\n", pid, regs[parts[1]])
			} else {
				sendTo <- val
				//fmt.Printf("%d: send %d\n", pid, val)
			}
			cnt[pid]++

		case "rcv":
			parts := strings.Split(instr, " ")
			ensureReg(parts[1])
			regs[parts[1]] = <-receiveFrom
			//fmt.Printf("%d: recv %d\n", pid, regs[parts[1]])

		case "jgz":
			parts := strings.Split(instr, " ")
			val1, err1 := strconv.Atoi(parts[1])
			if err1 != nil {
				ensureReg(parts[1])
				val1 = regs[parts[1]]
			}
			if val1 > 0 {
				val2, err2 := strconv.Atoi(parts[2])
				if err2 != nil {
					ensureReg(parts[2])
					curr += regs[parts[2]]
				} else {
					curr += val2
				}
				continue Loop
			}
		}

		curr++
	}

	return 0
}
