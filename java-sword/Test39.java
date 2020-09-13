public class Test39 
{
	public static int findMid(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return -1;
		}
		int start=0;
		int end =arr.length-1;
		int mid=arr.length>>1;
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
		int result=arr[index];
		if (!moreThanHalf(arr,result))
		{
			result=-1;
		}
		return result;
    }
	public static boolean moreThanHalf(int[] arr,int result)
	{
		int count=0;
		for (int i=0;i<arr.length ;i++ )
		{
			if (arr[i]==result)
			{
				count++;
			}
		}
		if (count*2<=arr.length)
		{
			return false;
		}
		return true;
	}
	public static int findMid2(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return -1;
		}
		int times=1;
		int result=arr[0];
		for (int i=1;i<arr.length ;i++ )
		{
			if (times==0)
			{
				result=arr[i];
				times=1;
			}
			else if (arr[i]==result)
			{
				times++;
			}
			else
				times--;
		}
		if (!moreThanHalf(arr,result))
		{
			result=-1;
		}
		return result;
	}
	public static void main(String[] args) 
	{
		int[] arr=new int[]{1,1};
		System.out.println(findMid2(arr));
	}
}
