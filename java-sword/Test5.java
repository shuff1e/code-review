public class Test5 
{
	public static String replace(String str)
	{
		if (str==null)
		{
			return null;
		}
		if (str.length()==0)
		{
			return null;
		}
		char[] c=str.toCharArray();
		int count=0;
		for (int i=0;i<c.length ;i++ )
		{
			if (c[i]==' ')
			{
				count++;
			}
		}
		char[] result=new char[c.length+count*2];
		int index=result.length-1;
		for (int i=c.length-1;i>=0 ;i-- )
		{
			if (c[i]==' ')
			{
				result[index--]='0';
				result[index--]='2';
				result[index--]='%';
			}
			else
			{
				result[index--]=c[i];
			}
		}
		return new String(result);
	}
	public static void main(String[] args) 
	{
		String str="We will we will rock you";
		//System.out.println(replace(str));
		str="";
		System.out.println(str.length());
		str=new String();
		System.out.println(str.length());
		str=null;
		System.out.println(str);
	}
}
