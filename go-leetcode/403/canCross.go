package main

import (
	"fmt"
)

/*

403. 青蛙过河
一只青蛙想要过河。 假定河流被等分为 x 个单元格，
并且在每一个单元格内都有可能放有一石子（也有可能没有）。 青蛙可以跳上石头，但是不可以跳入水中。

给定石子的位置列表（用单元格序号升序表示），
请判定青蛙能否成功过河（即能否在最后一步跳至最后一个石子上）。
开始时， 青蛙默认已站在第一个石子上，并可以假定它第一步只能跳跃一个单位（即只能从单元格1跳至单元格2）。

如果青蛙上一步跳跃了 k 个单位，
那么它接下来的跳跃距离只能选择为 k - 1、k 或 k + 1个单位。 另请注意，青蛙只能向前方（终点的方向）跳跃。

请注意：

石子的数量 ≥ 2 且 < 1100；
每一个石子的位置序号都是一个非负整数，且其 < 231；
第一个石子的位置永远是0。
示例 1:

[0,1,3,5,6,8,12,17]

总共有8个石子。
第一个石子处于序号为0的单元格的位置, 第二个石子处于序号为1的单元格的位置,
第三个石子在序号为3的单元格的位置， 以此定义整个数组...
最后一个石子处于序号为17的单元格的位置。

返回 true。即青蛙可以成功过河，按照如下方案跳跃：
跳1个单位到第2块石子, 然后跳2个单位到第3块石子, 接着
跳2个单位到第4块石子, 然后跳3个单位到第6块石子,
跳4个单位到第7块石子, 最后，跳5个单位到第8个石子（即最后一块石子）。
示例 2:

[0,1,2,3,4,8,9,11]

返回 false。青蛙没有办法过河。
这是因为第5和第6个石子之间的间距太大，没有可选的方案供青蛙跳跃过去。

 */

func main() {
	arr := []int{0,1,3,5,6,8,12,17}
	fmt.Println(canCross3(arr))

	arr = []int{0,1,2,3,4,8,9,11}
	fmt.Println(canCross3(arr))
}

func canCross(stones []int) bool {
	memo := make([][]int,len(stones))
	for i := 0;i<len(memo);i++ {
		memo[i] = make([]int,len(stones))
		for j := 0;j<len(stones);j++ {
			memo[i][j] = -1
		}
	}
	result := help(stones,0,0,memo)
	return result == 1
}

func help(stones []int,index int,jumpsize int,memo [][]int) int {
	if index == len(stones)-1 {
		return 1
	}
	if memo[index][jumpsize] != -1 {
		return memo[index][jumpsize]
	}
	for i := index + 1;i<len(stones);i++ {
		gap := stones[i] - stones[index]
		if gap >= jumpsize-1 && gap <= jumpsize + 1 {
			if help(stones,i,gap,memo) == 1{
				memo[index][jumpsize] = 1
				return 1
			}
		}
	}
	return 0
}

func canCross2(stones []int) bool {
	dp := make([][]bool,len(stones))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]bool,len(stones) + 1)
	}

	dp[0][0] = true
	for i := 1;i<len(stones);i++ {
		for j := 0;j<i;j++ {
			gap := stones[i] - stones[j]
			if gap <= i {
				dp[i][gap] = dp[j][gap+1] || dp[j][gap] || dp[j][gap-1]
				if i == len(stones)-1 && dp[i][gap] {
					return true
				}
			}
		}
	}
	return false
}

func canCross3(stones []int) bool {
	dict := make(map[int]map[int]struct{})

	for i := 0;i<len(stones);i++ {
		dict[stones[i]] = make(map[int]struct{},0)
	}

	dict[stones[0]][0] = struct{}{}

	for i := 0;i<len(stones);i++ {
		for k,_ := range dict[stones[i]] {
			for step := k-1;step<=k+1;step++ {
				if step > 0 {
					if _,ok := dict[stones[i]+step];ok {
						dict[stones[i]+step][step] = struct{}{}
					}
				}
			}
		}
	}
	return len(dict[stones[len(stones)-1]]) > 0
}

/*

思路①、使用二维数组的动态规划
         动态规划
         dp[i][k] 表示能否由前面的某一个石头 j 通过跳 k 步到达当前这个石头 i ，这个 j 的范围是 [1, i - 1]
         当然，这个 k 步是 i 石头 和 j 石头之间的距离
         那么对于 j 石头来说，跳到 j 石头的上一个石头的步数就必须是这个 k - 1 || k || k + 1
         由此可得状态转移方程：dp[i][k] = dp[j][k - 1] || dp[j][k] || dp[j][k + 1]

class Solution {
    public boolean canCross(int[] stones) {

        int len = stones.length;

        if(stones[1] != 1){
            return false;
        }

        boolean[][] dp = new boolean[len][len + 1];
        dp[0][0] = true;
        for(int i = 1; i < len; i++){
            for(int j = 0; j < i; j++){
                int k = stones[i] - stones[j];
                //
                //	为什么有这么个判断？
                //	因为其他石头跳到第 i 个石头跳的步数 k 必定满足 k <= i
                //	这又是为什么？
                //	1、比如 nums = [0,1,3,5,6,8,12,17]
                //	   那么第 0 个石头跳到第 1 个石头，步数肯定为 1，然后由于后续最大的步数是 k + 1，因此第 1 个石头最大只能跳 2 个单位
                //	   因此如果逐个往上加，那么第 2 3 4 5 ... 个石头最多依次跳跃的步数是 3 4 5 6...
                //	2、 第 i 个石头能跳的最大的步数是 i + 1，那么就意味着其他石头 j 跳到第 i 个石头的最大步数只能是 i 或者 j + 1
                //	   而 这个 k 是其他石头跳到 i 石头上来的，因此 k 必须 <= i （或者是 k <= j + 1）
                //
                if(k <= i){
                    dp[i][k] = dp[j][k - 1] || dp[j][k] || dp[j][k + 1];
                    //提前结束循环直接返回结果
                    if(i == len - 1 && dp[i][k]){
                        return true;
                    }
                }
            }
        }
        return false;
    }
}

 */

/*

public class Solution {
    public boolean canCross(int[] stones) {
        HashMap<Integer, Set<Integer>> map = new HashMap<>();
        for (int i = 0; i < stones.length; i++) {
            map.put(stones[i], new HashSet<Integer>());
        }
        map.get(0).add(0);
        for (int i = 0; i < stones.length; i++) {
            for (int k : map.get(stones[i])) {
                for (int step = k - 1; step <= k + 1; step++) {
                    if (step > 0 && map.containsKey(stones[i] + step)) {
                        map.get(stones[i] + step).add(step);
                    }
                }
            }
        }
        return map.get(stones[stones.length - 1]).size() > 0;
    }
}

 */