package main

import (
	"fmt"
	"strconv"
)

/*

386. 字典序排数
给定一个整数 n, 返回从 1 到 n 的字典顺序。

例如，

给定 n =1 3，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。

请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。

 */

func main() {
	result := lexicalOrder2(13)
	fmt.Printf("%#v\n",result)
}

func lexicalOrder(n int) []int {
	result := []int{}
	str := strconv.Itoa(n)

	msb := int(str[0] - '0')
	length := len(str)

	for i := 1;i<msb;i++ {
		help(i,2,length,&result)
	}

	help2(msb,2,length,n,&result)

	for i := msb + 1;i<10;i++ {
		help(i,3,length,&result)
	}

	return result
}

func help(prefix,level,length int,result *[]int) {
	if level == length + 1 {
		*result = append(*result,prefix)
		return
	} else if level > length + 1 {
		return
	}
	*result = append(*result,prefix)
	for i := 0;i<10;i++ {
		help(prefix*10+i,level + 1,length,result)
	}
}

func help2(prefix,level,length,n int,result *[]int) {
	if level == length + 1 {
		if prefix <= n {
			*result = append(*result,prefix)
		}
		return
	} else if level > length + 1 {
		return
	}

	*result = append(*result,prefix)
	if prefix > n/Pow(10,length-level + 1) {
		if length - level >= 1 {
			for i := 0;i<10;i++ {
				help2(prefix*10+i,level+2,length,n,result)
			}
			return
		}
	}

	for i := 0;i<10;i++ {
		help2(prefix*10+i,level+1,length,n,result)
	}
}

func Pow(base,exp int) int {
	result := 1
	for exp > 0 {
		if exp & 1 > 0 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

func lexicalOrder2(n int) []int {
	result := []int{}
	for i := 1;i<10;i++ {
		dfs(i,n,&result)
	}
	return result
}

func dfs(prefix ,n int,result *[]int) {
	if prefix > n {
		return
	}
	*result = append(*result,prefix)
	for i := 0;i<10;i++ {
		dfs(prefix * 10 +i,n,result)
	}
}

/*

class Solution {
    public List<Integer> lexicalOrder(int n) {
        List<Integer> list = new ArrayList<>(n);

        for (int i = 1; i < 10; i++) { // 字典顺序从1开始
            if (i > n) {
                break;
            }

            list.add(i);
            addChildren(n, list, i);
        }

        return list;
    }

    private void addChildren(int n, List<Integer> list, int parent) {
        parent *= 10;

        for (int i = 0; i < 10; i++) { // 非最高位，取值范围为0 ~ 9
            int value = parent + i;
            if (value > n) {
                break;
            }

            list.add(value);
            addChildren(n, list, value);
        }
    }

}

class Solution {
    // public List<Integer> lexicalOrder(int n) {
    //     List<Integer> list = new ArrayList<>();
    //     int curr = 1;
    //     //10叉树的先序遍历
    //     for(int i=0;i<n;i++){
    //         list.add(curr);
    //         if(curr*10<=n){
    //             curr*=10;//进入下一层
    //         }else{
    //             if(curr>=n)   curr/=10;//如果这一层结束了
    //             curr+=1;
    //             while(curr%10==0) curr/=10;//如果>10就要返回上一层
    //         }
    //     }
    //     return list;
    // }
    public List<Integer> lexicalOrder(int n) {
        List<Integer> list = new ArrayList<>();
        for (int i = 1; i < 10; i++){
             dfs(n, i, list);
        }
        return list;
    }
    private void dfs(int n,int i,List<Integer>list){
        if(i>n){
            return ;
        }
        list.add(i);
        for(int j=0;j<=9;j++){
            dfs(n,i*10+j,list);
        }
    }

}

 */