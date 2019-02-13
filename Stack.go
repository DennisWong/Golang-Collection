package main

import (
	"fmt"
	"os"
)

type Stack interface {
	Push(node interface{})
	Pop() interface{}
	Top() interface{}
	Empty() bool
	Size() int
}

type BTStack struct {
	Data     []TreeNode
	ElmCount int
	SizeVal  int
}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func (s *BTStack) Push(node TreeNode) {
	if s.ElmCount >= s.SizeVal {
		fmt.Println("stack is full")
		os.Exit(-1)
	}
	fmt.Println("push", node.Val)
	s.Data = append(s.Data, node)
	s.ElmCount++
}

func (s *BTStack) Size() int {
	return s.SizeVal
}

func (s *BTStack) Pop() TreeNode {
	if s.Empty() == true {
		fmt.Println("stack is empty")
		os.Exit(-1)
	}
	node := s.Data[s.ElmCount-1]
	s.Data = s.Data[0 : s.ElmCount-1]
	fmt.Println("pop", node.Val)

	s.ElmCount--
	return node
}

func (s *BTStack) Top() TreeNode {
	return s.Data[s.ElmCount-1]
}

func (s *BTStack) Empty() bool {
	return s.ElmCount == 0
}

//func main() {
//	data1 := []TreeNode{}
//	stack := &BTStack{
//		Data:     data1,
//		ElmCount: 0,
//		SizeVal:  10,
//	}
//
//	var root *TreeNode = new(TreeNode)
//	root.Val = 1
//	var n1 *TreeNode = new(TreeNode)
//	n1.Val = 2
//	var n2 *TreeNode = new(TreeNode)
//	n2.Val = 3
//
//	root.Left = nil
//	root.Right = n1
//
//	n1.Left = n2
//	n1.Right = nil
//
//	n2.Left = nil
//	n2.Right = nil
//	stack.Push(*root)
//	stack.Push(*n1)
//	stack.Push(*n2)
//
//	fmt.Println("count", stack.ElmCount)
//
//	fmt.Println("data", stack.Data)
//
//	x1 := stack.Pop()
//	fmt.Println(x1.Val)
//	x2 := stack.Pop()
//	fmt.Println(x2.Val)
//	x3 := stack.Pop()
//	fmt.Println(x3.Val)
//
//	fmt.Println(stack.Empty())
//}
