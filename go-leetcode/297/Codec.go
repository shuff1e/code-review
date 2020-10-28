package main

import (
    "fmt"
    "strconv"
    "strings"
)

/*

297. 二叉树的序列化与反序列化
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例:

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。
你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。

 */

func main() {
    //            8
    //        6      10
    //       5 7    9  11
    node1 := &TreeNode{Val: 8}
    node2 := &TreeNode{Val: 6}
    node3 := &TreeNode{Val: 10}
    node4 := (*TreeNode)(nil)
    node5 := (*TreeNode)(nil)
    node6 := &TreeNode{Val: 9}
    node7 := &TreeNode{Val: 11}
    connectNodes(node1,node2,node3)
    connectNodes(node2,node4,node5)
    connectNodes(node3,node6,node7)
    c := Constructor()
    str := c.serialize(node1)
    fmt.Println(str)
    result := c.deserialize(str)
    fmt.Println(result.Val)
}

func connectNodes(p,left,right *TreeNode) {
    if left != nil {
        p.Left = left
    }
    if right != nil {
        p.Right = right
    }
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Codec struct {

}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
	    return "nil"
    }
    result := []string{}
    stack := []*TreeNode{}
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack,root)
            result = append(result,strconv.Itoa(root.Val))
            root = root.Left
        }
        if root == nil {
            result = append(result,"nil")
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        root = root.Right
    }
    return strings.Join(result,",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    if data == "nil" {
        return nil
    }

	arr := strings.Split(data,",")
	val,_ := strconv.Atoi(arr[0])
	result := &TreeNode{Val: val}
	help(result,arr,1)
	return result
}

func help(node *TreeNode,arr []string,index int) int {
    if index  >= len(arr) {
        return index
    }

    if arr[index] == "nil" {
        if index + 1 < len(arr) && arr[index+1] == "nil" {
            return index + 2
        } else if index + 1 < len(arr) {
            val,_ := strconv.Atoi(arr[index+1])
            node.Right = &TreeNode{Val: val}
            return help(node.Right,arr,index+2)
        }
    } else {
        val,_ := strconv.Atoi(arr[index])
        node.Left = &TreeNode{Val: val}
        leftIndex := help(node.Left,arr,index+1)
        if leftIndex < len(arr) && arr[leftIndex] == "nil" {
            return leftIndex + 1
        } else if leftIndex < len(arr) {
            val,_ := strconv.Atoi(arr[leftIndex])
            node.Right = &TreeNode{Val: val}
            rightIndex := help(node.Right,arr,leftIndex+1)
            return rightIndex
        }
    }
    return index
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */