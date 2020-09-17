package main

import "fmt"

// 60：n个骰子的点数
// 题目：把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s
// 的所有可能的值出现的概率。

// A：点数的和从n到6*n
// 第一种思路，计算每个点数出现的次数，这样需要对每个点数，进行n层遍历
// 这种类似换人民币的方法数

// 第二种思路，n层遍历，遍历到最后，出现的点数保存在一个数组中，这样是走了每一个可能的路径
// 如果要打印每个可能的路径，则需要这样

// 第三种思路，可以看到第一种思路中的，第一层，可能的值为 6*(n-1)+1 -> 6*n
// 然后下一层可能的值为6*n-3，或者6*(n-1)+3，这两者计算的过程中是会重复的
// 因此可以加一个memo

// 第四种思路，递归+memo的话，就是dp了，
// dp关系为dp[i] = dp[i-6] +...+ dp[i-1]
// 由于值只和左边的值有关，因此也不用n*maxSum大小的matrix，每次去填表
// 只需要一个大小为1*maxSum的array就可以了

const global_max = 6

func getNumber1(n ,level ,parentSum int,sumArr []int) {
	if level == n + 1 {
		sumArr[parentSum-n] += 1
		return
	}
	for i := 1;i<=global_max;i++ {
		getNumber1(n,level + 1,parentSum + i,sumArr)
	}
}

// 一共n->global_max*n
// n对应0
// sum对应的位置为sum-n
func printProbability1(n int) []float64 {
	sumArr := make([]int,global_max*n-n+1)
	getNumber1(n,1,0,sumArr)
	totalSum := Pow(global_max,n)

	result := make([]float64,len(sumArr))
	for i := 0;i<len(sumArr);i++ {
		fmt.Printf("%d %f\n",i+n,float64(sumArr[i])/float64(totalSum))
		result[i] = float64(sumArr[i])/float64(totalSum)
	}
	return result
}

func Pow(base,exp int) int {
	result := 1
	for exp > 0 {
		if exp & 1 > 0 {
			result = result*base
		}
		base = base*base
		exp = exp >> 1
	}
	return result
}

func main() {
	arr1 := printProbability1(8)
	arr2 := printProbability2(8)
	fmt.Println(checkSame(arr1,arr2))
}

func checkSame(arr1,arr2 []float64) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i,v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}

func printProbability2(n int) []float64 {
	dp := make([]int,global_max*n+1)
	// row

	// 记得初始化
	for i := 1;i<= global_max;i++ {
		dp[i] = 1
	}

	for i := 2;i<=n;i++ {
		// col
		// dp[i][j] = dp[i-1][j-1] + dp[i-1][j-2] + ... + dp[i-1][j-global_max]
		for j := i*global_max;j>=i*1;j-- {
			dp[j] = 0
			for k := 1;k<=global_max;k++ {
				// 比如现在是第i层，则i-1层，值的范围最小只能是(i-1)*1，小于i-1的位置都是0
				if j-k >= i-1 {
					dp[j] += dp[j-k]
				}
			}
		}
	}
	result := make([]float64,global_max*n-n+1)
	resultIndex := 0
	totalSum := Pow(global_max,n)
	for i := n;i<=global_max*n;i++ {
		fmt.Printf("%d %f\n",i,float64(dp[i])/float64(totalSum))
		result[resultIndex] = float64(dp[i])/float64(totalSum)
		resultIndex ++
	}
	return result
}