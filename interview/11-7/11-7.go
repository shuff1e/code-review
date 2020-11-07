package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// A ：1到n这个n个数字，按字典序排序

// 比如输入 13
// 返回 1,10,11,12,13,2,3,4,5,6,7,8,9

func main() {
	n := 5345
	//n = 9
	result := getArr(n)

	for i := 1;i<len(result);i++ {
		a := strconv.Itoa(result[i-1])
		b := strconv.Itoa(result[i])
		a = "125"
		b = "1234"
		if strings.Compare(a,b) != -1 {
			panic("err: " + a + " >= " + b)
		}
	}

	fmt.Printf("%#v\n",result)

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	fmt.Println(len(result))

	compare := []int{}
	for i := 1;i <= n;i++ {
		compare = append(compare,i)
	}
	fmt.Println(reflect.DeepEqual(result,compare))
}

func getArr(n int) []int {
	result := []int{}
	if n < 10 {
		for i := 1;i<=n;i++ {
			result = append(result,i)
		}
		return result
	}


	str := strconv.Itoa(n)

	msb := int(str[0] - '0')
	length := len(str)

	for i := 1;i<msb;i++ {
		// prefix，剩下的位数
		getArrHelp(i,length-1,&result)
	}

	getArrHelp2(msb,length-1,n,&result)

	for i := msb + 1;i<10;i++ {
		getArrHelp(i,length-2,&result)
	}
	return result
}

// 任意扩展，不需要有顾虑
func getArrHelp(prefix int,cnt int,result *[]int) {
	if cnt == 0 {
		*result = append(*result,prefix)
		return
	}
	*result = append(*result,prefix)
	for i := 0;i<10;i++ {
		getArrHelp(prefix * 10 + i,cnt - 1,result)
	}
}

func getArrHelp2(prefix int,cnt int,n int,result *[]int) {
	if cnt == 0 {
		if prefix <= n {
			*result = append(*result,prefix)
		}
		return
	}

	*result = append(*result,prefix)

	if prefix > n / Pow(10,cnt) {
		if cnt >= 2 {
			for i := 0;i<10;i++ {
				getArrHelp2(prefix * 10 + i,cnt - 2,n,result)
			}
		}
		return
	}

	for i := 0;i<10;i++ {
		getArrHelp2(prefix * 10 + i,cnt - 1,n,result)
	}
}

func Pow(base ,exp int) int {
	result := 1
	for exp > 0 {
		if exp & 1 > 0 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

// 使用子节点个数为10的字典树?
