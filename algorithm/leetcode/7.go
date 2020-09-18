package main

func reverse(x int) int {
	const INTMAX = (1<<31)-1
	const INTMIN = -(1<<31)
	recv := 0
	for x != 0 {
		pop := x %10
		x /= 10
		if recv > INTMAX/10 || (recv == INTMAX/10 && pop == INTMAX%10) {
			return 0
		}
		if recv < INTMIN/10 || (recv == INTMIN/10 && pop == INTMIN%10) {
			return 0
		}
		recv = recv*10 + pop
	}
	return recv
}

func main() {
	println(reverse(123))
}