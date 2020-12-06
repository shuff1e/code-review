package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*

1206. 设计跳表
不使用任何库函数，设计一个跳表。

跳表是在 O(log(n)) 时间内完成增加、删除、搜索操作的数据结构。跳表相比于树堆与红黑树，其功能与性能相当，并且跳表的代码长度相较下更短，其设计思想与链表相似。

例如，一个跳表包含 [30, 40, 50, 60, 70, 90]，然后增加 80、45 到跳表中，以下图的方式操作：


Artyom Kalinin [CC BY-SA 3.0], via Wikimedia Commons

跳表中有很多层，每一层是一个短的链表。在第一层的作用下，增加、删除和搜索操作的时间复杂度不超过 O(n)。跳表的每一个操作的平均时间复杂度是 O(log(n))，空间复杂度是 O(n)。

在本题中，你的设计应该要包含这些函数：

bool search(int target) : 返回target是否存在于跳表中。
void add(int num): 插入一个元素到跳表。
bool erase(int num): 在跳表中删除一个值，如果 num 不存在，直接返回false. 如果存在多个 num ，删除其中任意一个即可。
了解更多 : https://en.wikipedia.org/wiki/Skip_list

注意，跳表中可能存在多个相同的值，你的代码需要处理这种情况。

样例:

Skiplist skiplist = new Skiplist();

skiplist.add(1);
skiplist.add(2);
skiplist.add(3);
skiplist.search(0);   // 返回 false
skiplist.add(4);
skiplist.search(1);   // 返回 true
skiplist.erase(0);    // 返回 false，0 不在跳表中
skiplist.erase(1);    // 返回 true
skiplist.search(1);   // 返回 false，1 已被擦除
约束条件:

0 <= num, target <= 20000
最多调用 50000 次 search, add, 以及 erase操作。

 */

func main() {
	s := Constructor()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	fmt.Println(s.Search(0))

	s.Add(4)
	fmt.Println(s.Search(1))

	fmt.Println(s.Erase(0))

	fmt.Println(s.Erase(1))
	fmt.Println(s.Search(1))
}

type skipListNode struct {
	val int
	cnt int // 当前 val 出现的次数
	levels []*skipListNode // start from 0
}

func NewSkipListNode(maxLevel int) *skipListNode {
	return &skipListNode{
		levels: make([]*skipListNode,maxLevel),
	}
}

type Skiplist struct {
	p float64
	random *rand.Rand
	level int // 当前skiplist的高度（所有数字level数最大的）,保存此level有利于查询
	max_level int
	head *skipListNode // 头节点
}


func Constructor() Skiplist {
	return Skiplist{
		p : 0.5,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
		level: 0,
		max_level: 16,
		head: NewSkipListNode(16),
	}
}


func (this *Skiplist) Search(target int) bool {
	curNode := this.head
	for i := this.level - 1;i >= 0;i-- {
		for curNode.levels[i] != nil && curNode.levels[i].val < target {
			curNode = curNode.levels[i]
		}
	}
	curNode = curNode.levels[0]
	return curNode != nil && curNode.val == target
}


func (this *Skiplist) Add(num int)  {
	curNode := this.head
	// 记录每层能访问的最右节点
	levelTails := make([]*skipListNode,this.max_level)
	for i := this.level - 1;i >= 0;i-- {
		for curNode.levels[i] != nil && curNode.levels[i].val < num {
			curNode = curNode.levels[i]
		}
		levelTails[i] = curNode
	}

	curNode = curNode.levels[0]
	if curNode != nil && curNode.val == num {
		// 已存在，cnt加1
		curNode.cnt ++
	} else {
		// 插入
		newLevel := this.randomLevel()
		if newLevel > this.level {
			for i := this.level;i < newLevel;i++ {
				levelTails[i] = this.head
			}
			this.level = newLevel
		}
		newNode := NewSkipListNode(this.max_level)
		newNode.val = num
		newNode.cnt = 1
		for i := 0;i<this.level;i++ {
			newNode.levels[i] = levelTails[i].levels[i]
			levelTails[i].levels[i] = newNode
		}
	}
}


func (this *Skiplist) Erase(num int) bool {
	curNode := this.head
	levelTails := make([]*skipListNode,this.max_level)
	for i := this.level - 1;i >= 0;i-- {
		for curNode.levels[i] != nil && curNode.levels[i].val < num {
			curNode = curNode.levels[i]
		}
		levelTails[i] = curNode
	}
	curNode = curNode.levels[0]

	if curNode != nil && curNode.val == num {
		if curNode.cnt > 1 {
			curNode.cnt --
			return true
		}
		// 存在，删除
		for i := 0;i<this.level;i++ {
			if levelTails[i].levels[i] != curNode {
				break
			}
			levelTails[i].levels[i] = curNode.levels[i]
			curNode.levels[i] = nil
		}
		for this.level > 0 && this.head.levels[this.level - 1] == nil {
			this.level --
		}
		return true
	}
	return false
}

func (this *Skiplist) randomLevel() int {
	level := 1
	for this.random.Float64() < this.p && level < this.max_level {
		level ++
	}
	if level > this.max_level {
		return this.max_level
	}
	return level
}


/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */

/*

class Skiplist {

    class SkipListNode {
        int val;
        int cnt;  // 当前val出现的次数
        SkipListNode[] levels;  // start from 0
        SkipListNode() {
            levels = new SkipListNode[MAX_LEVEL];
        }
    }

    private double p = 0.5;
    private int MAX_LEVEL = 16;
    private SkipListNode head;  // 头结点
    private int level;  //
    private Random random;

    public Skiplist() {
        // 保存此level有利于查询（以及其他操作）
        level = 0;  // 当前 skiplist的高度（所有数字 level数最大的）
        head = new SkipListNode();
        random = new Random();
    }

    // 返回target是否存在于跳表中
    public boolean search(int target) {
        SkipListNode curNode = head;
        for (int i = level-1; i >= 0; i--) {
            while (curNode.levels[i] != null && curNode.levels[i].val < target) {
                curNode = curNode.levels[i];
            }
        }
        curNode = curNode.levels[0];
        return (curNode != null && curNode.val == target);
    }

    // 插入一个元素到跳表。
    public void add(int num) {
        SkipListNode curNode = head;


        // 记录每层能访问的最右节点
        SkipListNode[] levelTails = new SkipListNode[MAX_LEVEL];
        for (int i = level-1; i >= 0; i--) {
            while (curNode.levels[i] != null && curNode.levels[i].val < num) {
                curNode = curNode.levels[i];
            }
            levelTails[i] = curNode;
        }

        curNode = curNode.levels[0];

        if (curNode != null && curNode.val == num) {
            // 已存在，cnt 加1
            curNode.cnt++;
        } else {
            // 插入
            int newLevel = randomLevel();
            if (newLevel > level) {
                for (int i = level; i < newLevel; i++) {
                    levelTails[i] = head;
                }
                level = newLevel;
            }
            SkipListNode newNode = new SkipListNode();
            newNode.val = num;
            newNode.cnt = 1;
            for (int i = 0; i < level; i++) {
                newNode.levels[i] = levelTails[i].levels[i];
                levelTails[i].levels[i] = newNode;

            }
        }
    }

    private int randomLevel() {
        int level = 1;  // 注意思考此处为什么是 1 ？
        while (random.nextDouble() < p && level < MAX_LEVEL) {
            level++;
        }
        return level > MAX_LEVEL ? MAX_LEVEL : level;
    }

    // 在跳表中删除一个值，如果 num 不存在，直接返回false. 如果存在多个 num ，删除其中任意一个即可。
    public boolean erase(int num) {
        SkipListNode curNode = head;
        // 记录每层能访问的最右节点
        SkipListNode[] levelTails = new SkipListNode[MAX_LEVEL];

        for (int i = level-1; i >= 0; i--) {
            while (curNode.levels[i] != null && curNode.levels[i].val < num) {
                curNode = curNode.levels[i];
            }
            levelTails[i] = curNode;
        }
        curNode = curNode.levels[0];
        if (curNode != null && curNode.val == num) {
            if (curNode.cnt > 1) {
                curNode.cnt--;
                return true;
            }
            // 存在，删除
            for (int i = 0; i < level; i++) {
                if (levelTails[i].levels[i] != curNode) {
                    break;
                }
                levelTails[i].levels[i] = curNode.levels[i];
            }
            while (level > 0 && head.levels[level-1] == null) {
                level--;
            }
            return true;
        }
        return false;
    }
}

 */