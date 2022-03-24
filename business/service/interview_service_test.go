package service

import (
	"fmt"
	"testing"
)

func TestOutputStrAndNum(t *testing.T) {
	OutputStrAndNum()
}

func TestGetSquareSortedArr(t *testing.T) {
	arr := []int{-4,-3,0,1,2,5}
	result := GetSquareSortedArr(arr)
	fmt.Println(result)
}
