package service

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 1, 3, 7, 4, 8, 2, 9, 6}
	sortedArr := MergeSort(arr)
	fmt.Println(sortedArr)
}

func TestQuickSort(t *testing.T) {
	arr := []int{5, 1, 3, 7, 4, 8, 2, 9, 6}
	sortedArr := QuickSort(arr)
	fmt.Println(sortedArr)
}
