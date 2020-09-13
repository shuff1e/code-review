public class Test11 
{
	public static int min(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			throw new RuntimeException();
		}
		int index=0;
		int left=0;
		int right=arr.length-1;
		while (arr[left]>=arr[right])
		{
			if (right-left==1)
			{
				index=right;
				break;
			}
			int mid=(left+right)/2;
			if (arr[left]==arr[right]&&arr[left]==arr[mid])
			{
				return minHelper(arr,left,right);
			}
			if (arr[mid]>=arr[left])
			{
				left=mid;
			}
			else if (arr[mid]<=arr[right])
			{
				right=mid;
			}
		}
		return arr[index];
	}
	public static int minHelper(int[] arr,int p1,int p2)
	{
		int min=arr[p1];
		for (int i=p1+1;i<=p2 ;i++ )
		{
			if (min>arr[i])
			{
				min=arr[i];
			}
		}
		return min;
	}
	public static int search(int[] arr,int left,int right,int x)
	{
		if (right<left)
		{
			return -1;
		}
		int mid=(left+right)/2;
		if (arr[mid]==x)
		{
			return mid;
		}
		if (arr[left]<arr[mid])
		{
			if (arr[left]<=x&&x<arr[mid])
			{
				return search(arr,left,mid-1,x);
			}
			else
				return search(arr,mid+1,right,x);
		}
		else if (arr[left]>arr[mid])
		{
			if (arr[mid]<x&&x<=arr[right])
			{
				return search(arr,mid+1,right,x);
			}
			else
				return search(arr,left,mid-1,x);
		}
		else if (arr[left]==arr[mid])
		{
			if (arr[mid]!=arr[right])
			{
				return search(arr,mid+1,right,x);
			}
			else
			{
				int result=search(arr,left,mid-1,x);
				if (result==-1)
				{
					result=search(arr,mid+1,right,x);
				}
				return result;
			}
		}
		return -1;
	}
	public static void main(String[] args) 
	{
		int[] arr=new int[]{5,6,7,8,2,3,4};
		arr=new int[]{1,1,1,0,1};
		arr=new int[]{1,0,1,1,1};
		System.out.println(min(arr));
		//System.out.println(search(arr,0,arr.length-1,2));
	}
}
