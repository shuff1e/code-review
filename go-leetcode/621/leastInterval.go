package main

/*

621. 任务调度器
给你一个用字符数组 tasks 表示的 CPU 需要执行的任务列表。其中每个字母表示一种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。在任何一个单位时间，CPU 可以完成一个任务，或者处于待命状态。

然而，两个 相同种类 的任务之间必须有长度为整数 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

你需要计算完成所有任务所需要的 最短时间 。



示例 1：

输入：tasks = ["A","A","A","B","B","B"], n = 2
输出：8
解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B
     在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。
示例 2：

输入：tasks = ["A","A","A","B","B","B"], n = 0
输出：6
解释：在这种情况下，任何大小为 6 的排列都可以满足要求，因为 n = 0
["A","A","A","B","B","B"]
["A","B","A","B","A","B"]
["B","B","B","A","A","A"]
...
诸如此类
示例 3：

输入：tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
输出：16
解释：一种可能的解决方案是：
     A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> (待命) -> (待命) -> A -> (待命) -> (待命) -> A


提示：

1 <= task.length <= 104
tasks[i] 是大写英文字母
n 的取值范围为 [0, 100]

 */

func leastInterval(tasks []byte, n int) int {
	freq := make(map[byte]int)
	maxExec := 0
	for _,v := range tasks {
		freq[v] = freq[v] + 1
		maxExec = Max(maxExec,freq[v])
	}
	maxCount := 0
	for _,v := range freq {
		if v == maxExec {
			maxCount ++
		}
	}
	return Max((maxExec-1) * (n+1) + maxCount,len(tasks))
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

/*

一种容易想到的方法是，我们按照时间顺序，依次给每一个时间单位分配任务。

那么如果当前有多种任务不在冷却中，那么我们应该如何挑选执行的任务呢？直觉上，我们应当选择剩余执行次数最多的那个任务，将每种任务的剩余执行次数尽可能平均，使得 CPU 处于待命状态的时间尽可能少

class Solution {
    public int leastInterval(char[] tasks, int n) {
        Map<Character, Integer> freq = new HashMap<Character, Integer>();
        for (char ch : tasks) {
            freq.put(ch, freq.getOrDefault(ch, 0) + 1);
        }

        // 任务总数
        int m = freq.size();
        List<Integer> nextValid = new ArrayList<Integer>();
        List<Integer> rest = new ArrayList<Integer>();
        Set<Map.Entry<Character, Integer>> entrySet = freq.entrySet();
        for (Map.Entry<Character, Integer> entry : entrySet) {
            int value = entry.getValue();
            nextValid.add(1);
            rest.add(value);
        }

        int time = 0;
        for (int i = 0; i < tasks.length; ++i) {
            ++time;
            int minNextValid = Integer.MAX_VALUE;
            for (int j = 0; j < m; ++j) {
                if (rest.get(j) != 0) {
                    minNextValid = Math.min(minNextValid, nextValid.get(j));
                }
            }
            time = Math.max(time, minNextValid);
            int best = -1;
            for (int j = 0; j < m; ++j) {
                if (rest.get(j) != 0 && nextValid.get(j) <= time) {
                    if (best == -1 || rest.get(j) > rest.get(best)) {
                        best = j;
                    }
                }
            }
            nextValid.set(best, time + n + 1);
            rest.set(best, rest.get(best) - 1);
        }

        return time;
    }
}

 */

func leastInterval2(tasks []byte, n int) int {
	freq := make(map[byte]int)
	for _,v := range tasks {
		freq[v] = freq[v] + 1
	}

	// 多少个任务
	m := len(freq)
	// 每个任务开始的时间
	nextValid := []int{}
	// 每个任务的个数
	rest := []int{}

	for _,v := range freq {
		nextValid = append(nextValid,1)
		rest = append(rest,v)
	}

	time := 0
	for i := 0;i<len(tasks);i++ {
		time ++
		minNextValid := 0x7fffffff
		for j := 0;j<m;j++ {
			if rest[j] != 0 {
				minNextValid = Min(minNextValid,nextValid[j])
			}
		}
		time = Max(time,minNextValid)

		best := -1
		for j := 0;j<m;j++ {
			// 这个任务 是有效的
			if rest[j] != 0 && nextValid[j] <= time {
				// 找到valid中次数最大的
				if best == -1 || rest[j] > rest[best] {
					best = j
				}
			}
		}
		nextValid[best] = time + n + 1
		rest[best] --
	}
	return time
}