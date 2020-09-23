package main

// Q：把一个数组通过插入值的方式构造成一个回文数组，代价就是这个回文数组的所有元素的和，求最小代价是多少。
/*
A:
如果 arr[i] == arr[j]
dp[i][j] = dp[i+1][j-1] + arr[i] + arr[j]

如果 arr[i] != arr[j]
dp[i][j] = dp[i+1][j] + arr[i]*2
dp[i][j] = dp[i][j-1] + arr[j] * 2

初始化条件，dp[i] = arr[i]

 */