package main

import "fmt"

/*

 */

func main() {
	o := Constructor(5)
	fmt.Println(o.Insert(3,"c"))
	fmt.Println(o.Insert(1,"a"))
	fmt.Println(o.Insert(2,"b"))
	fmt.Println(o.Insert(5,"e"))
	fmt.Println(o.Insert(4,"d"))
}

type OrderedStream struct {
	data []string
	index int
}


func Constructor(n int) OrderedStream {
	return OrderedStream{
		data: make([]string,n+1),
		index: 1,
	}
}


func (this *OrderedStream) Insert(id int, value string) []string {
	this.data[id] = value
	if id == this.index {
		for this.index < len(this.data) && this.data[this.index] != "" {
			this.index ++
		}
		result := []string{value}
		temp := id + 1
		for temp < len(this.data) && this.data[temp] != "" {
			result = append(result,this.data[temp])
			temp ++
		}
		return result
	}
	return []string{}
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(id,value);
 */