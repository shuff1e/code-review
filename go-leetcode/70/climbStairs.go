package main

import "fmt"

/*
70. 爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
 */

// f(n) = f(n-1) + f(n-2)

// f(n-1),f(n-2) | 1 1 | = f(n),f(n-1)
//               | 1 0 |

// f(2),f(1) | 1 1 |^(n-2) = f(n),f(n-1)
//           | 1 0 |

// multiply

// pow

// // 矩阵乘法
//func multiMatrix(m [][]int,n [][]int) [][]int {
//	result := make([][]int,len(m))
//	for i := 0;i<len(result);i++ {
//		result[i] = make([]int,len(n[0]))
//	}
//
//	for i := 0;i<len(m);i++ {
//		for j := 0;j<len(n[0]);j++ {
//			// len(m[0].length) == len(n)
//			for k := 0;k < len(n);k++ {
//				result[i][j] += m[i][k] * n[k][j]
//			}
//		}
//	}
//
//	return result
//}
//
// 矩阵的n次方
// 例如10的13次方，10^13 = 10^8 * 10^4 * 10^1
// 13对应二进制的1101
//func matrixPower(matrix [][]int, p int) [][]int {
//	result := make([][]int,len(matrix))
//	for i := 0;i<len(matrix);i++ {
//		result[i] = make([]int,len(matrix[0]))
//	}
//	// 首先result初始化为单位矩阵，相当于1
//	for i := 0;i<len(matrix);i++ {
//		for j := 0;j<len(matrix[0]);j++ {
//			result[i][i] = 1
//		}
//	}
//	for ;p>0;p>>=1 {
//		if p & 1 == 1 {
//			result = multiMatrix(result,matrix)
//		}
//		matrix = multiMatrix(matrix,matrix)
//	}
//	return result
//}

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	prepre := 1
	pre := 2
	for i := 2;i<n;i++ {
		cur := pre + prepre
		prepre = pre
		pre = cur
	}
	return pre
}