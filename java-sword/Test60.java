public class Test60
{
	public static int g_maxValue=6;
	// 和最小为n，最大为6*n
	// 所以数组长度为6n-n+1
	//  然后和为s出现的次数，放在index为s-n的地方
	public static void print(int n)
	{
		if (n<1)
		{
			return;
		}
		int maxSum=n*g_maxValue;
		int[] arr=new int[maxSum-n+1];
		probability(n,arr);
		int total=(int)Math.pow(g_maxValue,n);
		for (int i=n;i<=maxSum ;i++ )
		{
			double ratio=(double)arr[i-n]/total;
			System.out.println(ratio);
		}
	}
	// n是骰子的个数
	public static void probability(int n,int[] arr)
	{
		for (int i=1;i<=g_maxValue ;i++ )
		{
			probability(n,n,i,arr);
		}
	}
	// cur是当前的level
	// sum就是和，存放在index为sum-n的数组元素中
	public static void probability(int n,int cur,int sum,int[] arr)
	{
		if (cur==1)
		{
			arr[sum-n]++;
		}
		else
		{
			for (int i=1;i<=g_maxValue ;i++ )
		    {
			probability(n,cur-1,i+sum,arr);
		    }
		}
	}
	public static void main(String[] args) 
	{
		print(2);
		System.out.println("-------------");
		print2(2);
	}
	public static void print2(int n)
	{
		if (n<1)
		{
			return;
		}
		int[][] arr=new int[2][];
		arr[0]=new int[g_maxValue*n+1];
		arr[1]=new int[g_maxValue*n+1];
		int flag=0;
		for (int i=1;i<=g_maxValue ;i++ )
		{
			arr[flag][i]=1;
		}
		// f(i) = f(i-1) + ... + f(i-6)
		for (int k=2;k<=n ;k++ )
		{
			for (int i=0;i<k ;i++ )
			{
				arr[1-flag][i]=0;
			}
			// 可取值范围,k到g_maxValue*k
			for (int i=k;i<=g_maxValue*k ;i++ )
			{
				arr[1-flag][i]=0;
				for (int j=1;j<=i&&j<=g_maxValue ;j++ )
				{
					arr[1-flag][i]+=arr[flag][i-j];
				}
			}
			flag=1-flag;
		}
		int total=(int)Math.pow(g_maxValue,n);
		for (int i=n;i<=g_maxValue*n ;i++ )
		{
			double ratio=(double)arr[flag][i]/total;
			System.out.println(ratio);
		}
	}
}
