package problem14

import (
	"fmt"
	"strings"
	"strconv"
)

type testCase struct {
	input  string
	result int
}

var testCasesPartOne = []testCase{
	{input: "flqrgnkx", result: 8108},
	{input: "hfdlxzhv", result: -1},
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
	cnt := 0
	for i := 0; i < 128; i++ {
		row := hexToBin(knotHash(tc.input + "-" + strconv.Itoa(i)))
		for _, rowItem := range row {
			if string(rowItem) == "1" {
				cnt++
			}
		}
	}
	return cnt
}

func (tc *testCase) solvePartTwo() int {
	return 0
}

func knotHash(inp string) string {
	input := []int32{}
	for _, x := range inp {
		input = append(input, x)
	}
	inputAdd := []int32{17, 31, 73, 47, 23}
	for _, x := range inputAdd {
		input = append(input, x);
	}

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

	denseHash := denseHashCalc(lst)
	return strings.Replace(formatHash(denseHash), " ", "0", -1)
}

func reverseInts32(input []int32) []int32 {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts32(input[1:]), input[0])
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

func hexToBin(hex string) string {
	lookup := [16]string{
		"0000", "0001", "0010", "0011", "0100", "0101", "0110", "0111",
		"1000", "1001", "1010", "1011", "1100", "1101", "1110", "1111",
	}
	result := ""
	for _, val := range hex {
		vl, _ := strconv.ParseUint(string(val), 16, 64)
		result += lookup[vl]
	}
	return result
}
