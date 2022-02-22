package service

import "math/rand"

type studyService struct {
}

func NewStudyService() *studyService {
	return &studyService{}
}

func (this *studyService) Sort() ([]int, error) {

	return nil, nil
}

//归并排序
func MergeSort(arr []int) []int {
	return MergeSortC(arr, 0, len(arr)-1)
}

func MergeSortC(arr []int, start, end int) []int {
	//递归调用的终止条件
	if start >= end {
		return arr[start : end+1]
	}

	middle := (start + end) / 2

	left := MergeSortC(arr, start, middle)
	right := MergeSortC(arr, middle+1, end)

	return MergeTwoSortdArr(left, right)
}

//合并两个有序数组
func MergeTwoSortdArr(arr1 []int, arr2 []int) []int {
	sorted := make([]int, 0, len(arr1) + len(arr2))
	p1, p2 := 0, 0

	for {
		if p1 == len(arr1) {
			sorted = append(sorted, arr2[p2:]...)
			break
		}

		if p2 == len(arr2) {
			sorted = append(sorted, arr1[p1:]...)
			break
		}

		if arr1[p1] < arr2[p2] {
			sorted = append(sorted, arr1[p1])
			p1++
		} else {
			sorted = append(sorted, arr2[p2])
			p2++
		}
	}

	return sorted
}

//快速排序
func QuickSort(arr []int) []int {
	QuickSortC(arr, 0, len(arr)-1)

	return arr
}

func QuickSortC(arr []int, start, end int) {
	if start >= end {
		return
	}

	//获取分区点
	q := Partition(arr, start, end)

	QuickSortC(arr, start, q-1)
	QuickSortC(arr, q+1, end)
}



//二分查找(非递归)
func BinarySearch1(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := (left + right) / 2
		if arr[middle] == target {
			return middle
		} else if arr[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}

//二分查找递归
func BinarySearch2(arr []int, target int) int {
	return BinarySearchC(arr, target, 0, len(arr)-1)
}

func BinarySearchC(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	middle := (left + right) / 2

	if arr[middle] == target {
		return middle
	}

	if arr[middle] < target {
		return BinarySearchC(arr, target, middle+1, right)
	}

	return BinarySearchC(arr, target, left, middle-1)
}

//反转链表
//https://leetcode-cn.com/problems/reverse-linked-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func getArr(maxNum int) []int {
	arr := make([]int, maxNum, maxNum)

	for i := 0; i < maxNum; i++ {
		arr[i] = rand.Intn(maxNum)
	}

	return arr
}
func For10000() {
	for j := 0; j < 800000; j++ {
		if j < 80000 {

		}
	}
}

func TopK (arr[]int , k int) []int{
	return arr[TopKPartition(arr, k, 0, len(arr) - 1):]
}

func TopKPartition(arr []int, k ,start ,end int) int {
	if len(arr) - start <= k {
		return start
	}

	q := Partition(arr, start, end)

	if q == len(arr) - k {
		return q
	}

	if q > len(arr) - k {
		return TopKPartition(arr, k, start, q-1)
	}

	return TopKPartition(arr, k, q + 1, end)
}

//获取分区点
func Partition(arr []int, start, end int) int {
	i := start

	for j := start; j <= end-1; j++ {
		if arr[j] < arr[end] {
			//互换
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	return i
}
