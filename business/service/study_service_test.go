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

func TestBinarySearch1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}
	target := 10
	t.Log(BinarySearch1(arr, target))
}

func TestBinarySearch2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}
	target := 10
	t.Log(BinarySearch2(arr, target))
}
