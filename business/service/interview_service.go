package service

import "fmt"

var chan1 = make(chan int, 1)
var chan2 = make(chan int)
var finishChan= make(chan bool)

func OutputStrAndNum() {
	chan1 <- 1
	go f1()
	go f2()
	<-finishChan
}

func f1() {
	for i:=1;i<11;i++ {
		<-chan1
		fmt.Println(i)
		chan2<-1
	}
}

func f2() {
	str := "ABCDEFGHIJ"
	for i:=0;i<10;i++ {
		<-chan2
		fmt.Println(string(str[i]))
		chan1<-1
	}
	close(chan1)
	close(chan2)
	close(finishChan)
}

//给定一个有序数组，返回平方后依然有序的数组
func GetSquareSortedArr(arr []int) []int {
	n := len(arr)

	result := make([]int, n)
	i,j := 0,n - 1

	for index:= n-1;index >=0;index-- {
		if v,w := arr[i] * arr[i],arr[j] * arr[j];v > w {
			result[index] = v
			i++
		} else {
			result[index] = w
			j--
		}
	}

	return result
}


