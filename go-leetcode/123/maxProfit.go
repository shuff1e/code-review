package main

import "fmt"

/*
我们不用递归思想进行穷举，而是利用「状态」进行穷举。
我们具体到每一天，看看总共有几种可能的「状态」，再找出每个「状态」对应的「选择」。
我们要穷举所有「状态」，穷举的目的是根据对应的「选择」更新状态。
听起来抽象，你只要记住「状态」和「选择」两个词就行，下面实操一下就很容易明白了。

for 状态1 in 状态1的所有取值：
	for 状态2 in 状态2的所有取值：
		for ...
			dp[状态1][状态2][...] = 择优(选择1，选择2...)

比如说这个问题，每天都有三种「选择」：买入、卖出、无操作，我们用 buy, sell, rest 表示这三种选择。
但问题是，并不是每天都可以任意选择这三种选择的，因为 sell 必须在 buy 之后，buy 必须在 sell 之后。
那么 rest 操作还应该分两种状态，一种是 buy 之后的 rest（持有了股票），一种是 sell 之后的 rest（没有持有股票）。
而且别忘了，我们还有交易次数 k 的限制，就是说你 buy 还只能在 k > 0 的前提下操作。

这个问题的「状态」有三个，第一个是天数，第二个是允许交易的最大次数，
第三个是当前的持有状态（即之前说的 rest 的状态，我们不妨用 1 表示持有，0 表示没有持有）。
然后我们用一个三维数组就可以装下这几种状态的全部组合：

dp[i][k][0 or 1]
0 <= i <= n-1, 1 <= k <= K
n 为天数，大 K 为最多交易数
此问题共 n × K × 2 种状态，全部穷举就能搞定。

for 0 <= i < n:
	for 1 <= k <= K:
		for s in {0, 1}:
			dp[i][k][s] = max(buy, sell, rest)

而且我们可以用自然语言描述出每一个状态的含义，比如说 dp[3][2][1] 的含义就是：
今天是第三天，我现在手上持有着股票，至今最多进行 2 次交易。再比如 dp[2][3][0] 的含义：
今天是第二天，我现在手上没有持有股票，至今最多进行 3 次交易。很容易理解，对吧？

我们想求的最终答案是 dp[n - 1][K][0]，即最后一天，最多允许 K 次交易，最多获得多少利润。
读者可能问为什么不是 dp[n - 1][K][1]？因为 [1] 代表手上还持有股票，[0] 表示手上的股票已经卖出去了，
很显然后者得到的利润一定大于前者。

dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
              max(   选择 rest  ,             选择 sell      )

解释：今天我没有持有股票，有两种可能：
要么是我昨天就没有持有，然后今天选择 rest，所以我今天还是没有持有；
要么是我昨天持有股票，但是今天我 sell 了，所以我今天没有持有股票了。

dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
              max(   选择 rest  ,           选择 buy         )

解释：今天我持有着股票，有两种可能：
要么我昨天就持有着股票，然后今天选择 rest，所以我今天还持有着股票；
要么我昨天本没有持有，但今天我选择 buy，所以今天我就持有股票了。
这个解释应该很清楚了，如果 buy，就要从利润中减去 prices[i]，如果 sell，就要给利润增加 prices[i]。
今天的最大利润就是这两种可能选择中较大的那个。而且注意 k 的限制，我们在选择 buy 的时候，把 k 减小了 1，
很好理解吧，当然你也可以在 sell 的时候减 1，一样的。

现在，我们已经完成了动态规划中最困难的一步：状态转移方程。**如果之前的内容你都可以理解，那么你已经可以秒杀所有问题了，
只要套这个框架就行了。**不过还差最后一点点，就是定义 base case，即最简单的情况。

dp[-1][k][0] = 0
解释：因为 i 是从 0 开始的，所以 i = -1 意味着还没有开始，这时候的利润当然是 0 。
dp[-1][k][1] = -infinity
解释：还没开始的时候，是不可能持有股票的，用负无穷表示这种不可能。
dp[i][0][0] = 0
解释：因为 k 是从 1 开始的，所以 k = 0 意味着根本不允许交易，这时候利润当然是 0 。
dp[i][0][1] = -infinity
解释：不允许交易的情况下，是不可能持有股票的，用负无穷表示这种不可能。
把上面的状态转移方程总结一下：


base case：
dp[-1][k][0] = dp[i][0][0] = 0
dp[-1][k][1] = dp[i][0][1] = -infinity

状态转移方程：
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
 */
























// dp[-1][k][0] = dp[i][0][0] = 0
// dp[-1][k][1] = dp[i][0][1] = -Inf

// dp[i][k][0] = max(dp[i-1][k][0],dp[i-1][k][1] + prices[i])
// dp[i][k][1] = max(dp[i-1][k][1],dp[i-1][k-1][0] - prices[i]

func main() {
	arr := []int{3,3,5,0,0,3,1,4}
	fmt.Println(maxProfit(arr))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxK := 2
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

func maxProfit2(prices []int) int {
	maxK := 2
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

func maxProfit3(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxK := 2
	dp := make([][][]int,len(prices))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([][]int,maxK + 1)
		for j :=0;j<len(dp[i]);j++ {
			dp[i][j] = make([]int,2)
		}
	}

	for k := 0;k<=maxK;k++ {
		dp[0][k][0] = 0
		dp[0][k][1] = -prices[0]
	}

	for i := 1;i<len(prices);i++ {
		for k := maxK;k>=1;k-- {
			// 这个地方不对，第K - 1次卖出使用之前不能第k次买入
			dp[i][k-1][1] = max(dp[i-1][k-1][1],dp[i-1][k-1][0] - prices[i])
			// k表示最多允许k次卖出
			dp[i][k][0] = max(dp[i-1][k][0],dp[i-1][k-1][1] + prices[i])
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
