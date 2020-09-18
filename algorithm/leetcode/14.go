package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	result := ""

	OUTER:
	for i := 0;;i++ {
		if len(strs[0])-1 < i {
			break
		}
		temp := strs[0][i]
		for j := 1;j<len(strs);j++ {
			if len(strs[j])-1 < i {
				break OUTER
			}
			if temp != strs[j][i] {
				break OUTER
			}
		}
		result += string(temp)
	}

	return result
}

func main() {
	strs := []string{"flower","flow","flight"}
	println(longestCommonPrefix(strs))
}