public class Test44 
{
	public static int digitAtIndex(int index)
	{
		if (index<0)
		{
			return -1;
		}
		int digit=1;
		while (true)
		{
			int number=numberAtIndex(digit);
			if (index<number)
			{
				return digitAtIndex(index,digit);
			}
			index -=number;
			digit++;
		}
	}
	public static int numberAtIndex(int digit)
	{
		if (digit==1)
		{
			return 10;
		}
		return 9*(int)Math.pow(10,digit-1);
	}
	public static int digitAtIndex(int index,int digit)
	{
		int number=beginNumber(digit)+index/digit;
		int countFromRight=digit-index%digit;
		for (int i=1;i<countFromRight ;i++ )
		{
			number/=10;
		}
		return number%10;
	}
	public static int beginNumber(int digit)
	{
		return (int)Math.pow(10,digit-1);
	}
	public static void main(String[] args) 
	{
		System.out.println(digitAtIndex(1));
	}
}
