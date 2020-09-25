package _23

/*
223. 矩形面积
在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。

每个矩形由其左下顶点和右上顶点坐标表示，如图所示。
Rectangle Area
示例:

输入: -3, 0, 3, 4, 0, -1, 9, 2
输出: 45
说明: 假设矩形面积不会超出 int 的范围。

   -----------------(3,4)
   |		 |------|-----------------------(9,2)
   |---------|------|					     |
(-3,0)       |-------------------------------|
		  (0,-1)
 */

// A：假设(A,B)在最左边 既A<E


// 如果D<=F，说明rectangle2 在 rectangle1 的上方
// 如果 B >= H，说明rectangle2 在 rectangle1 的下方
// 如果C <= E，说明rectangle2 在 rectangle1 的右方
// 这三种情况下 都没有交集

// 有交集，交集的左下角纵坐标Max(B,F)
// 左下角横坐标Max(A,E)

// 右上角纵坐标 Min(D,H)
// 右上角横坐标 Min(C,G)

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if A > E {
		return computeArea(E,F,G,H,A,B,C,D)
	}

	if  D <= F || B >= H || C <= E {
		return Abs(C-A)*Abs(D-B) + Abs(G-E) * Abs(H-F)
	}

	AA := Max(B,F)
	BB := Max(A,E)

	CC := Min(D,H)
	DD := Min(C,G)

	return Abs(C-A)*Abs(D-B) + Abs(G-E) * Abs(H-F) - Abs(CC-AA) * Abs(DD - BB)
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}