package main

func myMaxArea(height []int) int {
	l,r := 0,len(height)-1
	ans := 0

	for l <r {
		area := Min(height[l],height[r])*(r-l)
		ans = Max(ans,area)
		if height[l] <= height[r] {
			l++
		} else {
				r--
		}
	}
	return ans
}

func main() {

}