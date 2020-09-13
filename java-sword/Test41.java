import java.util.*;
public class Test41 
{
	private static TreeSet<Integer> max=new TreeSet<>();
	private static TreeSet<Integer> min=new TreeSet<>();
	public static void insert(int data)
	{
		// 当size的和是奇数，插入最大堆
		if (((max.size()+min.size())&1)==0)
		{
			if (min.size()>0&&data>min.first())
			{
				min.add(data);
				max.add(min.first());
				min.remove(min.first());
			}
			else
			{
				max.add(data);
			}
		}
		// 当size的和是偶数，插入最大堆
		else
		{
			if (max.size()>0&&max.last()>data)
			{
				max.add(data);
				min.add(max.last());
				max.remove(max.last());
			}
			else
			{
				min.add(data);
			}
		}
	}
	public static int getMedian()
	{
		int size=min.size()+max.size();
		if (size==0)
		{
			throw new RuntimeException("empty");
		}
		if ((size&1)==0)
		{
			return (max.last()+min.first())>>1;
		}
		else
		{
			return max.last();
		}
	}
	public static void main(String[] args) 
	{
		int[] arr=new int[]{1,2,3,4,5,6};
		arr=new int[]{};
		for (int i=0;i<arr.length ;i++ )
		{
			Test41.insert(arr[i]);
		}
		System.out.println(Test41.getMedian());
		/*
		for (int value:max )
		{
			System.out.println(value);
		}
		*/
	}
}
