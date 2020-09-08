package main

import "fmt"

// 10：斐波那契数列
// 题目：写一个函数，输入n，求斐波那契（Fibonacci）数列的第n项。

// F(1)=1，F(2)=1, F(n)=F(n - 1)+F(n - 2)（n ≥ 3，n ∈ N*）
// 1、1、2、3、5、8、13、21、34、

// 递归
func f1(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		return f1(n-1) + f1(n-2)
	}
}

// 迭代
func f2(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	prev,prevprev,cur := 1,1,-1

	for i := 3;i<=n;i++ {
		cur = prev + prevprev
		prevprev = prev
		prev = cur
	}
	return cur
}

// 矩阵
// f(n) = f(n-1) + f(n-2)
// [f(n-2),f(n-1)] * |0 1｜ = [f(n-1),f(n)]
// 					 ｜1 1｜
// f(1) = 1,f(2) = 1
// [1,1]*|0 1|^(n-2) = [f(n-1),f(n)]
// 		 |1 1|

func f3(n int) int {
	if n == 1 || n ==2 {
		return 1
	}
	matrix := [][]int{{0,1},{1,1}}
	matrix = matrixPower(matrix,n-2)
	return matrix[0][1] + matrix[1][1]
}

// 矩阵乘法
func multiMatrix(m [][]int,n [][]int) [][]int {
	result := make([][]int,len(m))
	for i := 0;i<len(result);i++ {
		result[i] = make([]int,len(n[0]))
	}

	for i := 0;i<len(m);i++ {
		for j := 0;j<len(n[0]);j++ {
			// len(m[0].length) == len(n)
			for k := 0;k < len(n);k++ {
				result[i][j] += m[i][k] * n[k][j]
			}
		}
	}

	return result
}

// 矩阵的n次方
// 例如10的13次方，10^13 = 10^8 * 10^4 * 10^1
// 13对应二进制的1101
func matrixPower(matrix [][]int, p int) [][]int {
	result := make([][]int,len(matrix))
	for i := 0;i<len(matrix);i++ {
		result[i] = make([]int,len(matrix[0]))
	}
	// 首先result初始化为单位矩阵，相当于1
	for i := 0;i<len(matrix);i++ {
		for j := 0;j<len(matrix[0]);j++ {
			result[i][i] = 1
		}
	}
	for ;p>0;p>>=1 {
		if p & 1 == 1 {
			result = multiMatrix(result,matrix)
		}
		matrix = multiMatrix(matrix,matrix)
	}
	return result
}


func main() {
	fmt.Println(f1(8))
	fmt.Println(f2(8))
	fmt.Println(f3(8))
	m := [][]int{{1,2},{3,4}}
	n := [][]int{{2,3,4,5},{6,7,8,9}}
	result := multiMatrix(m,n)
	for i := 0;i<len(result);i++ {
		fmt.Printf("%#v\n",result[i])
	}
	result = matrixPower(m,3)
	for i := 0;i<len(result);i++ {
		fmt.Printf("%#v\n",result[i])
	}
}