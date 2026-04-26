package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

//上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//示例:
//
//输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6

// A：找到一个非0但是右边是0的数，找到一个左边是0但是自己非0的数字
// 计算完这一层之后，全都减1，再次计算

func trap(height []int) int {
	return getWater(height)
}

func getWater(arr []int) int {
	if len(arr) <= 2 {
		return 0
	}
	sum := 0
	base := 0
	max := findMax(arr)
	for base < max {
		index := 0
		for {
			left := getLeft(arr,index,base)
			if left == -1 {
				break
			}
			right := getRight(arr,left + 2,base)
			if right == -1 {
				break
			}
			sum += right - left - 1
			index = right
		}
		base ++
	}
	return sum
}

func getLeft(arr []int,start int,base int) int {
	for i := start;i<len(arr)-1;i++ {
		if arr[i] - base > 0 && arr[i + 1] - base <= 0 {
			return i
		}
	}
	return -1
}

func getRight(arr []int,start int,base int) int {
	for i := start ;i<len(arr);i++ {
		if arr[i] - base > 0 && arr[i - 1] - base <= 0 {
			return i
		}
	}
	return -1
}

func findMax(arr []int) int {
	max := arr[0]
	for i := 0;i<len(arr);i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	arr := []int{0,7,1,4,6}
	fmt.Println(getWater3(arr))
	fmt.Println(getWaterBetter(arr))
	arr = []int{0,1,1,1,1,0,2,0,3,7,1,0,2,4,5}
	fmt.Println(getWater3(arr))
	fmt.Println(getWaterBetter(arr))
	arr = []int{}
	fmt.Println(getWater3(arr))
	fmt.Println(getWaterBetter(arr))
}

//单调栈
//单调栈分为单调递增栈和单调递减栈
//11. 单调递增栈即栈内元素保持单调递增的栈
//12. 同理单调递减栈即栈内元素保持单调递减的栈
//
//操作规则（下面都以单调递增栈为例）
//21. 如果新的元素比栈顶元素大，就入栈
//22. 如果新的元素较小，那就一直把栈内元素弹出来，直到栈顶比新元素小
//
//加入这样一个规则之后，会有什么效果
//31. 栈内的元素是递增的
//32. 当元素出栈时，说明这个新元素是出栈元素向后找第一个比其小的元素
//
//举个例子，配合下图，现在索引在 6 ，栈里是 1 5 6 。
//接下来新元素是 2 ，那么 6 需要出栈。
//当 6 出栈时，右边 2 代表是 6 右边第一个比 6 小的元素。
//
//当元素出栈后，说明新栈顶元素是出栈元素向前找第一个比其小的元素
//当 6 出栈时，5 成为新的栈顶，那么 5 就是 6 左边第一个比 6 小的元素。
//

// A：如果能遍历元素的时候，找到前面比该元素小的元素，
// 就能一次遍历，找到所有的雨滴

//思路
//使用单调栈，【单调栈入门】
//
//单调递减栈
//
//理解题目，参考图解，注意题目的性质，当后面的柱子高度比前面的低时，是无法接雨水的
//当找到一根比前面高的柱子，就可以计算接到的雨水
//所以使用单调递减栈
//对更低的柱子入栈
//
//更低的柱子以为这后面如果能找到高柱子，这里就能接到雨水，所以入栈把它保存起来
//平地相当于高度 0 的柱子，没有什么特别影响
//当出现高于栈顶的柱子时
//41. 说明可以对前面的柱子结算了
//42. 计算已经到手的雨水，然后出栈前面更低的柱子
//
//计算雨水的时候需要注意的是
//
//雨水区域的右边 r 指的自然是当前索引 i
//底部是栈顶 st.top() ，因为遇到了更高的右边，所以它即将出栈，使用 cur 来记录它，并让它出栈
//左边 l 就是新的栈顶 st.top()
//雨水的区域全部确定了，水坑的高度就是左右两边更低的一边减去底部，宽度是在左右中间
//使用乘法即可计算面积

// 单调栈，单调递减，如果arr[i]比栈中的数大，出栈，然后结算
// 模板
/*
stack<int> st;
for(int i = 0; i < nums.size(); i++)
{
	while(!st.empty() && st.top() < nums[i])
	{
		st.pop();
	}
	st.push(nums[i]);
}
 */

func getWaterBetter(arr []int) int {
	res := 0
	stack := linkedliststack.New()
	for i := 0;i<len(arr);i++ {
		for {
			if stack.Empty() {
				break
			}
			temp,_ := stack.Peek()
			if arr[temp.(int)] >= arr[i] {
				break
			}
			stack.Pop()

			if stack.Empty() {
				break
			}
			temp2,_ := stack.Peek()
			height := Min(arr[temp2.(int)],arr[i]) - arr[temp.(int)]
			res += height * (i - 1 - temp2.(int))
		}
		stack.Push(i)
	}
	return res
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

// A：实际上是计算每个柱子的储水
// 就是柱子左边的最大值和右边的最大值，减去当前值
// 如果能记录这个值，就可以不需要每次都去查找这个值
func getWater1(arr []int) int {
	if len(arr) == 0{
		return 0
	}
	leftArray := make([]int,len(arr))
	leftArray[0] = arr[0]
	leftMax := arr[0]

	for i := 1;i<len(arr);i++ {
		leftMax = Max(leftMax,arr[i])
		leftArray[i] = leftMax
	}

	rightArray := make([]int,len(arr))
	rightMax := arr[len(arr)-1]
	rightArray[len(arr)-1] = rightMax

	for i := len(arr) - 2;i>=0;i-- {
		rightMax = Max(rightMax,arr[i])
		rightArray[i] = rightMax
	}

	sum := 0
	for i := 0;i<len(arr);i++ {
		sum += Min(leftArray[i],rightArray[i]) - arr[i]
	}
	return sum
}

// A：取决于left_max和right_max
// 如果left_max < right_max（可能有更大的），说明从左到右的方向上，应该为left_max-arr[i]
// 如果left_max（可能有更大的）> right_max，说明从右到左的方向上，应该为right_max-arr[i]

func getWater3(arr []int) int {
	if len(arr) <= 2 {
		return 0
	}
	left_max := arr[0]
	right_max := arr[len(arr)-1]
	left := 1
	right := len(arr) - 2
	result := 0
	for left <= right {
		if left_max < right_max {
			left_max = Max(arr[left],left_max)
			result += left_max - arr[left]
			left ++
		} else {
			right_max = Max(arr[right],right_max)
			result += right_max - arr[right]
			right --
		}
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
