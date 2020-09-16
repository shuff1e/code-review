package main

import "fmt"

// 50（二）：字符流中第一个只出现一次的字符
// 题目：请实现一个函数用来找出字符流中第一个只出现一次的字符。例如，当从
// 字符流中只读出前两个字符"go"时，第一个只出现一次的字符是'g'。当从该字
// 符流中读出前六个字符"google"时，第一个只出现一次的字符是'l'。


















const SIZE rune = 256

type myStruct struct {
	// map
	// key就是index
	// data 的 value记录 char的位置
	data [SIZE]int
	index int
}

func New() *myStruct {
	data := [SIZE]int{}
	// 没有char，所以初始的位置就是-1
	// index为0
	for i := 0;i<len(data);i++ {
		data[i] = -1
	}
	return &myStruct{
		data: data,
		index: 0,
	}
}

func (m *myStruct) Insert(c rune) {
	if m.data[c] == -1 {
		m.data[c] = m.index
	} else {
		m.data[c] = -2
	}
	m.index ++
}

func (m *myStruct) getFirst() rune {
	// index是char，是key
	// value是出现的顺序
	result := '0'
	minOccur := 1<<31-1
	for i := rune(0);i<SIZE;i++ {
		if m.data[i] >= 0 {
			if m.data[i] < minOccur {
				minOccur = m.data[i]
				result = i
			}
		}
	}
	return result
}

func main() {
	chars := New()
	Test("Test1", chars, '0');

	chars.Insert('g');
	Test("Test2", chars, 'g');

	chars.Insert('o');
	Test("Test3", chars, 'g');

	chars.Insert('o');
	Test("Test4", chars, 'g');

	chars.Insert('g');
	Test("Test5", chars, '0');

	chars.Insert('l');
	Test("Test6", chars, 'l');

	chars.Insert('e');
	Test("Test7", chars, 'l');
	chars.Insert('0')
	Test("Test8", chars, 'l');
}

func Test(name string,m *myStruct,c rune) {
	fmt.Printf("%s %c %c\n",name,m.getFirst(),c)
	if m.getFirst() != c {
		panic("fuck")
	}
}