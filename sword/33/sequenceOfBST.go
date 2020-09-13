package main

import "fmt"

// 33：二叉搜索树的后序遍历序列
// 题目：输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。
// 如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。

//            10
//         /      \
//        6        14
//       /\        /\
//      4  8     12  16

// 4, 8, 6, 12, 16, 14, 10是上述二叉搜索树的后序遍历序列
// A：最后一个节点是根节点，左子树都小于根节点

// 右子树开始的位置
func findRightStart(arr []int) int {
	for i := 0;i<len(arr)-1;i++ {
		if arr[i] > arr[len(arr)-1] {
			return i
		}
	}
	return len(arr) - 1
}

func checkRightValid(arr []int,rightStart int) bool {
	for i := rightStart;i<len(arr)-1;i++ {
		if arr[i] <= arr[len(arr)-1] {
			return false
		}
	}
	return true
}

func isValid(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	rightStart := findRightStart(arr)
	if !checkRightValid(arr,rightStart) {
		return false
	}
	if rightStart == 0 || rightStart == len(arr) - 1 {
		return isValid(arr[0:len(arr)-1])
	} else {
		return isValid(arr[0:rightStart]) && isValid(arr[rightStart:len(arr)-1])
	}
}

func main() {
	//            10
	//         /      \
	//        6        14
	//       /\        /\
	//      4  8     12  16
	Test("Test1",[]int{4, 8, 6, 12, 16, 14, 10},true)
	//           5
	//          / \
	//         4   7
	//            /
	//           6
	Test("Test2",[]int{4, 6, 7, 5},true)
	//               5
	//              /
	//             4
	//            /
	//           3
	//          /
	//         2
	//        /
	//       1
	Test("Test3",[]int{1, 2, 3, 4, 5},true)
	// 1
	//  \
	//   2
	//    \
	//     3
	//      \
	//       4
	//        \
	//         5
	Test("Test4",[]int{5, 4, 3, 2, 1},true)
	Test("name5",[]int{5},true)
	Test("name6",[]int{7, 4, 6, 5},false)
	Test("name7",[]int{4, 6, 12, 8, 16, 14, 10},false)
}

func Test(name string,arr []int,expected bool) {
	fmt.Println(name,isValid(arr) == expected)
}