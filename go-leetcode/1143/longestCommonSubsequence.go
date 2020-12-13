package main

import "fmt"

/*
1143. 最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。

一个字符串的 子序列 是指这样一个新的字符串：
它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。

示例 1:

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace"，它的长度为 3。
示例 2:

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc"，它的长度为 3。
示例 3:

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0。


提示:

1 <= text1.length <= 1000
1 <= text2.length <= 1000
输入的字符串只含有小写英文字符。

 */

// A：以str1[i],str2[j]结尾的公共子序列

// dp[i][j] = dp[i-1][j-1] + 1 if str1[i] == str2[j]
// dp[i][j] = max(dp[i-1][j],dp[i][j-1] if str1[i] != str2[j]

// dp[i][0] = 1 或者 0
//
// 	str1 := "1A2C3D4B56"
//	str2 := "B1D123CA45B6A"
//

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int,len(text1))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,len(text2))
	}

	if text1[0] == text2[0] {
		dp[0][0] = 1
	}

	for i := 1;i<len(text1);i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 1;j<len(text2);j++ {
		if text1[0] == text2[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1;i<len(text1);i++ {
		for j := 1;j<len(text2);j++ {
			dp[i][j] = Max(dp[i-1][j],dp[i][j-1])
			if text1[i] == text2[j] {
				dp[i][j] = Max(dp[i-1][j-1] + 1,dp[i][j])
			}
		}
	}

	return dp[len(text1)-1][len(text2)-1]
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	str1 := "ABCABBA"
	str2 := "CBABAC"
	fmt.Println(longestCommonSubsequence2(str1,str2))
	fmt.Println(longestCommonSubsequence(str1,str2))
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	str1 := text1
	str2 := text2
	_,node := buildPath(str1,str2)
	minus,_ := buildDiff(node,str1,str2)
	return len(str1)-minus
}

type PathNode interface {
	isSnake() bool
	getI() int
	getJ() int
	getPrev() PathNode
}

type Snake struct {
	i int
	j int
	prev PathNode
}

func (*Snake) isSnake() bool {
	return true
}

func (s *Snake) getI() int {
	return s.i
}

func (s *Snake) getJ() int {
	return s.j
}

func (s *Snake) getPrev() PathNode {
	return s.prev
}

type DiffNode struct {
	i int
	j int
	prev PathNode
}

func (*DiffNode) isSnake() bool {
	return false
}

func (d *DiffNode) getI() int {
	return d.i
}

func (d *DiffNode) getJ() int {
	return d.j
}

func (d *DiffNode) getPrev() PathNode {
	return d.prev
}

func buildPath(str1,str2 string) ([]PathNode,PathNode) {
	n := len(str1)
	m := len(str2)
	max := n + m + 1
	size := 1 + 2 *max
	middle := size / 2

	diagonal := make([]PathNode,size)
	diagonal[middle + 1] = &Snake{i:0,j:-1,prev: nil}

	for d := 0;d < max;d ++ {
		for k := -d;k <= d;k += 2 {
			kmiddle := middle + k
			kplus := kmiddle + 1
			kminus := kmiddle - 1

			i := 0
			prev := (PathNode)(nil)

			if k == -d || k != d && diagonal[kminus].getI() < diagonal[kplus].getI() {
				i  = diagonal[kplus].getI()
				prev = diagonal[kplus]
			} else {
				i = diagonal[kminus].getI() + 1
				prev = diagonal[kminus]
			}

			j := i-k

			diagonal[kminus] = nil
			node := (PathNode)(nil)
			node = &DiffNode{i,j,prev}
			for i < n && j < m && str1[i] == str2[j] {
				i ++
				j ++
			}

			if i > node.getI() {
				node = &Snake{i,j,node}
			}
			diagonal[kmiddle] = node
			if i >= n && j >= m {
				return diagonal,diagonal[kmiddle]
			}
		}
	}
	return nil,nil
}

func buildDiff(node PathNode,str1,str2 string) (int,[]string) {
	result := []string{}
	minus := 0
	for node != nil && node.getPrev() != nil && node.getPrev().getJ() >= 0 {
		if node.isSnake() {
			endi := node.getI()
			begini := node.getPrev().getI()
			for i := endi - 1;i >= begini;i-- {
				result = append(result, "  " + string(str1[i]))
			}
		} else {
			i := node.getI()
			j := node.getJ()
			prei := node.getPrev().getI()
			if prei < i {
				result = append(result,"- " + string(str1[i-1]))
				minus ++
			} else {
				result = append(result,"+ " + string(str2[j-1]))
			}
		}
		node = node.getPrev()
	}
	reverse(result)
	return minus,result
}

func reverse(arr []string) {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left],arr[right] = arr[right],arr[left]
		left ++
		right --
	}
}

/*

https://chenshinan.github.io/2019/05/02/git%E7%94%9F%E6%88%90diff%E5%8E%9F%E7%90%86%EF%BC%9AMyers%E5%B7%AE%E5%88%86%E7%AE%97%E6%B3%95/
git生成diff原理：Myers差分算法

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public abstract class PathNode {
    public final int i;
    public final int j;
    public final PathNode prev;

    public PathNode(int i, int j, PathNode prev) {
        this.i = i;
        this.j = j;
        this.prev = prev;
    }

    public abstract Boolean isSnake();

    @Override
    public String toString() {
        StringBuffer buf = new StringBuffer("[");
        PathNode node = this;
        while (node != null) {
            buf.append("(");
            buf.append(Integer.toString(node.i));
            buf.append(",");
            buf.append(Integer.toString(node.j));
            buf.append(")");
            node = node.prev;
        }
        buf.append("]");
        return buf.toString();
    }
}

final class Snake extends PathNode {
    public Snake(int i, int j, PathNode prev) {
        super(i, j, prev);
    }

    @Override
    public Boolean isSnake() {
        return true;
    }
}

final class DiffNode extends PathNode {
    public DiffNode(int i, int j, PathNode prev) {
        super(i, j, prev);
    }

    @Override
    public Boolean isSnake() {
        return false;
    }
}

interface Equalizer<T> {
    boolean equals(T ori,T rev);
}

class MyersDiff<T> {

    private final Equalizer<T> DEFAULT_EQUALIZER = (original, revised) -> original.equals(revised);
    private final Equalizer<T> equalizer;

    public MyersDiff() {
        equalizer = DEFAULT_EQUALIZER;
    }

    public MyersDiff(Equalizer<T> equalizer) {
        this.equalizer = equalizer;
    }

    public PathNode buildPath(List<T> orig, List<T> rev) throws Exception {
        if (orig == null)
            throw new IllegalArgumentException("original sequence is null");
        if (rev == null)
            throw new IllegalArgumentException("revised sequence is null");
        final int N = orig.size();
        final int M = rev.size();
        //最大步数（先全减后全加）
        final int MAX = N + M + 1;
        final int size = 1 + 2 * MAX;
        final int middle = size / 2;
        //构建纵坐标数组（用于存储每一步的最优路径位置）
        final PathNode diagonal[] = new PathNode[size];
        //用于获取初试位置的辅助节点
        diagonal[middle + 1] = new Snake(0, -1, null);
        //外层循环步数
        for (int d = 0; d < MAX; d++) {
            //内层循环所处偏移量，以2为步长，因为从所在位置走一步，偏移量只会相差2
            for (int k = -d; k <= d; k += 2) {
                //找出对应偏移量所在的位置，以及它上一步的位置（高位与低位）
                final int kmiddle = middle + k;
                final int kplus = kmiddle + 1;
                final int kminus = kmiddle - 1;
                //若k为-d，则一定是从上往下走，即i相同
                //若diagonal[kminus].i < diagonal[kplus].i，则最优路径一定是从上往下走，即i相同
                int i;
                PathNode prev;
                if ((k == -d) || (k != d && diagonal[kminus].i < diagonal[kplus].i)) {
                    i = diagonal[kplus].i;
                    prev = diagonal[kplus];
                } else {
                    //若k为d，则一定是从左往右走，即i+1
                    //若diagonal[kminus].i = diagonal[kplus].i，则最优路径一定是从左往右走，即i+1
                    i = diagonal[kminus].i + 1;
                    prev = diagonal[kminus];
                }
                //根据i与k，计算出j
                int j = i - k;
                //上一步的低位数据不再存储在数组中（每个k只清空低位即可全部清空）
                diagonal[kminus] = null;
                //当前是diff节点
                PathNode node = new DiffNode(i, j, prev);
                //判断被比较的两个数组中，当前位置的数据是否相同，相同，则去到对角线位置
                while (i < N && j < M && equals(orig.get(i), rev.get(j))) {
                    i++;
                    j++;
                }
                //判断是否去到对角线位置，若是，则生成snack节点，前节点为diff节点
                if (i > node.i)
                    node = new Snake(i, j, node);
                //设置当前位置到数组中
                diagonal[kmiddle] = node;
                //达到目标位置，返回当前node
                if (i >= N && j >= M) {
                    return diagonal[kmiddle];
                }
            }
        }
        throw new Exception("could not find a diff path");
    }

    private boolean equals(T orig, T rev) {
        return equalizer.equals(orig, rev);
    }


    public void buildDiff(PathNode path, List<T> orig, List<T> rev) {
        List<String> result = new ArrayList<>();
        if (path == null)
            throw new IllegalArgumentException("path is null");
        if (orig == null)
            throw new IllegalArgumentException("original sequence is null");
        if (rev == null)
            throw new IllegalArgumentException("revised sequence is null");
        while (path != null && path.prev != null && path.prev.j >= 0) {
            if (path.isSnake()) {
                int endi = path.i;
                int begini = path.prev.i;
                for (int i = endi - 1; i >= begini; i--) {
                    result.add("  " + orig.get(i));
                }
            } else {
                int i = path.i;
                int j = path.j;
                int prei = path.prev.i;
                if (prei < i) {
                    result.add("- " + orig.get(i - 1));
                } else {
                    result.add("+ " + rev.get(j - 1));
                }
            }
            path = path.prev;
        }
        Collections.reverse(result);
        for (String line : result) {
            System.out.println(line);
        }
    }
    public static void main(String[] args) {
        String oldText = "A\nB\nC\nA\nB\nB\nA";
        String newText = "C\nB\nA\nB\nA\nC";
        List<String> oldList = Arrays.asList(oldText.split("\\n"));
        List<String> newList = Arrays.asList(newText.split("\\n"));
        MyersDiff<String> myersDiff = new MyersDiff<>();
        try {
            PathNode pathNode = myersDiff.buildPath(oldList, newList);
            System.out.println(pathNode);
            myersDiff.buildDiff(pathNode, oldList, newList);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

 */