package service

import (
	"fmt"
	"math/rand"
	"strings"
)

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
	if start >= len(arr) - k {
		return start
	}

	q := Partition(arr, start, end)

	if q == len(arr) - k {
		return q
	}

	if q > len(arr) - k {
		return TopKPartition(arr, k, start, q - 1)
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

//解析字符串
func ParseString(str string) {
	//按逗号分割成数组
	splitedArr := strings.Split(str, ",")
	//数组下标与引号个数的映射
	indexQuoteCountMap := make(map[int]int)
	//非独立字段的起始点映射
	startEndMap := make(map[int]int)
	resultArr := make([]string, 0)
	start := -1

	for i, v := range splitedArr {
		//获取引号的数量
		indexQuoteCountMap[i] = GetQuoteCount([]rune(v))
		//获取两个相邻的引号数量为奇数的坐标，构成一个区间map
		if indexQuoteCountMap[i]%2 == 0 {
			continue
		}

		//非独立字段起点
		if start == -1 {
			start = i
		} else {
			//非独立字段终点
			startEndMap[start] = i
			start = -1
		}
	}

	end := -1
	//需要拼接的字段
	var needJoinColumn string
	for i, v := range splitedArr {
		if startEndMap[i] > 0 {
			end = startEndMap[i]
		}

		//独立字段直接追加
		if end == -1 {
			resultArr = append(resultArr, ReplaceQuote(v))
			continue
		}

		//非独立字段需要拼接
		if end != i {
			needJoinColumn = needJoinColumn +  ReplaceQuote(v) + ","
			continue
		}

		//非独立字段部分结尾
		needJoinColumn = needJoinColumn + ReplaceQuote(v)
		resultArr = append(resultArr, needJoinColumn)

		//拼接结束，还原
		needJoinColumn = ""
		end = -1
	}
	fmt.Println(strings.Join(resultArr, "\t"))
}

func ReplaceQuote(str string) string {
	strRuneArr := []rune(str)
	result := make([]rune, 0)
	//连续引号的个数
	continuousQuote := 0
	for _, v := range strRuneArr {
		if string(v) == `"` {
			continuousQuote++
			if continuousQuote%2 == 0 {
				result = append(result, v)
			}
		} else {
			continuousQuote = 0
			result = append(result, v)
		}
	}
	return string(result)
}

func GetQuoteCount(params []rune) (count int) {
	for _, v := range params {
		if string(v) == `"` {
			count++
		}
	}
	return count
}
