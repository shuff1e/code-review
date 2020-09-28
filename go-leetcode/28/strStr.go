package main

import "fmt"

/*
28. 实现 strStr()
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，
在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

 */

func main() {
	haystack := "ababcaababcaabc"
	needle := "ababcaabc"
	fmt.Println(RabinKarp(haystack,needle))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0;i<=len(haystack) - len(needle);i++ {
		j := i
		for ;j-i <len(needle);j++ {
			if haystack[j] != needle[j-i] {
				break
			}
		}
		if j - i == len(needle) {
			return i
		}
	}
	return -1
}

// KMP

// Rabin-Karp

// aaaab
// aab

/*
Rabin Karp - 常数复杂度
有一种最坏时间复杂度也为 O(N)的算法。思路是这样的，先生成窗口内子串的哈希码，
然后再跟 needle字符串的哈希码做比较。

这个思路有一个问题需要解决，如何在常数时间生成子串的哈希码？

滚动哈希：常数时间生成哈希码

生成一个长度为 L 数组的哈希码，需要 O(L) 时间。

如何在常数时间生成滑动窗口数组的哈希码？利用滑动窗口的特性，每次滑动都有一个元素进，一个出。

由于只会出现小写的英文字母，因此可以将字符串转化成值为 0 到 25 的整数数组： arr[i] = (int)S.charAt(i) - (int)'a'。
按照这种规则，abcd 整数数组形式就是 [0, 1, 2, 3]，转换公式如下所示。

h0 = 0 x 26^3 + 1 x 26^2 + 2 x 26^1 + 3 x 26^0

下面来考虑窗口从 abcd 滑动到 bcde 的情况。这时候整数形式数组从 [0, 1, 2, 3] 变成了 [1, 2, 3, 4]，数组最左边的 0 被移除，
同时最右边新添了 4。滑动后数组的哈希值可以根据滑动前数组的哈希值来计算，计算公式如下所示。

h1 = (h0 - 0 x 26^3) x 26 + 4 x 26^0

如何避免溢出

h1可能是一个很大的数字，因此需要设置数值上限来避免溢出。设置数值上限可以用取模的方式，即用 h % modulus 来代替原本的哈希值。

理论上，modules 应该取一个很大数，但具体应该取多大的数呢? 详见这篇文章，对于这个问题来说 2^{31}就足够了。

算法

计算子字符串 haystack.substring(0, L) 和 needle.substring(0, L) 的哈希值。

从起始位置开始遍历：从第一个字符遍历到第 N - L 个字符。

根据前一个哈希值计算滚动哈希。

如果子字符串哈希值与 needle 字符串哈希值相等，返回滑动窗口起始位置。

返回 -1，这时候 haystack 字符串中不存在 needle 字符串。

 */

func RabinKarp(str1,str2 string) int {
	if len(str2) > len(str1) {
		return -1
	}
	module := 1<<30
	base := 26
	pivot := 0
	result := 0
	for i := 0;i<len(str2);i++ {
		pivot = (pivot + int(str2[i] - 'a'))*base%module
		result = (result + int(str1[i] - 'a'))*base%module
	}
	if result == pivot {
		return 0
	}

	EXP := Pow(base,len(str2),module)
	for start := 1;start<= len(str1)-len(str2);start++ {
		// 有减法运算，这里有可能是负的
		result = (result - int(str1[start-1] - 'a')*EXP + int(str1[start + len(str2)-1] - 'a') + module)*base%module
		if result == pivot {
			return start
		}
	}
	return -1
}

func Pow(base,exp,module int) int {
	result := 1
	for exp > 0 {
		if exp & 1 > 0 {
			result = result * base % module
		}
		base = base * base
		exp = exp >> 1
	}
	return result
}

// kmp

// http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html

/*
举例来说，有一个字符串"BBC ABCDAB ABCDABCDABDE"，我想知道，里面是否包含另一个字符串"ABCDABD"？

针对搜索词，算出一张《部分匹配表》（Partial Match Table）

移动位数 = 已匹配的字符数 - 最后一个匹配的字符的对应的部分匹配值

下面介绍《部分匹配表》是如何产生的。
首先，要了解两个概念："前缀"和"后缀"。
"前缀"指除了最后一个字符以外，一个字符串的全部头部组合；"后缀"指除了第一个字符以外，一个字符串的全部尾部组合。

"部分匹配值"就是"前缀"和"后缀"的最长的共有元素的长度。以"ABCDABD"为例，

　　－　"A"的前缀和后缀都为空集，共有元素的长度为0；

　　－　"AB"的前缀为[A]，后缀为[B]，共有元素的长度为0；

　　－　"ABC"的前缀为[A, AB]，后缀为[BC, C]，共有元素的长度0；

　　－　"ABCD"的前缀为[A, AB, ABC]，后缀为[BCD, CD, D]，共有元素的长度为0；

　　－　"ABCDA"的前缀为[A, AB, ABC, ABCD]，后缀为[BCDA, CDA, DA, A]，共有元素为"A"，长度为1；

　　－　"ABCDAB"的前缀为[A, AB, ABC, ABCD, ABCDA]，后缀为[BCDAB, CDAB, DAB, AB, B]，共有元素为"AB"，长度为2；

　　－　"ABCDABD"的前缀为[A, AB, ABC, ABCD, ABCDA, ABCDAB]，后缀为[BCDABD, CDABD, DABD, ABD, BD, D]，共有元素的长度为0。

"部分匹配"的实质是，有时候，字符串头部和尾部会有重复。比如，"ABCDAB"之中有两个"AB"，那么它的"部分匹配值"就是2（"AB"的长度）。搜索词移动的时候，第一个"AB"向后移动4位（字符串长度-部分匹配值），就可以来到第二个"AB"的位置。
*/