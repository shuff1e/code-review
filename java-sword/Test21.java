import java.util.*;
public class Test21 
{
	public static void reOrder(int[] arr)
	{
		if (arr==null)
		{
			return;
		}
		int start=0;
		int end=arr.length-1;
		while (start<end)
		{
			while ((start<end)&&((arr[start]&1)==1))
			{
				start++;
			}
			while ((start<end)&&((arr[end]&1)==0))
			{
				end--;
			}
			if (start<end)
			{
				arr[start]=arr[end]+(arr[end]=arr[start])*0;
			}
		}
	}
	public static void reOrder(int[] arr,Checker checker)
    {
		if (arr==null)
		{
			return;
		}
		int start=0;
		int end=arr.length-1;
		while (start<end)
		{
			while ((start<end)&&checker.check(arr[start]))
			{
				start++;
			}
			while ((start<end)&&!checker.check(arr[end]))
			{
				end--;
			}
			if (start<end)
			{
				arr[start]=arr[end]+(arr[end]=arr[start])*0;
			}
		}
	}
	public static void main(String[] args) 
	{
		int[] arr=new int[]{1,5,3,4,-4,6,7,91};
		arr=new int[]{};
		reOrder(arr,new Checker()
		{
			public boolean check(int val)
			{
				return (val&1)==1;
			}
		});
		System.out.println(Arrays.toString(arr));
	}
}
interface Checker
{
	boolean check(int val);
}
