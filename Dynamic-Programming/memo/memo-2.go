package memo

// n的阶乘

func Fact(n int) int {
	memo := make([]int,n+1)
	return f2(n,memo)
}

func f2(n int ,memo []int) int{
	if n == 1 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = n*f2(n-1,memo)
	return memo[n]
}