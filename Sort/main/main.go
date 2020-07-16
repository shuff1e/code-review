package main

import (
	"fmt"
	"github.com/shuff1e/code-review/Sort"
	"github.com/shuff1e/code-review/util"
)

func main() {
	// partition1
	arr := util.GenerateSlice(10)
	fmt.Printf("%#v\n",arr)
	Sort.QuickSortNetherlandsFlag(arr)
	fmt.Printf("%#v\n",arr)
}
