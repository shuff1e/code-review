package main

import "fmt"

/*

322. 零钱兑换
给定不同面额的硬币 coins 和一个总金额 amount。
编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

你可以认为每种硬币的数量是无限的。

示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0
示例 4：

输入：coins = [1], amount = 1
输出：1
示例 5：

输入：coins = [1], amount = 2
输出：2


提示：

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104

 */

// f(n) = Min(f(n-k)+1) k属于arr

func main() {
	coins := []int{86,210,29,22,402,140,16,466}
	amount := 3219
	result := coinChange(coins,amount)
	fmt.Println(result)
	fmt.Println(coinChangeBetter(coins,amount))
}

func coinChange(coins []int, amount int) int {
	memo := make([][]int,len(coins))
	for i := 0;i<len(memo);i++ {
		memo[i] = make([]int,amount+1)
		for j := 0;j<len(memo[i]);j++ {
			memo[i][j] = -1
		}
	}

	help(coins,0,amount,memo)
	return memo[0][amount]
}

func help(coins []int,index int,amount int,memo [][]int) int {
	if index == len(coins) {
		if amount == 0 {
			return 0
		}
		return -1
	}

	if amount < 0 {
		return -1
	}

	if memo[index][amount] != -1 {
		return memo[index][amount]
	}

	if amount == 0 {
		memo[index][amount] = 0
		return 0
	}


	result := 0x7fffffff

	// 消耗掉当前位置，而且我一张不用
	temp := help(coins,index+1,amount,memo)
	if temp >= 0 {
		result = Min(result,temp)
	}
	// 不消耗当前位置，而且我用一张
	if amount >= coins[index] {
		temp = help(coins,index,amount-coins[index],memo)
		if temp >= 0 {
			result = Min(result,temp+1)
		}
	}

	if result == 0x7fffffff {
		memo[index][amount] = -1
		return -1
	}
	memo[index][amount] = result
	return result
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

// dp[i][j] = Min(dp[i-1][j-k*arr[i]]+k) (0<=k<=j/arr[i])
//          = Min(dp[i-1][j], Min(dp[i-1][j-k*arr[i]]+k)) (1<=k<=j/arr[i])
//          												// 将k替换为y+1
//          = Min(dp[i-1][j], Min(dp[i-1][j-arr[i]-y*arr[i]]+y+1)) (0<=y<=j/arr[i]-1)
//
//
// dp[i][j-arr[i]] = Min(dp[i-1][j-arr[i]-y*arr[i]]+y) (0<=y<=j/arr[i]-1)
//
// 则，dp[i][j]= Min(dp[i-1][j],dp[i-1][j-arr[i]] + 1)

// 当前index位置的钱一张不用，并且消耗掉当前index
// 或者当前index位置的钱用掉一张，但是不消耗掉当前index(当前需要amount>=arr[index])

// 换钱的方法数
// dp[i][j] = Min(dp[i-1][j],dp[i][j-arr[i]]+1)

func coinChangeBetter(coins []int, amount int) int {
	arr := coins
	const fuck = 0x7fffffff

	dp := make([][]int,len(coins))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,amount+1)
		for j := 0;j<amount+1;j++ {
			dp[i][j] = fuck
		}
	}

	for i := 0;i<len(dp);i++ {
		dp[i][0] = 0
	}

	for j := 0;j<amount+1;j++ {
		if j >= arr[0] {
			if dp[0][j-arr[0]] != fuck {
				dp[0][j] = dp[0][j-arr[0]] + 1
			}
		}
	}
	for i := 1;i<len(dp);i++ {
		for j := 1;j<len(dp[0]);j++ {
			dp[i][j] = dp[i-1][j]
			if j >= arr[i] && dp[i][j-arr[i]] != fuck {
				dp[i][j] = Min(dp[i][j],dp[i][j-arr[i]]+1)
			}
		}
	}
	if dp[len(coins)-1][amount] == fuck {
		return -1
	}
	return dp[len(coins)-1][amount]
}

/*

public class Solution {
    public int coinChange(int[] coins, int amount) {
        return coinChange(0, coins, amount);
    }

    private int coinChange(int idxCoin, int[] coins, int amount) {
        if (amount == 0) {
            return 0;
        }
        if (idxCoin < coins.length && amount > 0) {
            int maxVal = amount / coins[idxCoin];
            int minCost = Integer.MAX_VALUE;
            for (int x = 0; x <= maxVal; x++) {
                if (amount >= x * coins[idxCoin]) {
                    int res = coinChange(idxCoin + 1, coins, amount - x * coins[idxCoin]);
                    if (res != -1) {
                        minCost = Math.min(minCost, res + x);
                    }
                }
            }
            return (minCost == Integer.MAX_VALUE)? -1: minCost;
        }
        return -1;
    }
}

// Time Limit Exceeded




public class Solution {
    public int coinChange(int[] coins, int amount) {
        if (amount < 1) {
            return 0;
        }
        return coinChange(coins, amount, new int[amount]);
    }

    private int coinChange(int[] coins, int rem, int[] count) {
        if (rem < 0) {
            return -1;
        }
        if (rem == 0) {
            return 0;
        }
        if (count[rem - 1] != 0) {
            return count[rem - 1];
        }
        int min = Integer.MAX_VALUE;
        for (int coin : coins) {
            int res = coinChange(coins, rem - coin, count);
            if (res >= 0 && res < min) {
                min = 1 + res;
            }
        }
        count[rem - 1] = (min == Integer.MAX_VALUE) ? -1 : min;
        return count[rem - 1];
    }
}




public class Solution {
    public int coinChange(int[] coins, int amount) {
        int max = amount + 1;
        int[] dp = new int[amount + 1];
        Arrays.fill(dp, max);
        dp[0] = 0;
        for (int i = 1; i <= amount; i++) {
            for (int j = 0; j < coins.length; j++) {
                if (coins[j] <= i) {
                    dp[i] = Math.min(dp[i], dp[i - coins[j]] + 1);
                }
            }
        }
        return dp[amount] > amount ? -1 : dp[amount];
    }
}

 */

// 如果是硬币只能选一次
// dp[i][j] = Min(dp[i-1][j],dp[i-1][j-arr[i]]+1)