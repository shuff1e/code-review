package main

import "fmt"

func romanToInt(s string) int {
	data := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL":40,
		"X":10,
		"IX":9,
		"V":5,
		"IV":4,
		"I":1,
	}
	result := 0
	for len(s) > 0 {
		fmt.Println(s)
		if len(s) >= 2 {
			if value,ok := data[s[len(s)-2:len(s)]];ok {
				result += value
				s = s[0:len(s)-2]
				continue
			}
		}

		if value,ok := data[s[len(s)-1:len(s)]];ok {
			result += value
			s = s[0:len(s)-1]
		}

	}
	return result
}

func main() {
	println(romanToInt("III"))
}
