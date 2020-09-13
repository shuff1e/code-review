import java.util.*;
public class MyQueue
{
	private Deque<Integer> major=new ArrayDeque<>();
	private Deque<Integer> helper=new ArrayDeque<>();
	public void offer(int val)
	{
		major.push(val);
	}
	public int poll()
	{
		while (!major.isEmpty())
		{
			helper.push(major.pop());
		}
		if (helper.size()>0)
		{
			int val=helper.pop();
			while (!helper.isEmpty())
			{
				major.push(helper.pop());
			}
			return val;
		}
		return Integer.MIN_VALUE;
	}
	public static void main(String[] args) 
	{
		MyQueue mq=new MyQueue();
		for (int i=10;i>0 ;i-- )
		{
			mq.push(i);
		}
		System.out.println(mq.pop());
		System.out.println(mq.pop());
	}
	public void push(int val)
	{
		if (major.isEmpty())
		{
			major.push(val);
			return ;
		}
		while (!major.isEmpty())
		{
			int temp=major.peekLast();
			if (temp>val)
			{
				helper.push(major.pop());
			}
			else
			{
				major.push(val);
				break;
			}
		}
		if (major.isEmpty())
		{
			major.push(val);
		}
		while (!helper.isEmpty())
		{
			major.push(helper.pop());
		}
	}
	public int pop()
	{
		return major.pop();
	}
}
