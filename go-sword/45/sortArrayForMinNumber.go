package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 45：把数组排成最小的数
// 题目：输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼
// 接出的所有数字中最小的一个。例如输入数组{3, 32, 321}，则打印出这3个数
// 字能排成的最小数字321323。

// A：取出最高位最小的
// 如果最高位都一样，看下一位，没有下一位的，使用最高位
// 取下一位最小的，
// 这样得到了两个数之间比较的法则
// 然后排序就可以了

// ab如果小于ba，则a小于b

//type Interface interface {
	// Len is the number of elements in the collection.
	//Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	//Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	//Swap(i, j int)
//}

// 32 323232

// 322232
// 323222

type myData []int

func (m myData) Len() int {
	return len(m)
}

func (m myData) Less(i,j int) bool {
	mij := strconv.Itoa(m[i]) + strconv.Itoa(m[j])
	mji := strconv.Itoa(m[j]) + strconv.Itoa(m[i])
	return strings.Compare(mij,mji) < 0
}

func (m myData) Swap(i,j int) {
	temp := m[i]
	m[i] = m[j]
	m[j] = temp
}

func (m myData) String() string {
	result := ""
	for _,v := range m {
		result += strconv.Itoa(v)
	}
	return result
}

func getMinNumber(arr []int) string {
	data := myData(arr)
	sort.Sort(data)
	return data.String()
}

func main() {
	Test("Test1",[]int{3, 5, 1, 4, 2},"12345")
	Test("Test2",[]int{3, 32, 321},"321323")
	Test("Test3",[]int{3, 323, 32123},"321233233")
	Test("Test4",[]int{1, 11, 111},"111111")
	Test("Test5",[]int{321},"321")
	Test("Test6",nil,"")
}

func Test(name string,arr []int,expected string) {
	fmt.Println(name)
	fmt.Println(getMinNumber(arr) == expected)
}

