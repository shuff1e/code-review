package main

import "fmt"

type stack []string

func NewStack() stack {
	return stack([]string{})
}

func (s *stack) push(str string) {
	*s = append(*s,str)
}

func (s *stack) pop() string {
	result := ""
	if len(*s) > 0 {
		result = (*s)[len(*s)-1]
		*s = (*s)[0:len(*s)-1]
	}
	return result
}

func isValid(s string) bool {
	data := map[string]string{
		"}":"{",
		"]":"[",
		")":"(",
	}

	st := NewStack()
	fmt.Println(&st)

	for i := 0;i<len(s);i ++ {
		if temp,ok := data[string(s[i])];ok {
			if st.pop() != temp {
				return false
			}
		} else {
			st.push(string(s[i]))
		}
	}
	return len(st) == 0
}

func main() {
	str := "()"
	fmt.Println(isValid(str))
}