package help

func merge(input []int,left ,middle,right int, temp []int) {
	start := left
	start1 := left
	end1 := middle
	start2 := middle + 1
	end2 := right
	for start1 <= end1 && start2 <= end2 {
		if input[start1] < input[start2] {
			temp[start] = input[start1]
			start1 ++
		} else {
			temp[start] = input[start2]
			start2 ++
		}
		start ++
	}
	for start1 <= end1 {
		temp[start] = input[start1]
		start1 ++
		start ++
	}
	for start2 <= end2 {
		temp[start] = input[start2]
		start2 ++
		start ++
	}
	for left <= right {
		input[left] = temp[left]
		left ++
	}
}

func mergeSort(input []int,left,right int,temp []int) {
	if left < right {
		middle := (left+right)/2
		mergeSort(input,left,middle,temp)
		mergeSort(input,middle+1,right,temp)
		merge(input,left,middle,right,temp)
	}
}
func MergeSort(input []int) {
	temp := make([]int,len(input))
	mergeSort(input,0,len(input)-1,temp)
}

func Max(x,y int) int{
	if x > y {
		return x
	}
	return y
}

func Min(x,y int) int {
	if x > y {
		return y
	}
	return x
}