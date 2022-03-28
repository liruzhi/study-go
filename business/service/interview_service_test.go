package service

import (
	"fmt"
	"testing"
)

func TestOutputStrAndNum(t *testing.T) {
	//for i:='A';i<='Z';i++ {
	//	fmt.Println(string(i))
	//}
	OutputStrAndNum()
}

func TestDumpNum(t *testing.T) {
	DumpNum()
}

func TestDumpLevelTree(t *testing.T) {
	DumpLevelTree()
}

func TestGetSquareSortedArr(t *testing.T) {
	arr := []int{-4, -3, 0, 1, 2, 5}
	result := GetSquareSortedArr(arr)
	fmt.Println(result)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	//str := "abcabcdefghijklmffffabcdegh"
	//str := "pwwkew"
	//str := "pwwpkew"
	//str := "abcabcbb"
	str := "bbbbb"

	fmt.Println(LengthOfLongestSubstring(str))
}
