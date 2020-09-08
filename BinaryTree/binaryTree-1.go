package BinaryTree

import "fmt"

//请把一段纸条竖着放在桌子上，然后从纸条的下边向上方对折1次，压 出折痕后展开。此时折痕是凹下去的，即折痕突起的方向指向纸条的背 面。如果从纸条的下边向上方连续对折2次，压出折痕后展开，此时有 三条折痕，从上到下依次是下折痕、下折痕和上折痕。给定一个输入参 数N，代表纸条都从下边向上方连续对折N次，请从上到下打印所有折痕 的方向。

//例如:N=1时，打印:
//down
//N=2时，打印:
//down
//down
//up
//N = 3.1，打印
//down
//down
//up
//
//down
//
//down
//up
//up
//
//N=4，打印
//down
//down
//up
//
//down
//
//down
//up
//up
//
//down
//
//down
//down
//up
//
//up
//
//down
//up
//up

// 分析可以得到上述折纸方法会得到一颗二叉树

/*                          down
                   down               up
            down          up    down       up
 */

// 所以本质上是二叉树的中序遍历

func MidOrder2(n int) {
	f(1,n,true)
}

func f(current,all int,isDown bool) {
	if current > all {
		return
	}
	f(current+1,all,true)
	if isDown {
		fmt.Println("down")
	} else {
		fmt.Println("up")
	}
	f(current+1,all,false)
}
