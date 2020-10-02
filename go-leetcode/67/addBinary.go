package main

import (
	"fmt"
	"strings"
)

/*
67. 二进制求和
给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。



示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"


提示：

每个字符串仅由字符 '0' 或 '1' 组成。
1 <= a.length, b.length <= 10^4
字符串如果不是 "0" ，就都不含前导零。
 */

// (a & b)<<1 就是carry
// a ^ b 就是

func main() {
	fmt.Println(strings.Trim("111222","1"))
	fmt.Println(addBinary("1111","101"))
}

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		return addBinary(b,a)
	}

	for {
		carry := ""
		result := ""
		for i,j := len(a)-1,len(b)-1;i>=0;i-- {
			if a[i] == '1' && b[j] == '1' {
				carry = "1"+carry
				result = "0" + result
			} else if a[i] == '0' && b[j] == '0' {
				carry = "0" +carry
				result = "0" + result
			} else {
				carry = "0" +carry
				result =   "1" + result
			}
			j--
		}

		result = b[0:len(b)-len(a)]+result

		if carry != "" {
			carry = carry + "0"
		}

		//
		carry = trimleftZero(carry)
		result = trimleftZero(result)
		if carry == "" && result == "" {
			return "0"
		}
		if carry == "" {
			return result
		}
		if result == "" {
			return carry
		}

		a = carry
		b = result
		if len(a) > len(b) {
			temp := a
			a = b
			b= temp
		}
	}
}

func trimleftZero(str string) string {
	for i := 0;i<len(str);i++ {
		if str[i] != '0' {
			return str[i:]
		}
	}
	return ""
}