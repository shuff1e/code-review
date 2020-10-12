package main

import (
	"fmt"
)

/*
187. 重复的DNA序列
所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。
在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。

编写一个函数来查找目标子串，目标子串的长度为 10，且在 DNA 字符串 s 中出现次数超过一次。

示例：

输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出：["AAAAACCCCC", "CCCCCAAAAA"]
 */

// rabin-karp

func main() {
	s := "TGCTCCTGTCACAACTTCTTTACCAGCCTGTTTTTCTAGAGTCGGCTCAAAACCTGCCTTTATGCGCAGCTGTCCACGAGAATTTCATGTTATCGAGGACCGCGATATACCCAATCGCGCGCCCCAGAAAAAAGAGTCTTACCAGATGTATACGGTGACGACCCAGTGGGTAAGACCGCTCTGCTCAGCGACCCGTCCATACCCACAGTCAGCCATGTGTGACATATCAGCGTGCATTCTTGATCTGTATGGGTGCGCTGCCCCCGCACTTGATGGGGTATGTGATGACTCCGCTCGGTAAGCAAGACCCTGGGGGTTCGGACGTAGGGTATACCCGAACTTCACGTATGCGGACACCAACGCACGTGCCAATTTATCTAACGTATGTCTCCATGCCGCCCAGAAGGTTAAAGTGGACCGCCGTTCGTATACTGTTTCTGCAATTGTGTGCGGCAGCACCAGGTAGAGAGCATTCTATTTCGCTAGCTAGTAAATCTACTTCACCGAGTCTGGAAGCTCCAATCGCTGTTTACAAACTTTTTGCCCCTGCGTGGGTCAGGCCATGTCCCGTTCCCGAGGATTCTAGCACTGACCTAGCCCTATATCACGAGCCGGGTTTTCTTAAAATAGAGATCGGGACGTTAAGGTCTTATGAACGGCTTCAGCTATCTTCCGCTTACCAACTGAGCCGAACTATCTCCGGGTGTTACATGGATCCTAAAATGCTCTCCAATTTTGCCCCTGCATGGTATTTCTCTTGAGACTACTGGATCTACCTGGGTTGTGCATGTTTCGTGTCTCTTCCGACGTTCGACAATTGGGGGCGACGCTTTAAGTTCTACTACGGTGAGATGCACATCCCACGGACGCCCTTTTCCTTTGGCTCTTCCTACGTTCGCGAGCGGTCCTGTAGGACAGTTGCTTTATGCCAACTTTTACGAGGGTGGAATACAGTATCGCCATGACACTCTGAAAAAGGATGGAAGACCTGAGATTCACC"
	result := findRepeatedDnaSequences2(s)
	fmt.Printf("%#v\n",result)
}

func findRepeatedDnaSequences(s string) []string {
	L := 10
	if len(s) <= L {
		return nil
	}
	result := map[string]struct{}{}
	a := 4
	aL := Pow(a,L)

	dict := map[byte]int{
		'A':1,
		'C':2,
		'G':3,
		'T':4,
	}

	hashDict := map[int]struct{}{}
	h := 0
	for i :=0;i<len(s)-L+1;i++ {
		if i == 0 {
			for j := 0;j<L;j++ {
				h = h*a + dict[s[j]]
			}
		} else {
			h = h*a - dict[s[i-1]]*aL + dict[s[i+L-1]]
		}
		if _,ok := hashDict[h];ok {
			result[s[i:i+L]] = struct{}{}
		} else {
			hashDict[h] = struct{}{}
		}
	}
	temp := []string{}
	for k,_ := range result {
		temp = append(temp,k)
	}
	return temp
}

func Pow(base ,exp int) int {
	result := 1
	for exp > 0 {
		if exp & 1 != 0 {
			result = result * base
		}
		exp = exp >> 1
		base = base*base
	}
	return result
}

func findRepeatedDnaSequences2(s string) []string {
	L := 10
	if len(s) <= L {
		return nil
	}
	result := map[string]struct{}{}

	dict := map[byte]int{
		'A':0,
		'C':1,
		'G':2,
		'T':3,
	}

	hashDict := map[int]struct{}{}
	h := 0
	for i :=0;i<len(s)-L+1;i++ {
		if i == 0 {
			for j := 0;j<L;j++ {
				h = h << 2
				h = h | dict[s[j]]
			}
		} else {
			h = h << 2
			h = h | dict[s[i+L-1]]
			h = h & ^(3 << (2*L))
		}
		if _,ok := hashDict[h];ok {
			result[s[i:i+L]] = struct{}{}
		} else {
			hashDict[h] = struct{}{}
		}
	}
	temp := []string{}
	for k,_ := range result {
		temp = append(temp,k)
	}
	return temp
}

// 将十进制数字转化为二进制字符串
func convertToBin(num int) string {
	s := ""
	exp := 1
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