package main

func lengthOfLongestSubstring(s string) int {
	i := 0
	j := 0
	ans := 0
	window := map[byte]int{}

	for i < len(s) && j < len(s) {
		if index,ok := window[s[j]];ok {
			i = Max(index+1,i)

		}
		window[s[j]] = j
		j += 1
		ans = Max(ans,j-i)
	}
	return ans
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
