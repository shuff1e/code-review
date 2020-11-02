package main

import "fmt"

/*

363. 矩形区域不超过 K 的最大数值和
给定一个非空二维矩阵 matrix 和一个整数 k，找到这个矩阵内部不大于 k 的最大矩形和。

示例:

输入: matrix = [[1,0,1],[0,-2,3]], k = 2
输出: 2
解释: 矩形区域 [[0, 1], [-2, 3]] 的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
说明：

矩阵内的矩形区域面积必须大于 0。
如果行数远大于列数，你将如何解答呢？

 */

// 1 0 1
// 0 -2 3

func main() {
	arr := [][]int{
		{1,0,1},
		{0,-2,3},
	}
	fmt.Println(maxSumSubmatrix(arr,2))
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	result := -0x80000000
	for l := 0;l<len(matrix[0]);l++ {
		rowSum := make([]int,len(matrix))
		for r := l;r <len(matrix[0]);r++ {
			for i := 0;i<len(matrix);i++ {
				rowSum[i] += matrix[i][r]
			}

			temp := dpMax(rowSum,k)
			if temp == k {
				return k
			}

			result = Max(result,temp)
		}
	}
	return result
}

func dpMax(arr []int,K int) int {
	result := arr[0]
	sum := arr[0]
	for i:=1;i<len(arr);i++ {
		if sum >=0 {
			sum += arr[i]
		} else {
			sum = arr[i]
		}
		if sum == K {
			return K
		}
		result = Max(result,sum)
	}
	if result <= K {
		return result
	}
	result = -0x80000000
	for l := 0;l<len(arr);l++ {
		sum := 0
		for r := l;r<len(arr);r++ {
			sum += arr[r]
			if sum == K {
				return sum
			}
			if sum < K && sum > result {
				result = sum
			}
		}
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

/*

状态转移方程为 dp(i1,j1,i2,j2) = dp(i1,j1,i2 - 1,j2) + dp(i1,j1,i2,j2 - 1) - dp(i1,j1,i2 - 1,j2 - 1) + matrix[i2 - 1][j2 - 1];

public int maxSumSubmatrix(int[][] matrix, int k) {
    int rows = matrix.length, cols = matrix[0].length, max = Integer.MIN_VALUE;
    int[][][][] dp = new int[rows + 1][cols + 1][rows + 1][cols + 1]; // from (i1,j1) to (i2,j2)
    for (int i1 = 1; i1 <= rows; i1++) {
        for (int j1 = 1; j1 <= cols; j1++) {
            dp[i1][j1][i1][j1] = matrix[i1 - 1][j1 - 1];
            for (int i2 = i1; i2 <= rows; i2++) {
                for (int j2 = j1; j2 <= cols; j2++) {
                    dp[i1][j1][i2][j2] = dp[i1][j1][i2 - 1][j2] + dp[i1][j1][i2][j2 - 1] - dp[i1][j1][i2 - 1][j2 - 1] + matrix[i2 - 1][j2 - 1];
                    if (dp[i1][j1][i2][j2] <= k && dp[i1][j1][i2][j2] > max) max = dp[i1][j1][i2][j2];
                }
            }
        }
    }
    return max;
}


public int maxSumSubmatrix(int[][] matrix, int k) {
    int rows = matrix.length, cols = matrix[0].length, max = Integer.MIN_VALUE;
    for (int i1 = 1; i1 <= rows; i1++) {
        for (int j1 = 1; j1 <= cols; j1++) {
            int[][] dp = new int[rows + 1][cols + 1]; // renew  // from (i1,j1) to (i2,j2)
            dp[i1][j1] = matrix[i1 - 1][j1 - 1];
            for (int i2 = i1; i2 <= rows; i2++) {
                for (int j2 = j1; j2 <= cols; j2++) {
                    dp[i2][j2] = dp[i2 - 1][j2] + dp[i2][j2 - 1] - dp[i2 - 1][j2 - 1] + matrix[i2 - 1][j2 - 1];
                    if (dp[i2][j2] <= k && dp[i2][j2] > max) max = dp[i2][j2];
                }
            }
        }
    }
    return max;
}


public int maxSumSubmatrix(int[][] matrix, int k) {
// 附上完整代码
public int maxSumSubmatrix(int[][] matrix, int k) {
    int rows = matrix.length, cols = matrix[0].length, max = Integer.MIN_VALUE;
    // O(cols ^ 2 * rows)
    for (int l = 0; l < cols; l++) { // 枚举左边界
        int[] rowSum = new int[rows]; // 左边界改变才算区域的重新开始
        for (int r = l; r < cols; r++) { // 枚举右边界
            for (int i = 0; i < rows; i++) { // 按每一行累计到 dp
                rowSum[i] += matrix[i][r];
            }
            max = Math.max(max, dpmax(rowSum, k));
            if (max == k) return k; // 尽量提前
        }
    }
    return max;
}
// 在数组 arr 中，求不超过 k 的最大值
private int dpmax(int[] arr, int k) {
    int rollSum = arr[0], rollMax = rollSum;
    // O(rows)
    for (int i = 1; i < arr.length; i++) {
        if (rollSum > 0) rollSum += arr[i];
        else rollSum = arr[i];
        if (rollSum > rollMax) rollMax = rollSum;
    }
    if (rollMax <= k) return rollMax;
    // O(rows ^ 2)
    int max = Integer.MIN_VALUE;
    for (int l = 0; l < arr.length; l++) {
        int sum = 0;
        for (int r = l; r < arr.length; r++) {
            sum += arr[r];
            if (sum > max && sum <= k) max = sum;
            if (max == k) return k; // 尽量提前
        }
    }
    return max;
}

 */

/*

前提还是一样的，采用固定好左右两边(列)，再对行进行求和操作
但是有没有不再像法一那样的暴力O(n^2)方法了呢?可以优化到O(nlogn)

首先想起来前缀和求法sum[j]-sum[i]，可以理解成 大面积-小面积
结合起题目来就是 大-小<=k 而稍微变化一下就是 小>=大-k
我们主要找的就是符合的小面积，而且要想最逼近k，在大面积一定情况下，小面积越小越好
首先定好大：其实就是暴力中的j即sum[j]
而小，我们要存好暴力中前面所有的结果都存起来即sum[i]
而找小>=大-k 中这个小
想起了lower_bound(elem)这个二分：找有序排列中第一个>=elem的元素的位置
问题来了有序？ 那就用set自动排序存
刚好set有lower_bound函数，一切都是刚刚好
我们就通过不断存入小sum.at(i)，在通过lower_bound(大-k)来找到对应的小，即当前小面积集合中，符合大面积减去小面积小于等于k的更小集合中，小面积最小的那一个
再用 大面积-小面积 即可能答案
注意：当大面积为[0]时，因为没有小面积，所以要先预存一个0进入set中作为假的小面积

 */

/*

可否进一步优化连二分都不做了呢？想起了法二中通过优化53题最大子序和的结论优化成O(n)
再大胆一点，是不是在构建row_sum时候直接就可以嵌入53题的代码了呢
就更加提前的剪枝，不用构建row_sum一个O(n), 求最大子序和一个O(n),然后二分O(nlog(n))

直接构建row_sum+最大子序和O(n)，倘若最大子序和max_sum都符合题意<=k，那就没必要再二分了，
因为二分的最大结果也只是max_sum

 */

/*

class Solution {
public:
    int maxSumSubmatrix(vector<vector<int>>& matrix, int k) {
        int row=matrix.size();
        if (row==0)
            return 0;
        int column=matrix.at(0).size();
        int ans=INT_MIN;
        for (int left=0;left<column;++left)
        {
            vector<int> row_sum(row,0);
            for (int right=left;right<column;++right)
            {
                for (int i=0;i<row;++i)
                    row_sum.at(i)+=matrix.at(i).at(right);
                set<int> helper_set;
                helper_set.insert(0);
                int prefix_row_sum=0;
                for (int i=0;i<row;++i)
                {
                    prefix_row_sum+=row_sum.at(i);
                    auto p=helper_set.lower_bound(prefix_row_sum-k);
                    helper_set.insert(prefix_row_sum);
                    if (p==helper_set.end())
                        continue;
                    else
                    {
                        int temp=prefix_row_sum-(*p);
                        if (temp>ans)
                            ans=temp;
                    }
                }
                if (ans==k)
                    return k;
            }
        }
        return ans;
    }
};

 */

/*

class Solution {
public:
    int maxSumSubmatrix(vector<vector<int>>& matrix, int k) {
        int row=matrix.size();
        if (row==0)
            return 0;
        int column=matrix.at(0).size();
        int ans=INT_MIN;
        for (int left=0;left<column;++left)
        {
            vector<int> row_sum(row,0);
            for (int right=left;right<column;++right)
            {
                //直接在构建row_sum时嵌入求最大子序和
                int max_sub_sum=0;
                int max_sum=INT_MIN;
                for (int i=0;i<row;++i)
                {
                    row_sum.at(i)+=matrix.at(i).at(right);
                    max_sub_sum+=row_sum.at(i);
                    if (max_sub_sum<=k && ans<max_sub_sum)
                        ans=max_sub_sum;
                    if (max_sub_sum<0)
                        max_sub_sum=0;
                    max_sum=max(max_sum,max_sub_sum);
                }
                if (ans==k)
                    return k;
                if (max_sum<=k)
                    continue;
                set<int> helper_set;
                helper_set.insert(0);
                int prefix_row_sum=0;
                for (int i=0;i<row;++i)
                {
                    prefix_row_sum+=row_sum.at(i);
                    auto p=helper_set.lower_bound(prefix_row_sum-k);
                    helper_set.insert(prefix_row_sum);
                    if (p==helper_set.end())
                        continue;
                    else
                    {
                        int temp=prefix_row_sum-(*p);
                        if (temp>ans)
                            ans=temp;
                    }
                }
                if (ans==k)
                    return k;
            }
        }
        return ans;
    }
};

 */