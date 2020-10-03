package main

import "fmt"

/*
74. 搜索二维矩阵
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
[1,   3,  5,  7],
[10, 11, 16, 20],
[23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
[1,   3,  5,  7],
[10, 11, 16, 20],
[23, 30, 34, 50]
]
target = 13
输出: false
 */

// A：二分查找

func main() {
	matrix := [][]int{{1}}
	fmt.Println(searchMatrix(matrix,2))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0])== 0{
		return false
	}
	l,r := 0,len(matrix) -1
	for l <= r {
		mid := (l+r)/2
		if matrix[mid][len(matrix[0])-1] == target {
			l = mid
			break
		} else if matrix[mid][len(matrix[0])-1] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	row := l
	if row >= len(matrix) {
		return false
	}

	l,r = 0,len(matrix[0]) - 1
	for l <= r {
		mid := (l+r)/2
		if matrix[row][mid] == target {
			l = mid
			break
		} else if matrix[row][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l >= len(matrix[0]) {
		return false
	}
	return matrix[row][l] == target
}


/*

方法：二分查找

注意到输入的 m x n 矩阵可以视为长度为 m x n的有序数组。

由于该 虚 数组的序号可以由下式方便地转化为原矩阵中的行和列 (我们当然不会真的创建一个新数组) ，该有序数组非常适合二分查找。

row = idx / n ， col = idx % n。

算法

这是一个标准二分查找算法 :

初始化左右序号
left = 0 和 right = m x n - 1。

While left < right :

选取虚数组最中间的序号作为中间序号: pivot_idx = (left + right) / 2。

该序号对应于原矩阵中的 row = pivot_idx / n行 col = pivot_idx % n 列, 由此可以拿到中间元素pivot_element。
该元素将虚数组分为两部分。

比较 pivot_element 与 target 以确定在哪一部分进行进一步查找。

class Solution {
  public boolean searchMatrix(int[][] matrix, int target) {
    int m = matrix.length;
    if (m == 0) return false;
    int n = matrix[0].length;

    // 二分查找
    int left = 0, right = m * n - 1;
    int pivotIdx, pivotElement;
    while (left <= right) {
      pivotIdx = (left + right) / 2;
      pivotElement = matrix[pivotIdx / n][pivotIdx % n];
      if (target == pivotElement) return true;
      else {
        if (target < pivotElement) right = pivotIdx - 1;
        else left = pivotIdx + 1;
      }
    }
    return false;
  }
}
 */