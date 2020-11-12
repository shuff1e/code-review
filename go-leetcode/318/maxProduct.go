package main

import "fmt"

/*

318. 最大单词长度乘积
给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。

示例 1:

输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
输出: 16
解释: 这两个单词为 "abcw", "xtfn"。
示例 2:

输入: ["a","ab","abc","d","cd","bcd","abcd"]
输出: 4
解释: 这两个单词为 "ab", "cd"。
示例 3:

输入: ["a","aa","aaa","aaaa"]
输出: 0
解释: 不存在这样的两个单词。

 */

func main() {
	words := []string{"a","ab","abc","d","cd","bcd","abcd"}
	fmt.Println(maxProduct(words))
}

func maxProduct(words []string) int {
	dict := map[int]int{}
	for i := 0;i<len(words);i++ {
		bitMask := 0
		for j := 0;j<len(words[i]);j++ {
			bitMask |= 1 << bitNumber(words[i][j])
		}
		dict[bitMask] = Max(dict[bitMask],len(words[i]))
	}

	result := 0
	for k1,v1 := range dict {
		for k2,v2 := range dict {
			if k1 & k2 == 0 {
				result = Max(result,v1*v2)
			}
		}
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func bitNumber(x byte) int {
	return int(x - 'a')
}

/*

class Solution {
  public boolean noCommonLetters(String s1, String s2){
    // TODO
  }

  public int maxProduct(String[] words) {
    int n = words.length;

    int maxProd = 0;
    for (int i = 0; i < n; ++i)
      for (int j = i + 1; j < n; ++j)
        if (noCommonLetters(words[i], words[j]))
          maxProd = Math.max(maxProd, words[i].length() * words[j].length());

    return maxProd;
  }
}

 */

/*

public int bitNumber(char ch) {
  return (int)ch - (int)'a';
}

public boolean noCommonLetters(String s1, String s2) {
  int bitmask1 = 0, bitmask2 = 0;
  for (char ch : s1.toCharArray())
    bitmask1 |= 1 << bitNumber(ch);
  for (char ch : s2.toCharArray())
    bitmask2 |= 1 << bitNumber(ch);

  return (bitmask1 & bitmask2) == 0;
}

 */

/*

class Solution {
  public int bitNumber(char ch) {
    return (int)ch - (int)'a';
  }

  public int maxProduct(String[] words) {
    int n = words.length;
    int[] masks = new int[n];
    int[] lens = new int[n];

    int bitmask = 0;
    for (int i = 0; i < n; ++i) {
      bitmask = 0;
      for (char ch : words[i].toCharArray()) {
        // add bit number bit_number in bitmask
        bitmask |= 1 << bitNumber(ch);
      }
      masks[i] = bitmask;
      lens[i] = words[i].length();
    }

    int maxVal = 0;
    for (int i = 0; i < n; ++i)
      for (int j = i + 1; j < n; ++j)
        if ((masks[i] & masks[j]) == 0)
          maxVal = Math.max(maxVal, lens[i] * lens[j]);

    return maxVal;
  }
}

 */

/*

class Solution {
  public int bitNumber(char ch){
    return (int)ch - (int)'a';
  }

  public int maxProduct(String[] words) {
    Map<Integer, Integer> hashmap = new HashMap();

    int bitmask = 0, bitNum = 0;
    for (String word : words) {
      bitmask = 0;
      for (char ch : word.toCharArray()) {
        // add bit number bitNumber in bitmask
        bitmask |= 1 << bitNumber(ch);
      }
      // there could be different words with the same bitmask
      // ex. ab and aabb
      hashmap.put(bitmask, Math.max(hashmap.getOrDefault(bitmask, 0), word.length()));
    }

    int maxProd = 0;
    for (int x : hashmap.keySet())
      for (int y : hashmap.keySet())
        if ((x & y) == 0) maxProd = Math.max(maxProd, hashmap.get(x) * hashmap.get(y));

    return maxProd;
  }
}

 */