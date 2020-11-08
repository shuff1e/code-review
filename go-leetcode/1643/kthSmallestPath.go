package main

import (
	"fmt"
	"strings"
)

/*

1643. 第 K 条最小指令
Bob 站在单元格 (0, 0) ，想要前往目的地 destination ：(row, column) 。他只能向 右 或向 下 走。你可以为 Bob 提供导航 指令 来帮助他到达目的地 destination 。

指令 用字符串表示，其中每个字符：

'H' ，意味着水平向右移动
'V' ，意味着竖直向下移动
能够为 Bob 导航到目的地 destination 的指令可以有多种，例如，如果目的地 destination 是 (2, 3)，"HHHVV" 和 "HVHVH" 都是有效 指令 。

然而，Bob 很挑剔。因为他的幸运数字是 k，他想要遵循 按字典序排列后的第 k 条最小指令 的导航前往目的地 destination 。k  的编号 从 1 开始 。

给你一个整数数组 destination 和一个整数 k ，请你返回可以为 Bob 提供前往目的地 destination 导航的 按字典序排列后的第 k 条最小指令 。



示例 1：



输入：destination = [2,3], k = 1
输出："HHHVV"
解释：能前往 (2, 3) 的所有导航指令 按字典序排列后 如下所示：
["HHHVV", "HHVHV", "HHVVH", "HVHHV", "HVHVH", "HVVHH", "VHHHV", "VHHVH", "VHVHH", "VVHHH"].
示例 2：



输入：destination = [2,3], k = 2
输出："HHVHV"
示例 3：



输入：destination = [2,3], k = 3
输出："HHVVH"


提示：

destination.length == 2
1 <= row, column <= 15
1 <= k <= nCr(row + column, row)，其中 nCr(a, b) 表示组合数，即从 a 个物品中选 b 个物品的不同方案数。

 */

/*

方法一：优先确定高位 + 组合计数
思路与算法

当字符串中每一种字符的数量固定时（例如对于本题，我们需要在字符串中放入 h 个 H 和 v 个 V，如果需要求出字典序第 k小的字符串，可以考虑从高位高位向低位依次确定每一个位置的字符。

如果我们在最高位放置了 H，那么剩余的 (h−1,v) 就是一个规模减少的相同问题；同理如果我们在最高位放置了 V，
那么剩余的 (h,v−1) 也是一个规模减少的相同问题。
V。由于后者的字典序较大，因此如果最高位放 V，那么所有最高位为 H 的字符串的字典序都比它小，这样的字符串共有


o=( h−1
    h+v−1
)

个。也就是确定了最高位为 H，剩余 h+v−1 个位置中选择 h−1 个放入 H，其余位置自动放入 V 的方案数。因此：

如果 k 大于这个组合数 o，那么最高位一定是 V。我们将 v 减少 11，并且需要将 k 减少 o，这是因为剩余部分应当是包含 (h,v−1) 的字典序第 k−o 小的字符串；

如果 k 小于 o，那么最高位是 H。我们将 h 减少 1，但我们不需要改变 k 的值，这是因为剩余部分就是包含 (h−1,v) 的字典序第 k 小的字符串。

这样一来，我们就可以从高位开始，依次确定每一个位置的字符了。需要注意的是，当 h=0 时，我们只能放 V，无需进行判断。

代码

对于 Python 语言，可以使用 math.comb() 方便地求出组合数。但对于 C++ 而言，由于本题会导致乘法溢出，因此可以考虑使用组合数的递推式

c[n][k] = c[n-1][k-1] + c[n-1][k]

预处理处所有可能需要用到的组合数。

本题中，可能需要计算的最大组合数为
( 14
29
​)，在 C++ 语言中，直接通过先乘法后除法的方法计算该组合数，在乘法过程中就会超出 64 位无符号整数的上限。





class Solution {
public:
    string kthSmallestPath(vector<int>& destination, int k) {
        int h = destination[1];
        int v = destination[0];

        // 预处理组合数
        vector<vector<int>> comb(h + v, vector<int>(h));
        comb[0][0] = 1;
        for (int i = 1; i < h + v; ++i) {
            comb[i][0] = 1;
            for (int j = 1; j <= i && j < h; ++j) {
                comb[i][j] = comb[i - 1][j - 1] + comb[i - 1][j];
            }
        }

        string ans;
        for (int i = 0, imax = h + v; i < imax; ++i) {
            if (h > 0) {
                int o = comb[h + v - 1][h - 1];
                if (k > o) {
                    ans += 'V';
                    --v;
                    k -= o;
                }
                else {
                    ans += 'H';
                    --h;
                }
            }
            else {
                ans += 'V';
                --v;
            }
        }
        return ans;
    }
};

 */

func main()  {
	fmt.Println(kthSmallestPath2([]int{2,3},6))
}

func kthSmallestPath2(destination []int, k int) string {

	// 打表
	h,v := destination[1],destination[0]

	comb := make([][]int,h+v)
	for i := 0;i<len(comb);i++ {
		comb[i] = make([]int,h)
	}

	// comb[i][j] 表示一共i个位置，j个位置放h
	comb[0][0] = 1
	for i := 1;i<len(comb);i++ {
		comb[i][0] = 1
		for j := 1;j<=i && j <h;j++ {
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
		}
	}

	result := ""
	for i := 0;i<len(comb);i++ {
		if h > 0 {
			o := comb[h-1+v][h-1]
			if k > o {
				result += "V"
				v --
				k -= o
			} else {
				result += "H"
				h --
			}
		} else {
			result += "V"
			v--
		}
	}
	return result
}

func kthSmallestPath(destination []int, k int) string {
	temp := []string{}
	result := help(destination[1],destination[0],&k,&temp)
	return result
}

func help(hcnt,vcnt int,count *int,temp *[]string) string {
	if hcnt == 0 && vcnt == 0 {
		*count = *count - 1
		if *count == 0 {
			return strings.Join(*temp,"")
		} else {
			return ""
		}
	}

	if hcnt > 0 {
		*temp = append(*temp,"H")
		result := help(hcnt - 1,vcnt,count,temp)
		*temp = (*temp)[:len(*temp)-1]
		if result != "" {
			return result
		}
	}

	if vcnt > 0 {
		*temp = append(*temp,"V")
		result := help(hcnt,vcnt-1,count,temp)
		*temp = (*temp)[:len(*temp)-1]
		return result
	}

	return ""
}
