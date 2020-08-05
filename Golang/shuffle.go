package Golang

import "math/rand"

// fisher-yates 算法
// 每次随机挑选一个值，放在数组末尾
// 然后在n-1个元素的数组中再随机挑选一个值
// 放在数组末尾
// 以此类推

func Shuffle(indexes []int) {
	for i := len(indexes);i>0;i-- {
		lastIdx := i-1
		idx := rand.Intn(i)
		indexes[lastIdx],indexes[idx] = indexes[idx],indexes[lastIdx]
	}
}
