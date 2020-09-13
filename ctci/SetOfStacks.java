import java.util.*;
public class SetOfStacks 
{
	public static ArrayList<Deque<Integer>> stacks=new ArrayList<>();
	public static int capacity=8;
	public Deque<Integer> getLastStack()
	{
		if (stacks.size()>=1)
		{
			return stacks.get(stacks.size()-1);
		}
		else
			return null;
	}
	public void push(int value)
	{
		Deque<Integer> last=getLastStack();
		if (last!=null&&!(last.size()==capacity))
		{
			last.push(value);
		}
		else
		{
			Deque<Integer> stack=new ArrayDeque<Integer>(capacity);
			stacks.add(stack);
		}
	}
	public int pop()
	{
		Deque<Integer> last=getLastStack();
		int value=last.pop();
		if (last.isEmpty())
		{
			stacks.remove(stacks.size()-1);
		}
		return value;
	}
	public static void main(String[] args) 
	{
		SetOfStacks sos=new SetOfStacks();
		for (int i=0;i<50 ;i++ )
		{
			sos.push(i);
		}
		System.out.println(stacks.size());
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
	}
}
