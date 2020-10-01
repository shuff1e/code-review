package main

import "fmt"

/*
52. N皇后 II
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。



上图为 8 皇后问题的一种解法。

给定一个整数 n，返回 n 皇后不同的解决方案的数量。

示例:

输入: 4
输出: 2
解释: 4 皇后问题存在如下两个不同的解法。
[
[".Q..",  // 解法 1
"...Q",
"Q...",
"..Q."],

["..Q.",  // 解法 2
"Q...",
"...Q",
".Q.."]
]


提示：

皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一或 N-1 步，可进可退。（引用自 百度百科 - 皇后 ）

 */

func main() {
	fmt.Println(totalNQueens(1))
}

func totalNQueens(n int) int {
	totalNum := 0
	help(n,0,&totalNum,0,0,0)
	return totalNum
}

func help(n int,level int,totalNum *int,colLimit,leftLimit,rightLimit int) {
	if level == n {
		*totalNum = *totalNum + 1
		return
	}
	positions := ( (1 << n) - 1) & (^(colLimit|leftLimit|rightLimit))
	for positions > 0 {
		rightMost := positions & (-positions)
		positions = positions & (positions - 1)
		help(n,level+1,totalNum,colLimit|rightMost,(leftLimit|rightMost) << 1,(rightLimit | rightMost) >> 1)
	}
}

