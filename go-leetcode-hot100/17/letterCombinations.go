package main

import "fmt"

/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

 */

func main() {
	fmt.Printf("%#v\n",letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	dict := map[byte][]string{
		2:{"a","b","c"},
		3:{"d","e","f"},
		4:{"g","h","i"},
		5:{"j","k","l"},
		6:{"m","n","o"},
		7:{"p","q","r","s"},
		8:{"t","u","v"},
		9:{"w","x","y","z"},
	}
	result := []string{}
	help(digits,"",&result,dict)
	return result
}

func help(digits string,temp string,result *[]string,dict map[byte][]string) {
	if len(digits) == 0 {
		if len(temp) > 0 {
			*result = append(*result,temp)
		}
		return
	}
	for _,v := range dict[digits[0]-'0'] {
		help(digits[1:],temp + v,result,dict)
	}
}