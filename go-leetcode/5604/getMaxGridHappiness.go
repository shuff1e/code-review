package main

import "fmt"

/*

5604. 最大化网格幸福感
给你四个整数 m、n、introvertsCount 和 extrovertsCount 。有一个 m x n 网格，和两种类型的人：内向的人和外向的人。总共有 introvertsCount 个内向的人和 extrovertsCount 个外向的人。

请你决定网格中应当居住多少人，并为每个人分配一个网格单元。 注意，不必 让所有人都生活在网格中。

每个人的 幸福感 计算如下：

内向的人 开始 时有 120 个幸福感，但每存在一个邻居（内向的或外向的）他都会 失去  30 个幸福感。
外向的人 开始 时有 40 个幸福感，每存在一个邻居（内向的或外向的）他都会 得到  20 个幸福感。
邻居是指居住在一个人所在单元的上、下、左、右四个直接相邻的单元中的其他人。

网格幸福感 是每个人幸福感的 总和 。 返回 最大可能的网格幸福感 。



示例 1：


输入：m = 2, n = 3, introvertsCount = 1, extrovertsCount = 2
输出：240
解释：假设网格坐标 (row, column) 从 1 开始编号。
将内向的人放置在单元 (1,1) ，将外向的人放置在单元 (1,3) 和 (2,3) 。
- 位于 (1,1) 的内向的人的幸福感：120（初始幸福感）- (0 * 30)（0 位邻居）= 120
- 位于 (1,3) 的外向的人的幸福感：40（初始幸福感）+ (1 * 20)（1 位邻居）= 60
- 位于 (2,3) 的外向的人的幸福感：40（初始幸福感）+ (1 * 20)（1 位邻居）= 60
网格幸福感为：120 + 60 + 60 = 240
上图展示该示例对应网格中每个人的幸福感。内向的人在浅绿色单元中，而外向的人在浅紫色单元中。
示例 2：

输入：m = 3, n = 1, introvertsCount = 2, extrovertsCount = 1
输出：260
解释：将内向的人放置在单元 (1,1) 和 (3,1) ，将外向的人放置在单元 (2,1) 。
- 位于 (1,1) 的内向的人的幸福感：120（初始幸福感）- (1 * 30)（1 位邻居）= 90
- 位于 (2,1) 的外向的人的幸福感：40（初始幸福感）+ (2 * 20)（2 位邻居）= 80
- 位于 (3,1) 的内向的人的幸福感：120（初始幸福感）- (1 * 30)（1 位邻居）= 90
网格幸福感为 90 + 80 + 90 = 260
示例 3：

输入：m = 2, n = 2, introvertsCount = 4, extrovertsCount = 0
输出：240


提示：

1 <= m, n <= 5
0 <= introvertsCount, extrovertsCount <= min(m * n, 6)

 */

func main() {
	m := 3
	n := 1
	introvertsCount := 2
	extrovertsCount := 1
	fmt.Println(getMaxGridHappiness2(m,n,introvertsCount,extrovertsCount))
}

func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
	// 3^6 = 729
	// 最多6位
	// 每一个mask的三进制 表示
	mask_span := make([][]int,729)
	for i := 0;i<len(mask_span);i++ {
		mask_span[i] = make([]int,6)
	}

	// 上一行的mask，当前处理到的行，剩余的内向人数，剩余的外向人数
	dp := make([][][][]int,729)
	for i := 0;i<len(dp);i++ {
		dp[i] = make([][][]int,6)
		for j := 0;j<len(dp[i]);j++ {
			dp[i][j] = make([][]int,7)
			for k := 0;k<len(dp[i][j]);k++ {
				dp[i][j][k] = make([]int,7)
				for h := 0;h<len(dp[i][j][k]);h++ {
					dp[i][j][k][h] = -1
				}
			}
		}
	}

	// 每个mask中内向的人的个数
	nx_inner := make([]int,729)
	// 每个mask中外向的人的个数
	wx_inner := make([]int,729)

	// 行内得分，只统计 mask 本身的得分，不包括它与上一行的）
	score_inner := make([]int,729)

	// 行外得分
	score_outer := make([][]int,729)
	for i := 0;i<len(score_outer);i++ {
		score_outer[i] = make([]int,729)
	}

	n3 := Pow(3,n)

	for mask := 0 ; mask < n3;mask++ {
		temp := mask
		for i := 0;i<n;i++ {
			mask_span[mask][i] = temp%3
			temp /= 3
		}

		for i := 0;i<n;i++ {
			if mask_span[mask][i] != 0 {
				if mask_span[mask][i] == 1 {
					nx_inner[mask] ++
					score_inner[mask] += 120
				} else if mask_span[mask][i] == 2 {
					wx_inner[mask] ++
					score_inner[mask] += 40
				}
				if i >= 1 {
					score_inner[mask] += calc(mask_span[mask][i],mask_span[mask][i-1])
				}
			}
		}

		// 行外得分
		for mask0 := 0;mask0<n3;mask0 ++ {
			for mask1 := 0;mask1<n3;mask1 ++ {
				score_outer[mask0][mask1] = 0
				for i := 0;i<n;i++ {
					score_outer[mask0][mask1] += calc(mask_span[mask0][i],mask_span[mask1][i])
				}
			}
		}
	}

	return dfs(dp,0,0,m,n3,introvertsCount,extrovertsCount,
		nx_inner,wx_inner,score_inner,score_outer)
}

func dfs(dp [][][][]int,mask_last int,row int,allRow,n3 int,nx int,wx int,
	nx_inner,wx_inner []int,score_inner []int,score_outer [][]int) int {
	if row == allRow || nx + wx == 0 {
		return 0
	}
	if dp[mask_last][row][nx][wx] != -1 {
		return dp[mask_last][row][nx][wx]
	}

	best := 0
	for mask := 0;mask < n3 ;mask ++ {
		if nx_inner[mask] > nx || wx_inner[mask] > wx {
			continue
		}
		score := score_inner[mask] + score_outer[mask_last][mask]
		best = Max(best,score + dfs(dp,mask,row+1,allRow,n3,nx - nx_inner[mask],wx - wx_inner[mask],
			nx_inner,wx_inner,score_inner,score_outer))
	}
	dp[mask_last][row][nx][wx] = best
	return best
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}


func calc(x,y int) int {
	if x == 0 || y == 0 {
		return 0
	}
	if x == 1 && y == 1 {
		return -60
	}
	if x == 2 && y == 2 {
		return 40
	}
	return -10
}

func Pow(base,exp int) int {
	result := 1
	for exp > 0 {
		if exp & 1 > 0 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

/*

class Solution {
private:
    // 预处理：每一个 mask 的三进制表示
    int mask_span[729][6];
    // dp[上一行的 mask][当前处理到的行][剩余的内向人数][剩余的外向人数]
    int dp[729][6][7][7];
    // 预处理：每一个 mask 包含的内向人数，外向人数，行内得分（只统计 mask 本身的得分，不包括它与上一行的），行外得分
    int nx_inner[729], wx_inner[729], score_inner[729], score_outer[729][729];
    // n3 = n^3
    int m, n, n3;

public:
    // 如果 x 和 y 相邻，需要加上的分数
    inline int calc(int x, int y) {
        if (x == 0 || y == 0) {
            return 0;
        }
        // 例如两个内向的人，每个人要 -30，一共 -60，下同
        if (x == 1 && y == 1) {
            return -60;
        }
        if (x == 2 && y == 2) {
            return 40;
        }
        return -10;
    }

    int getMaxGridHappiness(int m, int n, int nx, int wx) {
        this->m = m;
        this->n = n;
        this->n3 = pow(3, n);

        // 预处理
        for (int mask = 0; mask < n3; ++mask) {
            for (int mask_tmp = mask, i = 0; i < n; ++i) {
                mask_span[mask][i] = mask_tmp % 3;
                mask_tmp /= 3;
            }
            nx_inner[mask] = wx_inner[mask] = score_inner[mask] = 0;
            for (int i = 0; i < n; ++i) {
                if (mask_span[mask][i] != 0) {
                    // 个人分数
                    if (mask_span[mask][i] == 1) {
                        ++nx_inner[mask];
                        score_inner[mask] += 120;
                    }
                    else if (mask_span[mask][i] == 2) {
                        ++wx_inner[mask];
                        score_inner[mask] += 40;
                    }
                    // 行内分数
                    if (i - 1 >= 0) {
                        score_inner[mask] += calc(mask_span[mask][i], mask_span[mask][i - 1]);
                    }
                }
            }
        }
        // 行外分数
        for (int mask0 = 0; mask0 < n3; ++mask0) {
            for (int mask1 = 0; mask1 < n3; ++mask1) {
                score_outer[mask0][mask1] = 0;
                for (int i = 0; i < n; ++i) {
                    score_outer[mask0][mask1] += calc(mask_span[mask0][i], mask_span[mask1][i]);
                }
            }
        }

        memset(dp, -1, sizeof(dp));
        return dfs(0, 0, nx, wx);
    }


    // dfs(上一行的 mask，当前处理到的行，剩余的内向人数，剩余的外向人数）
    int dfs(int mask_last, int row, int nx, int wx) {
        // 边界条件：如果已经处理完，或者没有人了
        if (row == m || nx + wx == 0) {
            return 0;
        }
        // 记忆化
        if (dp[mask_last][row][nx][wx] != -1) {
            return dp[mask_last][row][nx][wx];
        }

        int best = 0;
        for (int mask = 0; mask < n3; ++mask) {
            if (nx_inner[mask] > nx || wx_inner[mask] > wx) {
                continue;
            }
            int score = score_inner[mask] + score_outer[mask][mask_last];
            best = max(best, score + dfs(mask, row + 1, nx - nx_inner[mask], wx - wx_inner[mask]));
        }

        return dp[mask_last][row][nx][wx] = best;
    }
};

 */

func getMaxGridHappiness2(m int, n int, introvertsCount int, extrovertsCount int) int {
	// 预处理：每一个 mask 的三进制表示
	mask_span := make([][]int,729)
	for i := 0;i<len(mask_span);i++ {
		mask_span[i] = make([]int,6)
	}

	// dp[位置][轮廓线上的 mask][剩余的内向人数][剩余的外向人数]
	dp := make([][][][]int,25)
	for i := 0;i<len(dp);i++ {
		dp[i] = make([][][]int,729)
		for j := 0;j<len(dp[i]);j++ {
			dp[i][j] = make([][]int,7)
			for k := 0;k<len(dp[i][j]);k++ {
				dp[i][j][k] = make([]int,7)
				for h := 0;h<len(dp[i][j][k]);h++ {
					dp[i][j][k][h] = -1
				}
			}
		}
	}

	// 预处理：每一个 mask 去除最高位、乘 3、加上新的最低位的结果
	truncate := make([][]int,729)
	for i := 0;i<len(truncate);i++ {
		truncate[i] = make([]int,3)
	}

	n3 := Pow(3,n)

	highest := n3/3


	for mask := 0;mask < n3;mask ++ {
		temp := mask
		for i := 0;i<n;i++ {
			// 与方法一不同的是，这里需要反过来存储，这样 [0] 对应最高位，[n-1] 对应最低位
			mask_span[mask][n-i-1] = temp%3
			temp/=3
		}
		truncate[mask][0] = mask%highest*3
		truncate[mask][1] = mask%highest*3 + 1
		truncate[mask][2] = mask%highest*3 + 2
	}

	return dfs2(0,0,introvertsCount,extrovertsCount,
		m,n,dp,truncate,mask_span)
}

func dfs2(pos int,borderline int,nx int,wx int,m,n int,dp [][][][]int,truncate [][]int,mask_span [][]int) int {
	if pos == m*n || nx + wx == 0 {
		return 0
	}

	if dp[pos][borderline][nx][wx] != -1 {
		return dp[pos][borderline][nx][wx]
	}

	y := pos%n

	// 什么都不做
	best := dfs2(pos+1,truncate[borderline][0],nx,wx,m,n,dp,truncate,mask_span)

	// 放一个内向的人
	if nx > 0 {
		temp := 120 + calc(1,mask_span[borderline][0])
		if y != 0 {
			temp += calc(1,mask_span[borderline][n-1])
		}
		temp += dfs2(pos+1,truncate[borderline][1],nx-1,wx,m,n,dp,truncate,mask_span)
		best = Max(best,temp)
	}
	// 放一个外向的人
	if wx > 0 {
		temp := 40 + calc(2,mask_span[borderline][0])
		if y > 0 {
			temp += calc(2,mask_span[borderline][n-1])
		}
		temp += dfs2(pos+1,truncate[borderline][2],nx,wx-1,m,n,dp,truncate,mask_span)
		best = Max(best,temp)
	}

	dp[pos][borderline][nx][wx] = best

	return best
}