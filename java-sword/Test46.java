import java.util.*;
public class Test46 
{
	public static int translate(int n)
	{
		if (n<0)
		{
			return 0;
		}
		char[] arr=String.valueOf(n).toCharArray();
		return getTranslationNumber(arr);
	}
	public static int getTranslationNumber(char[] arr)
	{
		int count=0;
		int[] helper=new int[arr.length];
		int len=arr.length;
		for (int i=len-1;i>=0 ;i-- )
		{
			count=0;
			if (i==len-1)
			{
				count+=1;
			}
			else
				count+=helper[i+1];
			if (i<len-1)
			{
				int digit1=arr[i]-'0';
				int digit2=arr[i+1]-'0';
				int num=digit1*10+digit2;
				if (num>=10&&num<=25)
				{
					if (i<len-2)
					{
						count+=helper[i+2];
					}
					else
					{
						count+=1;
					}
				}
			}
			helper[i]=count;
		}
		return helper[0];
	}
	public static void main(String[] args) 
	{
		System.out.println(translate(0));
		System.out.println(translateDP(0));
	}
	public static int translateDP(int n)
	{
		if (n<0)
		{
			return 0;
		}
		char[] arr=String.valueOf(n).toCharArray();
		HashMap<String,Integer> cache=new HashMap<>();
		return getNumber(arr,0,cache);
	}
	public static int getNumber(char[] arr,int index,HashMap<String,Integer> cache)
	{
		if (index>arr.length-1)
		{
			return 1;
		}
		String str=new String(arr,index,arr.length-index);
		if (cache.containsKey(str))
		{
			return cache.get(str);
		}
		int count=0;
		count+=getNumber(arr,index+1,cache);
		if (index<arr.length-1)
		{
			int digit1=arr[index]-'0';
			int digit2=arr[index+1]-'0';
			int num=digit1*10+digit2;
			if (num>=10&&num<=25)
			{
				count+=getNumber(arr,index+2,cache);
			}
		}
		cache.put(str,count);
		return count;
	}
}
