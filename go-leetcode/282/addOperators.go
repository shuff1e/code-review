package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*

282. 给表达式添加运算符
给定一个仅包含数字 0-9 的字符串和一个目标值，在数字之间添加二元运算符（不是一元）+、- 或 * ，返回所有能够得到目标值的表达式。

示例 1:

输入: num = "123", target = 6
输出: ["1+2+3", "1*2*3"]
示例 2:

输入: num = "232", target = 8
输出: ["2*3+2", "2+3*2"]
示例 3:

输入: num = "105", target = 5
输出: ["1*0+5","10-5"]
示例 4:

输入: num = "00", target = 0
输出: ["0+0", "0-0", "0*0"]
示例 5:

输入: num = "3456237490", target = 9191
输出: []

 */

func main() {
	num := "105"
	target := 5
	fmt.Println(addOperators(num,target))
}

func addOperators(num string, target int) []string {
	result := make([]string,0)
	temp := make([]string,0)
	help(num,target,0,0,0,temp,&result)
	return result
}

func help(num string,target int,start int,val int,pre int,temp []string,result *[]string) {
	if start == len(num){
		if target == val {
			*result = append(*result,strings.Join(temp,""))
		}
		return
	}
	for i := start;i<len(num);i++ {
		if num[start] == '0' && i > start {
			break
		}
		cur,_ := strconv.Atoi(num[start:i+1])

		length := len(temp)
		if start == 0 {
			help(num,target,i+1,cur,cur,append(temp,strconv.Itoa(cur)),result)
			temp = temp[:length]
		} else {
			help(num,target,i+1,val + cur,cur,
				append(temp,"+",strconv.Itoa(cur)),result)
			temp = temp[:length]
			help(num,target,i+1,val - cur,-cur,
				append(temp,"-",strconv.Itoa(cur)),result)
			temp = temp[:length]
			help(num,target,i+1,val - pre + pre * cur,pre*cur,
				append(temp,"*",strconv.Itoa(cur)),result)
			temp = temp[:length]
		}
	}
}

/*

public List<String> addOperators(String num, int target) {
    List<String> result = new ArrayList<>();
    addOperatorsHelper(num, target, result, "", 0, 0, 0);
    return result;
}

private void addOperatorsHelper(String num, int target, List<String> result, String path, int start, long eval, long pre) {
    if (start == num.length()) {
        if (target == eval) {
            result.add(path);
        }
        return;

    }
    for (int i = start; i < num.length(); i++) {
        // 数字不能以 0 开头
        if (num.charAt(start) == '0' && i > start) {
            break;
        }
        long cur = Long.parseLong(num.substring(start, i + 1));
        if (start == 0) {
            addOperatorsHelper(num, target, result, path + cur, i + 1, cur, cur);
        } else {
            // 加当前值
            addOperatorsHelper(num, target, result, path + "+" + cur, i + 1, eval + cur, cur);
            // 减当前值
            addOperatorsHelper(num, target, result, path + "-" + cur, i + 1, eval - cur, -cur);

            //乘法有两点不同

            //当前表达式的值就是 先减去之前的值，然后加上 当前值和前边的操作数相乘
            //eval - pre + pre * cur

            //另外 addOperatorsHelper 函数传进 pre 参数需要是 pre * cur
            //比如前边讲的 2+ 3 * 4 * 5 这种连乘的情况
            addOperatorsHelper(num, target, result, path + "*" + cur, i + 1, eval - pre + pre * cur, pre * cur);
        }
    }
}

 */