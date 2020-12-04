package main

import (
	"fmt"
	"strings"
)

/*

331. 验证二叉树的前序序列化
序列化二叉树的一种方法是使用前序遍历。
当我们遇到一个非空节点时，我们可以记录下这个节点的值。
如果它是一个空节点，我们可以使用一个标记值记录，例如 #。

     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
例如，上面的二叉树可以被序列化为字符串 "9,3,4,#,#,1,#,#,2,#,6,#,#"，其中 # 代表一个空节点。

给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。

每个以逗号分隔的字符或为一个整数或为一个表示 null 指针的 '#' 。

你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 "1,,3" 。

示例 1:

输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
输出: true
示例 2:

输入: "1,#"
输出: false
示例 3:

输入: "9,#,#,1"
输出: false

 */

func main() {
	str := "9,3,4,#,#,1,#,#,2,#,6,#,#"
	str = "9,#,#,1"
	str = "1,#"
	str =  "9,3,4,#,#,1,#,#,2,#,6,#,#"
	fmt.Println(isValidSerialization(str))
}

func isValidSerialization(preorder string) bool {
	index := 0
	arr := strings.Split(preorder,",")
	ok := help(arr,&index)
	return ok && index == len(arr)
}

func help(arr []string,index *int) bool {
	if *index >= len(arr) {
		return false
	}

	if arr[*index] == "#" {
		*index ++
		return true
	}

	*index ++

	leftOK := help(arr,index)
	rightOK := help(arr,index)
	return leftOK && rightOK
}

/*

首先不考虑最优性，从最简单的解法来讨论这个问题。

我们可以定义一个概念，叫做槽位，二叉树中任意一个节点或者空孩子节点都要占据一个槽位。二叉树的建立也伴随着槽位数量的变化。开始时只有一个槽位，如果根节点是空节点，就只消耗掉一个槽位，如果根节点不是空节点，除了消耗一个槽位，还要为孩子节点增加两个新的槽位。之后的节点也是同理。

有了上面的讨论，方法就很简单了。依次遍历前序序列化，根据节点是否为空，按照规则消耗/增加槽位。如果最后可以将所有的槽位消耗完，那么这个前序序列化就是合法的。

开始时只有一个可用槽位。

空节点和非空节点都消耗一个槽位。

空节点不增加槽位，非空节点增加两个槽位。

class Solution {
  public boolean isValidSerialization(String preorder) {
    // number of available slots
    int slots = 1;

    for(String node : preorder.split(",")) {
      // one node takes one slot
      --slots;

      // no more slots available
      if (slots < 0) return false;

      // non-empty node creates two children slots
      if (!node.equals("#")) slots += 2;
    }

    // all slots should be used up
    return slots == 0;
  }
}

class Solution {
  public boolean isValidSerialization(String preorder) {
    // number of available slots
    int slots = 1;

    int n = preorder.length();
    for(int i = 0; i < n; ++i) {
      if (preorder.charAt(i) == ',') {
        // one node takes one slot
        --slots;

        // no more slots available
        if (slots < 0) return false;

        // non-empty node creates two children slots
        if (preorder.charAt(i - 1) != '#') slots += 2;
      }
    }

    // the last node
    slots = (preorder.charAt(n - 1) == '#') ? slots - 1 : slots + 1;
    // all slots should be used up
    return slots == 0;
  }
}

 */

func isValidSerialization2(preorder string) bool {
	slots := 1
	for i := 0;i<len(preorder);i++ {
		if preorder[i] == ',' {
			slots --
			if slots < 0 {
				return false
			}
			if preorder[i-1] != '#' {
				slots += 2
			}
		}
	}
	return (preorder[len(preorder)-1] == '#') && (slots == 1)
}
