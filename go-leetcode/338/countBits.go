package main

import "fmt"

/*

338. 比特位计数
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

示例 1:

输入: 2
输出: [0,1,1]
示例 2:

输入: 5
输出: [0,1,1,2,1,2]
进阶:

给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
要求算法的空间复杂度为O(n)。
你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。

 */

func main() {
	n := 5
	result := countBits4(n)
	fmt.Printf("%#v\n",result)
	result = countBits(n)
	fmt.Printf("%#v\n",result)
}

func countBits(num int) []int {
	result := make([]int,num+1)
	for i := 0;i<=num;i++ {
		result[i] = getCount(i)
	}
	return result
}

func getCount(x int) int {
	result := 0
	for x > 0 {
		result ++
		x &= (x-1)
	}
	return result
}

// P(x+b)=P(x)+1,b=pow(2,m) > x
func countBits2(num int) []int {
	result := make([]int,num+1)
	i := 0
	b := 1
	for b <= num {
		for i < b && (i+b) <= num {
			result[i+b] = result[i] + 1
			i++
		}
		i = 0
		b <<= 1
	}
	return result
}

/*

public class Solution {
    public int[] countBits(int num) {
        int[] ans = new int[num + 1];
        int i = 0, b = 1;
        // [0, b) is calculated
        while (b <= num) {
            // generate [b, 2b) or [b, num) from [0, b)
            while(i < b && i + b <= num){
                ans[i + b] = ans[i] + 1;
                ++i;
            }
            i = 0;   // reset i
            b <<= 1; // b = 2b
        }
        return ans;
    }
}

 */

// P(x)=P(x/2)+(xmod2)

func countBits3(num int) []int {
	result := make([]int,num+1)
	for i := 1;i<=num;i++ {
		result[i] = result[i >> 1]
		if i % 2 == 1 {
			result[i] ++
		}
	}
	return result
}

/*

public class Solution {
  public int[] countBits(int num) {
      int[] ans = new int[num + 1];
      for (int i = 1; i <= num; ++i)
        ans[i] = ans[i >> 1] + (i & 1); // x / 2 is x >> 1 and x % 2 is x & 1
      return ans;
  }
}

 */

// P(x)=P(x&(x−1))+1;

func countBits4(num int) []int {
	result := make([]int,num+1)
	for i := 1;i<=num;i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}

/*

public class Solution {
  public int[] countBits(int num) {
      int[] ans = new int[num + 1];
      for (int i = 1; i <= num; ++i)
        ans[i] = ans[i & (i - 1)] + 1;
      return ans;
  }
}

*/
