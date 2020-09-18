package main

import "fmt"

func letterCombinations(digits string) []string {
	data := map[string][]string{
		"2":[]string{"a","b","c"},
		"3":[]string{"d","e","f"},
		"4":[]string{"g","h","i"},
		"5":[]string{"j","k","l"},
		"6":[]string{"m","n","o"},
		"7":[]string{"p","q","r","s"},
		"8":[]string{"t","u","v"},
		"9":[]string{"w","x","y","z"},
	}
	result := []string{}
	for i := 0;i<len(digits);i++ {
		letters := data[string(digits[i])]
		result = getResult(result,letters)
	}
	return result
}

func getResult(initial []string ,letters []string ) []string {
	if len(initial) == 0 {
		return letters
	}
	result := []string{}
	for _,v := range initial {
		for _,temp := range letters {
			result = append(result,v+temp)
		}
	}
	return result
}

func main() {
	str := "23"
	fmt.Println(letterCombinations(str))
}