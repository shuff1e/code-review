public class Test19 
{
	public static boolean match(String str,String pattern)
	{
		if (str==null||pattern==null)
		{
			return false;
		}
		return matchCore(str,pattern,0,0);
	}
	public static boolean matchCore(String str,String pattern,int index1,int index2)
	{
		int len1=str.length();
		int len2=pattern.length();
		if (index1==len1&&index2==len2)
		{
			return true;
		}
		if (index1>=len1||index2>=len2)
		{
			return false;
		}
		if ((index2+1)<len2&&pattern.charAt(index2+1)=='*')
		{
			if (str.charAt(index1)==pattern.charAt(index2)
		||pattern.charAt(index2)=='.')
			{
				return matchCore(str,pattern,index1,index2+2)//出现0次
					||matchCore(str,pattern,index1+1,index2+2)//出现一次
					||matchCore(str,pattern,index1+1,index2);//出现大于一次
			}
			return matchCore(str,pattern,index1,index2+2);
		}
		if (str.charAt(index1)==pattern.charAt(index2)
		||pattern.charAt(index2)=='.')
		{
			return matchCore(str,pattern,index1+1,index2+1);
		}
			return false;
	}
	public static void main(String[] args) 
	{
		String str="baaaaaaacsb";
		String pattern=".*";
		System.out.println(match(str,pattern));
	}
}
