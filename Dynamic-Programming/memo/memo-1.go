package memo

// 斐波那契数列
// 使用memo避免重复计算
// 时间复杂度O(n)，空间复杂度O(n)
func Fibonacci(n int) int {
	memo := make([]int,n+1)
	return f(n,memo)
}

func f(n int,memo []int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = f(n-1,memo) + f(n-2,memo)
	return memo[n]
}
