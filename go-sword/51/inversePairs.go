package main

import (
	"fmt"
	"sort"
)

// 51：数组中的逆序对
// 题目：在数组中的两个数字如果前面一个数字大于后面的数字，则这两个数字组
// 成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

// 在数组{7,5,6,4}中，一共存在5个逆序对
// [7,5],[7,6],[7,4],[5,4],[6,4]

// A：将数组分为两部分，如果两部分都是排序过的，
// 左边数组的指针为P1，右边数组的指针为P2
// 如果P1>P2，则可以一次得到两个数组之间的逆序对

// 类似归并排序的过程

func mergeSort(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	count := 0
	mid := len(arr)/2
	count1 := mergeSort(arr[0:mid])
	count2 := mergeSort(arr[mid:])
	help := make([]int,len(arr))
	helpIndex := len(help)-1
	for p1,p2 := mid-1,len(arr)-1;p1>=0 || p2>=mid; {
		if p1 == -1 {
			for p2 >=mid {
				help[helpIndex] = arr[p2]
				p2 --
				helpIndex --
			}
			break
		}
		if p2 == mid-1 {
			for p1 >= 0 {
				help[helpIndex] = arr[p1]
				p1 --
				helpIndex --
			}
			break
		}

		if arr[p1] > arr[p2] {
			help[helpIndex] = arr[p1]
			// 更新count
			count += p2 - mid + 1
			p1 --
			helpIndex--
		} else {
			help[helpIndex] = arr[p2]
			p2 --
			helpIndex --
		}
	}
	for i := 0;i<len(help);i++ {
		arr[i] = help[i]
	}
	return count + count1 + count2
}

func main() {
	Test([]int{2147483647,2147483647,-2147483647,-2147483647,-2147483647,2147483647},6)
	Test([]int{ 1, 2, 3, 4, 7, 6, 5 },3)
	Test([]int{ 6, 5, 4, 3, 2, 1 },15)
	Test([]int{ 1, 2, 3, 4, 5, 6 },0)
	Test([]int{ 1 },0)
	Test([]int{ 1, 2 },0)
	Test([]int{ 2, 1 },1)
	Test([]int{ 1, 2, 1, 2, 1 },3)
	Test([]int{},0)
}

func Test(arr []int,expected int) {
	arr2 := make([]int,len(arr))
	copy(arr2,arr)
	cnt := reversePairs(arr2)

	arr3 := make([]int,len(arr))
	copy(arr3,arr)
	cnt3 := reversePairsBit(arr3)

	count := mergeSort(arr)
	fmt.Printf("%#v\n",arr)
	fmt.Println(count,expected,cnt,cnt3)
	if count != expected {
		panic("fuck")
	}
}

//---------------------------------

func reversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	dict := map[int]struct{}{}

	for i := 0;i<len(nums);i++ {
		dict[nums[i]] = struct{}{}
	}

	temp := make([]int,0)
	for k,_ := range dict {
		temp = append(temp,k)
	}

	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	for i := 0;i<len(nums);i++ {
		nums[i] = lowBound(temp,nums[i])
	}

	s := NewSegTree(len(temp))
	result := 0
	for i := 0;i<len(nums);i++ {
		temp := s.Query(nums[i] + 1,s.n)
		result += temp
		s.Update(nums[i])
	}
	return result
}

// 1,1,2,2,9,13,15
// 查询2，返回index为3
// 查询5，也返回index为3
func lowBound(arr []int,x int ) int {
	left := 0
	right := len(arr)-1
	for left <= right {
		mid := (left + right)/2
		if x < arr[mid] {
			right = mid -1
		} else {
			left = mid + 1
		}
	}
	return right
}

type segTree struct {
	n int
	nodes []int
}

func NewSegTree(n int) *segTree {
	return &segTree{
		n: n,
		nodes: make([]int,4*n),
	}
}

func (s *segTree) Query(left, right int) int {
	return s.query(0,0,s.n,left,right)
}

func (s *segTree) Update(x int) {
	s.update(0,0,s.n,x)
}

func (s *segTree) query(index int,l,r int,ql,qr int) int {
	if qr < l || r < ql {
		return 0
	}
	if ql <= l && r <= qr {
		return s.nodes[index]
	}
	mid := (l + r) /2
	return s.query(2*index+1,l,mid,ql,qr) +
		s.query(2*index+2,mid+1,r,ql,qr)
}

func (s *segTree) update(index int,l,r int,x int) {
	if l > x || r < x {
		return
	}
	s.nodes[index] ++
	if l == r {
		return
	}
	mid := (l+r)/2
	s.update(2*index+1,l,mid,x)
	s.update(2*index+2,mid+1,r,x)
}

/*

https://my.oschina.net/u/4392886/blog/4361664

#include<algorithm>
#include<cstdio>
using namespace std;
typedef long long int ll;//十年OI一场空，不开long long见祖宗
const int maxn=500005;
int t[maxn],a[maxn],n;
ll ans;
void discretization(){//离散化
    scanf("%d",&n);//n是数据总量
    for(int i=1;i<=n;i++){
        scanf("%d",a+i);//输入元素的同时copy一份，用来排序
        t[i]=a[i];
    }
    sort(t+1,t+n+1);//从小到大快排
    int m=unique(t+1,t+n+1)-t-1;//去重，unique返回去重后数组长度(这个说法极其不准确，只是便于理解，如果您想了解更多，百度搜索C++ unique)
    for(int i=1;i<=n;i++)
        a[i]=lower_bound(t+1,t+m+1,a[i])-t;//寻找a[i]的下标并且用下标覆盖掉原来的a[i]
}
struct sag{//普通的自带大常数线段树(因为出题人没有卡常习惯QwQ)
    int l,r;//如果普通线段树GG了的话，可以考虑zkw线段树优化
    ll v;
    sag *ls,*rs;
    inline void push_up() { v=ls->v+rs->v; }//维护区间和
    inline bool in_range(const int L,const int R) { return (L<=l)&&(r<=R); }
    inline bool outof_range(const int L,const int R) { return (r<L)||(R<l); }
    void update(const int L,const int R){
        if(in_range(L,R)) v++;//找到这个叶子结点，线段树中的这个元素数量+1
        else if(!outof_range(L,R)){
            ls->update(L,R);
            rs->update(L,R);
            push_up();//从下往上更新逆序对数
        }
    }
    ll query(const int L,const int R){
        if(in_range(L,R)) return v;
        if(outof_range(L,R)) return 0;
        return ls->query(L,R)+rs->query(L,R);//返回区间和
    }
}*rot;
sag byte[maxn<<1],*pool=byte;//内存池建树
sag* New(const int L,const int R){
    sag *u=pool++;
    u->l=L,u->r=R;
    if(L==R){
        u->v=0;
        u->ls=u->rs=NULL;
    }else{
        int Mid=(L+R)>>1;
        u->ls=New(L,Mid);
        u->rs=New(Mid+1,R);
        u->push_up();
    }
    return u;
}
int main(){
    discretization();
    rot=New(1,n);//建立1-n的线段树
    for(int i=1;i<=n;i++){//枚举每个元素
        ans+=rot->query(a[i]+1,n);//因为相等元素不构成逆序对，所以a[i]+1
        rot->update(a[i],a[i]);//该元素数量++
    }
    printf("%lld",ans);
    return 0;
}

 */

type Bit struct {
	n int
	nodes []int
}

func lowBit(x int) int {
	return x & (-x)
}

// 从上往下
func (b *Bit) Query(x int) int {
	result := 0
	for x > 0 {
		result += b.nodes[x]
		x -= lowBit(x)
	}
	return result
}

func (b *Bit) Update(x int,delta int) {
	for x <= b.n {
		b.nodes[x] += delta
		x += lowBit(x)
	}
}

func NewBit(n int) *Bit {
	return &Bit{
		n: n,
		nodes: make([]int,n+1),
	}
}

func reversePairsBit(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	// 离散化处理
	dict := map[int]struct{}{}
	for _,v := range nums {
		dict[v] = struct{}{}
	}
	temp := []int{}
	for k,_ := range dict {
		temp = append(temp,k)
	}

	sort.Ints(temp)

	for i := 0;i<len(nums);i++ {
		nums[i] = lowBound(temp,nums[i]) + 1
	}

	//

	result := 0
	b := NewBit(len(temp))
	for i := 0;i<len(nums);i++ {
		result += b.Query(b.n) - b.Query(nums[i])
		b.Update(nums[i],1)
	}
	return result
}
