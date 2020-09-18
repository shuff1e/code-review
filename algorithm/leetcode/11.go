package main

func maxArea(height []int) int {
	a,b := 0,0
	max := 0
	for i:= 0;i<len(height)-1;i++ {
		for j:=i+1;j<len(height);j++ {
			area := (j-i)*Min(height[i],height[j])
			max = mymyMax(max,area)
			if max == area {
				a=i
				b=j
			}

		}
	}
	println(a,b)
	return max
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func mymyMax(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	a := []int{1,8,6,2,5,4,8,3,7}
	println(maxArea(a))
}