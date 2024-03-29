public class Test57 
{
	static class Wrapper
{
	public int k;
	public Wrapper(int k)
	{
		this.k=k;
	}
}
	public static boolean find(int[] arr,int sum,Wrapper wra1,Wrapper wra2)
	{
		if (arr==null||arr.length==0)
		{
			throw new RuntimeException("invalid exception");
		}
		int left=0;
		int right=arr.length-1;
		while (left<right)
		{
			int temp=arr[left]+arr[right];
			if (temp==sum)
			{
				wra1.k=left;
				wra2.k=right;
				return true;
			}
			else if (temp>sum)
			{
				right--;
			}
			else
				left++;
		}
		return false;
	}
	public static void main(String[] args) 
	{
		/*
		Wrapper wra1=new Wrapper(0);
		Wrapper wra2=new Wrapper(0);
		int[] arr=new int[]{1,2,4,7,11,15};
		System.out.println(find(arr,20,wra1,wra2));
		System.out.println(wra1.k+"->"+wra2.k);
		*/
		find(3);
	}
	// 这是个连续的正数序列
	// 2个n/2相加就是n
	// left又必须小于right
	public static void find(int n)
	{
		if (n<3)
		{
			return ;
		}
		int left=1;
		int right=2;
		int mid=(n+1)/2;
		int curSum=3;
		while (left<mid)
		{
			if (curSum==n)
			{
				print(left,right);
			}
			while (curSum>n&&left<mid)
			{
				curSum-=left;
				left++;
				if (curSum==n)
				{
					print(left,right);
				}
			}
			right++;
			curSum+=right;
		}
	}
	public static void print(int left,int right)
	{
		for (int i=left;i<=right ;i++ )
		{
			System.out.print(i+"->");
		}
		System.out.println();
	}
}