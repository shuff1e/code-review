package Catalan_number

// 一个栈(无穷大)的[进栈]序列为1，2，3，…，n，有多少个不同的[出栈]序列?

// 假设数K最后出栈
// 那么1到K-1的进栈序列为1，2...K-1，出栈序列为f(k-1)
// 同样K-1到n的进栈序列为K+1，K...n，出栈序列为f(n-K)
