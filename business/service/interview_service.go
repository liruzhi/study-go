package service

import (
	"fmt"
	"strings"
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

//3. 无重复字符的最长子串
//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/hua-dong-chuang-kou-by-powcai/ 答案
func LengthOfLongestSubstring(str string) int {
	indexMap := make(map[byte]int)
	left, ans:= 0, 0

	for i:=0;i<len(str);i++{
		if v,ok := indexMap[str[i]];ok && left < v + 1{
			left = v + 1
		}

		indexMap[str[i]] = i

		if ans < i + 1 - left {
			ans = i + 1 - left
		}
	}

	return ans
}

//非最优
func LengthOfLongestSubstring1(str string) int{
	var window string
	var max int
	var start int

	for i:=0;i<len(str);i++ {

		if index := strings.IndexByte(window, str[i]);index != -1 {
			start = start + index + 1
		}

		window  = str[start:i+1]

		if len(window) > max {
			max = len(window)
		}
	}

	return max
}



//class Solution {
//public int lengthOfLongestSubstring(String s) {
//if (s.length()==0) return 0;
//HashMap<Character, Integer> map = new HashMap<Character, Integer>();
//int max = 0;
//int left = 0;
//for(int i = 0; i < s.length(); i ++){
//if(map.containsKey(s.charAt(i))){
//left = Math.max(left,map.get(s.charAt(i)) + 1);
//}
//map.put(s.charAt(i),i);
//max = Math.max(max,i-left+1);
//}
//return max;
//
//}
//}
//
//作者：powcai
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/hua-dong-chuang-kou-by-powcai/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

//https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
//429. N 叉树的层序遍历
type Node struct {
     Val int
     Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := make([][]int,0)

	if root == nil {
		return res
	}

	return levelOrderC([]*Node{root}, res)
}

func levelOrderC(nodes []*Node, res [][]int) [][]int{
	if len(nodes) == 0 {
		return res
	}

	curLayerValue := make([]int, 0)
	nextLayerNode := make([]*Node, 0)

	for i := range nodes {
		curLayerValue = append(curLayerValue, nodes[i].Val)
		if len(nodes[i].Children) > 0 {
			nextLayerNode = append(nextLayerNode, nodes[i].Children...)
		}
	}

	res = append(res,curLayerValue)

	return levelOrderC(nextLayerNode, res)
}

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

//判断一个树是否为平衡二叉树
func IsBalanceTree(tree *Tree) bool {
	if tree == nil {
		return true
	}

	isBalance, _ := TreeLevel(tree)
	return isBalance
}

func TreeLevel(tree *Tree) (bool, int) {

	if tree == nil {
		return true, 0
	}
	leftIsBalance, leftLevel := TreeLevel(tree.Left)
	rightIsBalance, rightLevel := TreeLevel(tree.Right)

	if leftIsBalance && rightIsBalance && abs(leftLevel-rightLevel) < 2 {

		return true, max(leftLevel+1, rightLevel+1)
	}

	return false, 0
}

func abs(parmas int) int {
	if parmas >= 0 {
		return parmas
	}
	return 0 - parmas
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
