package main

import "fmt"

/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

 */

// A：dp[i][j] = dp[i+1][j-1] + 2 if s[i] == s[j] && dp[i+1][j-1] > 0
//  else 0

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
func longestPalindrome(s string) string {
	result := ""
	dp := make([][]int,len(s))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,len(s))
	}

	for i := 0;i<len(dp);i++ {
		dp[i][i] = 1
		result = s[i:i+1]
	}
	for i :=0;i+1<len(dp);i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
			result = s[i:i+2]
		} else {
			dp[i][i+1] = 0
		}
	}

	for length := 2;length < len(s);length ++ {
		for i := 0;i+length<len(s);i++ {
			if s[i] == s[i+length] && dp[i+1][i+length-1] > 0 {
				dp[i][i+length] = dp[i+1][i+length-1] + 2
				result = s[i:i+length+1]
			} else {
				dp[i][i+length] = 0
			}
		}
	}
	return result
}

/*

方法二：中心扩展算法
思路与算法

我们仔细观察一下方法一中的状态转移方程：



P(i,i)      = true
P(i,i+1)    = (s[i] == s[i+1])
P(i,j)      = P(i+1,j-1) && (s[i] == s[j])
​

找出其中的状态转移链：

P(i,j)←P(i+1,j−1)←P(i+2,j−2)←⋯←某一边界情况

可以发现，所有的状态在转移的时候的可能性都是唯一的。也就是说，我们可以从每一种边界情况开始「扩展」，也可以得出所有的状态对应的答案。

边界情况即为子串长度为 1 或 2 的情况。
我们枚举每一种边界情况，并从对应的子串开始不断地向两边扩展。如果两边的字母相同，我们就可以继续扩展，
例如从 P(i+1,j−1) 扩展到 P(i,j)；如果两边的字母不同，我们就可以停止扩展，因为在这之后的子串都不能是回文串了。

聪明的读者此时应该可以发现，「边界情况」对应的子串实际上就是我们「扩展」出的回文串的「回文中心」。
方法二的本质即为：我们枚举所有的「回文中心」并尝试「扩展」，直到无法扩展为止，此时的回文串长度即为此「回文中心」下的最长回文串长度。
我们对所有的长度求出最大值，即可得到最终的答案。


class Solution {
    public String longestPalindrome(String s) {
        if (s == null || s.length() < 1) {
            return "";
        }
        int start = 0, end = 0;
        for (int i = 0; i < s.length(); i++) {
            int len1 = expandAroundCenter(s, i, i);
            int len2 = expandAroundCenter(s, i, i + 1);
            int len = Math.max(len1, len2);
            if (len > end - start) {
                start = i - (len - 1) / 2;
                end = i + len / 2;
            }
        }
        return s.substring(start, end + 1);
    }

    public int expandAroundCenter(String s, int left, int right) {
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            --left;
            ++right;
        }
        return right - left - 1;
    }
}

 */
