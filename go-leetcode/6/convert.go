package main

import "fmt"

/*
6. Z 字形变换
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

 */

// A：找规律
// L EETCO D EISHI R ING
// 然后下一个要找的就是 E ETC O, E ISH I

func main() {
	s := "LEETCODEISHIRING"
	fmt.Println(convert(s,1))
}

func convert(s string, numRows int) string {
	if len(s) <= 1 {
		return s
	}
	if numRows == 1 {
		return s
	}
	start := 0
	result := ""
	for start < len(s) {
		next := start + (numRows-1)*2
		result = result + s[start:start+1]
		start = next
	}


	for i := 1;i<numRows-1;i++ {
		start := 0
		offset := i
		for start < len(s) {
			next := start+(numRows-1)*2
			if start+offset < len(s) {
				result = result + s[start+offset:start+offset+1]
			}
			if start+offset+ (numRows-1-i)*2 < len(s) {
				result = result + s[start+offset+ (numRows-1-i)*2:start+offset+(numRows-1-i)*2 + 1]
			}
			start = next
		}
	}

	start = 0
	offset := numRows - 1
	for start < len(s) {
		next := start + (numRows-1)*2
		if start+offset < len(s) {
			result = result + s[start+offset:start+offset+1]
		}
		start = next
	}
	return result
}