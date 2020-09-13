public class PlusTest
{
	public static int negate(int a)
	{
		int negate=0;
		int d=a<0?1:-1;
		while (a!=0)
		{
			negate +=d;
			a +=d;
		}
		return negate;
	}
	public static int minus(int a,int b)
	{
		return a+negate(b);
	}
	public static int abs(int a)
	{
		if (a>0)
		{
			return a;
		}
		else
		{
			return negate(a);
		}
	}
	public static int multiply(int a,int b)
	{
		if (a<b)
		{
			return multiply(b,a);
		}
		int sum=0;
		for (int i=abs(a);i>0 ;i-- )
		{
			sum +=b;
		}
		if (a<0)
		{
			sum=negate(sum);
		}
		return sum;
	}
	public int divide(int a,int b) throws java.lang.ArithmeticException
	{
		if (b==0)
		{
			throw new java.lang.ArithmeticException("ERROR");
		}
		int absa=abs(a);
		int absb=abs(b);
		int sum=0;
		int x=0;
		while (sum+absb<=absa)
		{
			sum +=absb;
			x++;
		}
		if ((a<0&&b<0)||(a>0&&b>0))
		{
			return x;
		}
		else
			return negate(x);
	}
	public static void main(String[] args) 
	{
		System.out.println("Hello World!");
	}
}
