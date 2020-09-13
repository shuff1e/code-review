package main

import (
	"fmt"
	"unsafe"
)

// 38：字符串的排列
// 题目：输入一个字符串，打印出该字符串中字符的所有排列。例如输入字符串abc，
// 则打印出由字符a、b、c所能排列出来的所有字符串abc、acb、bac、bca、cab和cba。

// 改成字符串中任选k个的全排列

// A：求全排列
// 第一次选一个数，有n种选择，第二次就只有n-1中选择，选k次
// 另外
// 求组合
// 每个字符都有选与不选两种选择

// int comm( int n ,int k)
// {
// if(k > n)
// return 0;
// else if( k == n || k == 0)
// return 1;
// else
// return comm(n-1,k) + comm(n-1 ,k-1);
// }

func doItBetter(str string,k int) {
	if len(str) < k {
		return
	}
	if k == 0 {
		return
	}
	permutationBetter(([]byte)(str),k,0)
}

func permutationBetter(str []byte,k,index int) {
	if index == k {
		fmt.Println(String(str[0:k]))
		return
	}
	for i := index;i<len(str);i++ {
		swap(str,index,i)
		permutationBetter(str,k,index+1)
		swap(str,index,i)
	}
}

func swap(str []byte,i,j int) {
	temp := str[i]
	str[i] = str[j]
	str[j] = temp

}






// 这种方法不好
func doIt(str string,k int) {
	if len(str) < k {
		return
	}
	if k == 0 {
		return
	}
	arr := make([]byte,0)
	chosen := make([]bool,len(str))

	permutation(str,k,0,chosen,arr)
}

func permutation(str string,k int,level int,chosen []bool,arr []byte) {

	if level == k {
		printArr(arr)
		return
	}

	for i := 0;i<len(str);i++ {
		// 是否已经选择过
		if !chosen[i] {
			// 选择
			arr = append(arr,str[i])
			chosen[i] = true
			permutation(str,k,level+1,chosen,arr)
			chosen[i] = false
			arr = arr[0:len(arr)-1]
		}
	}
}

func printArr(arr []byte) {
	fmt.Println(String(arr))
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}

func getRestLength(arr []bool) int {
	count := 0
	for _,v := range arr {
		if !v {
			count ++
		}
	}
	return count
}

func main() {
	Test("abc",2)
	Test("",0)
	Test("a",1)
	Test("ab",2)
}

func Test(str string,k int) {
	doItBetter(str,k)
}