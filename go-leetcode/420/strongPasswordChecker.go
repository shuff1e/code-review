package main

import (
	"container/heap"
	"fmt"
)

/*

420. 强密码检验器
一个强密码应满足以下所有条件：

由至少6个，至多20个字符组成。
至少包含一个小写字母，一个大写字母，和一个数字。
同一字符不能连续出现三次 (比如 "...aaa..." 是不允许的,
但是 "...aa...a..." 是可以的)。
编写函数 strongPasswordChecker(s)，s 代表输入字符串，
如果 s 已经符合强密码条件，则返回0；
否则返回要将 s 修改为满足强密码条件的字符串所需要进行修改的最小步数。

插入、删除、替换任一字符都算作一次修改。

 */

func main() {
	str := "aA123"
	fmt.Println(strongPasswordChecker(str))
}

func strongPasswordChecker(s string) int {
	length := len(s)
	if length == 0 {
		return 6
	}

	low := 1
	up := 1
	num := 1
	if s[0] <='z' && s[0] >= 'a' {
		low = 0
	}
	if s[0] <='Z' && s[0] >= 'A' {
		up = 0
	}
	if s[0] >= '1' && s[0] <= '9' {
		num = 0
	}

	h := &minHeap{}

	cnt := 1
	for i := 1;i<len(s);i++ {
		if low == 1 && s[i] <='z' && s[i] >= 'a' {
			low = 0
		} else if up == 1 && s[i] <='Z' && s[i] >= 'A' {
			up = 0
		} else if num == 1 && s[i] >= '1' && s[i] <= '9' {
			num = 0
		}

		if s[i] == s[i-1] {
			cnt ++
		} else {
			if cnt >=3 {
				heap.Push(h,cnt)
			}
			cnt = 1
		}
	}

	if cnt >=3 {
		heap.Push(h,cnt)
	}

	result := 0

	if length < 6 {
		return Max(6-length,low + up + num)
	}

	for h.Len() > 0 && length > 20 {
		curr := heap.Pop(h).(int)
		result ++
		length --
		curr --
		if curr >= 3 {
			heap.Push(h,curr)
		}
	}

	if length > 20 {
		return result + length - 20 + low + up + num
	}

	n := 0
	for h.Len() > 0 {
		n += heap.Pop(h).(int)/3
	}

	return result + Max(n,low+up+num)
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i,j int) bool {
	return h[i]%3 - h[j] % 3 < 0
}

func (h minHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h,x.(int))
}

func (h *minHeap) Pop() interface{} {
	result := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

/*

class Solution {
    public int strongPasswordChecker(String s) {
        int len = s.length();
        if (len == 0) return 6;
        char[] ss = s.toCharArray();

        // 记录所需小写字母、大写字母和数字的个数
        // cnt记录重复序列的字符计数
        int low = 1, up = 1, num = 1, cnt = 1;
        if (Character.isLowerCase(ss[0])) low = 0;
        else if (Character.isUpperCase(ss[0])) up = 0;
        else if (Character.isDigit(ss[0])) num = 0;

        // 优先队列，x%3小的先出队，为方便需要删除操作时先处理段的序列
        Queue<Integer> queue = new PriorityQueue<>((a, b) -> a % 3 - b % 3);

        for (int i = 1; i < len; i ++) {
            if (low == 1 && Character.isLowerCase(ss[i])) low = 0;
            else if (up == 1 && Character.isUpperCase(ss[i])) up = 0;
            else if (num == 1 && Character.isDigit(ss[i])) num = 0;

            // 对连续相同的序列计数，并存入优先队列
            if (ss[i] != ss[i - 1]) {
                if (cnt >= 3) queue.add(cnt);
                cnt = 1;
            } else {
                cnt ++;
            }
        }
        if (cnt >= 3) queue.add(cnt);

        int ans = 0, all = low + up + num;

        // 长度不足，则根据是否有重复序列进行替换、加添操作
        if (len < 6) return Math.max(6 - len, all);

        // 删除操作，从最短的连续序列开始处理
        while (!queue.isEmpty() && len > 20) {
            int cur = queue.remove();
            ans ++;
            len --;
            if (-- cur >= 3) queue.add(cur);
        }

        // 解决完重复序列后，字符串仍然过长
        if (len > 20) return ans + len - 20 + all;

        // 未处理完重复序列就已经达到长度要求，继续处理重复序列
        // 此时就只考虑替换操作就好了
        int n = 0;
        while (!queue.isEmpty()) {
            n += queue.remove() / 3;
        }

        return ans + Math.max(n, all);
    }

}

 */