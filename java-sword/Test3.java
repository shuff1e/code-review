import java.util.*;
public class Test3
{
	public static void find(int[] arr,HashSet<Integer> set)
	{
		if (arr==null)
		{
			return;
		}
		for (int i=0;i<arr.length ;i++ )
		{
			try{
			if ((arr[i]!=i)&&(arr[i]==arr[arr[i]]))
			{
				set.add(arr[i]);
				continue;
			}
			while (arr[i]!=i)
			{
				if (arr[i]==arr[arr[i]])
				{
					set.add(arr[i]);
					break;
				}
				swap(arr,i,arr[i]);
				
			}
			}
			catch(Exception e)
		{
			System.out.println("����Խ��"+arr[i]);
			
			
			e=new RuntimeException();
		}
		}
	}
	public static void swap(int[] arr,int a,int b)
	{
		arr[a]=arr[b]+(arr[b]=arr[a])*0;
	}
	public static void main(String[] args) 
	{
		int[] arr={1,2,3,4,5,6,7,8};
		int a=find2(arr);
		System.out.println(a);
	}
	public static int find2(int[] arr)
	{
		if (arr==null)
		{
			return -1;
		}
		Arrays.sort(arr);
		if (arr[arr.length-1]==arr.length)
		{
			return -1;
		}
		return find2Helper(arr,1,arr.length-1);
	}
	public static int find2Helper(int[] arr,int start,int end)
	{
		if (arr==null)
		{
			return -1;
		}
		if (start==end)
		{
			return start;
		}
		int mid=(start+end)/2;
		int count=0;
		for (int i=0;i<arr.length ;i++ )
		{
			if (arr[i]<=mid&&arr[i]>=start)
			{
				count++;
			}
		}
		if (count>mid-start+1)
		{
			return find2Helper(arr,start,mid);
		}
		else
		{
			return find2Helper(arr,mid+1,end);
		}
	}
}
