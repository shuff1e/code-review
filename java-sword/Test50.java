public class Test50 
{
	public static char getOnlyOnce(char[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return '\0';
		}
		int[] helper=new int[256];
		for (int i=0;i<arr.length ;i++ )
		{
			helper[arr[i]]++;
		}
		for (int i=0;i<helper.length ;i++ )
		{
			if (helper[i]==1)
			{
				return (char)i;
			}
		}
		return '\0';
	}
	public static void main(String[] args) 
	{
		String str="aabcddcgb";
		System.out.println(getOnlyOnce(str.toCharArray()));
		System.out.println(getOnlyOnce2(str.toCharArray()));
		String str1="We are students";
	    String str2="aeiou";
		System.out.println(deleteDuplicate(str1,str2));
		String str3="google";
		System.out.println(deleteDuplicate(str3));
	}

	// 类似于n皇后问题，n皇后中用一个record int[]数组标记是否访问过某个点
	// record用一个数字的二进制位代替

	// 这里如果char没出现过，把major的相应位置位
	// 如果char已经出现过，将helper置位（helper表示那些出现多次的char）
	// 将major和helper异或，这样就表示只出现一次的）
	public static char getOnlyOnce2(char[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return '\0';
		}
		int major=0;
		int helper=0;
		for (int i=0;i<arr.length ;i++ )
		{
			if ((major&(1<<(arr[i]-'a')))>0)
			{
				helper|=(1<<(arr[i]-'a'));
			}
			else
			{
				major|=(1<<(arr[i]-'a'));
			}
		}
		helper ^=major;
		for (int i=0;i<helper ;i++ )
		{
			if ((helper&(1<<i))>0)
			{
				return (char)(i+'a');
			}
		}
		return '\0';
	}
	public static String deleteDuplicate(String str1,String str2)
	{
		char[] arr1=str1.toCharArray();
		char[] arr2=str2.toCharArray();
		int[] helper=new int[256];
		for (int i=0;i<arr2.length ;i++ )
		{
			helper[arr2[i]]++;
		}
		for (int i=0;i<arr1.length ;i++ )
		{
			if (helper[arr1[i]]!=0)
			{
				arr1[i]='$';
			}
		}
		return new String(arr1).replace("$","");
	}
	public static String deleteDuplicate(String str1)
	{
		char[] arr1=str1.toCharArray();
		boolean[] helper=new boolean[256];
		for (int i=0;i<arr1.length ;i++ )
		{
			System.out.println(arr1[i]);
			if (helper[arr1[i]]==false)
			{
				helper[arr1[i]]=true;
			}
			else if (helper[arr1[i]]==true)
			{
				arr1[i]='$';
			}
		}
		return new String(arr1).replace("$","");
	}
}
