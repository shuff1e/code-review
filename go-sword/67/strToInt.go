package main

import "fmt"

// 67：把字符串转换成整数
// 题目：请你写一个函数StrToInt，实现把字符串转换成整数这个功能。当然，不
// 能使用atoi或者其他类似的库函数。

/**

A constant holding the minimum value an {@code int} can
have, -2^31.
*/
// @Native public static final int MIN_VALUE = 0x80000000;
/**

  A constant holding the maximum value an {@code int} can
  have, 2^31-1.
*/
// @Native public static final int MAX_VALUE = 0x7fffffff;

// A：考虑正负号，溢出，无效字符，空字符串
// 状态转换矩阵

//       |  +/-    |  0~9   |  a  | overflow |
// start |  symbol | number | end | end      |
// symbol|   end   | number | end | end      |
// number|   end   | number | end | end      |
// end   |    end  |   end  | end | end      |

func strToInt(str string) (int32,bool) {
	stateMatrix := map[string][]string {
		"start": {"symbol","num","end"},
		"symbol":{"end","num","end"},
		"num":{"end","num","end"},
		"end": {"end","end","end"},
	}
	var flag int32 = 1
	var result int64 = 0

	state := "start"
	for i := 0;i<len(str);i++ {
		if state == "end" {
			return 0,false
		}
		if str[i] == '+' || str[i] == '-' {
			if str[i] == '-' {
				flag = -1
			}
			state = stateMatrix[state][0]

		} else if str[i] >= '0' && str[i] <= '9' {
			result = result*10 + int64(str[i] - '0')
			if (flag == 1 && result > int64(0x7fffffff)) ||
				(flag == -1 && result > int64(0x80000000)) {
				state = stateMatrix[state][2]
			} else {
				state = stateMatrix[state][1]
			}

		} else {
			state = stateMatrix[state][2]
		}
	}
	if state == "num" {
		return flag*int32(result),true
	}
	return 0,false
}

func main() {
	Test("");

	Test("123");

	Test("+123");

	Test("-123");

	Test("1a33");

	Test("+0");

	Test("-0");

	//有效的最大正整数, 0x7FFFFFFF
	// 2147483647
	Test("+2147483647");

	Test("-2147483647");

	Test("+2147483648");

	//有效的最小负整数, 0x80000000
	// -2147483648
	Test("-2147483648");

	Test("+2147483649");

	Test("-2147483649");

	Test("+");

	Test("-");
}

func Test(str string) {
	num,ok := strToInt(str)
	if !ok {
		fmt.Println(str,"invalid")
	} else {
		fmt.Println(str,"valid",num)
	}
}