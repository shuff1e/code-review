import java.util.*;
public class Test51
{
	public static int getNumber(int[] arr)
	{
		if (arr==null||arr.length<=1)
		{
			return 0;
		}
		RankNode root=new RankNode(arr[0]);
		for (int i=1;i<arr.length ;i++ )
		{
			root.insert(arr[i]);
		}
		return number(root,root);
	}
	public static int number(RankNode fuck,RankNode root)
	{
		int sum=root.getRank(fuck.data);
		if (fuck.left!=null)
		{
			sum+=number(fuck.left,root);
		}
		if (fuck.right!=null)
		{
			sum+=number(fuck.right,root);
		}
		return sum;
	}
	public static void main(String[] args) 
	{
		int[] arr=new int[]{7};
		//System.out.println(getNumber(arr));
		//int a=inverse(arr);
		System.out.println(inverse(arr));
	}
	public static int inverse(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return 0;
		}
		int[] helper=new int[arr.length];
		return inverseHelper(arr,helper,0,arr.length-1);
	}
	public static int inverseHelper(int[] arr,int[] helper,int start,int end)
	{
		if (start==end)
		{
			return 0;
		}
		int mid=(start+end)/2;
		int left=inverseHelper(arr,helper,start,mid);
		int right=inverseHelper(arr,helper,mid+1,end);
		int i=mid;
		int j=end;
		int index=end;
		int count=0;
		while (i>=start&&j>=mid+1)
		{
			if (arr[i]>arr[j])
			{
				helper[index--]=arr[i--];
				count+=j-mid;
			}
			else
			{
				helper[index--]=arr[j--];
			}
		}
		while (i>=start)
		{
			helper[index--]=arr[i--];
		}
		while (j>=mid+1)
		{
			helper[index--]=arr[j--];
		}
		for (int k=start;k<=end ;k++ )
		{
			arr[k]=helper[k];
		}
		return left+right+count;
	}
}
