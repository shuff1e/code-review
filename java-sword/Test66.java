import java.util.*;
public class Test66 
{
	public static void multiply(int[] arr1,int[] arr2)
	{
		if (arr1.length==arr2.length&&(arr1.length>1))
		{
			arr2[0]=1;
			int temp1=1;
			for (int i=1;i<arr1.length ;i++ )
			{
				temp1*=arr1[i-1];
				arr2[i]=temp1;
			}
			int temp2=1;
			for (int i=arr1.length-2;i>=0 ;i-- )
			{
				temp2*=arr1[i+1];
				arr2[i]*=temp2;
			}
		}
	}
	// 这不是把迭代改成递归了吗
	public static void multiply2(int[] arr1,int[] arr2)
	{
		if (arr1.length==arr2.length&&arr1.length>1)
		{
			for (int i=0;i<arr2.length ;i++ )
		    {
			     int c=leftHelper(arr1,i);
		         int d=rightHelper(arr1,i);
			     arr2[i]=c*d;
		    }
		}
	}
	public static int leftHelper(int[] arr1,int index)
	{
		if (index==0)
		{
			return 1;
		}
		return leftHelper(arr1,index-1)*arr1[index-1];
	}
	public static int rightHelper(int[] arr1,int index)
	{
		if (index==arr1.length-1)
		{
			return 1;
		}
		return rightHelper(arr1,index+1)*arr1[index+1];
	}
	public static void main(String[] args) 
	{
		int[] arr1=new int[]{1,2,3,4,0,5,5,5};
		int[] arr2=new int[arr1.length];
		multiply(arr1,arr2);
		System.out.println(Arrays.toString(arr2));
		arr2=new int[arr1.length];
		multiply2(arr1,arr2);
		System.out.println(Arrays.toString(arr2));
	}
}
