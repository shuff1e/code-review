package main

/*

5562. 字符频次唯一的最小删除次数
如果字符串 s 中 不存在 两个不同字符 频次 相同的情况，就称 s 是 优质字符串 。

给你一个字符串 s，返回使 s 成为 优质字符串 需要删除的 最小 字符数。

字符串中字符的 频次 是该字符在字符串中的出现次数。例如，在字符串 "aab" 中，'a' 的频次是 2，而 'b' 的频次是 1 。



示例 1：

输入：s = "aab"
输出：0
解释：s 已经是优质字符串。
示例 2：

输入：s = "aaabbbcc"
输出：2
解释：可以删除两个 'b' , 得到优质字符串 "aaabcc" 。
另一种方式是删除一个 'b' 和一个 'c' ，得到优质字符串 "aaabbc" 。
示例 3：

输入：s = "ceabaacb"
输出：2
解释：可以删除两个 'c' 得到优质字符串 "eabaab" 。
注意，只需要关注结果字符串中仍然存在的字符。（即，频次为 0 的字符会忽略不计。）


提示：

1 <= s.length <= 105
s 仅含小写英文字母

*/
import "fmt"

func main() {
	s := "aaabbbcc"
	s = "ceabaacb"
	s = "abbcccddd"
	//s = "a"
	fmt.Println(minDeletions(s))
	fmt.Println(minDeletionsBetter(s))
	fmt.Println(minDeletionsBetter2(s))
}

func minDeletions(s string) int {
	length := 0
	dict := make(map[byte]int,0)
	for i := 0;i<len(s);i++ {
		dict[s[i]] = dict[s[i]] + 1
		length = Max(length,dict[s[i]])
	}
	arr := make([]int,length+1)
	for _,v := range dict {
		arr[v] = arr[v] + 1
	}
	return help(arr)
}

func help(arr []int) int {
	result := 0
	found := false
	for i := 1;i<len(arr);i++ {
		if arr[i] > 1 {
			arr[i] --
			arr[i-1] ++
			if ! found {
				found = true
				result = help(arr) + 1
			} else {
				result = Min(result,help(arr) + 1)
			}
			arr[i] ++
			arr[i-1] --
			break
		}
	}
	if !found {
		return 0
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func minDeletionsBetter(s string) int {
	arr := make([]int,26)
	for i := 0;i<len(s);i++ {
		arr[s[i]-'a'] ++
	}

	set := make(map[int]struct{})
	result := 0

	for i := 0;i<len(arr);i++ {
		if arr[i] != 0 {
			for {
				if _,ok := set[arr[i]];!ok {
					break
				}
				arr[i] --
				result ++
			}
			if arr[i] != 0 {
				set[arr[i]] = struct{}{}
			}
		}
	}
	return result
}

// p[i] = i
// 表示出现次数的为i的第一次出现
// 否则表示已经出现过了
// p[i]就表示i出现时， 应该减为多少

func minDeletionsBetter2(s string) int {
	arr := make([]int,26)
	for i := 0;i<len(s);i++ {
		arr[s[i]-'a'] ++
	}

	p := make([]int,100010)
	for i := 0;i<len(p);i++ {
		p[i] = i
	}

	result := 0
	for i := 0;i<len(arr);i++ {
		if arr[i] == 0 {
			continue
		}
		t := find(p,arr[i])
		result += arr[i] - t
		p[t] = Max(0,p[t]-1)
	}

	return result
}

func find(p []int,x int) int {
	if x != p[x] {
		p[x] = find(p,p[x])
	}
	return p[x]
}