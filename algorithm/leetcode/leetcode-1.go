package main

func twoSum(nums []int, target int) []int {
	// value -> index
	m := map[int]int{}
	for i,v := range nums {
		other := target -v
		if index,ok := m[other];ok {
			return []int{i,index}
		} else {
			m[v] = i
		}
	}
	return []int{}
}