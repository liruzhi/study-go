package service

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
	n := len(arr1)
	m := len(arr2)

	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0

	for {
		if p1 == n {
			sorted = append(sorted, arr2[p2:]...)
			break
		}

		if p2 == m {
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

	return
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
