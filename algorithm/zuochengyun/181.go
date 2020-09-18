package main

func f1(n int) int {
	if n < 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return f1(n-1) + f1(n-2)
}

func f2(n int) int{
	if n <1 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	pre := 1
	res := 1
	temp := 0
	for i :=3 ;i <=n;i++ {
		temp = res
		res = res + pre
		pre = temp
	}
	return res

}

func f3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	base := [][]int{{1,1},{1,0}}
	res := matrixPower(base,n-2)
	return res[0][0] + res[1][0]
}

func matrixPower(m [][]int ,p int) [][]int {
	result := make([][]int,len(m))
	for i := 0;i<len(m);i++ {
		result[i] = make([]int,len(m))
		result[i][i] = 1
	}

	temp := m
	for ;p>0;p>>=1 {
		if ((p&1)!=0) {
			result = muliMatrix(result,temp)
		}
		temp = muliMatrix(temp,temp)
	}
	return result
}

func muliMatrix(a [][]int,b [][]int) [][]int {
	result := make([][]int,len(a))
	for i := 0;i<len(a);i++ {
		result[i] = make([]int,len(b[0]))
	}

	for i:=0;i<len(a);i++ {
		for j:=0;j<len(b[0]);j++ {
			for k :=0;k<len(a[0]);k++ {
				result[i][j] += a[i][k]*b[k][j]
			}
		}
	}

	return result
}
func f4(n int) int {
	if  n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 3
	}

	one := 1
	two := 2
	three := 3
	result := 0

	for i := 0;i<n-3;i++ {
		result = three + one
		one = two
		two = three
		three = result
	}
	return result
}

func f5(n int) int {
	base := [][]int{
		{1,1,0},
		{0,0,1},
		{1,0,0},
	}
	res := matrixPower(base,n-3)
	return 3*res[0][0]+2*res[1][0]+res[2][0]
}

func main()  {
	n := 6
	println(f1(n))
	println(f2(n))
	println(f3(n))
	println(f4(n))
	println(f5(n))
}
