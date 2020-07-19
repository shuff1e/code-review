package Catalan_number

import "github.com/shuff1e/code-review/util/math"

// 已知一个字符串都是由左括号(和右括号)组成，返回最长有效括号子串的长度。

// 任何时刻左括号的数量都要大于等于右括号的数量
//  否则就是无效的字符串
func f2(str string,index int,maxCount *int) (int,int){
	if index == -1 {
		return 0,0
	}
	leftNumber,rightNumber := f2(str,index-1,maxCount)
	// 如果右括号比左括号还多，
	// 要从这里从新开始
	if leftNumber < rightNumber {
		if str[index] == '(' {
			return 1,0
		} else {
			return 0,1
		}
	} else if leftNumber == rightNumber {
		if str[index] == ')' {
			return 0,1
		} else {
			return leftNumber + 1,rightNumber
		}
	} else {
		if str[index] == '(' {
			return leftNumber + 1,rightNumber
		} else if str[index] == ')' {
			if leftNumber == rightNumber + 1 {
				*maxCount = math.Max(leftNumber*2,*maxCount)
			}
			return leftNumber,rightNumber + 1
		}
	}
	return 0,0
}

func GetMaxLength(str string) int {
	maxCount := 0
	f2(str,len(str)-1,&maxCount)
	return maxCount
}