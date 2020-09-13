import java.util.*;
public class Test30 
{
	public static Deque<Integer> major=new ArrayDeque<>();
	public static Deque<Integer> helper=new ArrayDeque<>();
	public static void push(int data)
	{
		major.push(data);
		if (helper.size()==0||data<=helper.peek())
		{
			helper.push(data);
		}
		else 
		{
			helper.push(helper.peek());
		}
	}
	public static int pop()
	{
		if (helper.size()>0&&major.size()>0)
		{
			helper.pop();
			return major.pop();
		}
		else
			return -1;
	}
	public static int min()
	{
		if (helper.size()>0)
		{
			return helper.peek();
		}
		return -1;
	}
	public static void main(String[] args) 
	{
		push(3);
		push(30);
		push(2);
		push(4);
		System.out.println(min());
		pop();
		System.out.println(min());
		pop();
		System.out.println(min());
	}
}
