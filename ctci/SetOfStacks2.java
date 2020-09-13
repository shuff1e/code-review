import java.util.*;
public class SetOfStacks2
{
	public static List<Deque<Integer>> stacks=new ArrayList<>();
	public static int capacity=8;
	public static Deque<Integer> getLastStack()
	{
		if (stacks.size()>=1)
		{
			return stacks.get(stacks.size()-1);
		}
		else
			return null;
	}
	public static void push(int value)
	{
		Deque<Integer> last=getLastStack();
		if (last!=null&&!(last.size()==capacity))
		{
			last.push(value);
		}
		else
		{
			Deque<Integer> stack=new ArrayDeque<Integer>(capacity);
			stack.push(value);
			stacks.add(stack);
		}
	}
	public static int pop()
	{
		Deque<Integer> last=getLastStack();
		if (last==null)
		{
			return Integer.MIN_VALUE;
		}
		int value=last.pop();
		if (last.isEmpty())
		{
			stacks.remove(stacks.size()-1);
		}
		return value;
	}
	public static int popAt(int index)
	{
		if (index<0)
		{
			return Integer.MIN_VALUE;
		}
		else if (index>stacks.size()-1)
		{
			return Integer.MAX_VALUE;
		}
		else
		{
			int value=stacks.get(index).pop();
			leftShift(index);
			deleteEmpty();
			return value;
		}
	}
	public static void leftShift(int index)
	{
		Deque<Integer> temp=new ArrayDeque<>(capacity);
		while (index<stacks.size()-1)
		{
			Deque<Integer> cur=stacks.get(index+1);
			while (!cur.isEmpty())
			{
				temp.push(cur.pop());
			}
			Deque<Integer> pre=stacks.get(index);
			while (pre.size()<capacity)
			{
				if (temp.isEmpty())
				{
					return ;
				}
				pre.push(temp.pop());
			}
			while (!temp.isEmpty())
			{
				cur.push(temp.pop());
			}
			index++;
		}
	}
	public static void deleteEmpty()
	{
		Deque<Integer> last=getLastStack();
		if (last==null)
		{
			return ;
		}
		if (last.isEmpty())
		{
			stacks.remove(stacks.size()-1);
		}
	}
	public static void main(String[] args) 
	{
		for (int i=0;i<50 ;i++ )
		{
			push(i);
		}
		/*
		System.out.println("--------------");
		while (stacks.size()>0)
		{
			Deque last=stacks.get(stacks.size()-1);
			while (!last.isEmpty())
		{
			System.out.println(last.pop());
		}
		stacks.remove(stacks.size()-1);
		System.out.println("--------------");
		}
		*/
		///*
		System.out.println(popAt(3));
		System.out.println(popAt(3));
		System.out.println(popAt(3));
		//*/
	}
}
