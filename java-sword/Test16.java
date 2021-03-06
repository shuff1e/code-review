public class Test16 
{
	public static double power(double base,int exp)
	{
		if ((base==0)&&(exp<0))
		{
			throw new ArithmeticException();
		}
		int temp=exp;
		if (exp<0)
		{
			temp=-exp;
		}
		double result=powerHelperBetter(base,exp);
		if (exp<0)
		{
			result=1/result;
		}
		return result;
	}
	public static double powerHelper(double base,int exp)
	{
		int result=1;
		for (int i=0;i<exp ;i++ )
		{
			result *=base;
		}
		return result;
	}
	public static double powerHelperBetter(double base,int exp)
	{
		if (exp==0)
		{
			return 1;
		}
		if (exp==1)
		{
			return base;
		}
		double result=powerHelperBetter(base,exp>>1);
		result *=result;
		if ((exp&1)==1)
		{
			result *=base;
		}
		return result;
	}
	public static void main(String[] args) 
	{
		System.out.println(power(3,3));
	}
}
