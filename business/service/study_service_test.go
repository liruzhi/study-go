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

func TestAppend11(t *testing.T) {
	a := make([]int, 0)

	a = append(a, 1, 2, 3)
	b := append(a, 6)
	c := append(a, 5)

	fmt.Println(a, b, c)
}


func BenchmarkMergeSort(b *testing.B) {

	//arr := []int{1,3,2,6,5,4,8,9,10,19,28,17,11,12,13,14,15,16,7,1,2,3,4,5,6,7,8,9,0}
	arr := getArr(100000)
	for i:=0; i< b.N; i++ {
		arr = MergeSort(arr)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	//arr := []int{1,3,2,6,5,4,8,9,10,19,28,17,11,12,13,14,15,16,7}
	arr := getArr(100000)
	for i:=0; i< b.N; i++ {
		QuickSort(arr)
	}
}

func TestGetArr1(t *testing.T) {
	result := getArr(10000)
	fmt.Println(result)
	fmt.Println(cap(result), len(result))
}

func TestFindTwoMaxNum(t *testing.T) {
	arr := []int{1,12,2,9,8,7,5,10,4,3,11}

	if len(arr) < 2 {
		return
	}
	var firstMax,secondMax int

	if arr[0] > arr[1] {
		firstMax, secondMax = arr[0], arr[1]
	} else {
		firstMax, secondMax = arr[1], arr[0]
	}

	for i:=2;i<len(arr);i++ {
		if arr[i] >= firstMax {
			firstMax, secondMax = arr[i], firstMax
			continue
		}
		if arr[i] > secondMax {
			secondMax = arr[i]
		}
	}

	fmt.Println(firstMax,secondMax)
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
	//arr := []int{13, 7, 14, 4, 7, 6, 2,8,12, 10, 6, 13, 15}
	arr := []int{2,1,1,1,3}
	result := TopK(arr, 2)

	fmt.Println(result)
}
