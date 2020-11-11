package main

import "fmt"

/*

514. 自由之路
视频游戏“辐射4”中，任务“通向自由”要求玩家到达名为“Freedom Trail Ring”的金属表盘，并使用表盘拼写特定关键词才能开门。

给定一个字符串 ring，表示刻在外环上的编码；给定另一个字符串 key，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。

最初，ring 的第一个字符与12:00方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00 方向对齐，然后按下中心按钮，以此逐个拼写完 key 中的所有字符。

旋转 ring 拼出 key 字符 key[i] 的阶段中：

您可以将 ring 顺时针或逆时针旋转一个位置，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00 方向对齐，并且这个字符必须等于字符 key[i] 。
如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key 的下一个字符（下一阶段）, 直至完成所有拼写。
示例：





输入: ring = "godding", key = "gd"
输出: 4
解释:
 对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
 对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
 当然, 我们还需要1步进行拼写。
 因此最终的输出是 4。
提示：

ring 和 key 的字符串长度取值范围均为 1 至 100；
两个字符串中都只有小写字符，并且均可能存在重复字符；
字符串 key 一定可以由字符串 ring 旋转拼出。

 */
// 0 1 2 3
// 1 2 3 4
// length - 1 - j + i + 1 = length -j + i

func main() {
	ring := "godding"
	key := "gd"
	fmt.Println(findRotateSteps2(ring,key))
}

func findRotateSteps(ring string, key string) int {
	dict := map[byte][]int{}
	for i := 0;i<len(ring);i++ {
		if _,ok := dict[ring[i]];!ok {
			dict[ring[i]] = []int{i}
		} else {
			dict[ring[i]] = append(dict[ring[i]],i)
		}
	}
	memo := make([][]int,len(key) + 1)
	for i := 0;i<len(memo);i++ {
		memo[i] = make([]int,len(ring) + 1)
	}
	return help(ring,key,0,0,dict,memo)
}

// 当前是ring中的initial 这个位置对准了12:00
// 匹配key中的字符，位置为index
// pos是所有字符的位置

func help(ring,key string,index int,initial int,pos map[byte][]int,memo [][]int) int {
	if index == len(key) {
		return 0
	}
	if memo[index][initial] != 0 {
		return memo[index][initial]
	}
	result := 0x7fffffff
	for _,v := range pos[key[index]] {
		// 顺时针，逆时针转都可以
		temp := help(ring,key,index + 1,v,pos,memo) +
			Min(Abs(v-initial),Abs(len(ring) - Abs(v-initial))) + 1
		result = Min(result,temp)
	}
	memo[index][initial] = result
	return result
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func findRotateSteps2(ring string, key string) int {
	dict := map[byte][]int{}
	for i := 0;i<len(ring);i++ {
		if _,ok := dict[ring[i]];!ok {
			dict[ring[i]] = []int{i}
		} else {
			dict[ring[i]] = append(dict[ring[i]],i)
		}
	}

	dp := make([][]int,len(key))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,len(ring))
		for j := 0;j<len(ring);j++ {
			dp[i][j] = 0x7fffffff
		}
	}

	for j := 0;j<len(ring);j++ {
		dp[0][j] = Min(j,len(ring)-j) + 1
	}

	for i := 1;i<len(key);i++ {
		for _,j := range dict[key[i]] {
			for _,k := range dict[key[i-1]] {
				dp[i][j] = Min(dp[i][j],
					dp[i-1][k] + Min(Abs(j-k),Abs(len(ring) - Abs(j-k))) + 1)
			}
		}
	}
	result := dp[len(key)-1][0]
	for j := 1;j<len(ring);j++ {
		result = Min(result,dp[len(key)-1][j])
	}
	return result
}

/*

方法一：动态规划
定义 dp[i][j] 表示从前往后拼写出 key 的第 i 个字符， ring 的第 j 个字符与 12:00 方向对齐的最少步数（下标均从 0 开始）。

显然，只有当字符串 ring 的第 j 个字符需要和 key 的第 i 个字符相同时才能拼写出 key 的第 i 个字符，
因此对于 key 的第 ii 个字符，需要考虑计算的 ring 的第 j 个字符只有 key[i] 在 ring 中出现的下标集合。
我们对每个字符维护一个位置数组 pos[i]，表示字符 i 在 ring 中出现的位置集合，用来加速计算转移的过程。

对于状态 dp[i][j]，需要枚举上一次与 12:00 方向对齐的位置 k，因此可以列出如下的转移方程：

dp[i][j]=
k∈pos[key[i−1]]
min
​
{dp[i−1][k]+min{abs(j−k),n−abs(j−k)}+1}

其中 min{abs(j−k),n−abs(j−k)}+1 表示在当前第 k 个字符与 12:00 方向对齐时第 j 个字符旋转到 12:00 方向并按下拼写的最少步数。

最后答案即为 min
i=0
n−1
​
 {dp[m−1][i]}。

这样的做法需要开辟 O(mn) 的空间来存放 dp 值。考虑到每次转移状态 dp[i][] 只会从 dp[i−1][] 转移过来，因此我们可以利用滚动数组优化第一维的空间复杂度，有能力的读者可以尝试实现。下面只给出最朴素的 O(mn)O(mn) 空间复杂度的实现。


dp[i][j] 表示当前是key的第i个字符，j是将ring的第j个字符 和 12:00对齐


dp[i][j] = Min(dp[i-1][j-k] + Min(j-k,n-(j-k)) + 1)

k属于pos[i-1]，之前12点肯定和pos[i-1]中的一个对齐
dp[0][0] = 1
dp[0][j] = Min(j,n-j) + 1

最后的答案为 Min(dp[len(key)-1][j])

class Solution {
    public int findRotateSteps(String ring, String key) {
        int n = ring.length(), m = key.length();
        List<Integer>[] pos = new List[26];
        for (int i = 0; i < 26; ++i) {
            pos[i] = new ArrayList<Integer>();
        }
        for (int i = 0; i < n; ++i) {
            pos[ring.charAt(i) - 'a'].add(i);
        }
        int[][] dp = new int[m][n];
        for (int i = 0; i < m; ++i) {
            Arrays.fill(dp[i], 0x3f3f3f);
        }

        for (int i : pos[key.charAt(0) - 'a']) {
            dp[0][i] = Math.min(i, n - i) + 1;
        }

        for (int i = 1; i < m; ++i) {
            for (int j : pos[key.charAt(i) - 'a']) {
                for (int k : pos[key.charAt(i - 1) - 'a']) {
                    dp[i][j] = Math.min(dp[i][j], dp[i - 1][k] + Math.min(Math.abs(j - k), n - Math.abs(j - k)) + 1);
                }
            }
        }
        return Arrays.stream(dp[m - 1]).min().getAsInt();
    }
}

 */