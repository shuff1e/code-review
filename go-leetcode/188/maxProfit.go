package main

/*
188. 买卖股票的最佳时机 IV
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [2,4,1], k = 2
输出: 2
解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2:

输入: [3,2,6,5,0,3], k = 2
输出: 7
解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
 */

func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxK := k
	dp := make([][][]int,len(prices))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([][]int,maxK + 1)
		for j :=0;j<len(dp[i]);j++ {
			dp[i][j] = make([]int,2)
		}
	}

	for i := 0;i<len(prices);i++ {
		for k := 1;k<=maxK;k++ {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[0]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0],dp[i-1][k][1] + prices[i])
			dp[i][k][1] = max(dp[i-1][k][1],dp[i-1][k-1][0] - prices[i])
		}
	}
	return dp[len(prices)-1][maxK][0]
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit2(k int, prices []int) int {
	maxK := k
	if len(prices) <= 1 {
		return 0
	}
	if maxK >= len(prices)/2 {
		dp_0,dp_1 := 0,-prices[0]
		for i := 1;i<len(prices);i++ {
			dp_0 = max(dp_0,dp_1 + prices[i])
			dp_1 = max(dp_1,dp_0 - prices[i])
		}
		return dp_0
	}

	dp := make([][]int,maxK+1)
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,2)
	}


	for i := 0;i<len(prices);i++ {
		for k := maxK;k >= 1;k-- {
			if i == 0 {
				dp[k][0] = 0
				dp[k][1] = -prices[0]
			}
			dp[k][0] = max(dp[k][0],dp[k][1] + prices[i])
			dp[k][1] = max(dp[k][1],dp[k-1][0] - prices[i])
		}
	}
	return dp[maxK][0]
}

/*
递归

class Solution {
    public int maxProfit(int k, int[] prices) {
        if(prices==null || prices.length==0) {
            return 0;
        }
        return dfs(0,0,0,k,prices);
    }
    //计算k次交易，index表示当前是哪天，status是买卖状态，coutnt为交易次数
    private int dfs(int index, int status, int count, int k, int[] prices) {
        if(index==prices.length || count==k) {
            return 0;
        }
        int a=0,b=0,c=0;
        //保持不动
        a = dfs(index+1,status,count,k,prices);
        if(status==1) {
            //卖一股，并将交易次数+1
            b = dfs(index+1,0,count+1,k,prices)+prices[index];
        } else {
            //买一股
            c = dfs(index+1,1,count,k,prices)-prices[index];
        }
        return Math.max(Math.max(a,b),c);
    }
}
 */