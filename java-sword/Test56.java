public class Test56 
{
	public static int[] findNumberAppearOnce(int[] arr)
	{
		if (arr==null||arr.length<2)
		{
			return null;
		}
		int temp=0;
		for (int i=0;i<arr.length ;i++ )
		{
			temp^=arr[i];
		}
		int indexOf1=findBit(temp);
		int[] result =new int[2];
		for (int i=0;i<arr.length ;i++ )
		{
			if (isBit(arr[i],indexOf1))
			{
				result[0]^=arr[i];
			}
			else
				result[1]^=arr[i];
		}
		return result;
	}
	public static boolean isBit(int n,int index)
	{
		n>>=index;
		return (n&1)>0;
	}
	public static int findBit(int n)
	{
		int index=0;
		while (((n&1)==0)&&index<32)
		{
			index++;
			n>>=1;
		}
		return index;
	}
	public static void main(String[] args) 
	{
		int[] arr=new int[]{2,2,3,3};
		/*
		int[] result=findNumberAppearOnce(arr);
		for (int i=0;i<result.length ;i++ )
		{
			System.out.println(result[i]);
		}
		*/
		arr=new int[]{-1,-1,-1,3,3,3,5,4,4,4};
		System.out.println(findNumberAppearOnce2(arr));
	}
	////////////
	////////////
	////////////
	public static int findNumberAppearOnce2(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			throw new RuntimeException("invalid input");
		}
		int[] bitSum=new int[32];
		for (int i=0;i<arr.length ;i++ )
		{
			int mask=1;
			for (int j=31;j>=0 ;j-- )
			{
				int temp=arr[i]&mask;
				if (temp!=0)
				{
					bitSum[j]++;
				}
				mask<<=1;
			}
		}
		/*
		for (int i=0;i<bitSum.length ;i++ )
		{
			System.out.println(bitSum[i]);
		}
		*/
		int result=0;
		for (int i=0;i<32 ;i++ )
		{
			result<<=1;
			result+=bitSum[i]%3;
		}
		return result;
	}
}
