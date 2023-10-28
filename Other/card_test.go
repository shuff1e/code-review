package Other

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestCardHeap(t *testing.T) {
	h := &cardHeap{
		//R
		{
			base:  12,
			count: 81,
			score: 81 / 12,
		},
		//SR
		{
			base:  6,
			count: 62,
			score: 62 / 6,
		},
		//SSR
		{
			base:  4,
			count: 21,
			score: 21 / 4,
		},
	}
	heap.Init(h)

	for i := 0; i < 23; i++ {
		ele := heap.Pop(h).(card)
		ele.count += 1
		ele.score = float64(ele.count) / float64(ele.base)
		heap.Push(h, ele)
	}

	fmt.Println(h)

	fmt.Println(heap.Pop(h).(card).score)
}
