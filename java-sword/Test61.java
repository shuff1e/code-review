import java.util.*;
public class Test61 
{
	public static boolean isContinuous(int[] arr)
	{
		if (arr==null||arr.length<1)
		{
			return false;
		}
		Arrays.sort(arr);
		int numberOfZeros=0;
		int numberOfGaps=0;
		for (int i=0;i<arr.length&&(arr[i]==0) ;i++ )
		{
			numberOfZeros++;
		}
		int small=numberOfZeros;
		int big=small+1;
		while (big<arr.length)
		{
			if (arr[small]==arr[big])
			{
				return false;
			}
			numberOfGaps+=arr[big]-arr[small]-1;
			small=big;
			big++;
		}
		return numberOfGaps<=numberOfZeros?true:false;
	}
	public static void main(String[] args) 
	{
		int[] arr=new int[]{1,2,3,4,6};
		System.out.println(isContinuous(arr));
	}
}
