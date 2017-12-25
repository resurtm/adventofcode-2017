package problem25

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"bufio"
	"strings"
)

type testCase struct {
	input  string
	result int
}

type stateDesc struct {
	whenZeroValue     byte
	whenZeroToLeft    bool
	whenZeroNextState byte

	whenOneValue     byte
	whenOneToLeft    bool
	whenOneNextState byte
}

var testCasesPartOne = []testCase{
	{input: "input0.txt", result: 3},
	{input: "input1.txt", result: 4385},
}

var testCasesPartTwo = []testCase{
	//{input: "input0.txt", result: -1},
	//{input: "input1.txt", result: -1},
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
	initial, steps, states := parseData(readInput(tc.input))

	const storageSize = 20000000
	storage := make([]byte, storageSize)
	curPos := storageSize / 2
	curState := initial

	for step := 0; step < steps; step++ {
		st := states[curState]
		if storage[curPos] == 0 {
			storage[curPos] = st.whenZeroValue
			if st.whenZeroToLeft {
				curPos--
			} else {
				curPos++
			}
			curState = st.whenZeroNextState
		} else {
			storage[curPos] = st.whenOneValue
			if st.whenOneToLeft {
				curPos--
			} else {
				curPos++
			}
			curState = st.whenOneNextState
		}
	}

	ones :=0
	for _, i:=range storage{
		if i == 1{
			ones++
		}
	}
	return ones
}

func (tc *testCase) solvePartTwo() int {
	fmt.Println(readInput(tc.input))
	return 0
}

func parseData(data []string) (initialState byte, steps int, states map[byte]stateDesc) {
	tempS := ""
	fmt.Sscanf(data[0], "Begin in state %s.\n", &tempS)
	initialState = tempS[0]

	fmt.Sscanf(data[1], "Perform a diagnostic checksum after %d steps.\n", &steps)

	stateCount := 0
	for _, elem := range data {
		if strings.Contains(elem, "In state ") {
			stateCount++
		}
	}

	states = make(map[byte]stateDesc)
	for i, cp := 0, 3; i < stateCount; i, cp = i+1, cp+10 {
		state := stateDesc{}

		fmt.Sscanf(data[cp], "In state %s:\n", &tempS)
		stateName := tempS[0]

		fmt.Sscanf(data[cp+2], "    - Write the value %d.\n", &state.whenZeroValue)

		fmt.Sscanf(data[cp+3], "    - Move one slot to the %s.\n", &tempS)
		state.whenZeroToLeft = tempS == "left."

		fmt.Sscanf(data[cp+4], "    - Continue with state %s.\n", &tempS)
		state.whenZeroNextState = tempS[0]

		fmt.Sscanf(data[cp+6], "    - Write the value %d.\n", &state.whenOneValue)

		fmt.Sscanf(data[cp+7], "    - Move one slot to the %s.\n", &tempS)
		state.whenOneToLeft = tempS == "left."

		fmt.Sscanf(data[cp+8], "    - Continue with state %s.\n", &tempS)
		state.whenOneNextState = tempS[0]

		states[stateName] = state
	}
	return
}

func readInput(fn string) []string {
	fp := getCurrDir()
	if fp[:5] == "/tmp/" {
		fp = "/home/resurtm/dev/go/src/github.com/resurtm/adventofcode-2017/problem25"
	} else {
		fp = filepath.Dir(fp)
	}
	fp += "/" + fn

	file, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func getCurrDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
