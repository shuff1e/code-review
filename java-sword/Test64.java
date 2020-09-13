public class Test64 
{
	public static int getSum(int n)
	{
		int sum=0;
		boolean flag=((n!=0)&&((sum=getSum(n-1))==0));
		sum+=n;
		return sum;
	}
	public static void main(String[] args) 
	{
		System.out.println(getSum(3));
	}
}
