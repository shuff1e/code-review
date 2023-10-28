package Other

/*
题目：
你拥有一些道具，万能道具，R级，SR级，SSR级道具， 稀有度SSR > SR > R。
万能道具可以当作R级，SR级，SSR级使用，1比1的汇率
现在想用道具换奖品， 比如A奖品， 需要消耗 12个R级道具，6个SR级道具， 4个SSR级道具

假如你拥有 23个万能道具，81个R级道具，62个SR级奖品，21个SSR级奖品。 根据拥有的道具数量计算出最多能换到几个奖品？



实现一个最小堆，
堆的元素为结构体 struct {
    base: A奖品需要多少张card
    count: card的数量
    score： card的数量/A奖品需要多少张card
}

比如对于R级卡
struct {
    base: 12
    count: 81
    score： 81/12
}

    for i := 0; i < 万能牌的数量; i++ {
        ele := 最小堆.pop()

        ele.count += 1
        ele.score = float64(ele.count) / float64(ele.base)

        heap.Push(h, ele) // 将该元素再重新放进最小堆
    }


    最小堆.pop().score 就是最多能兑换多少奖品
*/

type card struct {
	base  int
	count int
	score float64
}

type cardHeap []card

func (h cardHeap) Len() int           { return len(h) }
func (h cardHeap) Less(i, j int) bool { return h[i].score < h[j].score }
func (h cardHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *cardHeap) Push(x interface{}) {
	*h = append(*h, x.(card))
}

func (h *cardHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
