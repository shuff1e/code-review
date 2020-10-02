package main

import (
	"fmt"
	"strings"
)

/*
65. 有效数字
验证给定的字符串是否可以解释为十进制数字。

例如:

"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
" -90e3   " => true
" 1e" => false
"e3" => false
" 6e-1" => true
" 99e2.5 " => false
"53.5e93" => true
" --6 " => false
"-+3" => false
"95a54e53" => false

说明: 我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。这里给出一份可能存在于有效十进制数字中的字符列表：

数字 0-9
指数 - "e"
正/负号 - "+"/"-"
小数点 - "."
当然，在输入中，这些字符的上下文也很重要。

更新于 2015-02-10:
C++函数的形式已经更新了。如果你仍然看见你的函数接收 const char * 类型的参数，请点击重载按钮重置你的代码。
 */

// 去除trailing space
//        0,9         -+           .             a 空格
// start  number      symbol      point          end
// number number      end         np              end
// symbol number       end        point          end
// point  nPoint      end         end            end
// nPoint nPoint      end         end            end

func main() {
	fmt.Println(isNumber("6e6.5"))
}

func isNumber(s string) bool {
	matrix := map[string][]string{
		"start":{"num","sym","p","end"},
		"num":{"num","end","np","end"},
		"sym":{"num","end","p","end"},
		"p":{"np","end","end","end"},
		"np":{"np","end","end","end"},
		"end":{"end","end","end","end"},
	}
	s = strings.TrimSpace(s)
	index,count := FindE(s)
	if count > 1 {
		return false
	}
	if (count == 1) && (index == 0 || index == len(s) - 1) {
		return false
	}
	if (count == 1) && (s[index+1] == ' ' || s[index-1] == ' ') {
		return false
	}
	if (count == 1) && s[index+1] == '.' {
		return false
	}
	if count == 1 {
		left := getState(s[0:index],matrix)
		right := getState(s[index+1:],matrix)
		if (left == "np" || left == "num")&& (right == "num") {
			return true
		} else {
			return false
		}
	} else {
		state := getState(s,matrix)
		return state == "num" || state == "np"
	}
}

func getState(str string,matrix map[string][]string) string {
	state := "start"
	for i := 0;i<len(str);i++ {
		state = matrix[state][getCol(str[i])]
	}
	return state
}

func getCol( char byte) int {
	if char >= '0' && char <= '9' {
		return 0
	}
	if char == '+' || char == '-' {
		return 1
	}
	if char == '.' {
		return 2
	}
	return 3
}

func FindE(str string) (index ,count int) {
	for i := 0;i<len(str);i++ {
		if str[i] == 'e' {
			index = i
			count ++
		}
	}
	return index,count
}
