package problem10

import (
	"fmt"
	"strconv"
)

type testCase1 struct {
	input       []int
	stdListSize int
	result      int
}

type testCase2 struct {
	input  string
	result string
}

var testCasesPartOne = []testCase1{
	{input: []int{3, 4, 1, 5}, stdListSize: 5, result: 12},
	{input: []int{76, 1, 88, 148, 166, 217, 130, 0, 128, 254, 16, 2, 130, 71, 255, 229}, stdListSize: 256, result: -1},
}

var testCasesPartTwo = []testCase2{
	{input: ``, result: `a2582a3a0e66e6e86e3812dcb672a272`},
	{input: `AoC 2017`, result: `33efeb34ea91902bb2f59c9920caa6cd`},
	{input: `1,2,3`, result: `3efbe78a8d82f29979031a4aa0b16a9d`},
	{input: `1,2,4`, result: `63960835bcdc130f0b66d7ff4f6a5a8e`},
	{input: `76,1,88,148,166,217,130,0,128,254,16,2,130,71,255,229`, result: ``},
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
	fmt.Printf(">>> PART TWO <<<\n\n")
	for k, v := range testCasesPartTwo {
		fmt.Printf("test case : %d\n", k)
		if len(v.input) > 30 {
			fmt.Printf("input     : [TOO LONG VALUE]\n")
		} else {
			fmt.Printf("input     : %#v\n", v.input)
		}
		fmt.Printf("expected  : %s\n", v.result)
		fmt.Print("result    : ")
		fmt.Print(v.solvePartTwo())
		fmt.Print("\n\n")
	}
}

func (tc *testCase1) solvePartOne() int {
	lst := []int{}
	for i := 0; i < tc.stdListSize; i++ {
		lst = append(lst, i)
	}

	currPos := 0
	skipSize := 0

	for _, length := range tc.input {
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
	}

	return lst[0] * lst[1]
}

func (tc *testCase2) solvePartTwo() string {
	input := []int32{}
	for _, x := range tc.input {
		input = append(input, x)
	}
	inputAdd := []int32{17, 31, 73, 47, 23}
	for _, x := range inputAdd {
		input = append(input, x);
	}

	// ===================================

	lst := []int32{}
	for i := int32(0); i < 256; i++ {
		lst = append(lst, i)
	}

	currPos := int32(0)
	skipSize := int32(0)

	for times := 0; times < 64; times++ {
		for _, length := range input {
			// step 1
			tmpLst := []int32{}
			for i, j := int32(0), currPos; i < length; i, j = i+1, j+1 {
				tmpLst = append(tmpLst, lst[j])
				if j == int32(len(lst)-1) {
					j = -1
				}
			}
			tmpLst = reverseInts32(tmpLst)

			// step 2
			tmpLstIdx := currPos
			for _, tmpLstItem := range tmpLst {
				lst[tmpLstIdx] = tmpLstItem
				tmpLstIdx++
				if tmpLstIdx == int32(len(lst)) {
					tmpLstIdx = 0
				}
			}

			// step 3
			currPos += length + skipSize
			currPos %= int32(len(lst))
			skipSize++
		}
	}

	// ===================================

	//fmt.Printf("%#v\n", lst)
	//fmt.Printf("%#v\n", len(lst))

	denseHash := denseHashCalc(lst)
	return formatHash(denseHash)
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func reverseInts32(input []int32) []int32 {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts32(input[1:]), input[0])
}

func getAscii(x int) int {
	tmp := strconv.Itoa(x)
	return int([]rune(tmp)[0])
}

func denseHashCalc(inp []int32) []int32 {
	res := []int32{}
	for i := 0; i < 16; i++ {
		blk := []int32{}
		for j := 0; j < 16; j++ {
			blk = append(blk, inp[i*16+j])
		}
		tmp := blk[0]
		for j := 1; j < 16; j++ {
			tmp = tmp ^ blk[j]
		}
		res = append(res, tmp)
	}
	return res
}

func formatHash(inp []int32) string {
	res := ""
	for _, num := range inp {
		fmted := fmt.Sprintf("%2x", num)
		res += fmted
	}
	return res
}
