package service

import (
	"fmt"
	"time"
)

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

//两个协程交替打印奇偶数，一个channel
func DumpNum() {
	ch := make(chan int)

	go func() {
		for i:=1;i<=100;i++ {
			ch<-1
			if i%2 == 1 {
				fmt.Println("g1---", i)
			}
		}
	}()

	go func() {
		for i:=1;i<=100;i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println("g2---", i)
			}
		}
	}()

	time.Sleep(time.Second)
}



//https://blog.csdn.net/Lyon_Nee/article/details/119171977
type BinaryTree struct {
	Value     int         `json:"value,omitempty"`
	LeftNode  *BinaryTree `json:"left_node,omitempty"`
	RightNode *BinaryTree `json:"right_node,omitempty"`
}

//根节点
var rootNode *BinaryTree = &BinaryTree{
	Value: 1,
	LeftNode: &BinaryTree{
		Value: 2,
		LeftNode: &BinaryTree{
			Value: 4,
			RightNode: &BinaryTree{
				Value:     6,
				LeftNode:  &BinaryTree{Value: 7},
				RightNode: &BinaryTree{Value: 8},
			},
		},
	},
	RightNode: &BinaryTree{
		Value:     3,
		RightNode: &BinaryTree{Value: 5},
	},
}

func DumpLevelTree() {
	LevelTraversal(rootNode)
}

//按层遍历树
func LevelTraversal(tree *BinaryTree) {
	if tree == nil {
		return
	}
	LevelTraversalC([]*BinaryTree{tree})
}

func LevelTraversalC(treeSlice []*BinaryTree) {
	if len(treeSlice) == 0 {
		return
	}

	nextLevel := make([]*BinaryTree, 0)
	for _,v := range treeSlice {
		fmt.Println(v.Value)

		if v.LeftNode != nil {
			nextLevel = append(nextLevel, v.LeftNode)
		}

		if v.RightNode != nil {
			nextLevel = append(nextLevel, v.RightNode)
		}
	}
	LevelTraversalC(nextLevel)
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
