public class Test48
{
	public static int longest(String str)
	{
		if (str==null||str.length()==0)
		{
			return 0;
		}
		char[] arr=str.toCharArray();
		int[] position=new int[26];
		for (int i=0;i<26 ;i++ )
		{
			position[i]=-1;
		}
		int curLength=0;
		int maxLength=0;
		for (int i=0;i<arr.length ;i++ )
		{
			int prevIndex=position[arr[i]-'a'];
			if (prevIndex<0||i-prevIndex>curLength)
			{
				curLength++;
			}
			else
			{
				if (curLength>maxLength)
				{
					maxLength=curLength;
				}
				curLength=i-prevIndex;
			}
			position[arr[i]-'a']=i;
		}
		if (curLength>maxLength)
		{
			maxLength=curLength;
		}
		return maxLength;
	}
	public static void main(String[] args) 
	{
		String str="arabcacfr";
		System.out.println(longest(str));
	}
}
