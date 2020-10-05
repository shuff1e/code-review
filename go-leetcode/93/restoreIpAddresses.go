package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
93. 复原IP地址
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。



示例 1：

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
示例 2：

输入：s = "0000"
输出：["0.0.0.0"]
示例 3：

输入：s = "1111"
输出：["1.1.1.1"]
示例 4：

输入：s = "010010"
输出：["0.10.0.10","0.100.1.0"]
示例 5：

输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]


提示：

0 <= s.length <= 3000
s 仅由数字组成
 */

func main() {
	str := "25525511135"
	str = "0000"
	str = "101023"
	str = "1111"
	result := restoreIpAddresses(str)
	for i := 0;i<len(result);i++ {
		fmt.Println(result[i])
	}
}

func restoreIpAddresses(s string) []string {
	if len(s)<4 || len(s) > 14 {
		return nil
	}
	result := []string{}
	temp := []string{}
	help(s,0,&temp,&result)
	return result
}

func help(s string, index int,temp *[]string,result *[]string) {
	if index == len(s) && len(*temp) == 4 {
		*result = append(*result,strings.Join(*temp,"."))
		return
	}

	if index == len(s) {
		return
	}

	if len(*temp) == 4 {
		return
	}

	count := 0
	for i :=index + 1;i<=len(s) && count < 3;i++ {
		if checkValid(s,index,i) {
			*temp = append(*temp,s[index:i])
			help(s,i,temp,result)
			*temp = (*temp)[0:len(*temp)-1]
		}
		count ++
	}
}

func checkValid(str string,start,end int) bool {
	if start + 1 == end {
		return true
	} else if str[start] == '0' {
		return false
	} else {
		v,err := strconv.Atoi(str[start:end])
		if err != nil {
			return false
		}
		return v >= 10 && v <= 255
	}
}