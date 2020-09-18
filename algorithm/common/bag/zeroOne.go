package main

import "fmt"

func fuck(volume []int,weight []int,v_max int) int{
	dp := make([][]int,len(volume))

	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int,v_max+1)
	}

	for i := 0;i<len(volume);i++ {
		dp[i][0] = 0
	}

	for j := 0;j<=v_max;j++ {
		if j >= volume[0] {
			dp[0][j] = weight[0]
		} else {
			dp[0][j] = 0
		}
	}
	for i := 0;i<len(volume);i++ {
		// dp[i][v] = max(dp[i-1][v],dp[i-1][v-volume[i-1]]+weight[i])

		for i := 1;i<len(volume);i++ {
			for j := 1;j<=v_max;j++ {
				if j < volume[i] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = Max(dp[i-1][j],dp[i-1][j-volume[i]]+weight[i])
				}
			}
		}
	}
	return dp[len(volume)-1][v_max]
}

func Max(x,y int) int{
	if x > y{
		return x
	}
	return y
}

func fuck2(volume []int,weight []int,v_max int) int {
	array := make([]int,v_max+1)

	for j := 0;j<=v_max;j++ {
		if j >= volume[0] {
			array[j] = weight[0]
		} else {
			array[j] = 0
		}
	}
	//fmt.Printf("%#v\n",array)

	for i := 1;i<len(volume);i++ {
		for j := v_max;j>=0;j-- {
			if j < volume[i] {
				array[j] = array[j]
			} else {
				array[j] = Max(array[j-volume[i]] + weight[i],array[j])
			}
		}
		//fmt.Printf("%#v\n",array)
	}

	return array[v_max]
}

func main() {
	volume := []int{5,4,7,2,6}
	weight := []int{12,3,10,3,6}
	v_max := 10

	w_max := fuck(volume,weight,v_max)
	fmt.Println(w_max)

	w_max = fuck2(volume,weight,v_max)
	fmt.Println(w_max)

}