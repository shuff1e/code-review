package main

import (
	"fmt"
	"unsafe"
)

/*
168. Excel表列名称
给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，

1 -> A
2 -> B
3 -> C
...
26 -> Z
27 -> AA
28 -> AB
...
示例 1:

输入: 1
输出: "A"
示例 2:

输入: 28
输出: "AB"
示例 3:

输入: 701
输出: "ZY"
 */

func main() {
	fmt.Println(convertToTitle(1510))
	fmt.Println('B'-'A'+1)
	fmt.Println(6+4*26)
	fmt.Println(2 + 6*26 + 2*26*26)
}

// AZ
// 26 + 1 * 26

// DF
// 6 + 4 *26 = 110

// BFB
// 2 + 6*26 + 2*26*26 = 1510

// 因为 Excel 取值范围为 1~26，故可将 26 进制 逻辑上的 个位、十位、百位…均减 1 映射到 0~25 即可，最后转换为字符

func convertToTitle(n int) string {
	const base = 26
	var pivot byte = 'A'
	result := []byte{}
	for n > 0 {
		n --
		result = append([]byte{byte(n%base) + pivot},result...)
		n = n/base
	}
	return String(result)
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}