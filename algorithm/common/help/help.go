package help

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Generates a slice of size, size filled with random numbers
func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func Swap(a []int,left ,right int) {
	temp := a[left]
	a[left] = a[right]
	a[right] = temp
}

type MyStack struct {
	lock sync.Mutex
	s []interface{}
}

func NewMyStack() *MyStack {
	return &MyStack{
		sync.Mutex{},
		make([]interface{},0),
	}
}

func (s *MyStack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.s = append(s.s,v)
}

func (s *MyStack) Pop() (interface{},error) {
	s.lock.Lock()
	defer s.lock.Unlock()


	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *MyStack) Peek() (interface{},error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	return s.s[l-1],nil
}

func (s *MyStack) Length() int {
	return len(s.s)
}

type Stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s []int
}

func NewStack() *Stack {
	return &Stack {sync.Mutex{}, make([]int,0), }
}

func (s *Stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()


	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *Stack) Peek() (int,error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	return s.s[l-1],nil
}

func (s *Stack) Length() int {
	return len(s.s)
}

type Node struct {
	Value int
	Next *Node
}

func GenerateNList(n int) *Node {
	slice := GenerateSlice(n)
	head := &Node{}
	temp := head
	for _,v := range slice {
		temp.Next = &Node{Value: v}
		temp = temp.Next
	}
	return head.Next
}

func PrintList(head *Node) {
	for head != nil {
		fmt.Print(head.Value)
		fmt.Print(" ")
		head=head.Next
	}
	fmt.Println("")
}

type DoubleNode struct {
	Value int
	Next *DoubleNode
	Prev *DoubleNode
}

type MyQueue struct {
	stackPush *MyStack
	stackPop *MyStack
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		NewMyStack(),
		NewMyStack(),
	}
}

func (q *MyQueue) Add(v interface{}) {
	q.stackPush.Push(v)
}

func (q *MyQueue) Poll() (interface{},error) {
	if q.stackPop.Length() == 0 {
		for q.stackPush.Length() != 0 {
			v,_ := q.stackPush.Pop()
			q.stackPop.Push(v)
		}
	}
	return q.stackPop.Pop()
}

func (q *MyQueue) Peek() (interface{},error) {
	if q.stackPop.Length() == 0 {
		for q.stackPush.Length() != 0 {
			v,_ := q.stackPush.Pop()
			q.stackPop.Push(v)
		}
	}
	return q.stackPop.Peek()
}

func (q *MyQueue) Length() int {
	return q.stackPush.Length() + q.stackPop.Length()
}

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

type TreeNode2 struct {
	Value int
	Left *TreeNode2
	Right *TreeNode2
	Num int
}

func Build(pre []int,mid []int) *TreeNode{
	if len(pre) == 0 {
		return nil
	}
	value := pre[0]
	index := findIndex(value,mid) //0
	leftLen := index //0
	rightLen := len(mid)-1-index//0

	left := build(pre[1:1+leftLen],mid[0:leftLen])
	right := build(pre[len(pre)-rightLen:len(pre)],mid[index+1:])
	root := &TreeNode{Value: value,Left: left,Right: right}
	return root
}

func build(pre []int,mid []int) *TreeNode{
	if len(pre) == 0 {
		return nil
	}
	value := pre[0]
	index := findIndex(value,mid) //0
	leftLen := index //0
	rightLen := len(mid)-1-index//0

	left := build(pre[1:1+leftLen],mid[0:leftLen])
	right := build(pre[len(pre)-rightLen:len(pre)],mid[index+1:])
	root := &TreeNode{Value: value,Left: left,Right: right}
	return root
}
func findIndex(value int,mid []int) int {
	for index,v := range mid {
		if value == v {
			return index
		}
	}
	return -1
}

func PreOrder(node *TreeNode) {
	root := node
	stack := NewMyStack()
	for root != nil || stack.Length() > 0 {
		// 入栈的时候，转移为左子节点
		for root != nil {
			fmt.Print(root.Value)
			fmt.Print(" ")
			stack.Push(root)
			root = root.Left
		}
		if stack.Length() > 0 {
			temp,_ := stack.Pop()
			root = temp.(*TreeNode)
			//出栈的时候转为右子节点
			root = root.Right
		}
	}
}

func MidOrder(node *TreeNode) {
	root := node
	stack := NewMyStack()
	for root != nil || stack.Length() > 0 {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}

		temp,_ := stack.Pop()
		root = temp.(*TreeNode)
		fmt.Print(root.Value)
		fmt.Print(" ")
		root = root.Right

	}
}

func DoubleQueueOrder(node *TreeNode) {
	root := node
	stackData := NewMyStack()
	stackPrint := NewMyStack()
	for root != nil || stackData.Length() > 0 {
		for root != nil {
			stackData.Push(root)
			stackPrint.Push(root)
			root = root.Right
		}
		if stackData.Length() > 0 {
			temp,_ := stackData.Pop()
			root = temp.(*TreeNode)
			root = root.Left
		}
	}

	for stackPrint.Length() > 0 {
		temp,_ := stackPrint.Pop()
		fmt.Print(temp.(*TreeNode).Value)
		fmt.Print(" ")
	}
}

func Cengcibianli(node *TreeNode) {
	queue := NewMyQueue()
	queue.Add(node)
	for queue.Length() > 0 {
		temp,_ := queue.Poll()
		node := temp.(*TreeNode)
		fmt.Print(node.Value)
		fmt.Print(" ")

		if node.Left != nil {
			queue.Add(node.Left)
		}
		if node.Right != nil {
			queue.Add(node.Right)
		}
	}
}

func Abs(x int) int {
	if x >0 {
		return x
	}
	return -x
}