package main

import (
	"fmt"
	"unsafe"
)

/*

402. 移掉K位数字
给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。

注意:

num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。
示例 1 :

输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
示例 2 :

输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 :

输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。

 */

/*

给定一个数字序列 [D_1D_2D_3…D_n]，如果数字 D_2小于其左邻居 D_1，则我们应该删除左邻居（D_1），以获得最小结果。

算法：

上述的规则使得我们通过一个接一个的删除数字，逐步的接近最优解。

这个问题可以用贪心算法来解决。上述规则阐明了我们如何接近最终答案的基本逻辑。一旦我们从序列中删除一个数字，剩下的数字就形成了一个新的问题，我们可以继续使用这个规则。

我们会注意到，在某些情况下，规则对任意数字都不适用，即单调递增序列。在这种情况下，我们只需要删除末尾的数字来获得最小数。

我们可以利用栈来实现上述算法，存储当前迭代数字之前的数字。

对于每个数字，如果该数字小于栈顶部，即该数字的左邻居，则弹出堆栈，即删除左邻居。否则，我们把数字推到栈上。
我们重复上述步骤（1），直到任何条件不再适用，例如堆栈为空（不再保留数字）。或者我们已经删除了 k 位数字。

 */

func main() {
	str := "10200"
	fmt.Println(removeKdigits(str,2))
}

func removeKdigits(num string, k int) string {
	stack := []byte{}
	count := 0
	for i:=0;i<len(num);i++ {
		for len(stack) > 0 && count < k && stack[len(stack)-1] > num[i] {
			count ++
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,num[i])
	}

	for ;count < k;count++{
		stack = stack[:len(stack)-1]
	}

	index := 0
	for ;index <len(stack);index ++ {
		if stack[index] != '0' {
			break
		}
	}
	if index == len(stack) {
		return "0"
	}
	return String(stack[index:])
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}