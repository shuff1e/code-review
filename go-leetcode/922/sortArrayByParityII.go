package main

import "fmt"

/*

922. 按奇偶排序数组 II
给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。

对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

你可以返回任何满足上述条件的数组作为答案。



示例：

输入：[4,2,5,7]
输出：[4,5,2,7]
解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。


提示：

2 <= A.length <= 20000
A.length % 2 == 0
0 <= A[i] <= 1000

 */

func main() {
	arr := []int{2,4,6,1,3,8,5,10,7,9}
	sortArrayByParityII(arr)
	fmt.Printf("%#v\n",arr)
}

func sortArrayByParityII(A []int) []int {
	for i := 0;i<len(A);i++ {
		mark := i
		for (A[i] % 2 == 1 && i %2 ==0) || (A[i] % 2==0 && i % 2 == 1) {
			mark ++
			A[i],A[mark] = A[mark],A[i]
		}
	}
	return A
}

// 双指针

/*

class Solution {
    public int[] sortArrayByParityII(int[] A) {
        int i=0,j=A.length-1,temp;
        while(i<A.length-1&&j>0)
        {
            while(i<A.length-1&&A[i]%2==0){i+=2;}     //定位到序号为偶数，值为奇数的地方（与要求不符）
            while(j>0&&A[j]%2!=0){j-=2;}      //定位到序号为奇数，值为偶数的地方（与要求不符）
            if(i<A.length-1&&j>0)     //交换定位到的两个值
            {
                temp=A[i];
                A[i]=A[j];
                A[j]=temp;
                i+=2;
                j-=2;
            }
        }
        return A;
    }
}

 */

/*

class Solution {
    public int[] sortArrayByParityII(int[] A) {
        int j = 1;
        for (int i = 0; i < A.length; i += 2)
            if (A[i] % 2 == 1) {
                while (A[j] % 2 == 1)
                    j += 2;

                // Swap A[i] and A[j]
                int tmp = A[i];
                A[i] = A[j];
                A[j] = tmp;
            }

        return A;
    }
}

 */