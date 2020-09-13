import java.util.*;
public class Test31 
{
	public static boolean isPopOrder(int[] pushArr,int[] popArr)
	{
		if (pushArr==null||popArr==null)
		{
			return false;
		}
		int len=pushArr.length;
		if (len==0||len!=popArr.length)
		{
			return false;
		}
		int indexPush=0;
		int indexPop=0;
		Deque<Integer> stack=new ArrayDeque<>();
		while (indexPop<len)
		{
			while (stack.isEmpty()||stack.peek()!=popArr[indexPop])
			{
				if (indexPush>=len)
				{
					return false;
				}
				stack.push(pushArr[indexPush]);
				indexPush++;
			}
			stack.pop();
			indexPop++;
		}
		return true;
	}
	public static void main(String[] args) 
	{
		int[] popArr=new int[]{1,2};
		int[] pushArr=new int[]{1,2};
		System.out.println(isPopOrder(popArr,pushArr));
	}
}
