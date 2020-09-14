package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/binaryheap"
)

// 41：数据流中的中位数
// 题目：如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么
// 中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，
// 那么中位数就是所有数值排序之后中间两个数的平均值。

// A：数据容器，
// 无序数组，插入时间复杂度为O(1)，找到中位数的时间复杂度为O(n)
// 有序数组，插入时，可能移动O(n)个数，找到中位数的时间复杂度为O(1)
// 排序链表，插入时，找到插入的点，时间复杂度为O(n)，如果用指针指向中间的节点，找到中位数的时间复杂度为O(1)
// 二叉搜索树可以将插入新数据的平均时间降低到O(logn)，极端情况下会退化为链表
// 为了得到中位数，可以在二叉树节点中，添加一个表示子树节点数目的字段，则可以在O(logn)时间找到中间节点
// 平衡二叉树，即AVL树，通常AVL树的平衡因子为左右子树的高度之差，可以改为左右子树的节点数目之差，在O(logn)时间内往AVL树添加一个新节点
// ，同时用O(1)时间可以得到中位数
//
// 如果数据已经排过序，中位数左边的数据都比中位数右边的小，平均数就是左边的最大值或者是右边的最小值
// 因此可以使用一个大根堆和一个小根堆，并保证大根堆和小根堆的数据量相等，并且大根堆的数据都小于小根堆的数据
// 第一次数据放大根堆，
// 第二次放小根堆，如果数据比大根堆的最大值要小，将大根堆的最大值放到小根堆，将该数字放到大根堆

type myStruct struct {
	flag bool
	maxHeap *binaryheap.Heap
	minHeap *binaryheap.Heap
}

func New() *myStruct {
	return &myStruct{
		maxHeap: binaryheap.NewWith(func(a,b interface{}) int {
			aAsserted := a.(int)
			bAsserted := b.(int)
			switch {
			case aAsserted < bAsserted:
				return 1
			case aAsserted > bAsserted:
				return -1
			default:
				return 0
			}
		}),
		// 默认是一个小根堆
		minHeap: binaryheap.NewWithIntComparator(),
		// false 放到大根堆
		// true 放小根堆
		flag: false,
	}
}

func (s *myStruct) Insert(value int) {
	// 放小根堆
	if s.flag {
		if s.maxHeap.Size() > 0 {
			temp,_ := s.maxHeap.Peek()
			if temp.(int) > value {
				s.maxHeap.Push(value)
				temp,_ = s.maxHeap.Pop()
				value = temp.(int)
			}
		}
		s.minHeap.Push(value)
	} else {
		// 如果value大于小根堆的最小值
		if s.minHeap.Size() > 0 {
			temp,_ := s.minHeap.Peek()
			if value > temp.(int) {
				s.minHeap.Push(value)
				temp,_ := s.minHeap.Pop()
				value = temp.(int)
			}
		}
		s.maxHeap.Push(value)
	}
	s.flag = !s.flag
}

func (s *myStruct) GetMedian() (float64,bool) {
	if s.maxHeap.Size() + s.minHeap.Size() == 0 {
		return -1,false
	}
	if (s.minHeap.Size() + s.maxHeap.Size()) % 2 == 1 {
		result,ok := s.maxHeap.Peek()
		return float64(result.(int)),ok
	} else {
		result1,_ := s.maxHeap.Peek()
		result2,_ := s.minHeap.Peek()
		return (float64(result1.(int)) + float64(result2.(int)))/2,true
	}
}

func main() {
	numbers := New()
	fmt.Println(numbers.GetMedian())
	numbers.Insert(5);
	Test("Test2", numbers, 5);

	numbers.Insert(2);
	Test("Test3", numbers, 3.5);

	numbers.Insert(3);
	Test("Test4", numbers, 3);

	numbers.Insert(4);
	Test("Test6", numbers, 3.5);

	numbers.Insert(1);
	Test("Test5", numbers, 3);

	numbers.Insert(6);
	Test("Test7", numbers, 3.5);

	numbers.Insert(7);
	Test("Test8", numbers, 4);

	numbers.Insert(0);
	Test("Test9", numbers, 3.5);

	numbers.Insert(8);
	Test("Test10", numbers, 4);
}

func Test(name string,s *myStruct,expetced float64) {
	fmt.Println(name)
	value,ok := s.GetMedian()
	if !ok || value != expetced {
		panic("fuck")
	}
}