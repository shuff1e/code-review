package main

import (
	"fmt"
	"strconv"
)

/*

421. 数组中两个数的最大异或值
给定一个非空数组，数组中元素为 a0, a1, a2, … , an-1，其中 0 ≤ ai < 231 。

找到 ai 和aj 最大的异或 (XOR) 运算结果，其中0 ≤ i,  j < n 。

你能在O(n)的时间解决这个问题吗？

示例:

输入: [3, 10, 5, 25, 2, 8]

输出: 28

解释: 最大的结果是 5 ^ 25 = 28.

 */

func main()  {
	arr := []int{2,4}
	fmt.Println(findMaximumXOR(arr))
	fmt.Println(findMaximumXOR2(arr))
}

func findMaximumXOR(nums []int) int {
	max := getMax(nums)
	if max == 0 {
		return 0
	}
	L := getMsb(max)

	maxXor,currXor := 0,0
	for i := L-1;i>=0;i-- {
		maxXor <<= 1
		currXor = maxXor | 1

		prefix := map[int]struct{}{}
		for _,num := range nums {
			prefix[num>>i] = struct{}{}
		}

		for k,_ := range prefix {
			if _,ok := prefix[k^currXor];ok {
				maxXor = currXor
			}
		}
	}
	return maxXor
}

func getMax(arr []int) int {
	result := arr[0]
	for i := 1;i<len(arr);i++ {
		result = Max(result,arr[i])
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func getMsb(x int) int {
	for x & (x - 1) > 0 {
		x = x & (x-1)
	}
	count := 1
	mask := 1
	for x & mask == 0 {
		mask <<= 1
		count ++
	}
	return count
}

type TrieNode struct {
	children map[byte]TrieNode
}

func findMaximumXOR2(nums []int) int {
	max := getMax(nums)
	if max == 0 {
		return 0
	}
	msb := getMsb(max)
	mask := 1 << msb

	strNums :=make([]string,len(nums))
	for i := 0;i<len(nums);i++ {
		temp := strconv.FormatInt(int64(nums[i] | mask),2)
		strNums[i] = temp[1:]
	}

	root := TrieNode{children: make(map[byte]TrieNode,0)}
	maxXor := 0
	for _,str := range strNums {
		node := root
		xorNode := root
		currXor := 0
		for i:=0;i<len(str);i++ {
			if _,ok := node.children[str[i]];ok {
				node = node.children[str[i]]
			} else {
				newNode := TrieNode{children: make(map[byte]TrieNode,0)}
				node.children[str[i]] = newNode
				node = node.children[str[i]]
			}

			toggleBit := byte('1')
			if str[i] == '1' {
				toggleBit = '0'
			}

			if _,ok := xorNode.children[toggleBit];ok {
				currXor = (currXor << 1) | 1
				xorNode = xorNode.children[toggleBit]
			} else {
				currXor = currXor << 1
				xorNode = xorNode.children[str[i]]
			}
		}
		maxXor = Max(maxXor,currXor)
	}
	return maxXor
}

