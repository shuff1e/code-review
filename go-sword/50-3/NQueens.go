package main

import (
	"fmt"
	"math"
)

// Q：在N*N的棋盘上，要摆N个皇后，要求任何两个皇后不同行，不同列，也不在同一条斜线上
// 给定一个整数n，返回n皇后的摆法有多少种

// A：在一行摆放一个皇后之后，位置为(i,j)
// 下一行皇后的位置为x,y，要满足x != i && y != j && |x-i| != |y-j|

func GetNumber(n int) int {
	record := make([]int,n)
	for i := 0;i<n;i++ {
		record[i] = -1
	}
	return getNumber(n,0,record)
}

func getNumber(n ,row int,record []int) int {
	if row == n {
		return 1
	}
	sum := 0
	for j := 0;j<n;j++ {
		if checkValid(row,j,record) {
			record[row] = j
			sum += getNumber(n,row +1 ,record)
			record[row] = -1
		}
	}
	return sum
}

func checkValid(row,j int,record []int) bool {
	// 之前的点，位置是(i,record[i])
	for i := 0;i< len(record);i++ {
		if record[i] == j {
			return false
		}
		if record[i] >=0 {
			if int(math.Abs(float64(row-i))) == int(math.Abs(float64(j-record[i]))) {
				return false
			}
		}
	}
	return true
}

func main() {
	Test(1,1)
	Test(2,0)
	Test(3,0)
	Test(8,92)
	Test(9,352)
	Test(10,724)
	Test(11,2680)
}

func Test(n ,expected int) {
	fmt.Println(GetNumber2(n),expected)
	if GetNumber2(n) != expected {
		panic("fuck")
	}
}

func GetNumber2(n int) int {
	upperLimit := (1<<n) -1
	return getNumber2(upperLimit,0,0,0)
}

func getNumber2(upperLimit,colLimit,leftLimit,rightLimit int) int {
	if colLimit == upperLimit {
		return 1
	}
	pos := 0
	rightMost := 0
	count := 0
	// 按位取反
	pos = upperLimit&(^(colLimit | leftLimit | rightLimit))
	for pos > 0 {
		// 最右边的1
		rightMost = pos & (^pos + 1)
		pos = pos - rightMost
		count += getNumber2(upperLimit,
			colLimit|rightMost,
			(leftLimit|rightMost)<<1,
			(rightLimit|rightMost)>>1)
	}
	return count
}