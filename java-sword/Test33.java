/*
注意变量什么含义一定要弄清楚，不要弄错了
是arr.length还是right
是arr[index]还是index
*/
public class Test33 
{
	public static boolean checkHelper(int[] arr,int left,int right)
	{
		if (left>=right)
		{
			return true;
		}
		int root=arr[right];
		int temp=left;
		//没有右子树
		while (temp<(right)&&arr[temp]<=root )
		{
			temp++;
		}
		//注意只有左子树时
		if (temp<right)
		{
			for (int i=temp;i<right ;i++ )
		   {
			if (arr[i]<=root)
			{
				return false;
			}
		   }
		   
		}
		return checkHelper(arr,left,temp-1)&&checkHelper(arr,temp,right-1);
	}
	public static boolean check(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return false;
		}
		return checkHelper(arr,0,arr.length-1);
	}
	public static boolean verify(int[] arr,int start,int end)
	{
		int i=start;
		int root=arr[end];
		for (;i<end ;i++ )
		{
			if (arr[i]>root)
			{
				break;
			}
		}
		for (int j=i;j<end ;j++ )
		{
			if (arr[j]<=root)
			{
				return false;
			}
		}
		boolean checkLeft=true;
		if (i>start)
		{
			checkLeft=verify(arr,start,i-1);
		}
		boolean checkRight=true;
		if (i<end)
		{
			checkRight=verify(arr,i,end-1);
		}
		return checkLeft&&checkRight;
	}
	public static boolean verify(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return false;
		}
		return verify(arr,0,arr.length-1);
	}
	public static void main(String[] args) 
	{
		int[] arr={5,7,6,9,11,10,8};
		arr=new int[]{1,2,3,1,1,6,4};
		arr=new int[]{7};
		System.out.println(verify(arr));
		System.out.println(check(arr));
	}
}
