import java.util.*;
public class Test9 
{
	public static class myQueue
	{
		public static Deque<Integer> stack1=new ArrayDeque<>();
		public static Deque<Integer> stack2=new ArrayDeque<>();
		public static void offer(int d)
		{
			stack1.push(d);
		}
		public static int poll()
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
	public static class myStack
	{
		public static Deque<Integer> queue1=new ArrayDeque<>();
		public static Deque<Integer> queue2=new ArrayDeque<>();
		public static void push(int d)
		{
			if (queue2.isEmpty())
			{
				queue1.offer(d);
			}
			else
				queue2.offer(d);
		}
		public static int pop()
		{
			if (queue1.isEmpty()&&queue2.isEmpty())
			{
				throw new RuntimeException("empty queue");
			}
			if (queue2.isEmpty())
			{
				while (!(queue1.size()==1))
				{
					queue2.offer(queue1.poll());
				}
				return queue1.poll();
			}
			else
			{
				while (!(queue2.size()==1))
				{
					queue1.offer(queue2.poll());
				}
				return queue2.poll();
			}
		}
	}
	public static void main(String[] args) 
	{
		myQueue stack=new myQueue();
		for (int i=0;i<5 ;i++ )
		{
			stack.offer(i);
		}
		System.out.println(stack.poll());
	}
}
