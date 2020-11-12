package main

import (
	"fmt"
	"strconv"
)

/*

306. 累加数
累加数是一个字符串，组成它的数字可以形成累加序列。

一个有效的累加序列必须至少包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。

给定一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是累加数。

说明: 累加序列里的数不会以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。

示例 1:

输入: "112358"
输出: true
解释: 累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
示例 2:

输入: "199100199"
输出: true
解释: 累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199
进阶:
你如何处理一个溢出的过大的整数输入?

 */

func main() {
	str := "199100199"
	fmt.Println(isAdditiveNumber(str))
}

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	index := 0
	for ;index < len(num);index ++ {
		if num[index] != '0' {
			break
		}
	}

	if index == len(num) {
		return true
	}
	if index > 1 {
		return false
	}

	if index == 1 {
		for j := index;j<len(num)-1;j++ {
			a := 0
			b,_ := strconv.Atoi(num[index:j+1])
			if help(num,j+1,a,b) {
				return true
			}
		}
		return false
	}

	// 确定前两个数
	for i := 0;i<len(num)-2;i++ {
		if num[i + 1] == '0' {
			a,_ := strconv.Atoi(num[:i+1])
			b := 0
			if help(num,i+2,a,b) {
				return true
			}
		} else {
			for j := i+1;j<len(num)-1;j++ {
				a,_ := strconv.Atoi(num[:i+1])
				b,_ := strconv.Atoi(num[i+1:j+1])
				if help(num,j+1,a,b) {
					return true
				}
			}
		}
	}
	return false
}

func help(str string,index int,a,b int) bool {
	if index == len(str) {
		return true
	}
	if str[index] == '0' {
		if a + b == 0 {
			return help(str,index+1,b,0)
		}
	} else {
		for i := index;i<len(str);i++ {
			t,_ := strconv.Atoi(str[index:i+1])
			if a + b == t {
				if help(str,i+1,b,t) {
					return true
				}
			}
		}
	}
	return false
}

/*

解题思路
1.这个题百思不得其解，一定要再做一遍。
查询题解，主要思路是DFS+剪枝。我看了一个代码数量比较少，看起来好理解的。
这个思路需要非常严谨，否则不知道哪里就会出错。
2.主要思想就是将数字从字符串截下来，转化为long double 型存到vector 里面。
在进行dfs时判断是否符合前两数之和，若不是就返回false，并且到上一层将vector里面的错误答案pop出来。

class Solution {
public:
    bool dfs(string  num,vector<long double > & res){

        int len_num = num.size();
        int len_res = res.size();

        if(len_res>2&&res[len_res-1]!=res[len_res-2]+res[len_res-3]) return false;
        if(len_num==0&&len_res>=3) return true;
        //上两行不能交换，若先判断2，那么111 这个样例就会返回true；

        for(int i=0;i<len_num;i++){
            if(num[0]=='0'&&i>0) break;//可以单个为0，但是不能是以0为开头的多位数字。
            long double a = stold(num.substr(0,i+1)); //stold = string-》long double；

            res.push_back(a);
            if(dfs(num.substr(i+1),res)){
                return true;
            }
            res.pop_back();
        }
        return false;
    }

    bool isAdditiveNumber(string num) {
        vector<long double > res;
        if( dfs(num,res)){
            //for(int i=0;i<res.size();i++){
           // cout<<res[i]<<" ";
        //}
            return true;
        }
        return false;
    }

};

 */