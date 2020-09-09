package main

import "fmt"

// 19：正则表达式匹配
// 题目：请实现一个函数用来匹配包含'.'和'*'的正则表达式。模式中的字符'.'
// 表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题
// 中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"
// 和"ab*ac*a"匹配，但与"aa.a"及"ab*a"均不匹配。























// A：.表示匹配任意一个字符，比较好处理
// 如果strIndex == patternIndex 或者 patternIndex == . ，则str和pattern都index+1
// *表示匹配0个，1个，或者多个

// 如果patternIndex+1 == *
// 如果strIndex == patternIndex，或者patternIndex == .
// 匹配0个，strIndex,patternIndex+2
// 匹配1个，strIndex+1,patternIndex+2
// 匹配多个，strIndex+1,patternIndex

// 如果strIndex != patternIndex
// 则只能匹配0个，strIndex,patternIndex+2

// 综上，每次都是strIndex+1，或者patternIndex增加

func match(str,pattern string,strIndex,patternIndex int) bool {
	if strIndex == len(str) && patternIndex == len(pattern) {
		return true
	}
	if strIndex == len(str) {
		if patternIndex + 1 < len(pattern) && pattern[patternIndex + 1] == '*' {
			return match(str,pattern,strIndex,patternIndex+2)
		} else {
			return false
		}
	}
	if patternIndex == len(pattern) {
		return false
	}
	if patternIndex + 1 < len(pattern) && pattern[patternIndex+1] == '*' {
		if str[strIndex] == pattern[patternIndex] || pattern[patternIndex] == '.' {
			return match(str,pattern,strIndex,patternIndex+2) ||
				match(str,pattern,strIndex+1,patternIndex+2) ||
				match(str,pattern,strIndex+1,patternIndex)
		} else {
			return match(str,pattern,strIndex,patternIndex+2)
		}
	}
	if str[strIndex] == pattern[patternIndex] || pattern[patternIndex] == '.' {
		return match(str,pattern,strIndex+1,patternIndex+1)
	}
	return false
}

func main() {
	Test("Test01", "", "", true);
	Test("Test02", "", ".*", true);
	Test("Test03", "", ".", false);
	Test("Test04", "", "c*", true);
	Test("Test05", "a", ".*", true);
	Test("Test06", "a", "a.", false);
	Test("Test07", "a", "", false);
	Test("Test08", "a", ".", true);
	Test("Test09", "a", "ab*", true);
	Test("Test10", "a", "ab*a", false);
	Test("Test11", "aa", "aa", true);
	Test("Test12", "aa", "a*", true);
	Test("Test13", "aa", ".*", true);
	Test("Test14", "aa", ".", false);
	Test("Test15", "ab", ".*", true);
	Test("Test16", "ab", ".*", true);
	Test("Test17", "aaa", "aa*", true);
	Test("Test18", "aaa", "aa.a", false);
	Test("Test19", "aaa", "a.a", true);
	Test("Test20", "aaa", ".a", false);
	Test("Test21", "aaa", "a*a", true);
	Test("Test22", "aaa", "ab*a", false);
	Test("Test23", "aaa", "ab*ac*a", true);
	Test("Test24", "aaa", "ab*a*c*a", true);
	Test("Test25", "aaa", ".*", true);
	Test("Test26", "aab", "c*a*b", true);
	Test("Test27", "aaca", "ab*a*c*a", true);
	Test("Test28", "aaba", "ab*a*c*a", false);
	Test("Test29", "bbbba", ".*a*a", true);
	Test("Test30", "bcbbabab", ".*a*a", false);
}

func Test(name,str,pattern string,isMatched bool) {
	fmt.Println(match(str,pattern,0,0)==isMatched)
}
