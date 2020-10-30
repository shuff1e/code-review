package main

import (
	"fmt"
	"unsafe"
)

/*

336. 回文对
给定一组 互不相同 的单词， 找出所有不同 的索引对(i, j)，
使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。

示例 1：

输入：["abcd","dcba","lls","s","sssll"]
输出：[[0,1],[1,0],[3,2],[2,4]]
解释：可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]
示例 2：

输入：["bat","tab","cat"]
输出：[[0,1],[1,0]]
解释：可拼接成的回文串为 ["battab","tabbat"]

 */

/*

假设存在两个字符串 s1 和 s2，s1 + s2 是一个回文串，记这两个字符串的长度分别为 len1 和 len2，
​
我们分三种情况进行讨论：

len1 = len2，这种情况下 s1是s2 的翻转。

len1 > len2，这种情况下将s1拆分为左右两部分，t1和t2，其中t1是s2的翻转，t2是一个回文串

len1 < len2,这种情况下将s2拆分为左右两部分，t1和t2，其中t2是s1的翻转，t1是一个回文串
​

也就是说，我们要枚举字符串 k 的每一个前缀和后缀，判断其是否为回文串。

如果是回文串，我们就查询其剩余部分的翻转是否在给定的字符串序列中出现即可。

注意到空串也是回文串，所以我们可以将 k 拆解为 k++∅ 或 ∅+k，这样我们就能将情况 1 也解释为特殊的情况 2 或情况 3。

而要实现这些操作，我们只需要设计一个能够在一系列字符串中查询「某个字符串的子串的翻转」是否存在的数据结构，有两种实现方法：

我们可以使用字典树存储所有的字符串。在进行查询时，我们将待查询串的子串逆序地在字典树上进行遍历，即可判断其是否存在。

我们可以使用哈希表存储所有字符串的翻转串。在进行查询时，我们判断带查询串的子串是否在哈希表中出现，就等价于判断了其翻转是否存在。

 */

func main() {
	arr := []string{"abcd","dcba","lls","s","sssll"}
	//arr = []string{"a",""}
	result := palindromePairs(arr)
	fmt.Println(result)
	str := "123"
	fmt.Println(str[3:])
}

func palindromePairs(words []string) [][]int {
	list := make([]string,len(words))
	for i := 0;i<len(list);i++ {
		list[i] = reverse(words[i])
	}

	dict := map[string]int{}
	for i := 0;i<len(list);i++ {
		dict[list[i]] = i
	}

	result := make([][]int,0)
	for i := 0;i<len(words);i++ {
		m := words[i]

		for j := 0;j<=len(m);j++ {
			if isPalindrome(m,j,len(m)-1) {
				if index,ok := dict[m[:j]];ok {
					if index != i {
						result = append(result,[]int{i,index})
					}
				}
			}
			if j!= 0 && isPalindrome(m,0,j-1) {
				if index,ok := dict[m[j:]];ok {
					if index != i {
						result = append(result,[]int{index,i})
					}
				}
			}
		}
	}
	return result
}

func reverse(str string) string {
	arr := make([]byte,len(str))
	for i := 0;i<len(str);i++ {
		arr[len(str)-1-i] = str[i]
	}
	return String(arr)
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}

func isPalindrome(str string,left,right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left ++
		right --
	}
	return true
}

/*

class Solution {
    List<String> wordsRev = new ArrayList<String>();
    Map<String, Integer> indices = new HashMap<String, Integer>();

    public List<List<Integer>> palindromePairs(String[] words) {
        int n = words.length;
        for (String word: words) {
            wordsRev.add(new StringBuffer(word).reverse().toString());
        }
        for (int i = 0; i < n; ++i) {
            indices.put(wordsRev.get(i), i);
        }

        List<List<Integer>> ret = new ArrayList<List<Integer>>();
        for (int i = 0; i < n; i++) {
            String word = words[i];
            int m = words[i].length();
            if (m == 0) {
                continue;
            }
            for (int j = 0; j <= m; j++) {
                if (isPalindrome(word, j, m - 1)) {
                    int leftId = findWord(word, 0, j - 1);
                    if (leftId != -1 && leftId != i) {
                        ret.add(Arrays.asList(i, leftId));
                    }
                }
                if (j != 0 && isPalindrome(word, 0, j - 1)) {
                    int rightId = findWord(word, j, m - 1);
                    if (rightId != -1 && rightId != i) {
                        ret.add(Arrays.asList(rightId, i));
                    }
                }
            }
        }
        return ret;
    }

    public boolean isPalindrome(String s, int left, int right) {
        int len = right - left + 1;
        for (int i = 0; i < len / 2; i++) {
            if (s.charAt(left + i) != s.charAt(right - i)) {
                return false;
            }
        }
        return true;
    }

    public int findWord(String s, int left, int right) {
        return indices.getOrDefault(s.substring(left, right + 1), -1);
    }
}

 */

/*

class Solution {
    class Node {
        int[] ch = new int[26];
        int flag;

        public Node() {
            flag = -1;
        }
    }

    List<Node> tree = new ArrayList<Node>();

    public List<List<Integer>> palindromePairs(String[] words) {
        tree.add(new Node());
        int n = words.length;
        for (int i = 0; i < n; i++) {
            insert(words[i], i);
        }
        List<List<Integer>> ret = new ArrayList<List<Integer>>();
        for (int i = 0; i < n; i++) {
            int m = words[i].length();
            for (int j = 0; j <= m; j++) {
                if (isPalindrome(words[i], j, m - 1)) {
                    int leftId = findWord(words[i], 0, j - 1);
                    if (leftId != -1 && leftId != i) {
                        ret.add(Arrays.asList(i, leftId));
                    }
                }
                if (j != 0 && isPalindrome(words[i], 0, j - 1)) {
                    int rightId = findWord(words[i], j, m - 1);
                    if (rightId != -1 && rightId != i) {
                        ret.add(Arrays.asList(rightId, i));
                    }
                }
            }
        }
        return ret;
    }

    public void insert(String s, int id) {
        int len = s.length(), add = 0;
        for (int i = 0; i < len; i++) {
            int x = s.charAt(i) - 'a';
            if (tree.get(add).ch[x] == 0) {
                tree.add(new Node());
                tree.get(add).ch[x] = tree.size() - 1;
            }
            add = tree.get(add).ch[x];
        }
        tree.get(add).flag = id;
    }

    public boolean isPalindrome(String s, int left, int right) {
        int len = right - left + 1;
        for (int i = 0; i < len / 2; i++) {
            if (s.charAt(left + i) != s.charAt(right - i)) {
                return false;
            }
        }
        return true;
    }

    public int findWord(String s, int left, int right) {
        int add = 0;
        for (int i = right; i >= left; i--) {
            int x = s.charAt(i) - 'a';
            if (tree.get(add).ch[x] == 0) {
                return -1;
            }
            add = tree.get(add).ch[x];
        }
        return tree.get(add).flag;
    }
}

 */