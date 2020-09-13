public class Test67 
{
	public static boolean flag;
	public static int strToInt(String str)
	{
		long num=0;
		if (str!=null&&str.length()>0)
		{
			char[] arr=str.toCharArray();
			boolean minus=false;
			int index=0;
			if (arr[index]=='+')
			{
				index++;
			}
			else if (arr[index]=='-')
			{
				minus=true;
				index++;
			}
			if (index<arr.length)
			{
				num=strToIntCore(arr,index,minus);
			}
		}
		return (int)num;
	}
	public static long strToIntCore(char[] arr,int index,boolean minus)
	{
		long num=0;
		int helper=minus?-1:1;
		while (index<arr.length)
		{
			if (arr[index]>='0'&&arr[index]<='9')
			{
				num=num*10+helper*(arr[index]-'0');
			}
			else
			{
				num=0;
				break;
			}
			if ((!minus&&num>0x7fff_ffff)||
				(minus&&num<0x8000_0000))
			{
				num=0;
				break;
			}
			index++;
		}
		if (index==arr.length)
		{
			flag=true;
		}
		return num;
	}
	public static void main(String[] args) 
	{
		System.out.println(flag+"->"+strToInt("999999999999999999999999999"));
	}
}
