package math

const INTEGER_MAX_VALUE = 0x7fffffff
const INTEGER_MIN_VALUE = 0x80000000

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}
