package main

import "fmt"

// 20：表示数值的字符串
// 题目：请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，
// 字符串“+100”、“5e2”、“-123”、“3.1416”及“-1E-16”都表示数值，但“12e”、
// “1a3.14”、“1.2.3”、“+-5”及“12e+5.4”都不是

// A：有限状态机 DFA
// 状态转换表为
//           | 5        |   +-   | .       | a    |
// start     | number   |  symbol| point   |  end |
// number    |number    | end    | numPoint|  end |
// symbol    |number    | end    | point   | end  |
// point     |numPoint  | end    | end     | end  |
// numPoint  |numPoint  | end    | end     |  end |
// end       | end      |  end   |   end   |  end |

// e或者E的状态比较复杂
// 字符串中多个e肯定是不规范的
// 一个e的情况下，以e分隔成两部分，每部分都需要是一个有效的数字

func isValid(str string) bool {
	eCount := 0
	eIndex := -1
	for i := 0;i<len(str);i++ {
		if str[i] == 'e' || str[i] == 'E' {
			eIndex = i
			eCount +=1
			if eCount >= 2 {
				return false
			}
		}
	}
	if eIndex == -1 {
		return stateValid(isValidHelper(str))
	} else if eIndex == 0 || eIndex == len(str)-1 {
		return false
	} else {
		left := isValidHelper(str[0:eIndex])
		right := isValidHelper(str[eIndex+1:])
		if right == "numPoint" {
			return false
		}
		return stateValid(left) && stateValid(right)
	}
}

func stateValid(state string) bool {
	return state == "number" || state == "numPoint"
}

func isValidHelper(str string) string {
	matrix := map[string][]string{
		"start":{"number","symbol","point","end"},
		"number":{"number","end","numPoint","end"},
		"symbol":{"number","end","point","end"},
		"point":{"numPoint","end","end","end"},
		"numPoint":{"numPoint","end","end","end"},
		"end":{"end","end","end","end"},
	}

	state := "start"
	for i := 0;i<len(str);i++ {
		if str[i] >= '0' && str[i] <= '9' {
			state = matrix[state][0]
		} else if str[i] == '+' || str[i] == '-' {
			state = matrix[state][1]
		} else if str[i] == '.' {
			state = matrix[state][2]
		} else {
			state = matrix[state][3]
		}
	}
	return state
}

func main() {
	Test("Test1", "100", true);
	Test("Test2", "123.45e+6", true);
	Test("Test3", "+500", true);
	Test("Test4", "5e2", true);
	Test("Test5", "3.1416", true);
	Test("Test6", "600.", true);
	Test("Test7", "-.123", true);
	Test("Test8", "-1E-16", true);
	Test("Test9", "1.79769313486232E+308", true);
	println("\n\n");

	Test("Test10", "12e", false);
	Test("Test11", "1a3.14", false);
	Test("Test12", "1+23", false);
	Test("Test13", "1.2.3", false);
	Test("Test14", "+-5", false);
	Test("Test15", "12e+5.4", false);
	Test("Test16", ".", false);
	Test("Test17", ".e1", false);
	Test("Test18", "e1", false);
	Test("Test19", "+.", false);
	Test("Test20", "", false);
}

func Test(name,str string,valid bool) {
	fmt.Println(name,isValid(str)==valid)
}
