import java.util.*;
public class Test59 
{
	public  MyQueue mq=new MyQueue();
	public  void offer(int data)
	{
		mq.offer(data);
	}
	public  int poll()
	{
		return mq.poll();
	}
	public  int max()
	{
		int temp1=mq.stack1.max();
		int temp2=mq.stack2.max();
		if ((temp1==-1)&&(temp2==-1))
		{
			return -1;
		}
		return temp1>temp2?temp1:temp2;
	}

	// 实现一个栈，O(1)时间内得到最大值
	// 使用两个栈，helper栈保存当前major的最大值
	class MyStack 
{
	public  Deque<Integer> major=new ArrayDeque<>();
	public  Deque<Integer> helper=new ArrayDeque<>();
	public  void push(int data)
	{
		major.push(data);
		if (helper.size()==0||data>=helper.peek())
		{
			helper.push(data);
		}
		else 
		{
			helper.push(helper.peek());
		}
	}
	public  int pop()
	{
		if (helper.size()>0&&major.size()>0)
		{
			helper.pop();
			return major.pop();
		}
		else
			return -1;
	}
	public  int max()
	{
		if (helper.size()>0)
		{
			return helper.peek();
		}
		return -1;
	}
	public  boolean isEmpty()
	{
		return major.size()==0;
	}
}

	// 两个stack，MyStack实现一个queue
	class MyQueue
	{
		public  MyStack  stack1=new MyStack();
		public  MyStack  stack2=new MyStack();
		public  void offer(int d)
		{
			stack1.push(d);
		}
		public  int poll()
		{
			if (stack2.isEmpty()&&stack1.isEmpty())
			{
				throw new RuntimeException("empty queue");
			}
			else if (stack2.isEmpty()&&!stack1.isEmpty())
			{
				while (!stack1.isEmpty())
				{
					stack2.push(stack1.pop());
				}
				return stack2.pop();
			}
			else
			{
				return stack2.pop();
			}
		}
	}
	public static void main(String[] args) 
	{
		Test59 t=new Test59();
		int[] arr=new int[]{2,2,2,2,2,2};
		for (int i=0;i<3 ;i++ )
		{
			t.offer(arr[i]);
		}
		//System.out.println(t.max());
		for (int i=3;i<arr.length ;i++ )
		{
			t.offer(arr[i]);
			t.poll();
			//System.out.println(t.max());
		}
		List<Integer> list=t.maxInWindows(arr,3);
		for (int temp:list )
		{
			System.out.println(temp);
		}
		//System.out.println(t.poll());
		//System.out.println(t.poll());
		//System.out.println(max());
		//System.out.println(poll());
		//System.out.println(poll());
		//System.out.println(poll());
		//System.out.println(max());
	}
	// arr[i]入queue之前，将所有小于arr[i]的都removeLast出去
	public List<Integer> maxInWindows(int[] arr,int size)
	{
		List<Integer> list=new ArrayList<>();
		if (arr!=null&&arr.length>=size&&size>=1)
		{
			Deque<Integer> queue=new ArrayDeque<>();
			for (int i=0;i<size ;i++ )
			{
				while (!queue.isEmpty()&&arr[i]>=arr[queue.peekLast()])
				{
					queue.removeLast();
				}
				queue.addLast(i);
			}
			for (int i=size;i<arr.length ;i++ )
			{ 
				list.add(arr[queue.peekFirst()]);
				while (!queue.isEmpty()&&arr[i]>=arr[queue.peekLast()])
				{
					queue.removeLast();
				}
				if (!queue.isEmpty()&&(queue.peekFirst()<=(i-size)))
				{
					queue.removeFirst();
				}
				queue.addLast(i);
			}
			list.add(arr[queue.peekFirst()]);
		}
		return list;
	}
}