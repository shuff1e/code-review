package main

import "fmt"

/*
87. 扰乱字符串
给定一个字符串 s1，我们可以把它递归地分割成两个非空子字符串，从而将其表示为二叉树。

下图是字符串 s1 = "great" 的一种可能的表示形式。

great
/    \
gr    eat
/ \    /  \
g   r  e   at
           / \
           a   t
在扰乱这个字符串的过程中，我们可以挑选任何一个非叶节点，然后交换它的两个子节点。

例如，如果我们挑选非叶节点 "gr" ，交换它的两个子节点，将会产生扰乱字符串 "rgeat" 。

rgeat
/    \
rg    eat
/ \    /  \
r   g  e   at
           / \
          a   t
我们将 "rgeat” 称作 "great" 的一个扰乱字符串。

同样地，如果我们继续交换节点 "eat" 和 "at" 的子节点，将会产生另一个新的扰乱字符串 "rgtae" 。

rgtae
/    \
rg    tae
/ \    /  \
r   g  ta  e
       / \
      t   a
我们将 "rgtae” 称作 "great" 的一个扰乱字符串。

给出两个长度相等的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。

示例 1:

输入: s1 = "great", s2 = "rgeat"
输出: true
示例 2:

输入: s1 = "abcde", s2 = "caebd"
输出: false
 */

func main() {
	fmt.Println(isScramble("abc","bca"))
	// abc acb
	// dbd
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	dp := make(map[point]bool)
	return isMatch(s1,s2,0,len(s1),0,len(s2),dp)
}

type point struct {
	start1 int
	end1 int
	start2 int
	end2 int
}

func isMatch(str1,str2 string,start1,end1,start2,end2 int,dp map[point]bool ) bool {
	dict := map[byte]int{}
	for i := 0;i<end1-start1;i++ {
		dict[str1[start1+i]] = dict[str1[start1+i]] + 1
		dict[str2[start2+i]] = dict[str2[start2+i]] - 1
	}
	for _,v := range dict {
		if v != 0 {
			return false
		}
	}
	if start1+1 == end1 {
		if str1[start1] == str2[start2] {
			return true
		}
		return false
	}
	if v,ok := dp[point{start1 ,end1,start2,end2}];ok {
		return v
	}

	result := false
	for length := 1;length < end1-start1;length++ {
		temp1 := isMatch(str1,str2,start1,start1+length,start2,start2+length,dp) &&
			isMatch(str1,str2,start1+length,end1,start2+length,end2,dp)

		temp2 := isMatch(str1,str2,start1,start1+length,end2-length,end2,dp) &&
			isMatch(str1,str2,start1+length,end1,start2,end2-length,dp)

		result = result || temp1 || temp2
		if result {
			dp[point{start1 ,end1,start2,end2}] = true
			return true
		}
	}
	dp[point{start1 ,end1,start2,end2}] = false
	return false
}

/*
还记得我们的定义吗？dp[i][j][len] 表示从
字符串 S 中 i 开始长度为 len 的字符串是否能变换为从字符串 T 中 j 开始长度为 len 的字符串，
所以答案是 dp[0][0][n]。 时间复杂度 O(N^4)，空间复杂度O(N^3)

class Solution {
    public boolean isScramble(String s1, String s2) {
        char[] chs1 = s1.toCharArray();
        char[] chs2 = s2.toCharArray();
        int n = s1.length();
        int m = s2.length();
        if (n != m) {
            return false;
        }
        boolean[][][] dp = new boolean[n][n][n + 1];
        // 初始化单个字符的情况
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                dp[i][j][1] = chs1[i] == chs2[j];
            }
        }

        // 枚举区间长度 2～n
        for (int len = 2; len <= n; len++) {
            // 枚举 S 中的起点位置
            for (int i = 0; i <= n - len; i++) {
                // 枚举 T 中的起点位置
                for (int j = 0; j <= n - len; j++) {
                    // 枚举划分位置
                    for (int k = 1; k <= len - 1; k++) {
                        // 第一种情况：S1 -> T1, S2 -> T2
                        if (dp[i][j][k] && dp[i + k][j + k][len - k]) {
                            dp[i][j][len] = true;
                            break;
                        }
                        // 第二种情况：S1 -> T2, S2 -> T1
                        // S1 起点 i，T2 起点 j + 前面那段长度 len-k ，S2 起点 i + 前面长度k
                        if (dp[i][j + len - k][k] && dp[i + k][j][len - k]) {
                            dp[i][j][len] = true;
                            break;
                        }
                    }
                }
            }
        }
        return dp[0][0][n];
    }
}
 */