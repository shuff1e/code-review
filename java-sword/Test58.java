public class Test58 
{
	public static void reverse(char[] arr,int left,int right)
	{
		if (arr==null||arr.length==0||left>=right)
		{
			return;
		}
		while (left<right)
		{
			char temp=arr[left];
			arr[left]=arr[right];
			arr[right]=temp;
			left++;
			right--;
		}
	}
	public static void reverseSentence(char[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return;
		}
		int left=0;
		int right=0;
		reverse(arr,left,arr.length-1);
		while (left<arr.length)
		{
			if (arr[left]==' ')
			{
				left++;
				right++;
			}
			else if (right==arr.length||arr[right]==' ')
			{
				reverse(arr,left,--right);
				left=++right;
			}
			else
				right++;
		}
	}
	public static void main(String[] args) 
	{
		String str="   ";
		/*
		char[] arr=str.toCharArray();
		reverseSentence(arr);
		str=new String(arr);
		System.out.println(str);
		System.out.println(str.length());
		*/
		str="abcdefg";
		char[] arr=str.toCharArray();
		rotateLeft(arr,6);
		str=new String(arr);
		System.out.println(str);
	}
	public static void rotateLeft(char[] arr,int k)
	{
		if (arr!=null)
		{
			int len=arr.length;
			if (len>0&&k>0&&k<len)
			{
				int first=0;
				int mid=k-1;
				int last=arr.length-1;
				reverse(arr,first,mid);
				reverse(arr,mid+1,last);
				reverse(arr,first,last);
			}
		}
	}
}
