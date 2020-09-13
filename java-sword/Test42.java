public class Test42 
{
	public static boolean isInvalidInput=false;
	public static int findMaxSum(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			isInvalidInput=true;
			return Integer.MIN_VALUE;
		}
		int curSum=0;
		int maxSum=Integer.MIN_VALUE;
		for (int i=0;i<arr.length ;i++ )
		{
			if (curSum<=0)
			{
				curSum=arr[i];
			}
			else
			{
				curSum+=arr[i];
			}
			if (curSum>maxSum)
			{
				maxSum=curSum;
			}
		}
		return maxSum;
	}
	public static void main(String[] args) 
	{
		int[] arr=new int[]{1,-2,3,10,-4,7,2,-5};
		//arr=new int[]{1};
		System.out.println(findMaxSum2(arr));
	}
	public static int DP(int[] arr,int index,int[] sum)
	{
		if (index==0)
		{
			return sum[index]=arr[index];
		}
		int temp=DP(arr,index-1,sum);
		if (temp>0)
		{
			return sum[index]=arr[index]+temp;
		}
		else
		{
			return sum[index]=arr[index];
		}
	}
	public static int findMaxSum2(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			isInvalidInput=true;
			return Integer.MIN_VALUE;
		}
		int[] sum=new int[arr.length];
		int max=Integer.MIN_VALUE;
		DP(arr,arr.length-1,sum);
		for (int i=0;i<sum.length ;i++ )
		{
			if (sum[i]>max)
			{
				max=sum[i];
			}
		}
		return max;
	}
}
