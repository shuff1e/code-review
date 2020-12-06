package main

import (
	"container/heap"
	"fmt"
)

/*

355. 设计推特
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

postTweet(userId, tweetId): 创建一条新的推文
getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
follow(followerId, followeeId): 关注一个用户
unfollow(followerId, followeeId): 取消关注一个用户
示例:

Twitter twitter = new Twitter();

// 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
twitter.postTweet(1, 5);

// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
twitter.getNewsFeed(1);

// 用户1关注了用户2.
twitter.follow(1, 2);

// 用户2发送了一个新推文 (推文id = 6).
twitter.postTweet(2, 6);

// 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
// 推文id6应当在推文id5之前，因为它是在5之后发送的.
twitter.getNewsFeed(1);

// 用户1取消关注了用户2.
twitter.unfollow(1, 2);

// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
// 因为用户1已经不再关注用户2.
twitter.getNewsFeed(1);

 */

func main() {
	obj := Constructor();

	obj.PostTweet(1,5);
	//param_2 := obj.GetNewsFeed(1);
	//fmt.Println(param_2)

	//obj.Follow(1,2);
	obj.PostTweet(1,3)
	//fmt.Println(obj.GetNewsFeed(1))

	//obj.Unfollow(1,2)
	fmt.Println(obj.GetNewsFeed(1))

	x := 1
	y := 2
	px := &x
	py := &y
	px = py
	fmt.Println(*px)
}

type Message struct {
	id int
	timeStamp int
	next *Message
}

type Twitter struct {
	posts map[int]*Message
	followings map[int]map[int]struct{}
	time int
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		posts: make(map[int]*Message),
		followings: make(map[int]map[int]struct{}),
		time: 0,
	}
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.time ++
	temp := &Message{
		id: tweetId,
		timeStamp: this.time,
	}
	if msg,ok := this.posts[userId];ok {
		temp.next = msg
		msg = temp
		this.posts[userId] = msg
	} else {
		this.posts[userId] = temp
	}
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	pq := NewHeap()

	if post,ok := this.posts[userId];ok {
		heap.Push(pq,post)
	}

	fs := this.followings[userId]
	for f := range fs {
		post := this.posts[f]
		if post != nil {
			heap.Push(pq,post)
		}
	}

	count := 0
	result := []int{}
	for pq.Len() > 0 && count < 10 {
		temp := heap.Pop(pq)
		result = append(result,temp.(*Message).id)
		if temp.(*Message).next != nil {
			heap.Push(pq,temp.(*Message).next)
		}
		count ++
	}
	return result
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	if followerId == followeeId {
		return
	}
	if dict,ok := this.followings[followerId];ok {
		dict[followeeId] = struct{}{}
	} else {
		this.followings[followerId] = map[int]struct{}{followeeId: {}}
	}
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if followerId == followeeId {
		return
	}
	if dict,ok := this.followings[followerId];!ok {
		return
	} else {
		delete(dict,followeeId)
	}
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

type maxHeap []*Message

func NewHeap() *maxHeap {
	result := make(maxHeap,0)
	return &result
}
func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i,j int) bool {
	return m[i].timeStamp > m[j].timeStamp
}

func (m maxHeap) Swap(i,j int) {
	m[i],m[j] = m[j],m[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m,x.(*Message))
}

func (m *maxHeap) Pop() interface{} {
	result := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return result
}

/*

class Twitter {
    private class Node {
        // 哈希表存储关注人的 Id
        Set<Integer> followee;
        // 用链表存储 tweetId
        LinkedList<Integer> tweet;

        Node() {
            followee = new HashSet<Integer>();
            tweet = new LinkedList<Integer>();
        }
    }

    // getNewsFeed 检索的推文的上限以及 tweetId 的时间戳
    private int recentMax, time;
    // tweetId 对应发送的时间
    private Map<Integer, Integer> tweetTime;
    // 每个用户存储的信息
    private Map<Integer, Node> user;

    public Twitter() {
        time = 0;
        recentMax = 10;
        tweetTime = new HashMap<Integer, Integer>();
        user = new HashMap<Integer, Node>();
    }

    // 初始化
    public void init(int userId) {
        user.put(userId, new Node());
    }

    public void postTweet(int userId, int tweetId) {
        if (!user.containsKey(userId)) {
            init(userId);
        }
        // 达到限制，剔除链表末尾元素
        if (user.get(userId).tweet.size() == recentMax) {
            user.get(userId).tweet.remove(recentMax - 1);
        }
        user.get(userId).tweet.addFirst(tweetId);
        tweetTime.put(tweetId, ++time);
    }

    public List<Integer> getNewsFeed(int userId) {
        LinkedList<Integer> ans = new LinkedList<Integer>();

        for (int it : user.getOrDefault(userId, new Node()).tweet) {
            ans.addLast(it);
        }

        // res从ans和it中线性归并
        // res变成ans

        for (int followeeId : user.getOrDefault(userId, new Node()).followee) {
            if (followeeId == userId) { // 可能出现自己关注自己的情况
                continue;
            }

            LinkedList<Integer> res = new LinkedList<Integer>();

            int tweetSize = user.get(followeeId).tweet.size();
            Iterator<Integer> it = user.get(followeeId).tweet.iterator();

            int i = 0;
            int j = 0;
            int curr = -1;
            // 线性归并
            if (j < tweetSize) {

                curr = it.next();

                while (i < ans.size() && j < tweetSize) {
                    if (tweetTime.get(curr) > tweetTime.get(ans.get(i))) {
                    // 从curr中取
                        res.addLast(curr);
                        ++j;
                        if (it.hasNext()) {
                            curr = it.next();
                        }
                    } else {
                    // 从ans中取
                        res.addLast(ans.get(i));
                        ++i;
                    }
                    // 已经找到这两个链表合起来后最近的 recentMax 条推文
                    if (res.size() == recentMax) {
                        break;
                    }
                }
            }

            for (; i < ans.size() && res.size() < recentMax; ++i) {
                res.addLast(ans.get(i));
            }
            if (j < tweetSize && res.size() < recentMax) {
                res.addLast(curr);
                for (; it.hasNext() && res.size() < recentMax;) {
                    res.addLast(it.next());
                }
            }
            ans = new LinkedList<Integer>(res);
        }
        return ans;
    }

    public void follow(int followerId, int followeeId) {
        if (!user.containsKey(followerId)) {
            init(followerId);
        }
        if (!user.containsKey(followeeId)) {
            init(followeeId);
        }
        user.get(followerId).followee.add(followeeId);
    }

    public void unfollow(int followerId, int followeeId) {
        user.getOrDefault(followerId, new Node()).followee.remove(followeeId);
    }
}

 */