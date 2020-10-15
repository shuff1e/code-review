package main

import "fmt"

/*
233. 数字 1 的个数
给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。

示例:

输入: 13
输出: 6
解释: 数字 1 出现在以下数字中: 1, 10, 11, 12, 13 。=
 */

func main() {
	fmt.Println(countDigitOne(13))
}

/*

总体思想就是分类，先求所有数中个位是 1 的个数，再求十位是 1 的个数，再求百位是 1 的个数...

假设 n = xyzdabc，此时我们求千位是 1 的个数，也就是 d 所在的位置。

那么此时有三种情况，

d == 0，那么千位上 1 的个数就是 xyz * 1000
d == 1，那么千位上 1 的个数就是 xyz * 1000 + abc + 1
d > 1，那么千位上 1 的个数就是 xyz * 1000 + 1000
为什么呢？

当我们考虑千位是 1 的时候，我们将千位定为 1，也就是 xyz1abc。

对于 xyz 的话，可以取 0,1,2...(xyz-1)，也就是 xyz 种可能。

当 xyz 固定为上边其中的一个数的时候，abc 可以取 0,1,2...999，也就是 1000 种可能。

这样的话，总共就是 xyz*1000 种可能。

注意到，我们前三位只取到了 xyz-1，那么如果取 xyz 呢？

此时就出现了上边的三种情况，取决于 d 的值。

d == 1 的时候，千位刚好是 1，此时 abc 可以取的值就是 0 到 abc ，所以多加了 abc + 1。

d > 1 的时候，d 如果取 1，那么 abc 就可以取 0 到 999，此时就多加了 1000。


如果n = 4560234
让我们统计一下千位有多少个 1
xyz 可以取 0 到 455, abc 可以取 0 到 999
4551000 to 4551999 (1000)
4541000 to 4541999 (1000)
4531000 to 4531999 (1000)
...
  21000 to   21999 (1000)
  11000 to   11999 (1000)
   1000 to    1999 (1000)
总共就是 456 * 1000

如果 n = 4561234
xyz 可以取 0 到 455, abc 可以取 0 到 999
4551000 to 4551999 (1000)
4541000 to 4541999 (1000)
4531000 to 4531999 (1000)
...
1000 to 1999 (1000)
xyz 还可以取 456, abc 可以取 0 到 234
4561000 to 4561234 (234 + 1)
总共就是 456 * 1000 + 234 + 1

如果 n = 4563234
xyz 可以取 0 到 455, abc 可以取 0 到 999
4551000 to 4551999 (1000)
4541000 to 4541999 (1000)
4531000 to 4531999 (1000)
...
1000 to 1999 (1000)
xyz 还可以取 456, abc 可以取 0 到 999
4561000 to 4561999 (1000)
总共就是 456 * 1000 + 1000

 */

func countDigitOne(n int) int {
	count := 0
	for k := 1;k<=n;k=k*10 {
		// xyz d abc
		abc := n%k
		xyzd := n/k

		d := xyzd % 10
		xyz := xyzd/10

		count += xyz * k
		if d == 1 {
			count += abc + 1
		}
		if d > 1 {
			count += k
		}
		if xyz == 0 {
			break
		}
	}
	return count
}