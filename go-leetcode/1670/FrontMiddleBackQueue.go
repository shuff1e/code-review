package main

import "fmt"

/*

1670. 设计前中后队列
请你设计一个队列，支持在前，中，后三个位置的 push 和 pop 操作。

请你完成 FrontMiddleBack 类：

FrontMiddleBack() 初始化队列。
void pushFront(int val) 将 val 添加到队列的 最前面 。
void pushMiddle(int val) 将 val 添加到队列的 正中间 。
void pushBack(int val) 将 val 添加到队里的 最后面 。
int popFront() 将 最前面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
int popMiddle() 将 正中间 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
int popBack() 将 最后面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
请注意当有 两个 中间位置的时候，选择靠前面的位置进行操作。比方说：

将 6 添加到 [1, 2, 3, 4, 5] 的中间位置，结果数组为 [1, 2, 6, 3, 4, 5] 。
从 [1, 2, 3, 4, 5, 6] 的中间位置弹出元素，返回 3 ，数组变为 [1, 2, 4, 5, 6] 。


示例 1：

输入：
["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle", "popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
[[], [1], [2], [3], [4], [], [], [], [], []]
输出：
[null, null, null, null, null, 1, 3, 4, 2, -1]

解释：
FrontMiddleBackQueue q = new FrontMiddleBackQueue();
q.pushFront(1);   // [1]
q.pushBack(2);    // [1, 2]
q.pushMiddle(3);  // [1, 3, 2]
q.pushMiddle(4);  // [1, 4, 3, 2]
q.popFront();     // 返回 1 -> [4, 3, 2]
q.popMiddle();    // 返回 3 -> [4, 2]
q.popMiddle();    // 返回 4 -> [2]
q.popBack();      // 返回 2 -> []
q.popFront();     // 返回 -1 -> [] （队列为空）


提示：

1 <= val <= 109
最多调用 1000 次 pushFront， pushMiddle， pushBack， popFront， popMiddle 和 popBack 。

 */

func main() {
	c := Constructor()
	c.PushFront(1)
	c.PushBack(2)
	c.PushMiddle(3)
	c.PushMiddle(4)

	fmt.Println(c.PopFront())
	fmt.Println(c.PopMiddle())
	fmt.Println(c.PopMiddle())
	fmt.Println(c.PopBack())
	fmt.Println(c.PopFront())
}

type ListNode struct {
	val int
	prev *ListNode
	next *ListNode
}

type FrontMiddleBackQueue struct {
	top *ListNode // 首部的哑节点
	tail *ListNode // 尾部的哑节点
	mid *ListNode // 记录中间节点位置的指针
	length int // 记录双向链表的长度
}


func Constructor() FrontMiddleBackQueue {
	result := FrontMiddleBackQueue{
		top: &ListNode{val: -1},
		tail: &ListNode{val: -1},
	}
	result.top.next = result.tail
	result.tail.prev = result.top
	return result
}


func (this *FrontMiddleBackQueue) PushFront(val int)  {
	temp := &ListNode{val: val}

	temp.next = this.top.next
	this.top.next.prev = temp

	this.top.next = temp
	temp.prev = this.top

	this.length ++

	if this.length == 1 {
		this.mid = this.top.next
	} else if this.length % 2 == 0 {
		this.mid = this.mid.prev
	}
}


func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
	if this.length == 0 {
		this.PushFront(val)
	} else {
		temp := &ListNode{val: val}
		if this.length % 2 == 0 {
			// 放在mid后面
			temp.next = this.mid.next
			temp.prev = this.mid

			this.mid.next.prev = temp
			this.mid.next = temp
		} else {
			// 放在mid前面
			temp.next = this.mid
			temp.prev = this.mid.prev

			this.mid.prev.next = temp
			this.mid.prev = temp
		}
		this.length ++

		if this.length%2 == 0 {
			this.mid = this.mid.prev
		} else {
			this.mid = this.mid.next
		}
	}
}


func (this *FrontMiddleBackQueue) PushBack(val int)  {
	temp := &ListNode{val: val}

	this.tail.prev.next = temp
	temp.prev = this.tail.prev

	temp.next = this.tail
	this.tail.prev = temp

	this.length ++

	if this.length == 1 {
		this.mid = temp
	} else if this.length % 2 == 1 {
		this.mid = this.mid.next
	}
}


func (this *FrontMiddleBackQueue) PopFront() int {
	if this.length == 0 {
		return -1
	}
	result := this.top.next

	this.top.next = result.next
	result.next.prev = this.top


	this.length --
	if this.length == 0 {
		this.mid = nil
	} else if this.length % 2 == 1 {
		this.mid = this.mid.next
	}

	result.next = nil
	result.prev = nil

	return result.val
}


func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.length == 0 {
		return -1
	}
	result := this.mid

	result.prev.next = result.next
	result.next.prev = result.prev

	this.length --
	if this.length == 0 {
		this.mid = nil
	} else if this.length % 2 == 0 {
		this.mid = this.mid.prev
	} else {
		this.mid = this.mid.next
	}

	result.prev = nil
	result.next = nil
	return result.val
}


func (this *FrontMiddleBackQueue) PopBack() int {
	if this.length == 0 {
		return -1
	}

	result := this.tail.prev

	result.prev.next = this.tail
	this.tail.prev = result.prev

	this.length --
	if this.length == 0 {
		this.mid = nil
	} else if this.length %  2 == 0 {
		this.mid = this.mid.prev
	}

	result.next = nil
	result.prev = nil
	return result.val
}


/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */

/*

class FrontMiddleBackQueue {
    //定义结点
    class Node{
        int val;
        Node next;
        Node(int val){
            this.val = val;
        }
    }
    Node dummy;//构造辅助结点
    Node tail;//尾结点
    int size;//结点个数

    //构造方法
    public FrontMiddleBackQueue() {
        dummy = new Node(-1);
        tail = null;
        size = 0;
    }
    //在头部添加结点
    public void pushFront(int val) {
        Node node = new Node(val);
        node.next = dummy.next;
        dummy.next = node;
        if(size == 0)tail = node;
        size++;
    }
    //在中间添加节点
    public void pushMiddle(int val) {
       //找到中间节点
       Node node = getMiddle();
       //分情况：如果有奇数个结点，插到中间节点之前;否则，插入到中间节点之后
       if(size % 2 != 0)node = getFrontNode(node);
       Node t = new Node(val);
       t.next = node.next;
       node.next = t;
       size++;
       if(size == 1)tail = t;
    }
    //在末尾添加节点
    public void pushBack(int val) {
        Node node = new Node(val);
        if(tail == null)dummy.next = node;
        else tail.next = node;
        tail = node;
        size++;
    }
    //在头部删除结点
    public int popFront() {
        if(tail == null)return -1;
        Node temp = dummy.next;
        int res = temp.val;
        dummy.next = dummy.next.next;
        temp.next = null;//让回收内存
        size--;
        if(size == 0)tail = null;//这里要注意，原链表只有一个结点，回收后要更新tail
        return res;
    }
    //在中间删除结点
    public int popMiddle() {
       if(size == 0)return -1;
       //找到中间节点
       Node node = getMiddle();
       //找到中间节点的前一个结点
       node = getFrontNode(node);
       //删除中间节点
       Node temp = node.next;
       node.next = node.next.next;
       int res = temp.val;
       temp = null;
       size--;
       if(size == 0)tail = null;
       return res;
    }
    //在尾部删除结点
    public int popBack() {
        if(tail == null)return -1;
        int res = tail.val;
        Node frontOfTail = getFrontNode(tail);
        frontOfTail.next = null;
        size--;
        if(size == 0)tail = null;
        else tail = frontOfTail;
        return res;
    }
    //求一个结点的前驱结点
    public Node getFrontNode(Node node){
        Node p = dummy;
        while(p.next != null){
            if(p.next == node)return p;
            p = p.next;
        }
        return null;
    }
    //获得中间节点并返回
    public Node getMiddle(){
        Node fast = dummy;
        Node slower = dummy;
        while(fast != null && fast.next != null){
            fast = fast.next.next;
            slower = slower.next;
        }
        return slower;
    }
}

 */

/*

class ListNode{
    int val;
    ListNode last;
    ListNode next;
    public ListNode(){};
    public ListNode(int val){this.val = val;};
}
class FrontMiddleBackQueue {

    ListNode top; // 首部的哑节点
    ListNode tail; // 尾部的哑节点
    ListNode mid; // 用于记录中间节点位置的指针
    int length; // 记录双向链表长度

    public FrontMiddleBackQueue() {
        // 创建首尾的哑节点，并连接在一起
        top = new ListNode(0);
        tail = new ListNode(0);
        top.next = tail;
        tail.last = top;

        length = 0;
    }

    public void pushFront(int val) {
        ListNode node = new ListNode(val);
        // 先该节点的next指向双向链变的第一个节点（即top.next）
        // 再双向链变的第一个节点（top.next）的last指向该节点
        node.next = top.next;
        top.next.last = node;
        // 该节点与首节点连接在一起
        node.last = top;
        top.next = node;

        length++;

        // 修改mid指针的位置
        if(length == 1){
            // 说明时第一个节点，那么mid就是它
            mid = node;
        }else if(length % 2 == 0){
            // 那么mid往前移一位
            mid = mid.last;
        }
        // System.out.println("pushFront:" + length);
    }

    public void pushMiddle(int val) {
        if(length == 0){
            pushFront(val);
        }else{
            ListNode node = new ListNode(val);

           if(length % 2 == 0){
                // 偶数放mid后面
                node.next = mid.next;
                mid.next.last = node;
                mid.next = node;
                node.last = mid;
            }else{
                // 奇数放mid前面
                node.last = mid.last;
                mid.last.next = node;
                node.next = mid;
                mid.last = node;
            }


            length++;

            // 修改mid指针位置
            if(length % 2 == 0){
                mid = mid.last;
            }else{
                mid = mid.next;
            }
            // System.out.println("pushMiddle:" + length);
        }
    }

    public void pushBack(int val) {
        ListNode node = new ListNode(val);
        // 该节点的last先指向双向链表的最后一个节点（即tail.last）
        // 双向链表的最后一个节点（tail.last）的next再指向该节点
        node.last = tail.last;
        tail.last.next = node;
        // 该节点再与尾节点连接在一起
        node.next = tail;
        tail.last = node;

        length++;

        // 修改mid的位置
        if(length == 1){
            mid = node;
        }else if(length % 2 != 0){
            // 那么mid往后移一位
            mid = mid.next;
        }
        // System.out.println("pushBack:" + length);
    }

    public int popFront() {
        int ret = -1;
        if(length != 0){
            // 当双向链表不为空时
            // 删除双向链表的第一个节点，只要让首节点与第二个节点连接在一起即可
            ret = top.next.val; // 记录返回值
            top.next.next.last = top;
            top.next = top.next.next;

            length--;

            // 修改mid指针位置
            if(length == 0){
                mid = null;
            }else if(length % 2 != 0){
                mid = mid.next;
            }
            // System.out.println("popFront:" + length);
        }

        return ret;
    }

    public int popMiddle() {
        int ret = -1;
        if(length != 0){
            ret = mid.val;
            mid.next.last = mid.last;
            mid.last.next = mid.next;

            length--;

            // 修改mid指针位置
            if(length % 2 == 0){
                mid = mid.last;
            }else{
                mid = mid.next;
            }
            // System.out.println("popMiddle:" + length);
        }

        return ret;
    }

    public int popBack() {
        int ret = -1;
        if(length != 0){
            // 当双向链表不为空时
            // 删除双向链表的最后一个节点，只需要让尾节点与倒数第二个节点连接在一起即可
            ret = tail.last.val;
            tail.last.last.next = tail;
            tail.last = tail.last.last;

            length--;

            // 修改mid指针位置
            if(length == 0){
                mid = null;
            }else if(length % 2 == 0){
                mid = mid.last;
            }
            // System.out.println("popBack:" + length);
        }

        return ret;
    }
}

 */