package Catalan_number

// n个节点的二叉搜索树(BST)总共有多少棵？

// n个节点的BST的数目称为卡特兰树，记为Cn

// 将树的节点从1到n编号
// 从这些节点中选择一个(i)作为根节点，然后将小于根节点的划分到左子树C(i-1)
// 将大于根节点的划分到右子树C(n-i)
// 由于两棵子树相互独立，所以总的个数位C(i-1)*C(n-i)

func BSTNumber(n int) int {
	memo := make([]int,n+1)
	return f(n,memo)
}

func f(n int,memo []int) int {
	if n == 0 {
		return 1
	}
	sum := 0
	for i := 1;i<=n;i++ {
		sum += f(i-1,memo)*f(n-i,memo)
	}
	memo[n] = sum
	return memo[n]
}

