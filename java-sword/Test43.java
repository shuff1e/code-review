public class Test43 
{
	public static int getNumberBetween1AndN(int n)
	{
		int count=0;
		for (int i=1;i<=n ;i++ )
		{
			count+=getNumberOf1(i);
		}
		return count;
	}
	public static int getNumberOf1(int num)
	{
		int count=0;
		while (num>0)
		{
			if (num%10==1)
			{
				count++;
			}
			num /=10;
		}
		return count;
	}
	public static void main(String[] args) 
	{
		System.out.println(getNumberBetween1AndN(1));
		System.out.println(getNumberBetween1AndN2(1));
	}


	
	public static int getNumberBetween1AndN2(int n)
	{
		if (n<=0)
		{
			return 0;
		}
		char[] arr=convertIntToCharArray(n);
		return getNumberOf1(arr,0);
	}
	public static char[] convertIntToCharArray(int n)
	{
		String str=String.valueOf(n);
		return str.toCharArray();
	}
	public static int getNumberOf1(char[] arr,int index)
	{
		if (index==arr.length-1&&arr[index]=='0')
		{
			return 0;
		}
		if (index==arr.length-1&&arr[index]>'0')
		{
			return 1;
		}
		int numberFirstDigit=0;
		if (arr[index]>'1')
		{
			numberFirstDigit+=Math.pow(10,arr.length-1-index);
		}
		else if (arr[index]=='1')
		{
			numberFirstDigit+=convertChaArrToInt(arr,index)-Math.pow(10,arr.length-1-index)+1;
		}
		int numberOtherDigits=(arr[index]-'0')*(arr.length-1-index)*
			(int)Math.pow(10,arr.length-2-index);
		int numberRecursive=getNumberOf1(arr,index+1);
		return numberFirstDigit+numberOtherDigits+numberRecursive;
	}
	public static int convertChaArrToInt(char[] arr,int index)
	{
		String str=new String(arr,index,arr.length-index);
		return Integer.valueOf(str);
	}
}
