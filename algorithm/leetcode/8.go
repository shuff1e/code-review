package main

import (
	"fmt"
)

var record = map[string][]string{
	"start": []string{"start","signed","innumber","end"},
	"signed": []string{"end","end","innumber","end"},
	"innumber": []string{"end","end","innumber","end"},
	"end": []string{"end","end","end","end"},
}

const INTMAX = (1<<31)-1
const INTMIN = -(1<<31)

func min(x,y int) int {
	if x >y {
		return y
	}
	return x
}

type automate struct {
	positive int
	ans int
	state string
}

func (auto *automate) get(c byte) {
	state := record[auto.state][auto.getcol(c)]
	auto.state = state
	if state == "innumber" {
		auto.ans = auto.ans*10 + int(c - '0')
		if auto.positive == 1 {
			auto.ans = min(auto.ans,INTMAX)
		} else {
			auto.ans = min(auto.ans,-INTMIN)
		}
	}
	if state == "signed" {
		if c == '+' {
			auto.positive = 1
		} else {
			auto.positive = -1
		}

	}

}

func (auto *automate) getcol(c byte ) int{
	if c == ' ' {
		return 0
	} else if c == '+' || c == '-' {
		return 1
	} else if c >= '0' && c <= '9' {
		return 2
	} else {
		return 3
	}
}

func myAtoi(str string) int {
	auto := automate{1,0,"start"}
	for i := 0;i<len(str);i++ {
		auto.get(str[i])
	}
	return auto.positive* auto.ans
}

func main() {
	fmt.Println(myAtoi("2147483648"))
}
