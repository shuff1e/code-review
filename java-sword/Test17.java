import java.util.*;
public class Test17 
{
	public static void printBad(int n)
	{
		int num=1;
		int i=0;
		while (i++<n)
		{
			num *=10;
		}
		for ( i=1;i<num ;i++ )
		{
			System.out.print(i+" ");
		}
	}
	public static boolean increment(char[] arr)
	{
		boolean isOverflow=false;
		int len=arr.length;
		int over=0;
		for (int i=len-1;i>=0 ;i-- )
		{
			int sum=arr[i]-'0'+over;
			if (i==len-1)
			{
				sum++;
			}
			if (sum>=10)
			{
				if (i==0)
				{
					isOverflow=true;
				}
				else
				{
					sum-=10;
					over=1;
					arr[i]=(char)('0'+sum);
				}
			}
			else
			{
				arr[i]=(char)('0'+sum);
				break;
			}
		}
		return isOverflow;
	}
	public static void printBetterHelper(char[] arr)
	{
		int len=arr.length;
		boolean isBegin0=true;
		for (int i=0;i<len ;i++ )
		{
			if (isBegin0&&arr[i]!='0')
			{
				isBegin0=false;
			}
			if (!isBegin0)
			{
				System.out.println(new String(Arrays.copyOfRange(arr,i,len)));
				break;
			}
		}
	}
	public static void printBetter(int n)
	{
		if (n<=0)
		{
			return ;
		}
		char[] arr=new char[n];
		for (int i=0;i<arr.length ;i++ )
		{
			arr[i]='0';
		}
		while (!increment(arr))
		{
			printBetterHelper(arr);
		}
	}
	public static void printBest(int n)
	{
		if (n<=0)
		{
			return;
		}
		char[] arr=new char[n];
		int len=arr.length;
		printBestHelper(arr,len,0);
	}
	public static void printBestHelper(char[] arr,int len,int index)
	{
		if (index==len)
		{
			printBetterHelper(arr);
			return;
		}
		for (int i=0;i<10 ;i++ )
		{
			arr[index]=(char)('0'+i);
			printBestHelper(arr,len,index+1);
		}
	}
	public static void main(String[] args) 
	{
		printBest(3);
	}
}
