public class Test49 
{
	public static boolean isUglyNumber(int number)
	{
		while (number%2==0)
		{
			number/=2;
		}
		while (number%3==0)
		{
			number /=3;
		}
		while (number%5==0)
		{
			number /=5;
		}
		return number==1?true:false;
	}
	public static int uglyNumber(int n)
	{
		if (n<=0)
		{
			return 0;
		}
		int count=0;
		int index=0;
		while (index<n)
		{
			count++;
			if (isUglyNumber(count))
			{
				index++;
			}
		}
		return count;
	}
	public static void main(String[] args) 
	{
		System.out.println(uglyNumber(15));
		System.out.println(uglyNumber2(15));
	}
	public static int uglyNumber2(int n)
	{
		int[] arr=new int[n];
		arr[0]=1;
		int cur=1;
		
		int multiply2=0;
		int multiply3=0;
		int multiply5=0;

		while (cur<n)
		{
			int min=min(arr[multiply2]*2,arr[multiply3]*3,arr[multiply5]*5);
			arr[cur]=min;
			while (arr[multiply2]*2<=min)
			{
				multiply2++;
			}
			while (arr[multiply3]*3<=min)
			{
				multiply3++;
			}
			while (arr[multiply5]*5<=min)
			{
				multiply5++;
			}
			cur++;
		}
		return arr[cur-1];
	}
	public static int min(int a,int b,int c)
	{
		int min=(a<b)?a:b;
		return min<c?min:c;
	}
}
