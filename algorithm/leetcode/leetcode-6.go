package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	if len(s) <= numRows {
		return s
	}
	result := make([][]byte,numRows)
	change := false
	row := 1
	hello := 1
	result[0] = append(result[0],s[0])
	for i := 1;i<len(s);i++ {
		result[row] = append(result[row],s[i])

		if row%(numRows-1) == 0 {
			change = !change
		}
		if change {
			hello = -1
		} else {
			hello = 1
		}

		row += hello

	}
	ans := ""
	for i:=0;i<len(result);i++ {
		ans += string(result[i])
	}
	return ans
}

func main() {
	fmt.Println(convert("LEETCODEISHIRING",40))
}
