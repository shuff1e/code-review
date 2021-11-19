/*

460. LFU 缓存
请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

实现 LFUCache 类：

LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
int get(int key) - 如果键存在于缓存中，则获取键的值，否则返回 -1。
void put(int key, int value) - 如果键已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量时，则应该在插入新项之前，使最不经常使用的项无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
注意「项的使用次数」就是自插入该项以来对其调用 get 和 put 函数的次数之和。使用次数会在对应项被移除后置为 0 。

为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

 */
package main

type ListNode struct {
	key int
	val int
	freq int
	prev *ListNode
	next *ListNode
}

type List struct {
	head *ListNode
	tail *ListNode
}

func NewList() *List {
	return &List{}
}

func (l *List) OfferFirst(node *ListNode) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		head := l.head
		head.prev = node
		node.next = head
		l.head = node
	}
}

func (l *List) PopLast() *ListNode {
	if l.tail == nil {
		return nil
	} else if l.head == l.tail {
		result := l.head
		l.head = nil
		l.tail = nil
		return result
	} else {
		tail := l.tail
		prev := tail.prev
		prev.next = nil
		tail.prev = nil
		l.tail = prev
		return tail
	}
}

func (l *List) Remove(node *ListNode) {
	if l.head == nil {
		return
	} else if l.head == l.tail {
		if l.head == node {
			l.head = nil
			l.tail = nil
		}
	} else if l.head == node {
		head := l.head
		next := head.next
		head.next = nil
		next.prev = nil
		l.head = next
	} else if l.tail == node {
		tail := l.tail
		prev := tail.prev
		tail.prev = nil
		prev.next = nil
		l.tail = prev
	} else {
		prev := node.prev
		next := node.next

		prev.next = next
		next.prev = prev

		node.prev = nil
		node.next = nil
	}
}

func (l *List) IsEmpty() bool {
	return l.head == nil
}


type LFUCache struct {
	freqTable map[int]*List
	keyTable map[int]*ListNode
	minFreq int
	capacity int
}


func Constructor(capacity int) LFUCache {
	return LFUCache{
		freqTable: make(map[int]*List),
		keyTable: make(map[int]*ListNode),
		minFreq: 0,
		capacity: capacity,
	}
}


func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}
	if _,ok := this.keyTable[key];!ok {
		return -1
	}

	value := this.keyTable[key]
	this.updateFreq(value)
	return value.val
}


func (this *LFUCache) Put(key int, value int)  {
	if this.capacity == 0 {
		return
	}
	if node,ok := this.keyTable[key];ok {
		node.val = value
		this.updateFreq(node)
	} else {
		if len(this.keyTable) == this.capacity {
			node := this.freqTable[this.minFreq].PopLast()
			delete(this.keyTable,node.key)
			if this.freqTable[this.minFreq].IsEmpty() {
				delete(this.freqTable,this.minFreq)
			}
		}
		if _,ok := this.freqTable[1];!ok {
			this.freqTable[1] = NewList()
		}
		node := &ListNode{
			key: key,
			val: value,
			freq: 1,
		}
		this.freqTable[1].OfferFirst(node)
		this.keyTable[key] = node
		this.minFreq = 1
	}
}

func (this *LFUCache) updateFreq(value *ListNode) {
	freq := value.freq
	value.freq += 1
	this.freqTable[freq].Remove(value)
	if _,ok := this.freqTable[freq+1];!ok {
		this.freqTable[freq+1] = NewList()
	}
	this.freqTable[freq+1].OfferFirst(value)

	if this.freqTable[freq].IsEmpty() {
		delete(this.freqTable,freq)
		if freq == this.minFreq {
			this.minFreq += 1
		}
	}
}

//func main() {
//	obj := Constructor(3)
//	obj.Put(1,100)
//	obj.Put(2,200)
//	obj.Put(3,300)
//
//	fmt.Println(obj.Get(3))
//	fmt.Println(obj.Get(3))
//	fmt.Println(obj.Get(3))
//
//
//	fmt.Println(obj.Get(2))
//	fmt.Println(obj.Get(2))
//
//	obj.Put(4,400)
//
//	fmt.Println(obj.Get(1))
//}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */


/*

方法一：哈希表 + 平衡二叉树
说明

本方法需要使用到「平衡二叉树」。在 C++ 语言中，我们可以直接使用 std::set 类作为平衡二叉树；同样在 Java 语言中，我们可以直接使用 TreeSet。但是在 Python 语言中，并没有内置的库可以用来模拟平衡二叉树。

思路和算法

首先我们定义缓存的数据结构：


struct Node {
    int cnt;
    int time;
    int key, value;

    // 我们需要实现一个 Node 类的比较函数
    // 将 cnt（使用频率）作为第一关键字，time（最近一次使用的时间）作为第二关键字
    // 下面是 C++ 语言的一个比较函数的例子
    bool operator< (const Node& rhs) const {
        return cnt == rhs.cnt ? time < rhs.time : cnt < rhs.cnt;
    }
};

其中 cnt 表示缓存使用的频率，time 表示缓存的使用时间，key 和 value 表示缓存的键值。

比较直观的想法就是我们用哈希表 key_table 以键 key 为索引存储缓存，建立一个平衡二叉树 S 来保持缓存根据 (cnt，time) 双关键字由于。在 C++ 中，我们可以使用 STL 提供的 std::set 类，set 背后的实现是红黑树：

对于 get(key) 操作，我们只要查看一下哈希表 key_table 是否有 key 这个键即可，有的话需要同时更新哈希表和集合中该缓存的使用频率以及使用时间，否则返回 -1。

对于 put(key, value) 操作，首先需要查看 key_table 中是否已有对应的键值。如果有的话操作基本等同于 get(key)，不同的是需要更新缓存的 value 值。如果没有的话相当于是新插入一个缓存，这时候需要先查看是否达到缓存容量 capacity，如果达到了的话，需要删除最近最少使用的缓存，即平衡二叉树中最左边的结点，同时删除 key_table 中对应的索引，最后向 key_table 和 S 插入新的缓存信息即可。

class LFUCache {
    // 缓存容量，时间戳
    int capacity, time;
    Map<Integer, Node> key_table;
    TreeSet<Node> S;

    public LFUCache(int capacity) {
        this.capacity = capacity;
        this.time = 0;
        key_table = new HashMap<Integer, Node>();
        S = new TreeSet<Node>();
    }

    public int get(int key) {
        if (capacity == 0) {
            return -1;
        }
        // 如果哈希表中没有键 key，返回 -1
        if (!key_table.containsKey(key)) {
            return -1;
        }
        // 从哈希表中得到旧的缓存
        Node cache = key_table.get(key);
        // 从平衡二叉树中删除旧的缓存
        S.remove(cache);
        // 将旧缓存更新
        cache.cnt += 1;
        cache.time = ++time;
        // 将新缓存重新放入哈希表和平衡二叉树中
        S.add(cache);
        key_table.put(key, cache);
        return cache.value;
    }

    public void put(int key, int value) {
        if (capacity == 0) {
            return;
        }
        if (!key_table.containsKey(key)) {
            // 如果到达缓存容量上限
            if (key_table.size() == capacity) {
                // 从哈希表和平衡二叉树中删除最近最少使用的缓存
                key_table.remove(S.first().key);
                S.remove(S.first());
            }
            // 创建新的缓存
            Node cache = new Node(1, ++time, key, value);
            // 将新缓存放入哈希表和平衡二叉树中
            key_table.put(key, cache);
            S.add(cache);
        } else {
            // 这里和 get() 函数类似
            Node cache = key_table.get(key);
            S.remove(cache);
            cache.cnt += 1;
            cache.time = ++time;
            cache.value = value;
            S.add(cache);
            key_table.put(key, cache);
        }
    }
}

class Node implements Comparable<Node> {
    int cnt, time, key, value;

    Node(int cnt, int time, int key, int value) {
        this.cnt = cnt;
        this.time = time;
        this.key = key;
        this.value = value;
    }

    public boolean equals(Object anObject) {
        if (this == anObject) {
            return true;
        }
        if (anObject instanceof Node) {
            Node rhs = (Node) anObject;
            return this.cnt == rhs.cnt && this.time == rhs.time;
        }
        return false;
    }

    public int compareTo(Node rhs) {
        return cnt == rhs.cnt ? time - rhs.time : cnt - rhs.cnt;
    }

    public int hashCode() {
        return cnt * 1000000007 + time;
    }
}


复杂度分析

时间复杂度：get 时间复杂度 O(logn)，put 时间复杂度 O(logn)，操作的时间复杂度瓶颈在于平衡二叉树的插入删除均需要 O(logn) 的时间。

空间复杂度：O(capacity)，其中 capacity 为 LFU 的缓存容量。哈希表和平衡二叉树不会存放超过缓存容量的键值对。



方法二：双哈希表
思路和算法

我们定义两个哈希表，第一个 freq_table 以频率 freq 为索引，每个索引存放一个双向链表，这个链表里存放所有使用频率为 freq 的缓存，缓存里存放三个信息，分别为键 key，值 value，以及使用频率 freq。第二个 key_table 以键值 key 为索引，每个索引存放对应缓存在 freq_table 中链表里的内存地址，这样我们就能利用两个哈希表来使得两个操作的时间复杂度均为 O(1)。同时需要记录一个当前缓存最少使用的频率 minFreq，这是为了删除操作服务的。

对于 get(key) 操作，我们能通过索引 key 在 key_table 中找到缓存在 freq_table 中的链表的内存地址，如果不存在直接返回 -1，否则我们能获取到对应缓存的相关信息，这样我们就能知道缓存的键值还有使用频率，直接返回 key 对应的值即可。

但是我们注意到 get 操作后这个缓存的使用频率加一了，所以我们需要更新缓存在哈希表 freq_table 中的位置。已知这个缓存的键 key，值 value，以及使用频率 freq，那么该缓存应该存放到 freq_table 中 freq + 1 索引下的链表中。所以我们在当前链表中 O(1) 删除该缓存对应的节点，根据情况更新 minFreq 值，然后将其O(1) 插入到 freq + 1 索引下的链表头完成更新。这其中的操作复杂度均为 O(1)。你可能会疑惑更新的时候为什么是插入到链表头，这其实是为了保证缓存在当前链表中从链表头到链表尾的插入时间是有序的，为下面的删除操作服务。

对于 put(key, value) 操作，我们先通过索引 key在 key_table 中查看是否有对应的缓存，如果有的话，其实操作等价于 get(key) 操作，唯一的区别就是我们需要将当前的缓存里的值更新为 value。如果没有的话，相当于是新加入的缓存，如果缓存已经到达容量，需要先删除最近最少使用的缓存，再进行插入。

先考虑插入，由于是新插入的，所以缓存的使用频率一定是 1，所以我们将缓存的信息插入到 freq_table 中 1 索引下的列表头即可，同时更新 key_table[key] 的信息，以及更新 minFreq = 1。

那么剩下的就是删除操作了，由于我们实时维护了 minFreq，所以我们能够知道 freq_table 里目前最少使用频率的索引，同时因为我们保证了链表中从链表头到链表尾的插入时间是有序的，所以 freq_table[minFreq] 的链表中链表尾的节点即为使用频率最小且插入时间最早的节点，我们删除它同时根据情况更新 minFreq ，整个时间复杂度均为 O(1)O(1)。


class LFUCache {
    int minfreq, capacity;
    Map<Integer, Node> key_table;
    Map<Integer, LinkedList<Node>> freq_table;

    public LFUCache(int capacity) {
        this.minfreq = 0;
        this.capacity = capacity;
        key_table = new HashMap<Integer, Node>();;
        freq_table = new HashMap<Integer, LinkedList<Node>>();
    }

    public int get(int key) {
        if (capacity == 0) {
            return -1;
        }
        if (!key_table.containsKey(key)) {
            return -1;
        }
        Node node = key_table.get(key);
        int val = node.val, freq = node.freq;
        freq_table.get(freq).remove(node);
        // 如果当前链表为空，我们需要在哈希表中删除，且更新minFreq
        if (freq_table.get(freq).size() == 0) {
            freq_table.remove(freq);
            if (minfreq == freq) {
                minfreq += 1;
            }
        }
        // 插入到 freq + 1 中
        LinkedList<Node> list = freq_table.getOrDefault(freq + 1, new LinkedList<Node>());
        list.offerFirst(new Node(key, val, freq + 1));
        freq_table.put(freq + 1, list);
        key_table.put(key, freq_table.get(freq + 1).peekFirst());
        return val;
    }

    public void put(int key, int value) {
        if (capacity == 0) {
            return;
        }
        if (!key_table.containsKey(key)) {
            // 缓存已满，需要进行删除操作
            if (key_table.size() == capacity) {
                // 通过 minFreq 拿到 freq_table[minFreq] 链表的末尾节点
                Node node = freq_table.get(minfreq).peekLast();
                key_table.remove(node.key);
                freq_table.get(minfreq).pollLast();
                if (freq_table.get(minfreq).size() == 0) {
                    freq_table.remove(minfreq);
                }
            }
            LinkedList<Node> list = freq_table.getOrDefault(1, new LinkedList<Node>());
            list.offerFirst(new Node(key, value, 1));
            freq_table.put(1, list);
            key_table.put(key, freq_table.get(1).peekFirst());
            minfreq = 1;
        } else {
            // 与 get 操作基本一致，除了需要更新缓存的值
            Node node = key_table.get(key);
            int freq = node.freq;
            freq_table.get(freq).remove(node);
            if (freq_table.get(freq).size() == 0) {
                freq_table.remove(freq);
                if (minfreq == freq) {
                    minfreq += 1;
                }
            }
            LinkedList<Node> list = freq_table.getOrDefault(freq + 1, new LinkedList<Node>());
            list.offerFirst(new Node(key, value, freq + 1));
            freq_table.put(freq + 1, list);
            key_table.put(key, freq_table.get(freq + 1).peekFirst());
        }
    }
}

class Node {
    int key, val, freq;

    Node(int key, int val, int freq) {
        this.key = key;
        this.val = val;
        this.freq = freq;
    }
}

 */
