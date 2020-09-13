import java.util.*;
public class Test14 
{
	public static int max(int n)
	{
		if (n<2)
		{
			return 0;
		}
		if (n==2)
		{
			return 1;
		}
		if (n==3)
		{
			return 2;
		}
		int[] result=new int[n+1];
		result[0]=0;
		result[1]=1;
		result[2]=2;
		result[3]=3;
		int max;
		for (int i=4;i<=n ;i++ )
		{
			max=0;
			for (int j=1;j<=i/2 ;j++ )
			{
				int temp=result[j]*result[i-j];
				if (max<temp)
				{
					max=temp;
				}
				result[i]=max;
			}
		}
		return result[n];
	}
	public static int maxRecursive(int n,HashMap<Integer,Integer> cache)
	{
		if (cache.containsKey(n))
		{
			return cache.get(n);
		}
		if (n==0)
		{
			return 0; 
		}
		int max=n;
		for (int i=1;i<=n/2 ;i++ )
		{
			int temp=maxRecursive(i,cache)*maxRecursive(n-i,cache);
			if (max<temp)
			{
				max=temp;
			}
		}
		cache.put(n,max);
		return max;
	}
	public static int maxGreedy(int n)
	{
		if (n<2)
		{
			return 0;
		}
		if (n==2)
		{
			return 1;
		}
		if (n==3)
		{
			return 2;
		}
		int timesOf3=n/3;
		if (n-timesOf3*3==1)
		{
			timesOf3--;
		}
		int timesOf2=(n-timesOf3*3)/2;
		return (int)(Math.pow(3,timesOf3)*Math.pow(2,timesOf2));
	}
	public static void main(String[] args) 
	{
		System.out.println(maxGreedy(8));
	}
	//两次买入卖出股票的问题
	public static int calMax(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return 0;
		}
		int len=arr.length;
		int max=0;
		for (int i=0;i<len ;i++ )
		{
			int temp=getMax(arr,0,i)+getMax(arr,i,len-1);
			if (max<temp)
			{
				max=temp;
			}
		}
		return max;
	}
	public static int getMax(int[] arr,int start,int end)
	{
		if (start>=end)
		{
			return 0;
		}
		int min=start;
		int max=0;
		for (int i=start;i<=end ;i++ )
		{
			int temp=arr[i]-arr[start];
			if (max<temp)
			{
				max=temp;
			}
			if (arr[i]<arr[min])
			{
				min=i;
			}
		}
		return max;
	}
}
