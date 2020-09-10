package main

// Q：编写代码，以给定值x为基准将链表分割成两部分，所有小于x的节点排在大于或等于x的节点之前
// 1-> 3-> 5-> 7-> 2-> 4-> 6
// x为4

// A：i <= x,swap(i,x),mark++
// 当i 到了2，mark为5，swap(2,5)
// 7.next = 5 5.next = 4
// 3.next = 2 2.next = 5

