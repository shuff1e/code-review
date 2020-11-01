package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*

726. 原子的数量
给定一个化学式formula（作为字符串），返回每种原子的数量。

原子总是以一个大写字母开始，接着跟随0个或任意个小写字母，表示原子的名字。

如果数量大于 1，原子后会跟着数字表示原子的数量。如果数量等于 1 则不会跟数字。
例如，H2O 和 H2O2 是可行的，但 H1O2 这个表达是不可行的。

两个化学式连在一起是新的化学式。例如 H2O2He3Mg4 也是化学式。

一个括号中的化学式和数字（可选择性添加）也是化学式。例如 (H2O2) 和 (H2O2)3 是化学式。

给定一个化学式，输出所有原子的数量。格式为：第一个（按字典序）原子的名子，
跟着它的数量（如果数量大于 1），然后是第二个原子的名字（按字典序），跟着它的数量（如果数量大于 1），以此类推。

示例 1:

输入:
formula = "H2O"
输出: "H2O"
解释:
原子的数量是 {'H': 2, 'O': 1}。
示例 2:

输入:
formula = "Mg(OH)2"
输出: "H2MgO2"
解释:
原子的数量是 {'H': 2, 'Mg': 1, 'O': 2}。
示例 3:

输入:
formula = "K4(ON(SO3)2)2"
输出: "K4N2O14S4"
解释:
原子的数量是 {'K': 4, 'N': 2, 'O': 14, 'S': 4}。
注意:

所有原子的第一个字母为大写，剩余字母都是小写。
formula的长度在[1, 1000]之间。
formula只包含字母、数字和圆括号，并且题目中给定的是合法的化学式。

 */

func main() {
	str := "Mg(OH)2"
	fmt.Println(countOfAtoms2(str))
}

// 遇到左括号，就递归调用
// 遇到右括号，就结束当前层次的递归，然后结算
// 在（）内的结果计算，返回给当前层之后，继续计算
func countOfAtoms(formula string) string {
	index := 0
	countsMap := parse(formula,&index)

	keySet := []string{}
	for k,_ := range countsMap {
		keySet = append(keySet,k)
	}
	sort.Slice(keySet, func(i, j int) bool {
		return strings.Compare(keySet[i],keySet[j]) < 0
	})

	result := ""
	for _,k := range keySet {
		result += k
		if countsMap[k] > 1 {
			result += strconv.Itoa(countsMap[k])
		}
	}
	return result
}

func parse(str string,index *int) map[string]int {
	countsMap := map[string]int{}

	for *index < len(str) && str[*index] != ')' {
		if str[*index] == '(' {
			*index ++
			tempMap := parse(str,index)
			for k,v := range tempMap {
				countsMap[k] += v
			}
		} else {
			istart := *index
			*index ++
			for *index < len(str) && (str[*index] >= 'a' && str[*index] <= 'z') {
				*index ++
			}
			name := str[istart:*index]
			istart = *index
			for *index < len(str) && str[*index] >= '0' && str[*index] <= '9' {
				*index ++
			}
			multiple := 1
			if *index > istart {
				multiple,_ = strconv.Atoi(str[istart:*index])
			}
			countsMap[name] = countsMap[name] + multiple
		}
	}
	*index ++
	istart := *index
	for *index < len(str) && (str[*index] >= '0' && str[*index] <= '9') {
		*index ++
	}
	if *index > istart {
		mutiple,_ := strconv.Atoi(str[istart:*index])
		for k,v := range countsMap {
			countsMap[k] = v *mutiple
		}
	}
	return countsMap
}

// 递归改成迭代
func countOfAtoms2(formula string) string {
	stack := []map[string]int{}
	stack = append(stack,make(map[string]int))

	for i := 0;i<len(formula); {
		if formula[i] == '(' {
			stack = append(stack,make(map[string]int))
			i ++
		} else if formula[i] == ')' {
			tempMap := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i ++
			istart := i
			for i < len(formula) && (formula[i] >= '0' && formula[i] <= '9') {
				i ++
			}
			if i > istart {
				mutiple,_ := strconv.Atoi(formula[istart:i])
				for k,v := range tempMap {
					tempMap[k] = v *mutiple
				}
			}
			temp2Map := stack[len(stack)-1]
			for k,v := range tempMap {
				temp2Map[k] += v
			}
		} else {
			istart := i
			i ++
			for i < len(formula) && (formula[i] >= 'a' && formula[i] <= 'z') {
				i ++
			}
			name := formula[istart:i]
			istart = i
			for i < len(formula) && formula[i] >= '0' && formula[i] <= '9' {
				i ++
			}
			multiple := 1
			if i > istart {
				multiple,_ = strconv.Atoi(formula[istart:i])
			}
			tempMap := stack[len(stack)-1]
			tempMap[name] = tempMap[name] + multiple
		}
	}

	tempMap := stack[len(stack)-1]
	keySet := []string{}
	for k,_ := range tempMap {
		keySet = append(keySet,k)
	}
	sort.Slice(keySet, func(i, j int) bool {
		return strings.Compare(keySet[i],keySet[j]) < 0
	})

	result := ""
	for _,k := range keySet {
		result += k
		if tempMap[k] > 1 {
			result += strconv.Itoa(tempMap[k])
		}
	}
	return result
}

/*

// A-Z a-z \d*
// 或者 (
// 或者 ) \d*

import java.util.regex.*;

class Solution {
    public String countOfAtoms(String formula) {
        Matcher matcher = Pattern.compile("([A-Z][a-z]*)(\\d*)|(\\()|(\\))(\\d*)").matcher(formula);
        Stack<Map<String, Integer>> stack = new Stack();
        stack.push(new TreeMap());

        while (matcher.find()) {
            String match = matcher.group();
            if (match.equals("(")) {
                stack.push(new TreeMap());
            } else if (match.startsWith(")")) {
                Map<String, Integer> top = stack.pop();
                int multiplicity = match.length() > 1 ? Integer.parseInt(match.substring(1, match.length())) : 1;
                for (String name: top.keySet()) {
                    stack.peek().put(name, stack.peek().getOrDefault(name, 0) + top.get(name) * multiplicity);
                }
            } else {
                int i = 1;
                while (i < match.length() && Character.isLowerCase(match.charAt(i))) {
                    i++;
                }
                String name = match.substring(0, i);
                int multiplicity = i < match.length() ? Integer.parseInt(match.substring(i, match.length())) : 1;
                stack.peek().put(name, stack.peek().getOrDefault(name, 0) + multiplicity);
            }
        }

        StringBuilder ans = new StringBuilder();
        for (String name: stack.peek().keySet()) {
            ans.append(name);
            final int count = stack.peek().get(name);
            if (count > 1) ans.append(String.valueOf(count));
        }
        return ans.toString();
    }
}

 */