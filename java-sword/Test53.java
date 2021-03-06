public class Test53 
{
	public static int getFirst(int[] arr,int start,int end,int k)
	{
		if (start>end)
		{
			return -1;
		}
		int mid=(start+end)/2;
		if (arr[mid]==k)
		{
			if (mid>start&&arr[mid-1]!=k||mid==start)
			{
				return mid;
			}
			else
				end=mid-1;
		}
		else if (arr[mid]<k)
		{
			start=mid+1;
		}
		else
			end=mid-1;
		return getFirst(arr,start,end,k);
	}
	public static int getLast(int[] arr,int start,int end,int k)
	{
		if (start>end)
		{
			return -1;
		}
		int mid=(start+end)/2;
		if (arr[mid]==k)
		{
			if (mid<end&&arr[mid+1]!=k||mid==end)
			{
				return mid;
			}
			else
				start=mid+1;
		}
		else if (arr[mid]>k)
		{
			end=mid-1;
		}
		else
			start=mid+1;
		return getLast(arr,start,end,k);
	}
	public static int getNumber(int[] arr,int k)
	{
		if (arr==null||arr.length==0)
		{
			return -1;
		}
		int number=-1;
		int first=getFirst(arr,0,arr.length-1,k);
		int last=getLast(arr,0,arr.length-1,k);
		if (first>-1&&last>-1)
		{
			number=last-first+1;
		}
		return number;
	}
	public static void main(String[] args) 
	{
		int[] arr={1,2,3,3,3,3,4,5};
		arr=null;
		arr=new int[]{0,1,2,3};
		//System.out.println(getMissingNumber(arr));
		arr=new int[]{-1,0,1,3};
		System.out.println(getNumberSameAsIndex(arr));
	}
	// 如果中间元素下标和中间元素的值相等，则缺失的数在右半边
	// 如果中间元素大于下标，但是前面一个数和下标相等，则这个就是缺失的数
	// 如果中间元素大于下标，前面一个数也大于下标，则缺失的数在左边
	public static int getMissingNumber(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return -1;
		}
		int left=0;
		int right=arr.length-1;
		while (left<=right)
		{
			int mid=(left+right)/2;
			if (arr[mid]!=mid)
			{
				if (mid==0||arr[mid-1]==mid-1)
				{
					return mid;
				}
				else
					right=mid-1;
			}
			else 
				left=mid+1;
		}
		if (left==arr.length)
		{
			return left;
		}
		return -1;
	}
	public static int getNumberSameAsIndex(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return -1;
		}
		int left=0;
		int right=arr.length-1;
		while (left<=right)
		{
			int mid=left+((right-left)>>1);
			if (arr[mid]==mid)
			{
				return mid;
			}
			else if (arr[mid]>mid)
			{
				right=mid-1;
			}
			else 
				left=mid+1;
		}
		return -1;
	}
}
