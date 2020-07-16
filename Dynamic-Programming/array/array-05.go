package array

// 换钱的方法数

// 给定数组arr，arr中的所有值得=都是正数且不重复。
// 每个值代表一种面值的货币
// 每种面值的货币可以使用任意张
// 再给定一个正数aim代表要找的钱数
// 求换钱有多少种方法

// arr=[5,10,25,1],aim=0,方法数为1
// arr=[5,10,25,1],aim=15,方法数为6
// arr=[3,5],aim=2,方法数为0

// 参考01背包的问题
// dp[i][j]表示当前在arr[i]，aim为j的情况下的换货币数
// 每张货币都可以不用，用1张，用2张，...用j/arr[i]张

// dp[i][j] = dp[i-1][j] + dp[i-1][j-arr[i] +...+ dp[i-1][j-k*arr[i]] k = j/arr[i]

func GetChange(arr []int,aim int) int {
	dp := make([][]int,len(arr))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,aim+1)
	}

	for i := 0;i<len(arr);i++ {
		dp[i][0] = 1
	}

	for j := 1;j<=aim;j++ {
		if j%arr[0] == 0 {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}

	for i := 1;i<len(arr);i++ {
		for j:=1;j<=aim;j++ {
			for k := 0;k<=j/arr[i];k++ {
				dp[i][j] += dp[i-1][j-k*arr[i]]
			}
		}
	}
	return dp[len(arr)-1][aim]
}