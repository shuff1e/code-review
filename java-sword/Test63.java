public class Test63 
{
	// 对于index为i，计算当以i卖出时，能获得的最大利润
	public static int getMax(int[] arr)
	{
		if (arr==null||arr.length<2)
		{
			return 0;
		}
		int min=arr[0];
		int maxDiff=arr[1]-min;
		for (int i=2;i<arr.length ;i++ )
		{
			if (arr[i-1]<min)
			{
				min=arr[i-1];
			}
			int curDiff=arr[i]-min;
			if (curDiff>maxDiff)
			{
				maxDiff=curDiff;
			}
		}
		return maxDiff;
	}
	public static void main(String[] args) 
	{
		int[] arr=new int[]{9,11,8,5,7,12,16,14};
		arr=new int[]{9,11};
		System.out.println(getMax(arr));
	}
}
