import java.util.*;
public class Test40
{
	public static int[] getMinK(int[] arr,int k)
	{
		if (arr==null||arr.length==0||k<1||k>arr.length)
		{
			return null;
		}
		int[] output=new int[k];
		int start=0;
		int end=arr.length-1;
		int mid=k;
		int index=SortTest.partition(arr,start,end);
		while (index!=mid)
		{
			if (index>mid)
			{
				end=index-1;
				index=SortTest.partition(arr,start,end);
			}
			else
			{
				start=index+1;
				index=SortTest.partition(arr,start,end);
			}
		}
		for (int i=0;i<k ;i++ )
		{
			output[i]=arr[i];
		}
		return output;
	}
	public static void main(String[] args) 
	{
		int[] arr=new int[]{4,5,1,6,1,1,3,8};
		int[] output=getMinK(arr,4);
		for (int i=0;i<output.length ;i++ )
		{
			System.out.println(output[i]);
		}
		Set<Integer> set=getMinK2(arr,1);
		for (int data:set )
		{
			System.out.println(data);
		}
	}
	public static Set<Integer> getMinK2(int[] arr,int k)
	{
		if (arr==null||arr.length==0||k<1||k>arr.length)
		{
			return null;
		}
		TreeSet<Integer> set=new TreeSet<>();
		for (int i=0;i<arr.length ;i++ )
		{
			if (set.size()<k)
			{
				set.add(arr[i]);
			}
			else
			{
				if (arr[i]<set.last())
				{
					set.remove(set.last());
					set.add(arr[i]);
				}
			}
		}
		return set;
	}
}
