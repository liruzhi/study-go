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

func TestA(t *testing.T) {
	var a uint = 1
	var b uint = 2

	fmt.Println(a-b)
}

func BenchmarkBinarySearch1(b *testing.B) {

}

func BenchmarkFor10000(b *testing.B)  {
	for i:=0;i<b.N;i++ {
		For10000()
	}
}

func TestCopy(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置

	fmt.Println(slice1, slice2)
}

func TestTopK(t *testing.T) {
	arr := []int{13, 7, 14, 4, 7, 6, 2,8,12, 10, 6, 13, 15}
	result := TopK(arr, 7)

	fmt.Println(result)
}
