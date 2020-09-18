package main

func isPalindrome(x int) bool {
	temp := x
	if x < 0 {
		return false
	}

	y := 0
	for x != 0 {
		y = y*10+x%10
		x /= 10
	}

	return temp == y
}

func main() {

	println(isPalindrome(121))
}