package main

/*

968. 监控二叉树
给定一个二叉树，我们在树的节点上安装摄像头。

节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。

计算监控树的所有节点所需的最小摄像头数量。



示例 1：
      0
     /
    0
   / \
  0   0

输入：[0,0,null,0,0]
输出：1
解释：如图所示，一台摄像头足以监控所有节点。
示例 2：
          0
         /
        0
       /
      0
     /
    0
     \
      0


输入：[0,0,null,0,null,0,null,null,0]
输出：2
解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。

提示：

给定树的节点数的范围是 [1, 1000]。
每个节点的值都是 0。

 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	count := 0
	status := help(root,&count)
	if status == 0 {
		count ++
	}
	return count
}

func help(root *TreeNode,count *int) int {
	if root == nil {
		return 2
	}
	left := help(root.Left,count)
	right := help(root.Right,count)

	if left == 2 && right == 2 {
		return 0
	}

	if left == 0 || right == 0 {
		*count ++
		return 1
	}

	if left == 1 || right == 1 {
		return 2
	}

	return -1
}

/*

这道题目其实不是那么好理解的，题目举的示例不是很典型，会误以为摄像头必须要放在中间，其实放哪里都可以只要覆盖了就行。

这道题目难在两点：

需要确定遍历方式
需要状态转移的方程
我们之前做动态规划的时候，只要最难的地方在于确定状态转移方程，至于遍历方式无非就是在数组或者二维数组上。

本题并不是动态规划，其本质是贪心，但我们要确定状态转移方式，而且要在树上进行推导，所以难度就上来了，一些同学知道这道题目难，但其实说不上难点究竟在哪。

需要确定遍历方式
首先先确定遍历方式，才能确定转移方程，那么该如何遍历呢？

在安排选择摄像头的位置的时候，我们要从底向上进行推导，因为尽量让叶子节点的父节点安装摄像头，这样摄像头的数量才是最少的 ，这也是本道贪心的原理所在！

如何从低向上推导呢？

就是后序遍历也就是左右中的顺序，这样就可以从下到上进行推导了。

后序遍历代码如下：

    int traversal(TreeNode* cur) {

        // 空节点，该节点有覆盖
        if (终止条件) return ;

        int left = traversal(cur->left);    // 左
        int right = traversal(cur->right);  // 右

        逻辑处理                            // 中

        return ;
    }


需要状态转移的方程
确定了遍历顺序，再看看这个状态应该如何转移，先来看看每个节点可能有几种状态：

可以说有如下三种：

该节点无覆盖
本节点有摄像头
本节点有覆盖
我们分别有三个数字来表示：

0：该节点无覆盖
1：本节点有摄像头
2：本节点有覆盖

大家应该找不出第四个节点的状态了。

一些同学可能会想有没有第四种状态：本节点无摄像头，其实无摄像头就是 无覆盖 或者 有覆盖的状态，所以一共还是三个状态。

那么问题来了，空节点究竟是哪一种状态呢？ 空节点表示无覆盖？ 表示有摄像头？还是有覆盖呢？

回归本质，为了让摄像头数量最少，我们要尽量让叶子节点的父节点安装摄像头，这样才能摄像头的数量最少。

那么空节点不能是无覆盖的状态，这样叶子节点就可以放摄像头了，空节点也不能是有摄像头的状态，这样叶子节点的父节点就没有必要放摄像头了，而是可以把摄像头放在叶子节点的爷爷节点上。

所以空节点的状态只能是有覆盖，这样就可以在叶子节点的父节点放摄像头了

接下来就是递推关系。

那么递归的终止条件应该是遇到了空节点，此时应该返回 2（有覆盖），原因上面已经解释过了。

class Solution {
private:
    int result;
    int traversal(TreeNode* cur) {

        // 空节点，该节点有覆盖
        if (cur == NULL) return 2;

        int left = traversal(cur->left);    // 左
        int right = traversal(cur->right);  // 右

        // 情况1
        // 左右节点都有覆盖
        if (left == 2 && right == 2) return 0;

        // 情况2
        // left == 0 && right == 0 左右节点无覆盖
        // left == 1 && right == 0 左节点有摄像头，右节点无覆盖
        // left == 0 && right == 1 左节点有无覆盖，右节点摄像头
        // left == 0 && right == 2 左节点无覆盖，右节点覆盖
        // left == 2 && right == 0 左节点覆盖，右节点无覆盖
        if (left == 0 || right == 0) {
            result++;
            return 1;
        }

        // 情况3
        // left == 1 && right == 2 左节点有摄像头，右节点有覆盖
        // left == 2 && right == 1 左节点有覆盖，右节点有摄像头
        // left == 1 && right == 1 左右节点都有摄像头
        // 其他情况前段代码均已覆盖
        if (left == 1 || right == 1) return 2;

        // 以上代码我没有使用else，主要是为了把各个分支条件展现出来，这样代码有助于读者理解
        // 这个 return -1 逻辑不会走到这里。
        return -1;
    }

public:
    int minCameraCover(TreeNode* root) {
        result = 0;
        // 情况4
        if (traversal(root) == 0) { // root 无覆盖
            result++;
        }
        return result;
    }
};

 */

/*

思路与算法

本题以二叉树为背景，不难想到用递归的方式求解。本题的难度在于如何从左、右子树的状态，推导出父节点的状态。

为了表述方便，我们约定：如果某棵树的所有节点都被监控，则称该树被「覆盖」。

假设当前节点为 root，其左右孩子为 left,right。如果要覆盖以 root 为根的树，有两种情况：

若在 root 处安放摄像头，则孩子 left,right 一定也会被监控到。此时，只需要保证 left 的两棵子树被覆盖，同时保证 right 的两棵子树也被覆盖即可。

否则， 如果 root 处不安放摄像头，则除了覆盖 root 的两棵子树之外，孩子 left,right 之一必须要安装摄像头，从而保证 root 会被监控到。
根据上面的讨论，能够分析出，对于每个节点 root ，需要维护三种类型的状态：

状态 a：root 必须放置摄像头的情况下，覆盖整棵树需要的摄像头数目。
状态 b：覆盖整棵树需要的摄像头数目，无论 root 是否放置摄像头。
状态 c：覆盖两棵子树需要的摄像头数目，无论节点 root 本身是否被监控到。
根据它们的定义，一定有 a≥b≥c。

对于节点 root 而言，设其左右孩子 left,right 对应的状态变量分别为 (l_a,l_b,l_c) 以及 (r_a,r_b,r_c)。

根据一开始的讨论，我们已经得到了求解 a,b 的过程：

a = l_c + r_c + 1

b=min(a,min(la + rb,ra+lb))

对于 c 而言，要保证两棵子树被完全覆盖，要么 root 处放置一个摄像头，需要的摄像头数目为 a；
要么 root 处不放置摄像头，此时两棵子树分别保证自己被覆盖，需要的摄像头数目为 l_b + r_b。

需要额外注意的是，对于 root 而言，如果其某个孩子为空，则不能通过在该孩子处放置摄像头的方式，监控到当前节点。
因此，该孩子对应的变量 a应当返回一个大整数，用于标识不可能的情形。

最终，根节点的状态变量 b 即为要求出的答案。

class Solution {
    public int minCameraCover(TreeNode root) {
        int[] array = dfs(root);
        return array[1];
    }

    public int[] dfs(TreeNode root) {
        if (root == null) {
            return new int[]{Integer.MAX_VALUE / 2, 0, 0};
        }
        int[] leftArray = dfs(root.left);
        int[] rightArray = dfs(root.right);
        int[] array = new int[3];
        array[0] = leftArray[2] + rightArray[2] + 1;
        array[1] = Math.min(array[0], Math.min(leftArray[0] + rightArray[1], rightArray[0] + leftArray[1]));
        array[2] = Math.min(array[0], leftArray[1] + rightArray[1]);
        return array;
    }
}

 */