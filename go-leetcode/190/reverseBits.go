package main

import "fmt"

/*
190. 颠倒二进制位
颠倒给定的 32 位无符号整数的二进制位。



示例 1：

输入: 00000010100101000001111010011100
输出: 00111001011110000010100101000000
解释: 输入的二进制串 00000010100101000001111010011100 表示无符号整数 43261596，
因此返回 964176192，其二进制表示形式为 00111001011110000010100101000000。
示例 2：

输入：11111111111111111111111111111101
输出：10111111111111111111111111111111
解释：输入的二进制串 11111111111111111111111111111101 表示无符号整数 4294967293，
因此返回 3221225471 其二进制表示形式为 10111111111111111111111111111111 。


提示：

请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 2 中，输入表示有符号整数 -3，输出表示有符号整数 -1073741825。
 */











// 11001
// result = 1 10  100 1001 10011
// mask = 10  100 1000 10000 100000

// 将十进制数字转化为二进制字符串
func convertToBin(num uint32) string {
	s := ""
	exp := uint32(1)
	for num != 0 {
		if num & exp != 0 {
			s = "1" + s
			num = num & (num-1)
		} else {
			s = "0" + s
		}
		exp = exp << 1
	}
	return s
}

func main() {
	num := uint32(0b00000010100101000001111010011100)
	fmt.Println(convertToBin(num))
	result := reverseBits3(num)
	fmt.Println(convertToBin(result))
}

func reverseBits(num uint32) uint32 {
	count := 0
	result := uint32(0)
	mask := uint32(1)
	for count < 32 {
		result = result << 1
		if num & mask != 0 {
			num = num & (num - 1)
			result = result + 1
		}
		mask = mask << 1
		count ++
	}
	return result
}

func reverseBits2(num uint32) uint32 {
	result := uint32(0)
	power := uint32(31)
	for num != 0 {
		result += (num&1) << power
		num >>= 1
		power --
	}
	return result
}

// c->1100
// a->1010
// 0011 -> 3
// 0101 -> 5
func reverseBits3(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8 ) | ((num & 0x00ff00ff) << 8 )
	num = ((num & 0xf0f0f0f0) >> 4 ) | ((num & 0x0f0f0f0f) << 4 )
	num = ((num & 0xcccccccc) >> 2 ) | ((num & 0x33333333) << 2 )
	num = ((num & 0xaaaaaaaa) >> 1 ) | ((num & 0x55555555) << 1 )
	return num
}
