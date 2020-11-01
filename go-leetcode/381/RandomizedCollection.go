package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
381. O(1) 时间插入、删除和获取随机元素 - 允许重复
设计一个支持在平均 时间复杂度 O(1) 下， 执行以下操作的数据结构。

注意: 允许出现重复元素。

insert(val)：向集合中插入元素 val。
remove(val)：当 val 存在时，从集合中移除一个 val。
getRandom：从现有集合中随机获取一个元素。每个元素被返回的概率应该与其在集合中的数量呈线性相关。
示例:

// 初始化一个空的集合。
RandomizedCollection collection = new RandomizedCollection();

// 向集合中插入 1 。返回 true 表示集合不包含 1 。
collection.insert(1);

// 向集合中插入另一个 1 。返回 false 表示集合包含 1 。集合现在包含 [1,1] 。
collection.insert(1);

// 向集合中插入 2 ，返回 true 。集合现在包含 [1,1,2] 。
collection.insert(2);

// getRandom 应当有 2/3 的概率返回 1 ，1/3 的概率返回 2 。
collection.getRandom();

// 从集合中删除 1 ，返回 true 。集合现在包含 [1,2] 。
collection.remove(1);

// getRandom 应有相同概率返回 1 和 2 。
collection.getRandom();
 */

func main() {
	collection := Constructor()
	collection.Insert(1)
	collection.Insert(1)
	collection.Insert(2)
	collection.Insert(3)
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	collection.Remove(1)
	collection.Remove(1)
	fmt.Println(collection.Remove(1))
	fmt.Println(collection.GetRandom())
}

// 为了在O(1)时间内，能够随机获取一个元素，将每个数值存储在一个列表nums中，
// 这样获取随机元素时，只需随机生成一个列表中的索引，就能够得到一个随机元素

// 列表中元素的顺序是无关紧要的，只要它们正确地存在于列表中即可。
// 因此，在删除元素时，我们可以将被删的元素与列表中最后一个元素交换位置，随后便可以在 O(1) 时间内，从列表中去除该元素。

// 因此，用一个map，key就是数字的值，value就是一个set，记录这个数字在nums中的索引index

// 删除一个元素时，找到该元素的一个index，如果index是len(nums)-1，直接nums的 length --
// 并从set中删除该index

// 否则，将该元素与最后一个位置的元素调换位置，最后一个元素的值的set 添加该index，并删除len(nums)-1
// 删除的数字的set删除掉该index

type RandomizedCollection struct {
	arr []int
	dict map[int]map[int]struct{}
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		arr: make([]int,0),
		dict: make(map[int]map[int]struct{},0),
	}
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	this.arr = append(this.arr,val)
	// list 是val在arr中的index
	if list,ok := this.dict[val];ok {
		list[len(this.arr)-1] = struct{}{}
		return false
	}
	this.dict[val] = map[int]struct{}{len(this.arr)-1: {}}
	return true
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if _,ok := this.dict[val];!ok {
		return false
	}

	if val == this.arr[len(this.arr)-1] {
		list := this.dict[val]
		delete(list,len(this.arr)-1)
		this.arr = this.arr[:len(this.arr)-1]

		if len(list) == 0 {
			delete(this.dict,val)
		}
		return true
	}

	// list 中删除该oneIndex
	list := this.dict[val]
	oneIndex := 0
	for k,_ := range list {
		oneIndex = k
		break
	}


	// lastSet删除旧的，添加新的
	lastSet := this.dict[this.arr[len(this.arr)-1]]
	delete(lastSet,len(this.arr)-1)
	lastSet[oneIndex] = struct{}{}

	// 交换位置，并删除最后一个元素
	this.arr[oneIndex],this.arr[len(this.arr)-1] = this.arr[len(this.arr)-1],this.arr[oneIndex]
	this.arr = this.arr[:len(this.arr)-1]

	// 从list中删除最后一个元素
	delete(list,oneIndex)
	if len(list) == 0 {
		delete(this.dict,val)
	}
	return true
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	result := this.arr[int(rand.Float64()*float64(len(this.arr)))]
	return result
}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */