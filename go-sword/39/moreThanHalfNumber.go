package main

// 39：数组中出现次数超过一半的数字
// 题目：数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。例
// 如输入一个长度为9的数组{1, 2, 3, 2, 2, 2, 5, 4, 2}。由于数字2在数组中
// 出现了5次，超过数组长度的一半，因此输出2。

// A：数组中出现次数超过一半
// 如果把这个数组排序，那么排序之后的数组中间的元素，一定是出现此事超过数组长度一半的数字
// 也就是说，该数字就是统计学上的中位数
// 即长度为n的数组中，索引为n/2的元素
// 转换为求长度为n的数组中第k大的数字

// 这里只是找到中位数
func doIt(arr []int,start,end int) int {
	mark := partition(arr,start,end)
	if mark == len(arr) / 2 {
		return arr[mark]
	} else if mark > len(arr)/2 {
		return doIt(arr,start,mark-1)
	} else {
		return doIt(arr,mark+1,end)
	}
}

// 还要检查中位数的出现次数
func getOccurimes(arr []int,k int) int {
	count := 0
	for _,v := range arr {
		if v == k {
			count ++
		}
	}
	return count
}

func moreThanHalf(arr []int) (result int,invalid bool) {
	if len(arr) == 0 {
		return 0,true
	}
	result = doIt(arr,0,len(arr)-1)
	times := getOccurimes(arr,result)
	if times <= len(arr)/ 2 {
		return 0,true
	}
	return result,false
}

func partition(arr []int,start,end int) int {
	if start >= end {
		return start
	}
	mark := start
	for i := start;i<=end;i++ {
		if arr[i] <= arr[start] {
			swap(arr,i,mark)
			mark ++
		}
	}
	swap(arr,start,mark-1)
	return mark -1
}

func swap(arr []int,i,j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func main() {
	Test("Test1",[]int{1, 2, 3, 2, 2, 2, 5, 4, 2},2,false)
	Test("Test2",[]int{1, 2, 3, 2, 4, 2, 5, 2, 3},0,true)
	Test("Test3",[]int{2, 2, 2, 2, 2, 1, 3, 4, 5},2,false)
	Test("Test4",[]int{1, 3, 4, 5, 2, 2, 2, 2, 2},2,false)
	Test("Test5",[]int{1},1,false)
	Test("Test6",[]int{},0,true)
	Test("Test7",[]int{1,1,1,2,2,2},0,true)
}

func Test(name string,arr []int,expected int,valid bool) {
	result,flag := moreThanHalf(arr)
	if result == expected && flag == valid {
	} else {
		panic("fuck")
	}
}

// 在遍历数组时，保存两个值
// 一个是数组中的数字，一个是改数字的次数
// 如果下一个数字和该数字相同，次数+1
// 如果下一个数字和该数字不同，次数-1
// 如果次数为0，保存下一个数字，并将次数设置为1

//
func moreThanHalf2(arr []int) (result int,invalid bool) {
	count := 0
	prev := 0
	for _,v := range arr {
		if count == 0 {
			count += 1
			prev = v
		} else if v == prev {
			 count ++
		} else {
			count --
		}
	}
	if count == 0 {
		return 0,true
	}
	times := getOccurimes(arr,prev)
	if times*2 <= len(arr) {
		return 0,true
	}
	return prev,false
}